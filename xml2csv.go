package main

import (
	"encoding/xml"

	"github.com/agrison/go-tablib"
)

type root struct {
	AndroidStrings []AndroidString `xml:"string"`
}

// AndroidString contains a single translation of a string
type AndroidString struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

// Xml2Csv converts android.xml to android.csv
func Xml2Csv(xmlData []byte) ([]byte, error) {
	var xmlRoot root

	if err := xml.Unmarshal(xmlData, &xmlRoot); err != nil {
		return []byte{}, err
	}

	csv := tablib.NewDataset([]string{"key", "translation"})

	for _, s := range xmlRoot.AndroidStrings {
		csv.AppendValues(s.Key, s.Value)
	}

	csvOutput, err := csv.CSV()
	if err != nil {
		return []byte{}, err
	}

	return csvOutput.Bytes(), nil
}
