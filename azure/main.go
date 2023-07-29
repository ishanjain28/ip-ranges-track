package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
)

type Azure struct {
	ChangeNumber int    `json:"changeNumber"`
	Cloud        string `json:"cloud"`
	Values       []struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		Properties struct {
			ChangeNumber    int      `json:"changeNumber"`
			Region          string   `json:"region"`
			RegionID        int      `json:"regionId"`
			Platform        string   `json:"platform"`
			SystemService   string   `json:"systemService"`
			AddressPrefixes []string `json:"addressPrefixes"`
			NetworkFeatures []string `json:"networkFeatures"`
		} `json:"properties"`
	} `json:"values"`
}

func main() {

	// 	https://www.microsoft.com/en-us/download/details.aspx?id=56519
	f, err := os.Open("./ServiceTags_Public_20230724.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := Azure{}
	err = json.NewDecoder(f).Decode(&input)
	if err != nil {
		panic(err)
	}

	totalIpv4 := 0
	totalIpv6 := 0
	smallestIpv4Prefix := 32
	smallestIpv6Prefix := 128

	for _, v := range input.Values {

		for _, addr := range v.Properties.AddressPrefixes {
			_, cidr, err := net.ParseCIDR(addr)
			if err != nil {
				panic(err)
			}

			ones, bits := cidr.Mask.Size()
			if bits == 128 {

				// Only track /64 blocks
				bits = 64
				totalIpv6 += int(math.Pow(2.0, float64(bits-ones)))

				if ones < smallestIpv6Prefix {
					smallestIpv6Prefix = ones
				}

			} else if bits == 32 {
				totalIpv4 += int(math.Pow(2.0, float64(bits-ones)))
				if ones < smallestIpv4Prefix {
					smallestIpv4Prefix = ones
				}
			} else {
				panic("unknown bits??")
			}

		}

	}

	fmt.Println("Total Ipv4 Address Count", totalIpv4, "Smallest Ipv4 Prefix", smallestIpv4Prefix)
	fmt.Println("Total Ipv6 Address Count", totalIpv6, "Smallest Ipv6 Prefix", smallestIpv6Prefix)
}
