package crawler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"granblue.orbb.li/crawler/jobs"
	"granblue.orbb.li/crawler/util"
)

func Crawling(w http.ResponseWriter, r *http.Request) {
	util.LoadConfig()
	defer func() {
		log.Printf("New cookie: ", util.Config.Cookie)
		util.SaveConfig()
	}()
	if !util.Config.Active {
		log.Printf("Inactive")
		return
	}
	loc, _ := time.LoadLocation("Japan")
	t := time.Now().In(loc)

	log.Printf("Collecting personal")
	if err := jobs.CollectP(t); err != nil {
		fmt.Fprint(w, err.Error())
	}
	log.Printf("Finish collect personal")

	log.Printf("Collecting crew")
	if err := jobs.CollectC(t); err != nil {
		fmt.Fprint(w, err.Error())
	}
	log.Printf("Finish collect crew")
	fmt.Fprintf(w, "Success")
}
