package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

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
	Bit32 Architecture = "x32"
	Bit64 Architecture = "x64"
)

var ArchitectureLookup = map[string]Architecture{
	"x86":    Bit32,
	"32-bit": Bit32,
	"x64":    Bit64,
	"64-bit": Bit64,
}

type DockerImage string

const (
	Ubuntu14 DockerImage = "ubuntu:trusty"
	Ubuntu16 DockerImage = "ubuntu:xenial"
	Ubuntu18 DockerImage = "ubuntu:bionic"
	CentOS6  DockerImage = "centos:6"
	CentOS7  DockerImage = "centos:7"
	Debian7  DockerImage = "debian:7"
	Debian8  DockerImage = "debian:8"
	Debian9  DockerImage = "debian:9"
	Fedora26 DockerImage = "fedora:26"
	Fedora27 DockerImage = "fedora:27"
	Fedora28 DockerImage = "fedora:28"
)

//DockerOSLookup used for taking an OS and returning all the available DockerOS containers to use
var DockerOSLookup = map[OS][]DockerImage{
	Ubuntu: {Ubuntu14, Ubuntu16, Ubuntu18},
	CentOS: {CentOS6, CentOS7},
	Debian: {Debian7, Debian8, Debian9},
	Fedora: {Fedora26, Fedora27, Fedora28},
}

type GameServerDetailModel struct {
	Name                    string
	Url                     string
	MinimumOperatingSystems *[]GameServerMinimumVersion
}

func (gs GameServerDetailModel) serverShortName() string {
	value := strings.Replace(gs.Url, "https://linuxgsm.com/lgsm/", "", -1)
	return strings.Replace(value, "/", "", -1)
}

type GameServerDependenciesModel struct {
	OperatingSystem OS
	Architecture    Architecture
	Addi386         bool
	Packages        []string
}

type GameServerMinimumVersion struct {
	OperatingSystem *OS
	Version         string
}

type GameServersModel struct {
	GameServer   GameServerDetailModel
	Dependencies []GameServerDependenciesModel
}

type AnsibleVariableFile struct {
	//for the filename and content - afterall we need to ensure we know the docker image we're using.
	ServerName             string
	Container              DockerImage
	OperatingSystem        OS
	Architecture           Architecture
	OperatingSystemVersion string
	//for the contents of the file
	Packages string
	IsI386   bool
}

func (value AnsibleVariableFile) generateFileName() string {
	s := string(value.Container)
	s = strings.Replace(s, ":", "-", -1)

	return fmt.Sprintf("%s-%s-%s.yml", value.ServerName, s, value.Architecture)
}
func (value AnsibleVariableFile) convertToBytes() []byte {
	t, err := template.ParseFiles("template/ansible.yml")
	check(err)

	var tpl bytes.Buffer
	err = t.Execute(&tpl, value)
	check(err)

	return tpl.Bytes()
}
