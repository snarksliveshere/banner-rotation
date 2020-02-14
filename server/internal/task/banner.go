package task

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/internal/database/models"
	"log"
)

type Banner struct {
	Id     int
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

func Run(db *pg.DB) {
	banners := getBannerStat(db)
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
			AudienceFK: 2,
			BannerFK:   uint64(v.Id),
			SlotFK:     2,
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
