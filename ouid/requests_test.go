package ouid

import (
	"net/url"
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
