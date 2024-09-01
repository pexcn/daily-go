package main

import (
	"daily/cmd"
)

func main() {
	cmd.Execute()

	// httpSniffer := &sniffer.HttpSniffer{Url: config.CHNROUTE_URL_APNIC}
	// fetchContent(httpSniffer)
}

// func fetchContent(s sniffer.Sniffer) {
// 	line, _ := s.Sniff(4096)
// 	if !isValid(line) {

// 		return
// 	}
// }

// func isValid(s string) bool {
// 	if lib.IsAPNICFormat(s) || lib.IsIP(s) || lib.IsCIDR(s) {
// 		return true
// 	}
// 	return false
// }
