package util

import (
	"cloud.google.com/go/firestore"
)

type (
	configStruct struct {
		Active    bool
		Teamraid  string
		UserAgent string
		Cookie    string
	}
)

var (
	Config *configStruct
)

func LoadConfig() {
	if data, err := FirestoreClient.Collection("configs").Doc("config").Get(FirestoreContext); err != nil {
		panic(err)
	} else {
		result := data.Data()
		Config = &configStruct{
			Active: result["active"].(bool),
			Teamraid: result["teamraid"].(string),
			UserAgent: result["useragent"].(string),
			Cookie: result["cookie"].(string),
		}
	}
}

func SaveConfig() {
	result := map[string]interface{}{
		"cookie": Config.Cookie,
	}
	if _, err := FirestoreClient.Collection("configs").Doc("config").Set(FirestoreContext, result, firestore.MergeAll); err != nil {
		panic(err)
	}
}
