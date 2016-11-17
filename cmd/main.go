package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/jpfielding/go-reso-ouid/ouid"
	"github.com/jpfielding/gowirelog/wirelog"
)

func main() {
	output := flag.String("filename", "", "File output")
	outputType := flag.String("type", "json", "Output type (xml | json)")

	flag.Parse()

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
	orgs := ouid.Organizations{}
	err := ouid.Process(ctx, cfg.Request(scope), func(org ouid.Organization, err error) error {
		orgs.Organization = append(orgs.Organization, org)
		return nil
	})
	if err != nil {
		panic(err)
	}
	out := os.Stdout
	if *output != "" {
		out, _ = os.Create(*output)
		defer out.Close()
	}

	switch strings.ToLower(*outputType) {
	case "xml":
		xml.NewEncoder(out).Encode(&orgs)
	default:
		raw, err := json.Marshal(&orgs.Organization)
		if err != nil {
			panic(err)
		}
		// formatted
		var buf bytes.Buffer
		json.Indent(&buf, raw, "", "\t")
		_, err = buf.WriteTo(out)
		if err != nil {
			panic(err)
		}
	}

}
