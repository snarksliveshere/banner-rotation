package main

import (
	"github.com/snarksliveshere/banner-rotation/models"
	"github.com/snarksliveshere/banner-rotation/task"
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/snarksliveshere/banner-rotation/configs"
)

type TestBanner struct {
	Banner   uint64
	Audience []uint64
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	//var conf configs.AppConfig
	//failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	//var loadedRows []*models.Statistics
	//dbInst := configs.DB{Conf: &conf}
	//db := dbInst.CreatePgConn()
	//
	//_ = getAudience(db)
	//
	////err := db.Model(&loadedRows).
	////	Column("id", "banner_id").
	////	Select()
	//err := db.Model(&loadedRows).
	//	Column("statistics.clicks", "statistics.shows", "statistics.id", "a2b.banner_fk").
	//	Join("RIGHT JOIN audience2banner a2b ON statistics.audience_fk = a2b.audience_fk").
	//	Where("a2b.audience_fk = ?", 1).
	//	Select()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var countS int
	//
	//var bsInit []Banner
	//for _, v := range loadedRows {
	//	countS = countS + int(v.Clicks)
	//	b := Banner{
	//		Id:     int(v.BannerFK),
	//		Trials: int(v.Clicks),
	//		Reward: int(v.Shows),
	//	}
	//	bsInit = append(bsInit, b)
	//}
	//bnrs = Banners{Count: countS, Banners: bsInit}

	//tbs := getTestBanners(loadedRows)
	//fmt.Println(tbs)
	//fmt.Println("olala")

	//err := r.db.Model(&r.rows).
	//	Column("event.time", "event.title", "event.description", "event.time", "event.id", "event.date_fk").
	//	Join("JOIN calendar.calendar ON event.date_fk = calendar.id").
	//	Where("calendar.date >= ?", from).
	//	Where("calendar.date <= ?", till).
	//	Select()
	//bs := []Banner{
	//	{
	//		Id:     1,
	//		Trials: 20,
	//		Reward: 2,
	//	},
	//	{
	//		Id:     2,
	//		Trials: 30,
	//		Reward: 4,
	//	},
	//	{
	//		Id:     3,
	//		Trials: 20,
	//		Reward: 10,
	//	},
	//}
	//bs := []Banner{
	//	{
	//		Id:     1,
	//		Trials: 0,
	//		Reward: 0,
	//	},
	//	{
	//		Id:     2,
	//		Trials: 0,
	//		Reward: 0,
	//	},
	//	{
	//		Id:     3,
	//		Trials: 0,
	//		Reward: 0,
	//	},
	//}
	//var count int
	//for _, v := range bs {
	//	count += v.Trials
	//}
	//
	//bnrs = Banners{Count: count, Banners: bs}

}

func getTestBanners(loadedRows []*models.Banner) []TestBanner {
	tbs := make([]TestBanner, 0, 10)

	for _, v := range loadedRows {
		var c bool
		for _, i := range tbs {
			if i.Banner == v.Id {
				c = true
			}
		}
		if c {
			continue
		}
		b := TestBanner{Banner: v.Id}
		for _, k := range loadedRows {
			if k.Id != v.Id {
				continue
			}
			b.Audience = append(b.Audience, k.Audience)
		}
		tbs = append(tbs, b)
	}
	return tbs
}

func main() {
	var conf configs.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	dbInst := configs.DB{Conf: &conf}
	db := dbInst.CreatePgConn()
	task.Run(db)
}

//func testT() {
//	mm := make(map[int]int, 3)
//	var id1, id2, id3 int
//	for i := 0; i < 100; i++ {
//		id, _ := getPerc()
//		if id == 1 {
//
//			id1++
//			mm[id] = id1
//		} else if id == 2 {
//			id2++
//			mm[id] = id2
//		} else {
//			id3++
//			mm[id] = id3
//		}
//
//	}
//	fmt.Println(bnrs)
//	fmt.Println(mm)
//	fmt.Printf("map: %#v", mm)
//}

// TODO: я должен показать баннер и учесть, что на него может быть склик
//val[i] = x_mean[i] + math.Sqrt(math.Log(float64(agent.Trials))/(2*float64(arm[i].Count)))
