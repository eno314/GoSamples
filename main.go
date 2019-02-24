package main

import (
	"fmt"

	"github.com/eno314/GoSamples/image"
)

func main() {
	duration, err := image.MeasureTimeByRequest("https://spnvcdn.sports-digican.com/tennis/images/photo/rect/15733.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("duration is %s\n", duration.String())
}
