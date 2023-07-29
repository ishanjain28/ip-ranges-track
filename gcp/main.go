package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"net/http"
)

type Gcp struct {
	SyncToken    string `json:"syncToken"`
	CreationTime string `json:"creationTime"`
	Prefixes     []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Service    string `json:"service"`
		Scope      string `json:"scope"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	} `json:"prefixes"`
}

func main() {

	resp, err := http.Get("https://www.gstatic.com/ipranges/cloud.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	input := Gcp{}
	err = json.NewDecoder(resp.Body).Decode(&input)
	if err != nil {
		panic(err)
	}

	total := 0
	smallestPrefix := 32
	for _, l := range input.Prefixes {
		if l.Ipv4Prefix == "" {
			continue
		}
		_, cidr, err := net.ParseCIDR(l.Ipv4Prefix)
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
	for _, l := range input.Prefixes {
		if l.Ipv6Prefix == "" {
			continue
		}
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
