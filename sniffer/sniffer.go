package sniffer

type Sniffer interface {
	Sniff(size int) (string, error)
}
