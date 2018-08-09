package main

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

const linuxGsmWebsite = "https://linuxgsm.com/servers/"

func processLinuxGSM() *[]GameServerDetailModel {
	res := getWebsite(linuxGsmWebsite)
	servers := getServerList(res)
	return &servers
}

func getWebsite(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}


	if res.StatusCode != 200 {
		log.Fatal("Status code error: %d %s", res.StatusCode, res.Status)
	}

	return res
}

func getServerList(res *http.Response) []GameServerDetailModel {
	var gameServers []GameServerDetailModel

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#menu1").Children().Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
		value, exists := s.Attr("value")
		if exists && len(value) > 0 {
			name := s.Text()
			gameServer := GameServerDetailModel{
				Name: name,
				Url:  value,
			}
			gameServers = append(gameServers, gameServer)
		}
	})

	return gameServers
}