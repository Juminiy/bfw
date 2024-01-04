package crl

import (
	"github.com/jackdanger/collectlinks"
	"net/http"
)

const (
	USER_AGENT_MACOS_SAFARI = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.35"
)

func GetLinksByExplicitURL(url string) ([]string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", USER_AGENT_MACOS_SAFARI)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return collectlinks.All(resp.Body), nil
}
