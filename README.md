# IP Ranges track


I wrote this to track the number of IPv4 and IPv6 addresses owned by the 3 major Cloud providers, AWS, Azure and GCP.


In case of AWS and GCP, It fetches the JSON listing the IP blocks from their public URL and counts the total number of IPv4 addresses and the total number of /64 IPv6 blocks.


Azure provides a file that is updated weekly. The URL to download page is in `azure/main.go`


As of 29/07/2023,

1. AWS has 136,014,590 IPv4 addresses and 105,140,683,236 /64 IPv6 blocks. Largest contiguous IPv4 mask is 11 and IPv6 mask is 32

2. Azure has 77,953,227 IPv4 addresses and 156,284,265 /64 IPv6 blocks. Largest IPv4 mask is 15 and IPv6 mask is 40

3. GCP has 12,179,200 IPv4 addresses and 43,057,152 /64 IPv6 blocks. Largest IPv4 mask is 14 and IPv6 mask is 44



## Limitations

This program does not attempt to fuse ranges to get a more accurate number. This will result in inflated numbers if these providers have overlapping ranges.
