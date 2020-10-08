package main

import (
	"fmt"
	"testing"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func TestIP(t *testing.T) {
	fmt.Println("err")
	region, err := ip2region.New("./../ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ip, err := region.MemorySearch("113.116.115.75")
	fmt.Println(ip, err)
	ip, err = region.BinarySearch("127.0.0.1")

	//CityId,Country,Region,Province,City,ISP
	fmt.Println(ip, err)
}
