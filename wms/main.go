package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	mode = flag.String("mode", "worldmeter", "")
	path = flag.String("json_path", "../covid19_live_update.json", "")
)

type docProcessor interface {
	Decode(r io.Reader) error
	Encode(path string) error
}

func processor() (docProcessor, error) {
	switch *mode {
	case "worldmeter":
		return &wmProcessor{}, nil
	}
	return nil, fmt.Errorf("no such mode: %q", *mode)
}

func url() (string, error) {
	switch *mode {
	case "worldmeter":
		return "https://www.worldometers.info/coronavirus/", nil
	}
	return "", fmt.Errorf("no such mode: %q", *mode)
}

func main() {
	flag.Parse()

	p, err := processor()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	url, err := url()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Bad status code: %v\n", err)
		os.Exit(1)
	}

	if err := p.Decode(res.Body); err != nil {
		fmt.Fprintf(os.Stderr, "Processing upstream data failed: %v\n", err)
		os.Exit(1)
	}

	if err := p.Encode(*path); err != nil {
		fmt.Fprintf(os.Stderr, "Updating local data failed: %v\n", err)
		os.Exit(1)
	}
}
