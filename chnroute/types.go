package chnroute

import "daily/sniffer"

type PenultimateLineResult struct {
	Line  string
	Error error
	sniffer.Sniffer
}
