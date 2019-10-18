package main

import "fmt"

func main() {
	fmt.Println("Device starts.. ")
	device := Device{}
	device.Name = "Gib Device"
	device.SerialNumber = "CT6012313213"
	fmt.Println(device.String())
}
