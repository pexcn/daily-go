package chnroute

import (
	"daily/config"
	"daily/lib"
	"daily/sniffer"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func fetchPenultimateLine(s sniffer.Sniffer, result chan<- PenultimateLineResult, wg *sync.WaitGroup) {
	defer wg.Done()
	line, err := s.Sniff()
	result <- PenultimateLineResult{
		Line:    line,
		Error:   err,
		Sniffer: s,
	}
}

func isAPNIC(line string) bool {
	for _, word := range strings.Split(line, "|") {
		if strings.ToLower(word) == "apnic" {
			return true
		}
	}
	return false
}

func isValid(line string) bool {
	return lib.IsIP(line) || lib.IsCIDR(line) || isAPNIC(line)
}

func preFetch(flag *config.ChnrouteFlag) error {
	// TODO: add skip prefetch option
	var wg sync.WaitGroup
	results := make(chan PenultimateLineResult, len(flag.Url)+len(flag.File))
	var coroutines int

	for _, url := range flag.Url {
		wg.Add(1)
		sniffer := sniffer.HttpSniffer{
			Url:  url,
			Size: config.DEFAULT_SNIFF_SIZE,
		}
		go fetchPenultimateLine(&sniffer, results, &wg)
		coroutines++
	}

	for _, file := range flag.File {
		wg.Add(1)
		sniffer := sniffer.FileSniffer{
			Path: file,
			Size: config.DEFAULT_SNIFF_SIZE,
		}
		go fetchPenultimateLine(&sniffer, results, &wg)
		coroutines++
	}

	// TODO: encapsulation log function
	log.Printf("start %d coroutines prefetching...", coroutines)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error == nil {
			if !isValid(result.Line) {
				switch s := result.Sniffer.(type) {
				case *sniffer.HttpSniffer:
					return fmt.Errorf("the url %s has invalid format", s.Url)
				case *sniffer.FileSniffer:
					return fmt.Errorf("the file %s has invalid format", s.Path)
				default:
					return errors.New("unknown sniffer type")
				}
			}
		} else {
			// TODO: check error type
			// if unsupport request partial content, skip and fetch full size check isValid()
			log.Fatal(result.Error)
		}

	}

	return nil
}

func Run(cmd *cobra.Command, flag *config.ChnrouteFlag) {
	err := preFetch(flag)
	if err != nil {
		log.Fatal(err)
	}
}

// func GetApnicList() {
// 	resp, err := http.Get(config.CHNROUTE_URL_APNIC)
// 	if err != nil {
// 		log.Fatalf("Cannot fetch APNIC_URL: %s", err)
// 	}
// 	defer resp.Body.Close()

// 	scanner := bufio.NewScanner(resp.Body)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.Contains(line, "ipv4") && strings.Contains(line, "CN") {
// 			fields := strings.Split(line, "|")
// 			if len(fields) < 5 {
// 				// without mask, skip
// 				continue
// 			}
// 			ip := fields[3]
// 			size, _ := strconv.Atoi(fields[4])
// 			mask := 32 - int(math.Log2(float64(size)))
// 			fmt.Printf("%s/%d\n", ip, mask)
// 		}
// 	}
// }

// func GetIpipList() {

// }

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
