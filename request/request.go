package request

import (
	"fmt"
	"net/http"

	"granblue.orbb.li/crawler/util"

	// "log"
	"compress/gzip"
	"encoding/json"
)

var (
	base string = "http://game.granbluefantasy.jp"
)

type (
	Requestor struct{}
)

func (r Requestor) Get(url string) map[string]interface{} {
	fullurl := fmt.Sprintf("%s/%s%s", base, util.Config.Teamraid, url)

	req, err := http.NewRequest("GET", fullurl, nil)
	req.Header = http.Header{}
	req.Header.Set("X-VERSION", xVersion)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8,zh-TW;q=0.7,zh;q=0.6,zh-CN;q=0.5,ja;q=0.4")
	req.Header.Set("User-Agent", util.Config.UserAgent)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", "http://game.granbluefantasy.jp/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", util.Config.Cookie)
	req.Header.Set("Connection", "keep-alive")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}

	for _, v := range resp.Header["Set-Cookie"] {
		oldcookie := util.FromCookiestring(util.Config.Cookie)
		setcookie := util.FromCookiestring(v)
		for k, v2 := range setcookie {
			oldcookie[k] = v2
		}
		util.Config.Cookie = util.ToCookiestring(oldcookie)
	}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}
	jsonDecoder := json.NewDecoder(reader)
	rt := new(map[string]interface{})
	if err = jsonDecoder.Decode(rt); err != nil {
		panic(err)
	}
	return *rt
}
