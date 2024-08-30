package sniffer

import (
	"bufio"
	"fmt"
	"net/http"
)

type HttpSniffer struct {
	Url string
}

func (s *HttpSniffer) Sniff(size int) (string, error) {
	req, err := http.NewRequest(http.MethodGet, s.Url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to parse %v: %w", s.Url, err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=0-%d", size-1))
	client := &http.Client{
		Transport: &http.Transport{},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return "", fmt.Errorf("server unsupport request partial content: %w", err)
	}

	var prevLine, curLine string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		prevLine = curLine
		curLine = scanner.Text()
	}

	return prevLine, nil
}
