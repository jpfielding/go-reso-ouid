package ouid

import "encoding/xml"

// <organization>
//   <name>Magnolia Board Of Realtors</name>
//   <ouid>A00000068</ouid>
//   <active>0</active>
//   <notes>Merged with El Dorado</notes>
//   <location>
//     <address>
//     105 West Calhoun
//     </address>
//     <city>Magnolia</city>
//     <state>AR</state>
//     <zip>71753</zip>
//   </location>
// </organization>

// Organizations is a simple wrapper for collecting Organizations
type Organizations struct {
	XMLName      xml.Name       `xml:"organizations,omitempty"`
	Organization []Organization `xml:"organization" json:"organization"`
}

func (o Organizations) Len() int {
	return len(o.Organization)
}
func (o Organizations) Swap(i, j int) {
	o.Organization[i], o.Organization[j] = o.Organization[j], o.Organization[i]
}
func (o Organizations) Less(i, j int) bool {
	return o.Organization[i].Name < o.Organization[j].Name
}

// Organization defines the basic identity of a RESO organization
type Organization struct {
	Name     string    `xml:"name" json:"name"`
	OuID     string    `xml:"ouid" json:"ouid"`
	Active   bool      `xml:"active" json"active"` // TODO need to marshal to int
	Notes    string    `xml:"notes,omitempty" json:"notes,omitempty"`
	Location *Location `xml:"location,omitempty" json:"location,omitempty"`
}

// Location is where its located
type Location struct {
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	City    string `xml:"city,omitempty" json:"city,omitempty"`
	State   string `xml:"state,omitempty" json:"state,omitempty"`
	Zip     string `xml:"zip,omitempty" json:"zip,omitempty"`
}
