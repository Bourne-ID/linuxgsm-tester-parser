package main

import "strings"

type OS string;
const (
	Ubuntu OS = "Ubuntu"
	Fedora OS = "Fedora"
	Debian OS = "Debian"
	CentOS OS = "CentOS"
)

type GameServerDetailModel struct {
	Name	string
	Url		string
}

func (gs GameServerDetailModel) serverShortName() string {
	value := strings.Replace(gs.Url,"https://linuxgsm.com/lgsm/","",-1)
	return strings.Replace(value,"/","",-1)
}

type GameServerDependenciesModel struct {
	Addi386		bool
	Packages[]	string
}

type GameServersModel struct {
	GameServer			GameServerDetailModel
	OperatingSystem		OS
	Version				string
	Dependencies		GameServerDependenciesModel
}
