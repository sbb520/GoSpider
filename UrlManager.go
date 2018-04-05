package main

type UrlManager struct {
	NewUrls []string
	OldUrls []string
}

func (u *UrlManager) addNewUrl(url string) error {
	if url == nil {
		return "error"
	}
	return nil
}
