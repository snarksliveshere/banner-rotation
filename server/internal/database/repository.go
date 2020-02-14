package database

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"go.uber.org/zap"
)

func InsertRowIntoStat(db *pg.DB, log *zap.SugaredLogger, loadedRow *models.Statistics) {
	_, err := db.Model(loadedRow).
		OnConflict("(audience_id, banner_id, slot_id) DO UPDATE").
		Set("clicks = EXCLUDED.clicks").
		Set("shows = EXCLUDED.shows").
		Insert()
	if err != nil {
		log.Error(err.Error())
	}
}

func GetBannerStat(db *pg.DB, audience, slot string) ([]*models.Statistics, error) {
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
		return nil, err
	}
	return loadedRows, nil
}
