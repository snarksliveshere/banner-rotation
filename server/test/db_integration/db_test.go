package db_integration_test

import (
	"github.com/go-pg/pg"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"github.com/snarksliveshere/banner-rotation/server/internal/database"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"github.com/snarksliveshere/banner-rotation/server/internal/task"
	"log"
	"os"
	"testing"
)

const (
	dummyAudience    = "male_adult"
	dummySlot        = "top_slot_id"
	dummyBanner      = "some_male2_adult_app_id"
	dummyBannerToAdd = "some_male2_kid_app_id"
)

var (
	db   *pg.DB
	conf configs.AppConfig
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func TestInsertRowIntoStat(t *testing.T) {
	// after dummy
	var row []*models.Statistics
	err := db.Model(&row).Select()
	if err != nil {
		t.Errorf("TestInsertRowIntoStat(), db select error:%v\n", err)
	}
	if len(row) != 3 {
		t.Errorf("TestInsertRowIntoStat(), db length error:%v\n", err)
	}
}

func TestAddClick(t *testing.T) {
	var row []*models.Statistics
	err := db.Model(&row).Column("clicks").
		Where("banner_id = ?", dummyBanner).
		Where("audience_id = ?", dummyAudience).
		Where("slot_id = ?", dummySlot).
		Select()
	if err != nil {
		t.Errorf("TestAddClick(), db select error:%v\n", err)
	}

	err = database.AddClick(db, dummyBanner, dummySlot, dummyAudience)
	if err != nil {
		t.Errorf("TestAddClick(), db AddClick error:%v\n", err)
	}
	var rowAfter []*models.Statistics
	err = db.Model(&rowAfter).Column("clicks").
		Where("banner_id = ?", dummyBanner).
		Where("audience_id = ?", dummyAudience).
		Where("slot_id = ?", dummySlot).
		Select()
	if err != nil {
		t.Errorf("TestAddClick(), db after select error:%v\n", err)
	}
	if len(row) == 0 || len(rowAfter) == 0 {
		t.Errorf("TestAddClick(), unappropriate length=row:%d,rowAfter=%d\n", len(row), len(rowAfter))
	}

	if row[0].Clicks == rowAfter[0].Clicks {
		t.Errorf("TestAddClick(), equal clicks:%d\n", row[0].Clicks)
	}

	if (row[0].Clicks + 1) != rowAfter[0].Clicks {
		t.Errorf("TestAddClick(), wrong num of clicks=init:%v,after:%v\n", row[0].Clicks, rowAfter[0].Clicks)
	}
}

func TestAddBannerToSlot(t *testing.T) {
	err := database.AddBannerToSlot(db, dummyBannerToAdd, dummySlot)
	if err != nil {
		t.Errorf("TestAddBannerToSlot(), db insert error:%v\n", err)
	}
	err = database.AddBannerToSlot(db, dummyBannerToAdd, dummySlot)
	if err == nil {
		t.Errorf("TestAddBannerToSlot(), db second insert without error\n")
	}

	var row models.Banner2Slot
	query := `SELECT banner_fk, slot_fk FROM banner2slot 
			  WHERE banner_fk = (SELECT id FROM banner WHERE banner_id = ?)
			  AND slot_fk = (SELECT id FROM slot WHERE slot_id = ?) 		
			  ;
			`

	_, err = db.Query(&row, query, dummyBannerToAdd, dummySlot)
	if err != nil {
		t.Errorf("TestAddBannerToSlot(), db select check with error:%v\n", err)
	}
	if row.BannerFK == 0 || row.SlotFK == 0 {
		t.Errorf("TestAddBannerToSlot(), bad result from select:\n")
	}
}

func TestDeleteBannerFromSlot(t *testing.T) {
	err := database.DeleteBannerFromSlot(db, dummyBannerToAdd, dummySlot)
	if err != nil {
		t.Errorf("TestDeleteBannerFromSlot(), db delete error:%v\n", err)
	}
	err = database.DeleteBannerFromSlot(db, dummyBannerToAdd, dummySlot)
	if err == nil {
		t.Errorf("TestDeleteBannerFromSlot(), db second delete without error\n")
	}

	var row models.Banner2Slot
	query := `SELECT banner_fk, slot_fk FROM banner2slot
			  WHERE banner_fk = (SELECT id FROM banner WHERE banner_id = ?)
			  AND slot_fk = (SELECT id FROM slot WHERE slot_id = ?)
			  ;
			`

	_, err = db.Query(&row, query, dummyBannerToAdd, dummySlot)
	if err != nil {
		t.Errorf("TestDeleteBannerFromSlot(), db select check with error:%v\n", err)
	}
	if row.BannerFK != 0 || row.SlotFK != 0 {
		t.Errorf("TestDeleteBannerFromSlot(), bad result from select:\n")
	}
}

func setup() {
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	dbInst := configs.DB{Conf: &conf}
	db = dbInst.CreatePgConn()

	for i := 0; i < 10; i++ {
		setDummy()
	}
}

func setDummy() {
	bannersRows, err := database.GetBannerStat(db, dummyAudience, dummySlot)
	if err != nil {
		log.Fatalf("setDummmy(), err on GetBannerStat:%v\n", err)
	}
	banners, err := getBanners(bannersRows)
	if err != nil {
		log.Fatalf("setDummmy(), err on getBanners:%v\n", err)
	}
	banner, err := task.GetBanner(&banners)
	if err != nil {
		log.Fatalf("setDummmy(), err on GetBanner:%v\n", err)
	}
	row := &models.Statistics{
		AudienceId: dummyAudience,
		BannerId:   banner.Id,
		SlotId:     dummySlot,
		Clicks:     uint64(banner.Clicks),
		Shows:      uint64(banner.Shows) + 1,
	}
	err = database.InsertRowIntoStat(db, row)
}

func shutdown() {
	defer func() { _ = db.Close() }()
}

func getBanners(loadedRows []*models.Statistics) (task.Banners, error) {
	var bsInit []task.Banner
	var count int
	for _, v := range loadedRows {
		count = count + int(v.Shows)
		b := task.Banner{
			Id:     v.BannerId,
			Shows:  int(v.Shows),
			Clicks: int(v.Clicks),
		}
		bsInit = append(bsInit, b)
	}
	return task.Banners{Count: count, Banners: bsInit}, nil
}
