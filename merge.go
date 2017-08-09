package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/agrison/go-tablib"
)

type mergedTranslations map[string]map[string]string

// MergeAllLocalizableCSV merges alls csv files into one
func MergeAllLocalizableCSV(config *CrowdinSheetsConfig) ([]byte, error) {
	mergedCsv := tablib.NewDataset([]string{"Source"})

	results := make(mergedTranslations)

	for _, language := range config.Languages {
		mergedCsv.AppendColumn(language, nil)
		results[language] = make(map[string]string)
	}

	for _, language := range config.Languages {
		csvPath := path.Join(config.OutputFolder, language, "Localizable.csv")
		csvBytes, err := ioutil.ReadFile(csvPath)
		if err != nil {
			continue
		}

		csvContent, err := tablib.LoadCSV(csvBytes)
		if err != nil {
			continue
		}

		var i int
		for {
			row, err := csvContent.Row(i)
			if err != nil {
				break
			}
			source, ok := row["Source"].(string)
			translation, ok := row["English"].(string)
			if !ok {
				i++
				continue
			}

			fmt.Printf("%s %s = %s", language, source, translation)

			results[language][source] = translation
			i++
		}

	}

	var language string
	for k := range results {
		if language == "" {
			language = k
		}
	}

	horizontaledResults := make([][]interface{}, len(results[language]))
	for source := range results[language] {
		columns := []string{}
		columns = append(columns, source)
		for _, lang := range config.Languages {
			columns = append(columns, results[lang][source])
		}

		columnsInterface := make([]interface{}, len(columns))
		for i, v := range columns {
			columnsInterface[i] = v
		}
		horizontaledResults = append(horizontaledResults, columnsInterface)
	}

	for _, row := range horizontaledResults {
		if err := mergedCsv.Append(row); err != nil {
			continue
		}
	}

	csvOutput, err := mergedCsv.CSV()
	return csvOutput.Bytes(), err
}

// func MergeAndroidCSVIntoStrings(androidCSVContent []byte, stringsCSV []byte) ([]byte, error) {

// }
