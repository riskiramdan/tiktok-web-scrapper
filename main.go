package main

import (
	"fmt"

	"github.com/riskiramdan/tiktok-web-scrapper/src"
	"github.com/sirupsen/logrus"
)

func main() {
	result, err := src.NewTiktok().GetVideoProperties("https://www.tiktok.com/@officialbmth/video/6976962464957189381?lang=en&is_copy_url=1&is_from_webapp=v1")
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(result)
}
