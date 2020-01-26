package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

var m1 map[int]int
var m2 map[int]int
var bnrs Banners

func init() {
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
	bs := []Banner{
		{
			Id:     1,
			Trials: 0,
			Reward: 0,
		},
		{
			Id:     2,
			Trials: 0,
			Reward: 0,
		},
		{
			Id:     3,
			Trials: 0,
			Reward: 0,
		},
	}
	var count int
	for _, v := range bs {
		count += v.Trials
	}

	bnrs = Banners{Count: count, Banners: bs}
	m1 = make(map[int]int, 1000)
	m2 = make(map[int]int, 1000)
}

func getBanners() Banners {
	return bnrs
}

func incBannersCount() {
	banners := getBanners()
	banners.Count++
	bnrs = banners
}

func main() {
	for i := 0; i < 1000; i++ {
		perc, num := getPercentage()
		id, rew := choose(perc, num)
		incBannerStat(id, rew)
	}
	fmt.Println(bnrs)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println("olala")

}

func incBannerStat(id int, rew bool) {

	for k, v := range bnrs.Banners {
		if v.Id == id {
			tr := v.Trials + 1
			b := Banner{
				Id:     v.Id,
				Trials: tr,
				Reward: v.Reward,
			}

			//v.Trials++

			if rew {
				r := v.Reward + 1
				b.Reward = r
			}
			bnrs.Banners[k] = b

		}
	}

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

func getPercentage() ([]Percentage, int) {
	var num int
	var percentage []Percentage
	for _, v := range bnrs.Banners {
		var rs float64
		if v.Trials == 0 {
			rs = 0.5
		} else {
			profit := float64(v.Reward) / float64(v.Trials)
			rs = profit + math.Sqrt(math.Log(float64(bnrs.Count))/float64(v.Trials))
		}

		//fmt.Println(rs, v.Id, int(rs*100))
		var p Percentage
		p.id = v.Id
		l := len(percentage)
		if l > 0 {
			p.start = percentage[l-1].end + 1
		} else {
			p.start = num
		}
		nn := int(rs * 100)
		if l > 0 {
			num += nn + 1
			p.end = num
		} else {
			num += nn
			p.end = num
		}

		percentage = append(percentage, p)

	}
	return percentage, num
}

func randomClick() bool {
	rand.Seed(time.Now().UnixNano())
	b := rand.Float64()
	if b > 0.5 {
		return true
	} else {
		return false
	}
}

func choose(percentage []Percentage, num int) (int, bool) {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(num + 1)
	var id int
	var reward bool
	for _, v := range percentage {
		if rnd >= v.start && rnd <= v.end {
			id = v.id
			m2[id] = m2[id] + 1

			if randomClick() {
				reward = true
				m1[id] = m1[id] + 1
			}

		}
	}

	incBannersCount()

	return id, reward
}

// TODO: я должен показать баннер и учесть, что на него может быть склик
//val[i] = x_mean[i] + math.Sqrt(math.Log(float64(agent.Trials))/(2*float64(arm[i].Count)))
