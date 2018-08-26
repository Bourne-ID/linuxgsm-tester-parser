package main

import (
	"io/ioutil"
	"strings"
)

func ProcessGSMToFile(servers []GameServersModel) {
	for _, v := range servers {
		serverShortName := v.GameServer.serverShortName()

		avf := AnsibleVariableFile{
			ServerName: serverShortName,
		}
		for _, dep := range v.Dependencies {

			avf.OperatingSystem = dep.OperatingSystem
			avf.Packages = strings.Join(dep.Packages, " ")
			avf.Architecture = dep.Architecture
			avf.IsI386 = dep.Addi386

			for _, vv := range DockerOSLookup[dep.OperatingSystem] {
				avf.Container = vv

				makeFile(avf)

			}
		}
	}
}
func makeFile(file AnsibleVariableFile) {
	data := file.convertToBytes()
	err := ioutil.WriteFile("out/"+file.generateFileName(), data, 0644)
	check(err)
}
