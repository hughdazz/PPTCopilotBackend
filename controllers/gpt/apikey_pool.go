package gpt

import (
	"backend/conf"
	"fmt"
	"math/rand"
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

	// 随机选择一个可用ApiKey
	var validKeys []string
	for _, key := range pool.ApiKeyList {
		if !pool.InUse[key] {
			validKeys = append(validKeys, key)
		}
	}
	if len(validKeys) > 0 {
		// 生成随机数
		randIndex := rand.Intn(len(validKeys))
		pool.InUse[validKeys[randIndex]] = true
		return validKeys[randIndex], nil
	}

	return "", fmt.Errorf("no api key available")
}

func ReleaseApiKey(key string) {
	mutex <- true
	defer func() { <-mutex }()

	pool.InUse[key] = false
}
