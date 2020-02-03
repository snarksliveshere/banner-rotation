package task

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/models"
	"log"
)

func getAudience(db *pg.DB) []uint64 {
	var loadedRows []*models.Audience
	err := db.Model(&loadedRows).
		Column("id").
		Select()
	if err != nil {
		log.Fatal(err)
	}
	sl := make([]uint64, 0, len(loadedRows))
	for _, v := range loadedRows {
		sl = append(sl, v.Id)
	}
	return sl
}

func getBannerStat(db *pg.DB) Banners {
	var loadedRows []*models.Statistics
	query := `SELECT DISTINCT statistics.clicks, statistics.shows, statistics.banner_fk
				FROM statistics
				RIGHT JOIN audience2banner a2b ON statistics.audience_fk = a2b.audience_fk
				WHERE a2b.audience_fk = ?
;`
	_, err := db.Query(&loadedRows, query, 1)
	//err := db.Model(&loadedRows).
	//	Column("statistics.clicks", "statistics.shows", "a2b.banner_fk").
	//	Join("RIGHT JOIN audience2banner a2b ON statistics.audience_fk = a2b.audience_fk").
	//	Where("a2b.audience_fk = ?", 1).
	//	Group("statistics.clicks", "statistics.shows", "a2b.banner_fk").
	//	Select()

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

func insertIntoStat(db *pg.DB) {
	var loadedRows []*models.Statistics

	for i, v := range m2 {
		var row models.Statistics
		row.Shows = uint64(v)
		row.Clicks = uint64(m1[i])
		row.BannerFK = uint64(i)
		row.AudienceFK = 1
		loadedRows = append(loadedRows, &row)
	}

	//err := db.Insert(&loadedRows)
	_, err := db.Model(&loadedRows).
		OnConflict("(audience_fk, banner_fk, slot_fk) DO UPDATE").
		Set("clicks = EXCLUDED.clicks").
		Set("shows = EXCLUDED.shows").
		Insert()

	if err != nil {
		log.Fatal(err)
	}

}
