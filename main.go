package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://storage.googleapis.com/panels-api/data/20240916/media-1a-i-p~s"

type ImageUrl struct {
    Dhd string
    Dsd string
    S string
}

type PanelResponse struct {
    Version int
    Data map[int]ImageUrl
}

func readToJson(url string) error {
    res, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("Couldn't visit %s", url)
    }

    b, err := io.ReadAll(res.Body)
    if err != nil {
        return fmt.Errorf("Couldn't read response body. Url: %s", url)
    }
    // sb := string(body)

    var f PanelResponse
    err = json.Unmarshal(b, &f)
    if err != nil {
        return fmt.Errorf("Couldn't read json from response body. Url: %s", url)
    }

    for k, v := range f.Data {
        if v == nil {
            fmt.Println("Element %d has no ...")
            continue
        }
        fmt.Printf("Element %d: %s\n", k, v.Dhd)
    }

    return nil
}


func main() {
    err := readToJson(URL)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }

}
