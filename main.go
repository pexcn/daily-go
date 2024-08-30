package main

import (
	"daily/lib"
	"daily/sniffer"
	"fmt"
)

func main() {
	httpSniffer := &sniffer.HttpSniffer{Url: CHNROUTE_URL_APNIC}

	checkFormat(httpSniffer)
}

func checkFormat(s sniffer.Sniffer) {
	line, _ := s.Sniff(4096)
	fmt.Println(lib.IsAPNICFormat(fmt.Sprintln(line)))
}
