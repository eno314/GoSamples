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

	doNormal(imageURLs)

	doWithChannel(imageURLs)
}

func doNormal(imageURLs []string) {
	startTime := time.Now()

	for _, imageURL := range imageURLs {
		duration, err := image.MeasureTimeByRequest(imageURL)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("duration is %s. URL : %s\n", duration.String(), imageURL)
	}

	fmt.Printf("Normal duration is %s\n", time.Now().Sub(startTime).String())
}

func doWithChannel(imageURLs []string) {
	startTime := time.Now()

	chanel := make(chan time.Duration)

	for _, imageURL := range imageURLs {
		go measureTimeImageRequestWithChannel(imageURL, chanel)
	}

	for i := 0; i < len(imageURLs); i++ {
		<-chanel
	}

	fmt.Printf("Goroutine duration is %s\n", time.Now().Sub(startTime).String())
}

func measureTimeImageRequestWithChannel(imageURL string, durationChan chan time.Duration) {
	duration, err := image.MeasureTimeByRequest(imageURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("duration is %s. URL : %s\n", duration.String(), imageURL)

	durationChan <- duration
}
