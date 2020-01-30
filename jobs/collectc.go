package jobs

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"granblue.orbb.li/crawler/request"
	"granblue.orbb.li/crawler/util"
)

var (
	crewRanks int = 5500
)

func CollectC(t time.Time) error {
	for i := 1; i <= crewRanks/10; i++ {
		if i%50 == 0 {
			log.Printf("Crawling %d", i)
		}
		batch := util.FirestoreClient.Batch()
		url := fmt.Sprintf("/rest/ranking/totalguild/detail/%d/0", i)
		resp := request.Requestor{}.Get(url)
		list := resp["list"].([]interface{})
		for _, v := range list {
			m := v.(map[string]interface{})
			ranking, err := strconv.Atoi(m["ranking"].(string))
			if err != nil {
				panic(err)
			}
			point, err := strconv.ParseFloat(m["point"].(string), 64)
			if err != nil {
				panic(err)
			}

			docref := util.FirestoreClient.Collection("crew").Doc(m["id"].(string))
			crew := map[string]interface{}{
				"name":         m["name"],
				"last_updated": t,
			}
			batch.Set(docref, crew, firestore.MergeAll)

			docref2 := docref.Collection("records").Doc(util.Config.Teamraid)
			record := map[string]interface{}{
				"ranking":  ranking,
				"point":    point,
				"datetime": t,
			}
			rt := map[string]interface{}{}
			old, err := docref2.Get(util.FirestoreContext)
			if err != nil {
				rt["records"] = []interface{}{record}
			} else {
				rt["records"] = append(old.Data()["records"].([]interface{}), record)
			}
			batch.Set(docref2, rt, firestore.MergeAll)
		}
		if _, err := batch.Commit(util.FirestoreContext); err != nil {
			log.Printf("An error has occurred: %s", err)
		}
	}
	return nil
}
