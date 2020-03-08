package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8080", "server port")
	flag.Parse()
}

func main() {
	r := gin.Default()
	r.GET("/ssd2ss", ssd2ssHandler())
	if err := r.Run(port); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func ssd2ssHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		if ssdUrl, ok := c.GetQuery("url"); ok {
			ss, err := ssd2ss(ssdUrl)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.String(200, ss)
			return
		}
		c.String(http.StatusInternalServerError, "param \"url\" is required")
		return
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

	var ssdConfig SsdConfig
	err = json.Unmarshal(decodedSsdSubscription, &ssdConfig)
	if err != nil {
		return "", errors.Wrap(err, "unmarshal ShadowsocksD subscription failed")
	}

	commonConfigList := CommonConfigList{}
	for _, sc := range ssdConfig.Servers {
		commonConfigList = append(commonConfigList, CommonConfig{
			Server:     sc.Server,
			ServerPort: ssdConfig.Port,
			Method:     ssdConfig.Encryption,
			Password:   ssdConfig.Password,
			Remarks:    sc.Remarks,
		})
	}

	return commonConfigList.getSubscription(), nil
}

type CommonConfig struct {
	Server     string
	ServerPort int
	Method     string
	Password   string
	Remarks    string
}

func (sc *CommonConfig) getSubscription() string {
	encodedUserInfo := base64.RawURLEncoding.EncodeToString([]byte(sc.Method + ":" + sc.Password))
	return "ss://" + encodedUserInfo + "@" + sc.Server + ":" + strconv.Itoa(sc.ServerPort) + "#" + url.QueryEscape(sc.Remarks)
}

type CommonConfigList []CommonConfig

func (scl *CommonConfigList) getSubscription() string {
	result := ""
	for _, v := range *scl {
		result += v.getSubscription() + "\n"
	}
	return base64.RawURLEncoding.EncodeToString([]byte(result))
}

type SsdConfig struct {
	Airport      string  `json:"airport"`
	Port         int     `json:"port"`
	Encryption   string  `json:"encryption"`
	Password     string  `json:"password"`
	TrafficUsed  float64 `json:"traffic_used"`
	TrafficTotal float64 `json:"traffic_total"`
	Expiry       string  `json:"expiry"`
	URL          string  `json:"url"`
	Servers      []struct {
		ID      int     `json:"id"`
		Server  string  `json:"server"`
		Ratio   float32 `json:"ratio"`
		Remarks string  `json:"remarks"`
	} `json:"servers"`
}
