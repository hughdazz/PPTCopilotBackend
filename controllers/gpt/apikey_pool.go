package gpt

import (
	"backend/conf"
	"fmt"
)

type ApiKeyPool struct {
	ApiKeyList []string
	InUse      map[string]bool
}

var pool = ApiKeyPool{
	ApiKeyList: conf.GetGptApiKeys(),
	InUse:      map[string]bool{},
}

var mutex = make(chan bool, 1)

func GetApiKey() (string, error) {
	mutex <- true
	defer func() { <-mutex }()

	for _, key := range pool.ApiKeyList {
		if !pool.InUse[key] {
			pool.InUse[key] = true
			return key, nil
		}
	}
	return "", fmt.Errorf("no api key available")
}

func ReleaseApiKey(key string) {
	mutex <- true
	defer func() { <-mutex }()

	pool.InUse[key] = false
}
