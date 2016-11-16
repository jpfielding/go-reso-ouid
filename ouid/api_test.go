package ouid

import (
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	testutils "github.com/jpfielding/gotest/testutils"
)

func TestSimple(t *testing.T) {
	var orgs []Organization
	each := func(each Organization, err error) error {
		orgs = append(orgs, each)
		return err
	}
	ctx := context.Background()
	req := func(ctx context.Context) (io.ReadCloser, error) {
		return ioutil.NopCloser(strings.NewReader(data)), nil
	}
	err := Process(ctx, req, each)
	testutils.Ok(t, err)
	testutils.Equals(t, 3, len(orgs))
	testutils.Equals(t, true, orgs[0].Active)
	testutils.Equals(t, "4205 Minnesota Drive", orgs[0].Location.Address)
	testutils.Equals(t, true, orgs[1].Active)
	testutils.Equals(t, "Anchorage", orgs[1].Location.City)
	testutils.Equals(t, true, orgs[2].Active)
	testutils.Equals(t, "99503", orgs[2].Location.Zip)
}

var data = `<?xml version="1.0" encoding="UTF-8"?>
<organizations>
 <organization>
  <name>ALASKA ASSOCIATION OF REALTORS? INC</name>
  <active>1</active>
  <location>
   <address>4205 Minnesota Drive</address>
   <city>Anchorage</city>
   <state>AK</state>
   <zip>99503</zip>
  </location>
 </organization>
 <organization>
  <name>Alaska MLS</name>
  <ouid>M00000001</ouid>
  <active>1</active>
  <location>
   <address>903 W. Northern Lights Blvd Suite 100</address>
   <city>Anchorage</city>
   <state>AK</state>
   <zip>99503</zip>
  </location>
 </organization>
 <organization>
  <name>Anchorage Board Of Realtors? Inc</name>
  <ouid>A00000002</ouid>
  <assoc2mls>M00000001</assoc2mls>
  <active>1</active>
  <location>
   <address>3340 Arctic Blvd. #101</address>
   <city>Anchorage</city>
   <state>AK</state>
   <zip>99503</zip>
  </location>
 </organization>
</organizations>`
