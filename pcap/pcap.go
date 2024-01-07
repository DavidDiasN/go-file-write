package main

import "fmt"

func main() {
	// version := pcap.Version()
	// fmt.Println(version)

	// devices, _ := pcap.FindAllDevs()
	// for _, dev := range devices {
	// 	fmt.Println(dev)
	// }

	slice1 := []int{1, 1}
	slice3 := append(slice1, 5)
	fmt.Println(slice3)
}
