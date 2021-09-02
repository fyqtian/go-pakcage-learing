package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"
)

//go:embed configs
//go:embed js
var selfFiles embed.FS

func main() {

	fName := []string{
		"configs/a.json",
		"configs/v.json",
		"js/a.js",
	}
	for _, f := range fName {
		f, err := selfFiles.Open(f)
		if err != nil {
			log.Fatalln(err)
		}
		data, _ := ioutil.ReadAll(f)
		fmt.Println(string(data))
	}

}
