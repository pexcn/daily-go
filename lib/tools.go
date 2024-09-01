package lib

import (
	"net"
	"strings"
)

// func HttpGet(url string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(body), nil
// }

func IsIP(line string) bool {
	return net.ParseIP(line) != nil
}

func IsCIDR(line string) bool {
	_, _, err := net.ParseCIDR(line)
	return err == nil
}

func IsAPNICFormat(line string) bool {
	for _, word := range strings.Split(line, "|") {
		if strings.ToLower(word) == "apnic" {
			return true
		}
	}
	return false
}
