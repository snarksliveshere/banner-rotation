package database

import (
	"errors"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
)

func InsertRowIntoStat(db *pg.DB, loadedRow *models.Statistics) error {
	_, err := db.Model(loadedRow).
		OnConflict("(audience_id, banner_id, slot_id) DO UPDATE").
		Set("clicks = EXCLUDED.clicks").
		Set("shows = EXCLUDED.shows").
		Insert()
	if err != nil {
		return err
	}
	return nil
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

func AddBannerToSlot(db *pg.DB, banner, slot string) error {
	var row *models.Banner2Slot
	query := `INSERT INTO banner2slot (banner_fk, slot_fk)
			  VALUES 
			  (
				(SELECT id FROM banner WHERE banner_id = ?),
				(SELECT id FROM slot WHERE slot_id = ?)
			  );
			`

	res, err := db.Query(row, query, banner, slot)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("there is no result in AddBannerToSlot method")
	}
	if res.RowsAffected() == 0 {
		return errors.New("there is no affected rows in AddBannerToSlot method")
	}

	return nil
}

func DeleteBannerFromSlot(db *pg.DB, banner, slot string) error {
	var row *models.Banner2Slot
	query := `DELETE FROM banner2slot
			  WHERE banner_fk = (SELECT id FROM banner WHERE banner_id = ?)  
			  AND slot_fk = (SELECT id FROM slot WHERE slot_id = ?)  
			  ;
			`

	res, err := db.Query(row, query, banner, slot)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("there is no result in DeleteBannerToSlot method")
	}
	if res.RowsAffected() == 0 {
		return errors.New("there is no affected rows in DeleteBannerToSlot method")
	}

	return nil
}
