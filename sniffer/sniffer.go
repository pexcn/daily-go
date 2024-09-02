package sniffer

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

type Sniffer interface {
	Sniff() (string, error)
}

type HttpSniffer struct {
	Url  string
	Size int
}

type FileSniffer struct {
	Path string
	Size int
}

func (s *HttpSniffer) Sniff() (penultimate string, err error) {
	req, err := http.NewRequest(http.MethodGet, s.Url, nil)
	if err != nil {
		err = fmt.Errorf("failed to parse %v: %w", s.Url, err)
		return
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=0-%d", s.Size-1))
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{},
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		err = fmt.Errorf("the url %s is unsupport request partial content", s.Url)
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

func (s *FileSniffer) Sniff() (penultimate string, err error) {
	return
}
