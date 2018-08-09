package main

import "strings"

type OS string
const (
	Ubuntu OS = "Ubuntu"
	Fedora OS = "Fedora"
	Debian OS = "Debian"
	CentOS OS = "CentOS"
)
var OSLookup = map[string]OS{
	"Ubuntu": Ubuntu,
	"Fedora": Fedora,
	"Debian": Debian,
	"CentOS": CentOS,
}

type Architecture string
const (
	Bit32 Architecture = "32"
	Bit64 Architecture = "64"
)

type GameServerDetailModel struct {
	Name	string
	Url		string
	MinimumOperatingSystems		*[]GameServerMinimumVersion
}

func (gs GameServerDetailModel) serverShortName() string {
	value := strings.Replace(gs.Url,"https://linuxgsm.com/lgsm/","",-1)
	return strings.Replace(value,"/","",-1)
}

type GameServerDependenciesModel struct {
	OperatingSystem OS
	Architecture *Architecture
	Addi386		bool
	Packages[]	string
}

type GameServerMinimumVersion struct {
	OperatingSystem *OS
	Version string
}

type GameServersModel struct {
	GameServer			*GameServerDetailModel
	Dependencies		*GameServerDependenciesModel
}
