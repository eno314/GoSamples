// Package image : 画像操作をするサンプルパッケージ
package image

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// MeasureTimeByRequest は指定したURLの画像のリクエストに掛かった時間を返す
func MeasureTimeByRequest(imageURL string) (time.Duration, error) {
	startTime := time.Now()
	err := Request(imageURL)
	if err != nil {
		fmt.Println("image request error. url : " + imageURL)
		return -1, err
	}
	endTime := time.Now()

	return endTime.Sub(startTime), nil
}

// Request は指定したURLの画像にHTTP.GETを行う
func Request(imageURL string) error {
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("image get is failed. url : " + imageURL)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("image status code is not 200. url : " + imageURL)
		return errors.New(resp.Status)
	}

	return nil
}
