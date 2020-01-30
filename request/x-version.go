package request

import (
	"bufio"
	"net/http"
	"strings"

	"granblue.orbb.li/crawler/util"
)

var (
	xVersion string = ""
)

func init() {
	resp, err := http.Get("http://game.granbluefantasy.jp")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	line := util.Grep(scanner, "Game.version")
	xVersion = strings.Split(line, "\"")[1]
}
