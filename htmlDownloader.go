package main

import(
	"net/http"
	"io/ioutil"
	"fmt"
)

type htmlDownloader struct {

}

func NewHtmlDownloader () *htmlDownloader {
	h := new(htmlDownloader)
	return h
}

func (h *htmlDownloader) download(url string) (string, error) {
	resp, err := http.Get(url)
	
    if err != nil {
        return "", err
    }
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{  
        return "", fmt.Errorf("%s %d", "error status", &resp.StatusCode)
	} 
	
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}