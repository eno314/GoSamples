package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/eno314/GoSamples/image"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please give imageURLs to argumens")
		return
	}

	doNormal(args)

	doWithChannel(args)

	doWithWaitGroup(args)
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

	fmt.Printf("channel duration is %s\n", time.Now().Sub(startTime).String())
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

func doWithWaitGroup(imageURLs []string) {
	startTime := time.Now()

	//  goルーチンで非同期に実行される処理を待つためにWaitGroupを使う
	wg := sync.WaitGroup{}

	for _, imageURL := range imageURLs {
		// goルーチンを実行する関数分Addする
		wg.Add(1)

		go func(innerImageUrl string) {
			// goルーチン内で関数を呼び出し、WaitGroupをデクリメント
			measureTimeByImageRequest(innerImageUrl)
			wg.Done()
		}(imageURL)
	}

	// goルーチンで実行される関数が終了するまで待つ
	wg.Wait()

	fmt.Printf("WaitGroup duration is %s\n", time.Now().Sub(startTime).String())
}

func measureTimeByImageRequest(imageURL string) {
	duration, err := image.MeasureTimeByRequest(imageURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("duration is %s. URL : %s\n", duration.String(), imageURL)
}
