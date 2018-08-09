package main

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func processGameServers(servers *[]GameServerDetailModel) []GameServersModel {
	var gameServerDetails []GameServersModel

	for _,server := range *servers {
		res := getWebsite(server.Url)
		processServerMinimumVersion(&server, res)
		processServerDependencies(&server, res)
		res.Body.Close()
	}

	return gameServerDetails
}

func processServerMinimumVersion(server *GameServerDetailModel, res *http.Response) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var minimumVersions []GameServerMinimumVersion

	doc.Find("h3:contains(' Minimum Recommended Distros')").Next().Children().Each(func(i int, s *goquery.Selection) {
		sep := strings.Split(strings.TrimLeft(s.Text(), " "), " ")
		os := OSLookup[sep[0]]
		version := sep[1]

		if os == "" || version == "" {
			log.Fatal("OS or Version is nil: %s", s.Text())
		}
		minimumVersion := GameServerMinimumVersion{
			OperatingSystem:&os,
			Version:version,
		}

		minimumVersions = append(minimumVersions, minimumVersion)
	})
		server.MinimumOperatingSystems = &minimumVersions
}

func processServerDependencies(server *GameServerDetailModel, res *http.Response) GameServersModel {
  //todo
	return GameServersModel{}
}
