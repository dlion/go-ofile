package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	const V = 1.0
	fmt.Printf("-------------------------------------\n"+
		"|Go-ofile v%.1f                       |\n"+
		"|Coded by Domenico (DLion) Luciani   |\n"+
		"|https://domenicoluciani.com         |\n"+
		"|github.com/DLion/go-ofile           |\n"+
		"-------------------------------------\n\n\n", V)
	domain := flag.String("d", "google.com", "domain to search")
	filetype := flag.String("f", "pdf", "filetype (ex. pdf)")
	flag.Parse()

	url := "https://google.com/search?num=1000&q=" + (*domain) + "+filetype:" + (*filetype)
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Searching in %s for %s\n==============================================\n\n\n", *domain, *filetype)
	r, err := regexp.Compile("\\/url\\?q=[a-zA-Z0-9_/.\\-:\r\n\t]*\\." + (*filetype))
	if err != nil {
		log.Panicln(err)
	}

	found := r.FindAllString(string(body), -1)
	if found != nil {
		fmt.Printf("Files found: \n==============================================\n\n\n")
		for _, file := range found {
			fmt.Println(strings.Replace(file, "/url?q=", " ", -1))
		}

	} else {
		fmt.Println("No results were found")
	}
	fmt.Println("==============================================")
}
