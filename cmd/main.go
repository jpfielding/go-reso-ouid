package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jpfielding/go-reso-ouid/ouid"
	"github.com/jpfielding/gowirelog/wirelog"
)

func main() {
	transport := wirelog.NewHTTPTransport()
	wirelog.LogToFile(transport, "/tmp/reso-ouid.log", true, true)
	client := http.Client{
		Transport: transport,
	}
	cfg := ouid.Config{
		HTTP:     &client,
		Endpoint: ouid.DefaultURL,
	}
	ctx := context.Background()
	scope := ouid.All()
	// city := ouid.ByCity("dallas")
	// active := ouid.ByActive(true)
	// scope := ouid.And(all, city, active)
	err := ouid.Process(ctx, cfg.Request(scope), func(org ouid.Organization, err error) error {
		fmt.Printf("%v\n", org)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
