package db_integration_test

import (
	"errors"
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
	dummyAudience = "male_adult"
	dummySlot     = "top_slot_id"
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

func TestAddClick(t *testing.T) {
	t.Errorf("TestGetBanner(), equal shows on distance:%d\n", 1)

}

func AddClick(db *pg.DB, banner, slot, audience string) error {
	var row *models.Statistics
	query := `  UPDATE statistics SET clicks = clicks + 1
				WHERE banner_id = ?
				AND slot_id = ?
				AND audience_id = ?;
			`

	res, err := db.Query(row, query, banner, slot, audience)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("there is no result in addClick method")
	}
	if res.RowsAffected() == 0 {
		return errors.New("there is no affected rows in addClick method")
	}

	return nil
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
