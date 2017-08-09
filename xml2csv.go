package main

import (
	"encoding/xml"
	"fmt"

	"github.com/agrison/go-tablib"
)

type root struct {
	AndroidStrings      []AndroidString      `xml:"string"`
	AndroidStringsArray []AndroidStringArray `xml:"string-array"`
	AndroidIntArray     []AndroidIntArray    `xml:"integer-array"`
}

// AndroidString contains a single translation of a string
type AndroidString struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

// AndroidArrayItem is just an item in a string-array or int-array
type AndroidArrayItem struct {
	Value string `xml:",innerxml"`
}

// AndroidStringArray contains multiple strings that represent the same key
type AndroidStringArray struct {
	Key   string             `xml:"name,attr"`
	Items []AndroidArrayItem `xml:"item"`
}

// AndroidIntArray contains multiple int items that represent the same key
type AndroidIntArray struct {
	Key   string             `xml:"name,attr"`
	Items []AndroidArrayItem `xml:"item"`
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

	for _, s := range xmlRoot.AndroidStringsArray {
		for i, item := range s.Items {
			csv.AppendValues(fmt.Sprintf("%s_%d", s.Key, i), item.Value)
		}
	}

	for _, s := range xmlRoot.AndroidIntArray {
		for i, item := range s.Items {
			csv.AppendValues(fmt.Sprintf("%s_int_%d", s.Key, i), item.Value)
		}
	}

	csvOutput, err := csv.CSV()
	if err != nil {
		return []byte{}, err
	}

	return csvOutput.Bytes(), nil
}
