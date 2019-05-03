package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/eno314/GoSamples/image"
	"github.com/eno314/GoSamples/random"
	"github.com/eno314/GoSamples/sort"
)

func main() {
	randomValues := random.RandomValues(10000, 10000)
	// fmt.Println(randomValues)

	doBubbleSort(randomValues)
	// fmt.Println(doBubbleSort(randomValues))
	doSelectionSort(randomValues)
	// fmt.Println(doSelectionSort(randomValues))
	doInsertionSort(randomValues)
	// fmt.Println(doInsertionSort(randomValues))
	doHeapSort(randomValues)
	// fmt.Println(doHeapSort(randomValues))
	doMergeSort(randomValues)
	// fmt.Println(doMergeSort(randomValues))
	doQuickSort(randomValues)
	// fmt.Println(doQuickSort(randomValues))
}

func doBubbleSort(values []int) []int {
	startTime := time.Now()
	sorted := sort.Bubble(values)
	printDuration("BubbleSort", startTime)
	return sorted
}

func doSelectionSort(values []int) []int {
	startTime := time.Now()
	sorted := sort.Selection(values)
	printDuration("SelectionSort", startTime)
	return sorted
}

func doInsertionSort(values []int) []int {
	startTime := time.Now()
	sorted := sort.Insertion(values)
	printDuration("InsertionSort", startTime)
	return sorted
}

func doHeapSort(values []int) []int {
	startTime := time.Now()
	sorted := sort.Heap(values)
	printDuration("HeapSort", startTime)
	return sorted
}

func doMergeSort(values []int) []int {
	startTime := time.Now()
	sorted := sort.Merge(values)
	printDuration("MergeSort", startTime)
	return sorted
}

func doQuickSort(values []int) []int {
	startTime := time.Now()
	sorted := sort.Quick(values)
	printDuration("QuickSort", startTime)
	return sorted
}

func printDuration(sortName string, startTime time.Time) {
	fmt.Printf("call took %v to run by %v.\n", time.Now().Sub(startTime), sortName)
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
