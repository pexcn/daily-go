package cmd

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/netip"
	"strconv"
	"strings"
)

const APNIC_URL = "https://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"
const IPIP_URL = "https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt"

func GetApnicList() {
	resp, err := http.Get(APNIC_URL)
	if err != nil {
		log.Fatalf("Cannot fetch APNIC_URL: %s", err)
	}
	defer resp.Body.Close()

	// convert to cidr format, like https://github.com/pexcn/daily/blob/f31f71d/scripts/chnroute/chnroute.sh#L25
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
	append()
}

func mergeCIDRs(cidrs []netip.Prefix) []netip.Prefix {
	var b netipx.IPSetBuilder
	for _, cidr := range cidrs {
		b.AddPrefix(cidr)
	}
	// ignore errors on purpose to avoid errors in single cidr causing fail output,
	// see comment for IPSetBuilder.IPSet()
	ipset, _ := b.IPSet()

	return ipset.Prefixes()
}
