package test

import (
	"local-build/internal/utils"
	"log"
	"net/url"
	"testing"
)

func TestXxx(t *testing.T) {
	utils.GetHomePath()

	str := "local-build.sqlite?_json=1"
	parsedUrl, err := url.Parse(str)
	if err != nil {
		panic(err)
	}

	query := parsedUrl.Query()
	log.Printf("query: %+v\n", query)

	query.Set("_json", "2")
	log.Printf("query: %+v\n", query)
}