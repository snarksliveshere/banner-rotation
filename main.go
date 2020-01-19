package main

import (
	"fmt"
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

func main() {
	bs := []Banner{
		{
			Id:     1,
			Trials: 20,
			Reward: 2,
		},
		{
			Id:     2,
			Trials: 30,
			Reward: 4,
		},
		{
			Id:     3,
			Trials: 20,
			Reward: 6,
		},
	}
	banners := Banners{Count: 70, Banners: bs}
	for _, v := range banners.Banners {
		profit := float64(v.Reward) / float64(v.Trials)
		rs := profit + math.Sqrt(math.Log(float64(banners.Count))/float64(v.Trials))
		fmt.Println(rs, v.Id)

	}
}

//val[i] = x_mean[i] + math.Sqrt(math.Log(float64(agent.Trials))/(2*float64(arm[i].Count)))
