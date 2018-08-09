package main

func main() {
	servers := processLinuxGSM()

	processGameServers(servers)
	//for _, e := range servers {
	//	fmt.Printf("%s %s", e.serverShortName(), e.Name)
	//}
}
