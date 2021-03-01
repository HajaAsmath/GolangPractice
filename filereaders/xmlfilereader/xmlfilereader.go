package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type CrewMember struct {
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearace,attr"`
	AccessCodes       []string `xml:"codes>code"`
}

type ShipInfo struct {
	XMLName   xml.Name   `xml:"Ship"`
	ShipID    int        `xml:"Ship>ShipId"`
	ShipClass string     `xml:"Ship>ShipClass"`
	Captain   CrewMember `xml:"Captain"`
}

func main() {
	codes := []string{"ALD", "ALL"}
	cm := CrewMember{ID: 11, Name: "Haja", SecurityClearance: 3311, AccessCodes: codes}
	shipinfo := ShipInfo{ShipID: 123, ShipClass: "AAA", Captain: cm}
	file, err := os.Create("xmlConfig.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, err := xml.MarshalIndent(shipinfo, " ", "	")
	if err != nil {
		fmt.Println(err)
		return
	}
	i, err := file.Write(marshal)
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	file1, err := os.Open("xmlConfig.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	bytereader := make([]byte, 222)
	file1.Read(bytereader)
	cm1 := new(ShipInfo)
	xml.Unmarshal(bytereader, cm1)
	fmt.Println(cm1.Captain.AccessCodes[0])
	file1.Close()
}
