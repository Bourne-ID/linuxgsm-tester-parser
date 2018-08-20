package main

import "strings"

//two different types of output:
// 1. to file for local use
// 2. compare to github for automatic branch and pull requesting

//output:
/*
---
- name: shortname
-
 */
func ProcessGSMToFile(servers []*GameServersModel) {
	for _, v := range servers {
		serverShortName := v.GameServer.serverShortName()
		//I have the OS and Arch, but no version. Therefore for each run through all the versions. We can mark success failure exectations later.


		for _, vv := range DockerOSLookup[v.Dependencies.OperatingSystem] {
			avf := AnsibleVariableFile{
				OperatingSystem:v.Dependencies.OperatingSystem,
				Packages:strings.Join(v.Dependencies.Packages, " "),
				Architecture:v.Dependencies.Architecture,
				IsI386:v.Dependencies.Addi386,
				Container:vv,
			}
		}

	}
}
