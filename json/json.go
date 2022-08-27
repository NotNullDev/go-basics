package main

import (
	"encoding/json"
	"encoding/xml"
)

type InnerTest struct {
	Aa string `json:"aa"`
}

type Test struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sample InnerTest
}

type TestXML struct {
	XmlName xml.Name
	Id      int
	Name    string
}

func main() {
	t := Test{
		Id:   1,
		Name: "LOL",
		Sample: InnerTest{
			Aa: "XD",
		},
	}

	tJSON, _ := json.MarshalIndent(t, "", "  ")

	println(string(tJSON))

	t1 := TestXML{
		XmlName: xml.Name{
			Local: "XDDLOLCAL",
			Space: "IDKSPACE",
		},
		Id:   1,
		Name: "OMAG",
	}

	t1XML, _ := xml.MarshalIndent(t1, "", "  ")

	println(string(t1XML))
}
