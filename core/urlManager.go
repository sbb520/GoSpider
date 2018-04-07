package core

import (
	"fmt"
)

type urlManager struct {
	NewUrls map[int]string
	OldUrls map[int]string
	newIndex int
	oldIndex int
}
func NewUrlManager() *urlManager {
	u := new(urlManager)
	u.NewUrls = make(map[int]string)
	u.OldUrls = make(map[int]string)
	
	u.newIndex = 0
	u.oldIndex = 0

	return u
    
}

func (u *urlManager) AddNewUrl(url string) error {
	if url == "" {
		return fmt.Errorf("%s", "error")
	}
	u.NewUrls[u.newIndex] = url
	u.newIndex++;
	return nil
}

func (u *urlManager) AddNewUrls(urls []string) error {
	for _, url := range urls {
		if url == "" {
			return fmt.Errorf("%s", "error")
		}
		u.AddNewUrl(url)
	}
	return nil
}


func (u *urlManager) HasNewUrl() bool {
	return len(u.NewUrls) != 0
}

func (u *urlManager) GetNewUrl() (string, error){
	if !u.HasNewUrl() {
		return "", fmt.Errorf("%s", "didn't have index")
	}
	url := u.NewUrls[u.oldIndex]
	u.OldUrls[u.oldIndex] = url
	u.oldIndex++;
	return url, nil;
}

