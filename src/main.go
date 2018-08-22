package main

import "log"

func main() {
	servers := processLinuxGSM()

	serverDetails := processGameServers(servers)

	ProcessGSMToFile(serverDetails)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
