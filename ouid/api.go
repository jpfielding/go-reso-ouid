package ouid

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/jpfielding/gominidom/minidom"

	"golang.org/x/net/context/ctxhttp"
)

// ContentType ...
const ContentType string = "Content-Type"

// DefaultURL is the advertised endpoint for RESO ouid service
const DefaultURL = "http://www.reso.org/ouid/"

// EachOrg is a callback for each found ouid.Organization with the option to return any errors
type EachOrg func(Organization, error) error

// Requester provides a common func for extracting Organizations
type Requester func(context.Context) (io.ReadCloser, error)

// Config is the request configuration
type Config struct {
	HTTP     *http.Client
	Endpoint string
	Decoder  func(io.Reader, bool) *xml.Decoder
}

// All requests all Organizations from the endpoint
func All(cfg Config) Requester {
	return func(ctx context.Context) (io.ReadCloser, error) {
		req, err := http.NewRequest("GET", cfg.Endpoint, nil)
		if err != nil {
			return nil, err
		}
		resp, err := ctxhttp.Do(ctx, cfg.HTTP, req)
		return DefaultReEncodeReader(resp.Body, resp.Header.Get(ContentType)), nil
		// return res.Body, err
	}
}

// Process all requested Orgnizations
func Process(ctx context.Context, requester Requester, each EachOrg) error {
	in, err := requester(ctx)
	defer in.Close()
	if err != nil {
		return err
	}
	parser := DefaultXMLDecoder(in, true)
	md := minidom.MiniDom{
		EndFunc: minidom.QuitAt("organizations"),
	}
	return md.Walk(parser, minidom.ByName("organization"), func(segment io.ReadCloser, err error) error {
		tmp := Organization{}
		xml.NewDecoder(segment).Decode(&tmp)
		return each(tmp, err)
	})
}
