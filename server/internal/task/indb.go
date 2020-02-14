package task

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"log"
)

func getBannerStat(db *pg.DB, audience, slot string) (Banners, error) {
	var loadedRows []*models.Statistics
	query := `SELECT banner.banner_id AS banner_id, a.audience_id, sl.slot_id,  shows, clicks
			FROM banner
				 JOIN audience2banner a2b ON banner.id = a2b.banner_fk
				 JOIN audience a ON a2b.audience_fk = a.id
				 JOIN banner2slot b2s ON a2b.banner_fk = b2s.banner_fk
				 JOIN slot sl ON sl.id = b2s.slot_fk
				 LEFT JOIN statistics s ON banner.banner_id = s.banner_id
			WHERE a.audience_id = ?
  			AND sl.slot_id = ?;`

	_, err := db.Query(&loadedRows, query, audience, slot)

	if err != nil {
		return Banners{}, err
	}
	var bsInit []Banner
	var count int
	for _, v := range loadedRows {
		count = count + int(v.Clicks)
		b := Banner{
			Id:     v.BannerId,
			Trials: int(v.Shows),
			Reward: int(v.Clicks),
		}
		bsInit = append(bsInit, b)
	}
	return Banners{Count: count, Banners: bsInit}, nil
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
