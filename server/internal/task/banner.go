package task

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"go.uber.org/zap"
	"log"
)

type Banner struct {
	Id     string
	Trials int
	Reward int
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

type Percentage struct {
	id    int
	start int
	end   int
}

func ReturnBanner(db *pg.DB, slog *zap.SugaredLogger, audience, slot string) (string, error) {
	banners := getBannerStat(db, audience, slot)
	fmt.Println(banners)
	return "", nil
}

func Run(db *pg.DB) {
	banners := getBannerStat(db, "", "")
	for i := 0; i < 100000; i++ {
		bId, err := getBanner(&banners)
		if err != nil {
			log.Fatal(err)
		}
		var rew bool
		if randomClick() {
			rew = true
		}
		incBannerStatistics(&banners, bId, rew)
	}
	var statRows []*models.Statistics
	for _, v := range banners.Banners {
		row := models.Statistics{
			AudienceId: "2",
			BannerId:   v.Id,
			SlotId:     "2",
			Clicks:     uint64(v.Reward),
			Shows:      uint64(v.Trials),
		}
		statRows = append(statRows, &row)
	}
	insertIntoStat(db, statRows)
	fmt.Println("finish")
}

//
//func getTestRes() []float64 {
//	var sl []float64
//	for _, v := range bnrs.Banners {
//		var rs float64
//		if v.Trials == 0 {
//			rs = 0.5
//		} else {
//			//profit :=
//			//rs = profit
//			rs = (float64(v.Reward) / float64(v.Trials)) + math.Sqrt(math.Log(float64(bnrs.Count))/float64(v.Trials))
//		}
//		sl = append(sl, rs)
//	}
//	//fmt.Println(sl)
//	incBannersCount()
//	return sl
//}
//
//func getBanners() Banners {
//	return bnrs
//}
