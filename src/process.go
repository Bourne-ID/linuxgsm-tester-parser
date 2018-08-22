package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"html"
)

func processGameServers(servers *[]GameServerDetailModel) []GameServersModel {
	var gameServerDetails []GameServersModel

	for _,server := range *servers {
		res := getWebsite(server.Url)
		doc, err := goquery.NewDocumentFromReader(res.Body)
		check(err)

		processServerMinimumVersion(&server, doc)
		serverDetails := processServerDependencies(&server, doc)
		res.Body.Close()
		gameServerDetails = append(gameServerDetails, serverDetails)
	}
	return gameServerDetails
}

func processServerMinimumVersion(server *GameServerDetailModel, doc *goquery.Document) {
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

func processServerDependencies(server *GameServerDetailModel, doc *goquery.Document) GameServersModel {
	var gsm GameServersModel
	gsm = GameServersModel{
		GameServer:*server,
	}
	doc.Find("h2:contains(' Dependencies')").Parent().Find("#myTabContent").Children().Each(func(i int, selection *goquery.Selection) {
		//get first element available, usually 64 bit


		gsm.Dependencies = append(gsm.Dependencies, extractInformation(selection.Children().Eq(0).Text(), selection.Children().Eq(2).Text()))

		if selection.Children().Length() > 3 {
			gsm.Dependencies = append(gsm.Dependencies, extractInformation(selection.Children().Eq(3).Text(), selection.Children().Eq(5).Text()))
		}
	})
	return gsm
}
func extractInformation(osTextRaw string, dependencies string) GameServerDependenciesModel {
	osText := strings.Split(osTextRaw, " ")
	os := osText[0]
	arch := osText[1]

	//get dependencies
	dependencies = html.UnescapeString(dependencies)
	isI386Required := isAddI386Present(dependencies)
	dependencyList := strings.Split(removeNonPackages(dependencies)," ")
	//Nuke non-packages and convert them into the object

	gsdm := GameServerDependenciesModel{
		OperatingSystem:OSLookup[os],
		Architecture:ArchitectureLookup[arch],
		Addi386:isI386Required,
		Packages:dependencyList,
	}
	return gsdm
}

func isAddI386Present(s string) bool {
	return strings.Contains(s, "--add-architecture i386")
}
func removeNonPackages(s string) string {
	return strings.NewReplacer("sudo ", "",
		"dnf ", "",
		"yum ", "",
		"install ", "",
		"apt ", "",
		"dpkg ", "",
		"update ","",
		"update; ","",
		"--add-architecture i386; ", "").Replace(s)
}