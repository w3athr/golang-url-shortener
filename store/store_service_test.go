package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://habr.com/ru/companies/ruvds/articles/464289/"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	SaveUrlMapping(shortUrl, initialLink, userUUID)

	retievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retievedUrl)
}