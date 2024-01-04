package hurun

import (
	my_net "bfw/wheel/net"
	"errors"
	"fmt"
	my_json "github.com/json-iterator/go"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	hurunBaseURL                = "https://www.hurun.net/zh-CN/Rank/HsRankDetailsList"
	hurunRESTAPIHTTPMethod      = "GET"
	hurunRESTAPIHTTPQueryNum    = "num"
	hurunRESTAPIHTTPQuerySearch = "search"
	hurunRESTAPIHTTPQueryOffset = "offset"
	hurunRESTAPIHTTPQueryLimit  = "limit"
	hurunWriteDir               = "testdata"
)

var (
	requestHurunHTTPCodeError = errors.New("request hurun.net HTTP Code is NOT 200 OK")
	requestHurunRankModeError = errors.New("request hurun.net Rank Mode not allowed")
)

type HurunIResp interface {
	TotalCNT() int
}

type HurunURL struct {
	url string
	qry map[string]string
}

func MakeHurunURL() *HurunURL {
	return &HurunURL{url: hurunBaseURL, qry: make(map[string]string)}
}

func (h *HurunURL) Request(iResp HurunIResp) {
	url := h.splice()
	time0 := time.Now()
	resp, err := my_net.RequestByURL(hurunRESTAPIHTTPMethod, url)
	fmt.Printf("request cost: %v\n", time.Since(time0))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(requestHurunHTTPCodeError)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = my_json.Unmarshal(respBody, &iResp)
	if err != nil {
		panic(err)
	}
}

func (h *HurunURL) Num(num string) {
	h.query(hurunRESTAPIHTTPQueryNum, num)
}

func (h *HurunURL) Search(search string) {
	h.query(hurunRESTAPIHTTPQuerySearch, search)
}

func (h *HurunURL) Range(from, to int) {
	if from < 0 {
		from = 0
	}
	if to < from {
		to = from
	}
	offset, limit := strconv.Itoa(from), strconv.Itoa(to-from+1)
	h.query(hurunRESTAPIHTTPQueryOffset, offset)
	h.query(hurunRESTAPIHTTPQueryLimit, limit)
}

func (h *HurunURL) validate() bool {
	return len(h.url) >= 50
}

func (h *HurunURL) make() {
	if !h.validate() {
		h.url = hurunBaseURL
	}
}

func (h *HurunURL) query(key, value string) {
	h.qry[key] = value
}

func (h *HurunURL) splice() string {
	h.make()
	url := h.url + "?"
	for key, value := range h.qry {
		url += key + "=" + value + "&"
	}
	return url
}

func getJSONFileName(dir, output string) string {
	return getFileName(dir, output, "json")
}

func getFileName(prefix, output, suffix string) string {
	return prefix + "/" + output + "." + suffix
}
