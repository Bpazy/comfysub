package main

import (
	"encoding/base64"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"log"
)

func main() {
	ssContent, err := ssd2ss()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Println(ssContent)
}

func ssd2ss(url string) (string, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return "", errors.Wrap(err, "get ShadowsocksD subscription failed")
	}

	var decodedSsdSubscription []byte
	decodedSsdSubscription, err = base64.RawURLEncoding.DecodeString(string(resp.Body())[6:])
	if err != nil {
		return "", errors.Wrap(err, "decode ShadowsocksD subscription failed")
	}
	return string(decodedSsdSubscription), nil
}
