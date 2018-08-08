package main

import "fmt"

func main() {
	servers := processLinuxGSM()

	for _, e := range servers {
		fmt.Printf("%s %s", e.serverShortName(), e.Name)
	}
}
