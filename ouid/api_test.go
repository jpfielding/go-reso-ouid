package ouid

import (
	"context"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	testutils "github.com/jpfielding/gotest/testutils"
)

func TestScopers(t *testing.T) {
	var v url.Values

	v = url.Values{}
	testutils.Ok(t, And(ByOuID("test"), ByActive(true))(v))
	testutils.Equals(t, "active=1&ouid=test", v.Encode())

	v = url.Values{}
	testutils.Ok(t, All()(v))
	testutils.Equals(t, "", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByOuID("test")(v))
	testutils.Equals(t, "ouid=test", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByAssocToMLS("test")(v))
	testutils.Equals(t, "assoc2mls=test", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByActive(true)(v))
	testutils.Equals(t, "active=1", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByName("test")(v))
	testutils.Equals(t, "org=test", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByCity("test")(v))
	testutils.Equals(t, "city=test", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByState("test")(v))
	testutils.Equals(t, "state=test", v.Encode())

	v = url.Values{}
	testutils.Ok(t, ByZip("test")(v))
	testutils.Equals(t, "zip=test", v.Encode())
}

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
	testutils.Equals(t, "", orgs[0].OuID)
	testutils.Equals(t, true, orgs[0].Active)
	testutils.Equals(t, "4205 Minnesota Drive", orgs[0].Location.Address)
	testutils.Equals(t, "M00000001", orgs[1].OuID)
	testutils.Equals(t, true, orgs[1].Active)
	testutils.Equals(t, "Anchorage", orgs[1].Location.City)
	testutils.Equals(t, "A00000002", orgs[2].OuID)
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
