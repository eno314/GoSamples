package main

import (
	"fmt"
	"time"

	"github.com/eno314/GoSamples/image"
)

func main() {
	imageURLs := []string{
		"https://spnvcdn.sports-digican.com/tennis/images/photo/rect/15733.jpg",
		"https://spnavi.c.yimg.jp/spnavi/photo/20170905/soccer/japan_700799.jpg",
		"https://sportsnavi.ht.kyodo-d.jp/golf/img/player_photo/K1015422.jpg",
	}

	totalDuration := time.Duration(0)
	for _, imageURL := range imageURLs {
		duration, err := image.MeasureTimeByRequest(imageURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("duration is %s. URL : %s\n", duration.String(), imageURL)

		totalDuration += duration
	}

	fmt.Printf("total duration is %s\n", totalDuration.String())
}
