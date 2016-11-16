package ouid

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

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

// Scoper limits the scope of the Requester
type Scoper func(url.Values) error

// Request requests all Organizations from the endpoint
func (cfg *Config) Request(scoper Scoper) Requester {
	return func(ctx context.Context) (io.ReadCloser, error) {
		url, err := url.Parse(cfg.Endpoint)
		if err != nil {
			return nil, err
		}
		values := url.Query()
		err = scoper(values)
		if err != nil {
			return nil, err
		}
		url.RawQuery = values.Encode()
		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			return nil, err
		}
		resp, err := ctxhttp.Do(ctx, cfg.HTTP, req)
		return DefaultReEncodeReader(resp.Body, resp.Header.Get(ContentType)), nil
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
