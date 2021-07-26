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
	personalRanks []int = append([]int{2000}, util.IntSlice(5000, 300000, 5000)...)
)

func CollectP(t time.Time) error {
	result := map[string]interface{}{}
	for _, v := range personalRanks {
		pageno := v / 10
		url := fmt.Sprintf("/rest_ranking_user/detail/%d/0", pageno)
		resp := request.Requestor{}.Get(url)
		list := resp["list"].([]interface{})
		for _, v2 := range list {
			m := v2.(map[string]interface{})
			if m["rank"] == strconv.Itoa(v) {
				point, err := strconv.ParseFloat(m["point"].(string), 64)
				if err != nil {
					panic(err)
				}
				result[m["rank"].(string)] = point
			}
		}
	}
	result["datetime"] = t

	collectionName := "personal_border"
	documentName := util.Config.Teamraid
	docref := util.FirestoreClient.Collection(collectionName).Doc(documentName)

	rt := map[string]interface{}{}
	old, err := docref.Get(util.FirestoreContext)
	if err != nil {
		rt["records"] = []interface{}{result}
	} else {
		rt["records"] = append(old.Data()["records"].([]interface{}), result)
	}
	_, err = docref.Set(util.FirestoreContext, rt, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return nil
}
