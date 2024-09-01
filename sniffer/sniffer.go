package sniffer

import (
	"bufio"
	"fmt"
	"net/http"
)

type Sniffer interface {
	Sniff(size int) (string, error)
}

type HttpSniffer struct {
	Url string
}

type FileSniffer struct {
	Path string
}

func (s *HttpSniffer) Sniff(size int) (penultimate string, err error) {
	req, err := http.NewRequest(http.MethodGet, s.Url, nil)
	if err != nil {
		err = fmt.Errorf("failed to parse %v: %w", s.Url, err)
		return
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=0-%d", size-1))
	client := &http.Client{
		Transport: &http.Transport{},
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		err = fmt.Errorf("server unsupport request partial content")
		return
	}

	var prevLine, curLine string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		prevLine = curLine
		curLine = scanner.Text()
	}

	return prevLine, nil
}

func (s *FileSniffer) Sniff(size int) (penultimate string, err error) {
	return
}
