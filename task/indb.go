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

func getBannerStat(db *pg.DB) {
	var loadedRows []*models.Statistics
	err := db.Model(&loadedRows).
		Column("statistics.clicks", "statistics.shows", "statistics.id", "a2b.banner_fk").
		Join("RIGHT JOIN audience2banner a2b ON statistics.audience_fk = a2b.audience_fk").
		Where("a2b.audience_fk = ?", 1).
		Select()

	if err != nil {
		log.Fatal(err)
	}
	var bsInit []Banner
	var count int
	for _, v := range loadedRows {
		count = count + int(v.Clicks)
		b := Banner{
			Id:     int(v.BannerFK),
			Trials: int(v.Clicks),
			Reward: int(v.Shows),
		}
		bsInit = append(bsInit, b)
	}
	bnrs = Banners{Count: count, Banners: bsInit}
}
