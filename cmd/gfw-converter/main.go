package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ssd2ss", ssd2ssHandler())
	err := r.Run()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func ssd2ssHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		ss, err := ssd2ss(c.Query("url"))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.String(200, ss)
	}
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
