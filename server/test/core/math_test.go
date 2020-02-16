package core_test

import (
	"github.com/snarksliveshere/banner-rotation/server/internal/task"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestGetBanner(t *testing.T) {
	banners := prepareBanners()
	control := prepareBanners()

	for i := 0; i < 1000000; i++ {
		banner, err := task.GetBanner(&banners)
		if err != nil {
			log.Fatal(err)
		}
		var rew bool
		if randomClick() {
			rew = true
		}
		incBannerStatistics(&banners, banner.Id, rew)
	}

	for k, v := range banners.Banners {
		if v.Shows == control.Banners[k].Shows {
			t.Errorf("TestGetBanner(), equal shows on distance:%d\n", v.Shows)
		}
	}
}

func prepareBanners() task.Banners {
	b1 := task.Banner{
		Id:     "popular_banner_1",
		Shows:  150,
		Clicks: 20,
	}
	b2 := task.Banner{
		Id:     "popular_banner_2",
		Shows:  400,
		Clicks: 60,
	}
	b3 := task.Banner{
		Id:     "low_banner_1",
		Shows:  400,
		Clicks: 5,
	}
	b4 := task.Banner{
		Id:     "extra_low_banner_1",
		Shows:  800,
		Clicks: 0,
	}
	b5 := task.Banner{
		Id:     "extra_low_banner_2",
		Shows:  1500,
		Clicks: 0,
	}

	bs := make([]task.Banner, 0, 5)
	bs = append(bs, b1, b2, b3, b4, b5)

	var count int
	for _, v := range bs {
		count += v.Shows
	}

	return task.Banners{
		Banners: bs,
		Count:   count,
	}
}

func incBannerStatistics(banners *task.Banners, id string, rew bool) {
	for k, v := range banners.Banners {
		if v.Id == id {
			tr := v.Shows + 1
			b := task.Banner{
				Id:     v.Id,
				Shows:  tr,
				Clicks: v.Clicks,
			}
			if rew {
				r := v.Clicks + 1
				b.Clicks = r
			}
			banners.Banners[k] = b
		}
	}
	banners.Count++
}

func randomClick() bool {
	b := getRandomFloat()
	if b < 0.05 {
		return true
	} else {
		return false
	}
}

func getRandomFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
