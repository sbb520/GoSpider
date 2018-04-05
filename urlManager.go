package main

import (
	"fmt"
)

type urlManager struct {
	NewUrls []string
	OldUrls []string
}

func (u *urlManager) addNewUrl(url string) error {
	if url == "" {
		return fmt.Errorf("%s", "error")
	}
	u.NewUrls = append(u.NewUrls, url)
	return nil
}

func (u *urlManager) addNewUrls(urls []string) error {
	for _, url := range urls {
		if url == "" {
			return fmt.Errorf("%s", "error")
		}
		u.addNewUrl(url)
	}
	return nil
}

func (u *urlManager) print() {
	fmt.Println(u.NewUrls)
}
