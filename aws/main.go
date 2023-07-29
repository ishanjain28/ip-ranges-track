package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"net/http"
)

type Aws struct {
	Prefixes []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
	Ipv6Prefixes []struct {
		Ipv6Prefix         string `json:"ipv6_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"ipv6_prefixes"`
}

func main() {

	resp, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	input := Aws{}
	err = json.NewDecoder(resp.Body).Decode(&input)
	if err != nil {
		panic(err)
	}

	total := 0
	smallestPrefix := 32
	for _, l := range input.Prefixes {
		_, cidr, err := net.ParseCIDR(l.IPPrefix)
		if err != nil {
			panic(err)
		}
		ones, bits := cidr.Mask.Size()
		if ones < smallestPrefix {
			smallestPrefix = ones
		}

		total += int(math.Pow(2.0, float64(bits-ones)))
	}

	fmt.Println("Total Ipv4 Address Count", total, "Smallest Ipv4 Prefix", smallestPrefix)
	var totalIpv6 int64 = 0
	smallestPrefix = 128
	for _, l := range input.Ipv6Prefixes {
		_, cidr, err := net.ParseCIDR(l.Ipv6Prefix)
		if err != nil {
			panic(err)
		}
		ones, bits := cidr.Mask.Size()
		if ones < smallestPrefix {
			smallestPrefix = ones
		}

		// Only count /64 blocks
		bits = 64
		totalIpv6 += int64(math.Pow(2.0, float64(bits-ones)))
	}

	fmt.Println("Total Ipv6 Address Count", totalIpv6, "Smallest Ipv6 Prefix", smallestPrefix)
}
