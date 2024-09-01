package chnroute

import (
	"bufio"
	"daily/config"
	"daily/lib"
	"daily/sniffer"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, flag *config.ChnrouteFlag) {
	// fmt.Println("URL:", flag.Url)
	// fmt.Println("File:", flag.File)
	// fmt.Println("Output:", flag.Output)
	// fmt.Println("IPv4:", flag.Ipv4)
	// fmt.Println("IPv6:", flag.Ipv6)
	// fmt.Println("Verbose:", flag.Verbose)

	var wg sync.WaitGroup

	for _, url := range flag.Url {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			sniffer := &sniffer.HttpSniffer{Url: u}
			line := fetchPenultimateLine(sniffer)
			if !isValid(line) {
				log.Fatal("there is a URL with invalid content format: ", u)
			}
		}(url)
	}

	// for _, file := range flag.File {
	// 	// 每启动一个 Goroutine，计数器加一
	// 	wg.Add(1)
	// 	// 启动 Goroutine
	// 	go func(f string) {
	// 		// 确保在 Goroutine 完成时计数器减一
	// 		defer wg.Done()
	// 		fmt.Println(f)
	// 		time.Sleep(1 * time.Second)
	// 	}(file)
	// }

	wg.Wait()

}

func fetchPenultimateLine(s sniffer.Sniffer) string {
	line, err := s.Sniff(4096)
	if err != nil {
		log.Fatal(err)
	}
	return line
}

func isValid(line string) bool {
	return lib.IsIP(line) || lib.IsCIDR(line) || lib.IsAPNICFormat(line)
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
