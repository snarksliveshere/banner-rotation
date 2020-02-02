package task

import (
	"math"
	"math/rand"
	"time"
)

func getBanner(banners *Banners) int {
	var rs float64
	var bId int
	for _, v := range banners.Banners {
		if v.Trials == 0 {
			bId = v.Id
			break
		} else {
			profit := float64(v.Reward) / float64(v.Trials)
			res := profit + math.Sqrt(math.Log(float64(banners.Count))/float64(v.Trials))
			if res > rs {
				bId = v.Id
				rs = res
			}
		}
	}
	return bId
}

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

func incBannersCount() {
	banners := getBanners()
	banners.Count++
	bnrs = banners
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

func incBannerStatistics(banners *Banners, id int, rew bool) {
	for k, v := range banners.Banners {
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
			banners.Banners[k] = b
			//bnrs.Banners[k] = b

		}
	}
	banners.Count++

}

//func incBannerStat(id int, rew bool) {
//	for k, v := range bnrs.Banners {
//		if v.Id == id {
//			tr := v.Trials + 1
//			b := Banner{
//				Id:     v.Id,
//				Trials: tr,
//				Reward: v.Reward,
//			}
//
//			//v.Trials++
//
//			if rew {
//				r := v.Reward + 1
//				b.Reward = r
//			}
//			bnrs.Banners[k] = b
//
//		}
//	}
//
//}

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

func randomClick() bool {
	rand.Seed(time.Now().UnixNano())
	b := rand.Float64()
	if b > 0.9 {
		return true
	} else {
		return false
	}
}
