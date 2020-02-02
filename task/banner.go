package task

import (
	"fmt"
	"github.com/go-pg/pg"
	"log"
	"math"
)

type Banner struct {
	Id     int
	Trials int
	Reward int
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

var m1 map[int]int // clicks
var m2 map[int]int // shows
var bannerMap map[int]Banner
var bnrs Banners

func Run(db *pg.DB) {
	banners := getBannerStat(db)

	m1 = make(map[int]int, 100)
	m2 = make(map[int]int, 100)
	bannerMap = make(map[int]Banner, 10)
	for i := 0; i < 1000; i++ {
		// TODO: len(banners)
		bId, err := getBanner(&banners)
		if err != nil {
			log.Fatal(err)
		}
		//if _, ok := m2[bId]; ok {
		//	m2[bId] = m2[bId] + 1
		//} else {
		//	m2[bId] = 1
		//}
		b := Banner{Id: bId}

		bm, ok := bannerMap[bId]
		if ok {
			b.Trials = bm.Trials + 1
		}

		var rew bool
		if randomClick() {
			rew = true
			if ok {
				b.Reward = bannerMap[bId].Reward + 1
			}
		}
		bannerMap[bId] = b
		//
		//sl := getTestRes()
		//if i == 1 || i == 9999999{
		//	fmt.Println(sl)
		//	fmt.Println(bnrs.Count)
		//}
		//perc, num := getPercentage()
		//id, rew := choose(perc, num)
		//incBannerStat(bId, rew)
		incBannerStatistics(&banners, bId, rew)
	}

	//insertIntoStat(db)
	fmt.Println(banners)
	fmt.Println(bannerMap)
	//fmt.Println(m2)
	fmt.Println("olala")

}

func getTestRes() []float64 {
	var sl []float64
	for _, v := range bnrs.Banners {
		var rs float64
		if v.Trials == 0 {
			rs = 0.5
		} else {
			//profit :=
			//rs = profit
			rs = (float64(v.Reward) / float64(v.Trials)) + math.Sqrt(math.Log(float64(bnrs.Count))/float64(v.Trials))
		}
		sl = append(sl, rs)
	}
	//fmt.Println(sl)
	incBannersCount()
	return sl
}

func getBanners() Banners {
	return bnrs
}
