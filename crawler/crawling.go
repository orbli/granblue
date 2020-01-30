package crawler

import (
	"net/http"
)

func Crawling(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test success"))
}
