package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

const URL = "https://storage.googleapis.com/panels-api/data/20240916/media-1a-i-p~s"

type ImageUrl struct {
	Dhd string
	Dsd string
	S   string
	E   string
}

type PanelResponse struct {
	Version int
	Data    map[int]ImageUrl
}

func readToJson(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Couldn't visit %s", url)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read response body. Url: %s", url)
	}

	var f PanelResponse
	err = json.Unmarshal(b, &f)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read json from response body. Url: %s", url)
	}

	var arr []string
	for _, v := range f.Data {
		if v.Dhd != "" {
			arr = append(arr, v.Dhd)
			continue
		}

		if v.Dsd != "" {
			arr = append(arr, v.Dsd)
			continue
		}

		if v.S != "" {
			arr = append(arr, v.S)
			continue
		}

		if v.E != "" {
			arr = append(arr, v.E)
			continue
		}
	}

	return arr, nil
}

func getFileName(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("Failed to parse URL: %s", s)
	}

	return path.Base(u.Path), nil
}

func main() {
	urls, err := readToJson(URL)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

    fmt.Println(urls)
	for _, s := range urls {
		fmt.Printf("Parsing url: %s\n", s)
		name, err := getFileName(s)
		if err != nil {
		    fmt.Fprintf(os.Stderr, "Error: %s", err)
            continue
		}

        fmt.Printf("Name: %s\n", name)
	}

}
