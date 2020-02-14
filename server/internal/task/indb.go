package task

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"log"
)

func getBanners(loadedRows []*models.Statistics) (Banners, error) {
	var bsInit []Banner
	var count int
	for _, v := range loadedRows {
		count = count + int(v.Shows)
		b := Banner{
			Id:     v.BannerId,
			Shows:  int(v.Shows),
			Clicks: int(v.Clicks),
		}
		bsInit = append(bsInit, b)
	}
	return Banners{Count: count, Banners: bsInit}, nil
}

func insertIntoStat(db *pg.DB, loadedRows []*models.Statistics) {
	_, err := db.Model(&loadedRows).
		OnConflict("(audience_id, banner_id, slot_id) DO UPDATE").
		Set("clicks = EXCLUDED.clicks").
		Set("shows = EXCLUDED.shows").
		Insert()

	if err != nil {
		log.Fatal(err)
	}
}

//func getAudience(db *pg.DB) []uint64 {
//	var loadedRows []*models.Audience
//	err := db.Model(&loadedRows).
//		Column("id").
//		Select()
//	if err != nil {
//		log.Fatal(err)
//	}
//	sl := make([]uint64, 0, len(loadedRows))
//	for _, v := range loadedRows {
//		sl = append(sl, v.Id)
//	}
//	return sl
//}
//
