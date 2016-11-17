package ouid

import (
	"fmt"
	"net/url"
)

// And will apply multple scope limiting operations
func And(scopes ...Scoper) Scoper {
	return func(v url.Values) error {
		for _, s := range scopes {
			err := s(v)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// All requests all organizations
func All() Scoper {
	return func(url.Values) error {
		return nil
	}
}

// ByOuID searches by Organizational Unique ID.
/// Example: http://www.reso.org/ouid/?ouid=A00000007
func ByOuID(ouid string) Scoper {
	return byParam("ouid", ouid)
}

// ByAssocToMLS searches by association to an MLS organization.
// Example: http://www.reso.org/ouid/?assoc2mls=M00000001
// Example: http://www.reso.org/ouid/?assoc2mls=M00000002
func ByAssocToMLS(assoc2mls string) Scoper {
	return byParam("assoc2mls", assoc2mls)
}

// ByActive searches for active or inactive organizations.
// Example: Active: http://www.reso.org/ouid/?active=1
func ByActive(active bool) Scoper {
	v := 1
	if !active {
		v = 0
	}
	return byParam("active", fmt.Sprintf("%d", v))
}

// ByName searches by an organization’s name or a portion of the name
// Example: http://www.reso.org/ouid/?org=mlslistings
func ByName(part string) Scoper {
	return byParam("org", part)
}

// ByCity searches by an organization’s city
// Example: http://www.reso.org/ouid/?city=dallas
func ByCity(city string) Scoper {
	return byParam("city", city)
}

// ByState searches by an organization’s state
// Example: http://www.reso.org/ouid/?state=AR
func ByState(state string) Scoper {
	return byParam("state", state)
}

// ByZip searches by an organization’s zip
// Example: http://www.reso.org/ouid/?zip=90210
func ByZip(zip string) Scoper {
	return byParam("zip", zip)
}

// byParam is a helper to wrap the updating of url.Values
func byParam(key, value string) Scoper {
	return func(v url.Values) error {
		v.Add(key, value)
		return nil
	}
}
