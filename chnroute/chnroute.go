package chnroute

import (
	"bufio"
	"daily/config"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, flag *config.ChnrouteFlag) {
	fmt.Println("URL:", flag.Url)
	fmt.Println("File:", flag.File)
	fmt.Println("Output:", flag.Output)
	fmt.Println("IPv4:", flag.Ipv4)
	fmt.Println("IPv6:", flag.Ipv6)
	fmt.Println("Verbose:", flag.Verbose)
}

func GetApnicList() {
	resp, err := http.Get(config.CHNROUTE_URL_APNIC)
	if err != nil {
		log.Fatalf("Cannot fetch APNIC_URL: %s", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ipv4") && strings.Contains(line, "CN") {
			fields := strings.Split(line, "|")
			if len(fields) < 5 {
				// without mask, skip
				continue
			}
			ip := fields[3]
			size, _ := strconv.Atoi(fields[4])
			mask := 32 - int(math.Log2(float64(size)))
			fmt.Printf("%s/%d\n", ip, mask)
		}
	}
}

func GetIpipList() {

}

// func mergeCIDRs(cidrs []netip.Prefix) []netip.Prefix {
// 	var b netipx.IPSetBuilder
// 	for _, cidr := range cidrs {
// 		b.AddPrefix(cidr)
// 	}
// 	// ignore errors on purpose to avoid errors in single cidr causing fail output,
// 	// see comment for IPSetBuilder.IPSet()
// 	ipset, _ := b.IPSet()

// 	return ipset.Prefixes()
// }
