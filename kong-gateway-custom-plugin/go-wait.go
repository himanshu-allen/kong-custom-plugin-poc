package main

import (
	"fmt"
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"time"
)

type Config struct {
	WaitTime int
}

func New1() interface{} {
	return &Config{}
}

var requests = make(map[string]time.Time)

func (config Config) Access(kong *pdk.PDK) {
	_ = kong.Response.SetHeader("x-wait-time", fmt.Sprintf("%d seconds", config.WaitTime))

	host, _ := kong.Request.GetHost()
	lastRequest, exists := requests[host]

	if exists && time.Now().Sub(lastRequest) < time.Duration(config.WaitTime)*time.Second {
		kong.Response.SetHeader("X-Custom-Header", "Limit Exceeded")
		kong.Response.Exit(400, "Maximum Requests Reached, Kindly wait and retry", make(map[string][]string))
	} else {
		requests[host] = time.Now()
	}
}

func main1() {
	Version := "1.1"
	Priority := 1000
	_ = server.StartServer(New1, Version, Priority)
}
