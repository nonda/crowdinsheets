package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	configFile := flag.String("conf", "crowdin.yml", "crowdin.yml config file")
	action := flag.String("action", "sync", "sync|csv2strings|xml2csv")

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
		filepath.Walk(config.OutputFolder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ".csv" {
				folder := filepath.Dir(path)

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
	} else if *action == "xml2csv" {

	}

}
