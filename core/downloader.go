package core

import(
	"net/http"
	"io"
	"fmt"
)

type downloader struct {

}

func NewDownloader () *downloader {
	h := new(downloader)
	return h
}

func (h *downloader) Download(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	
    if err != nil {
        return nil, err
	}
	
	if resp.StatusCode != http.StatusOK{  
        return nil, fmt.Errorf("%s %d", "error status", &resp.StatusCode)
	} 
    
    return resp.Body, nil
}