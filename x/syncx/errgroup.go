package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

func main() {
	var wg errgroup.Group
	var urls = []string{
		"https://www.baidu.com",
		"https://www.unkonw.com",
		"https://www.sina.com",
	}
	for _, url := range urls {
		u := url
		wg.Go(func() error {
			resp, err := http.Get(u)
			if err != nil {
				return err
			}
			log.Println(resp.Status, u)
			return nil
		})
	}
	err := wg.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("end")

	for range time.NewTicker(time.Second).C {
		log.Println(1)
	}
}
