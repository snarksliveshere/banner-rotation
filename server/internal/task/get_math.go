package task

import (
	"errors"
	"math"
)

func GetBanner(banners *Banners) (Banner, error) {
	if len(banners.Banners) == 0 {
		return Banner{}, errors.New("no banners")
	}
	var rs float64
	var banner Banner
	for _, v := range banners.Banners {
		if v.Shows == 0 {
			// Баннер еще не ротировался, даем ему сразу шанс (инициализация)
			banner = v
			break
		}
		profit := float64(v.Clicks) / float64(v.Shows)
		res := profit + math.Sqrt(math.Log(float64(banners.Count))/float64(v.Shows))
		if res > rs {
			banner = v
			rs = res
		}
	}

	return banner, nil
}
