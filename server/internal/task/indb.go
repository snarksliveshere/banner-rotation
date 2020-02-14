package task

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"log"
)

func getBannerStat(db *pg.DB) Banners {
	var loadedRows []*models.Statistics
	_ = `SELECT DISTINCT statistics.clicks, statistics.shows, statistics.banner_fk
				FROM statistics
				RIGHT JOIN audience2banner a2b ON statistics.audience_fk = a2b.audience_fk
				WHERE a2b.audience_fk = ?
;`
	query := `SELECT banner.id AS banner_fk, shows, clicks FROM banner
		JOIN audience2banner a2b ON banner.id = a2b.banner_fk
		JOIN banner2slot b2s ON a2b.banner_fk = b2s.banner_fk
		LEFT JOIN statistics s ON banner.id = s.banner_fk
		WHERE a2b.audience_fk = 2
		AND b2s.slot_fk = 2`
	_, err := db.Query(&loadedRows, query)

	if err != nil {
		log.Fatal(err)
	}
	var bsInit []Banner
	var count int
	for _, v := range loadedRows {
		count = count + int(v.Clicks)
		b := Banner{
			Id:     int(v.BannerFK),
			Trials: int(v.Shows),
			Reward: int(v.Clicks),
		}
		bsInit = append(bsInit, b)
	}
	return Banners{Count: count, Banners: bsInit}
}

func insertIntoStat(db *pg.DB, loadedRows []*models.Statistics) {
	_, err := db.Model(&loadedRows).
		OnConflict("(audience_fk, banner_fk, slot_fk) DO UPDATE").
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
