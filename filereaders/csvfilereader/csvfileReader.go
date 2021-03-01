package main

type CrewMember struct {
	ID                int      `xml:"id,omitempty",csv:"id"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearace,attr"`
	AccessCodes       []string `xml:"codes>code"`
}

type ShipInfo struct {
	ShipID    int        `xml:"Ship>ShipId"`
	ShipClass string     `xml:"Ship>ShipClass"`
	Captain   CrewMember `xml:"Captain"`
}

func main() {

}
