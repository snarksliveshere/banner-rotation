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

func main() {
	mm := make(map[int]int, 3)
	var id1, id2, id3 int
	for i := 0; i < 1000; i++ {
		id := getPerc()
		if id == 1 {
			id1++
			mm[id] = id1
		} else if id == 2 {
			id2++
			mm[id] = id2
		} else {
			id3++
			mm[id] = id3
		}

	}
	fmt.Printf("map: %#v", mm)
}

// выбрать баннер для показа. без учета дем.группы и слота
func getPerc() int {
	bs := []Banner{
		{
			Id:     1,
			Trials: 30,
			Reward: 4,
		},
		{
			Id:     2,
			Trials: 30,
			Reward: 3,
		},
		{
			Id:     3,
			Trials: 30,
			Reward: 6,
		},
	}
	banners := Banners{Count: 90, Banners: bs}
	var num int
	var percentage []Percentage
	for _, v := range banners.Banners {
		profit := float64(v.Reward) / float64(v.Trials)
		rs := profit + math.Sqrt(math.Log(float64(banners.Count))/float64(v.Trials))
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
	//fmt.Printf("percentage: %#v\n", percentage)
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(num + 1)
	//fmt.Println(num, rnd)
	var id int
	for _, v := range percentage {
		if rnd >= v.start && rnd <= v.end {
			id = v.id
		}
	}

	return id
}

// TODO: я должен показать баннер и учесть, что на него может быть склик
//val[i] = x_mean[i] + math.Sqrt(math.Log(float64(agent.Trials))/(2*float64(arm[i].Count)))
