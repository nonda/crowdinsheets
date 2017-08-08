package main

import (
	"fmt"
	"strings"

	"github.com/agrison/go-tablib"
)

func escapeTranslationString(src string) string {
	escaped := strings.Replace(src, "\n", "\\n", 0)
	escaped = strings.Replace(escaped, "\"", "\\\"", 0)
	return escaped
}

// Csv2Strings convert input CSV content to iOS strings file
func Csv2Strings(csvSource []byte) (string, error) {
	csv, err := tablib.LoadCSV(csvSource)
	if err != nil {
		return "", err
	}

	var results []string
	var currentRow map[string]interface{}

	var i int
	for {
		currentRow, err = csv.Row(i)
		if err != nil {
			break
		}

		source, ok := currentRow["Source"].(string)
		translation, ok := currentRow["English"].(string)
		if !ok {
			i++
			continue
		}

		source = escapeTranslationString(source)
		translation = escapeTranslationString(translation)
		line := fmt.Sprintf("\"%s\" = \"%s\";", source, translation)
		results = append(results, line)
		i++
	}
	return strings.Join(results, "\n"), nil
}
