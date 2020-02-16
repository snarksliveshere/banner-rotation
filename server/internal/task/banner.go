package task

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"github.com/snarksliveshere/banner-rotation/server/internal/database"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

type Banner struct {
	Id     string
	Shows  int
	Clicks int
}

type BannerStat struct {
	Id       int
	Trials   int
	Reward   int
	Slot     int
	Audience int
}

type Banners struct {
	Banners []Banner
	Count   int
}

type BannerStatistics struct {
	Type     string `json:"type"`
	Slot     string `json:"slot"`
	Audience string `json:"audience"`
	Banner   string `json:"banner"`
	Time     string `json:"time"`
}

func ReturnBanner(db *pg.DB, slog *zap.SugaredLogger, channel *amqp.Channel, audience, slot string) (string, error) {
	bannersRows, err := database.GetBannerStat(db, audience, slot)
	if err != nil {
		return "", err
	}
	banners, err := getBanners(bannersRows)
	if err != nil {
		return "", err
	}
	banner, err := GetBanner(&banners)
	if err != nil {
		return "", err
	}
	row := &models.Statistics{
		AudienceId: audience,
		BannerId:   banner.Id,
		SlotId:     slot,
		Clicks:     uint64(banner.Clicks),
		Shows:      uint64(banner.Shows) + 1,
	}
	err = database.InsertRowIntoStat(db, slog, row)
	if err != nil {
		return "", err
	}
	statToRabbit := &BannerStatistics{
		Type:     configs.BannerStatShow,
		Slot:     slot,
		Audience: audience,
		Banner:   banner.Id,
		Time:     time.Now().Format(configs.EventTimeLayout),
	}

	data, err := json.Marshal(statToRabbit)
	if err != nil {
		return "", err
	}
	err = bannerStatToRabbit(channel, configs.BannerStatQueue, data)
	if err != nil {
		return "", err
	}
	return banner.Id, nil
}

func AddClickToBanner(db *pg.DB, channel *amqp.Channel, banner, slot, audience string) error {
	err := database.AddClick(db, banner, slot, audience)
	if err != nil {
		return err
	}
	statToRabbit := &BannerStatistics{
		Type:     configs.BannerStatClick,
		Slot:     slot,
		Audience: audience,
		Banner:   banner,
		Time:     time.Now().Format(configs.EventTimeLayout),
	}
	data, err := json.Marshal(statToRabbit)
	if err != nil {
		return err
	}
	err = bannerStatToRabbit(channel, configs.BannerStatQueue, data)
	if err != nil {
		return err
	}
	return nil
}

func AddBannerToSlot(db *pg.DB, banner, slot string) error {
	err := database.AddBannerToSlot(db, banner, slot)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBannerFromSlot(db *pg.DB, banner, slot string) error {
	err := database.DeleteBannerToSlot(db, banner, slot)
	if err != nil {
		return err
	}
	return nil
}

func bannerStatToRabbit(ch *amqp.Channel, rk string, stat []byte) error {
	err := ch.Publish(
		"",    // exchange
		rk,    // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         stat,
		})
	if err != nil {
		return err
	}
	return nil
}

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
