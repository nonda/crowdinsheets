package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	configFile := flag.String("conf", "crowdin.yml", "crowdin.yml config file")
	action := flag.String("action", "sync", "sync|csv2strings|xml2csv|merge")

	flag.Parse()

	config, err := ReadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	if *action == "sync" {

		files := configToFiles(config)

		fmt.Println("Downloading...")
		DownloadTranslations(config, files)

	} else if *action == "csv2strings" {
		for _, lang := range config.Languages {
			folder := path.Join(config.OutputFolder, lang+".lproj")

			filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if filepath.Ext(path) == ".csv" {
					var baseName string
					n := strings.LastIndexByte(path, '.')
					if n >= 0 {
						baseName = path[:n]
					} else {
						baseName = path
					}
					stringsFilename := filepath.Join(folder, filepath.Base(baseName)+".strings")

					content, readError := ioutil.ReadFile(path)
					if readError != nil {
						return readError
					}

					stringsContent, convertError := Csv2Strings(content)
					if convertError != nil {
						return convertError
					}

					writeError := ioutil.WriteFile(stringsFilename, []byte(stringsContent), 0700)
					if writeError != nil {
						return writeError
					}
				}
				return nil
			})
		}
	} else if *action == "xml2csv" {
		filepath.Walk(config.OutputFolder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ".xml" {
				xmlContent, readError := ioutil.ReadFile(path)
				if readError != nil {
					panic(readError)
				}
				csvContent, convertError := Xml2Csv(xmlContent)
				if convertError != nil {
					panic(convertError)
				}

				folder := filepath.Dir(path)

				var baseName string
				n := strings.LastIndexByte(path, '.')
				if n >= 0 {
					baseName = path[:n]
				} else {
					baseName = path
				}
				csvFilename := filepath.Join(folder, filepath.Base(baseName)+".csv")

				writeError := ioutil.WriteFile(csvFilename, []byte(csvContent), 0700)
				if writeError != nil {
					return writeError
				}
			}
			return nil
		})
	} else if *action == "merge" {
		mergedCsv, err := MergeAllLocalizableCSV(&config)
		if err != nil {
			panic(err)
		}
		outputCsvPath := path.Join(config.OutputFolder, "LocalizableAll.csv")
		if err = ioutil.WriteFile(outputCsvPath, mergedCsv, 0700); err != nil {
			panic(err)
		}
	}

}
