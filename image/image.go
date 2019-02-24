// Package image : 画像操作をするサンプルパッケージ
package image

import (
	"errors"
	"fmt"
	"net/http"
)

// Request は指定したURLの画像にHTTP.GETを行う
func Request(imageUrl string) error {
	resp, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("image get is failed. url : " + imageUrl)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("image status code is not 200. url : " + imageUrl)
		return errors.New(resp.Status)
	}

	return nil
}
