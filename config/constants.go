package config

const CHNROUTE_URL_APNIC = "https://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"
const CHNROUTE_URL_IPIP = "https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt"

type GlobalFlag struct {
	Verbose bool
}

type ChnrouteFlag struct {
	Url    []string
	File   []string
	Output string
	Ipv4   bool
	Ipv6   bool
	GlobalFlag
}
