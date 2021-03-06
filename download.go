package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	crowdin "github.com/medisafe/go-crowdin"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func convertOsxCodeToCrowdinCode(language string) string {
	s := strings.Replace(language, "zh-Hans", "zh-CN", -1)
	s = strings.Replace(s, "zh-Hant", "zh-TW", -1)
	return s
}

func configToFiles(config CrowdinSheetsConfig) []crowdin.ExportFileOptions {
	options := []crowdin.ExportFileOptions{}

	folder := config.OutputFolder

	for _, lang := range config.Languages {
		for _, file := range config.Files {
			localFilename := path.Join(folder, lang+".lproj", file)

			options = append(options, crowdin.ExportFileOptions{
				Language:    convertOsxCodeToCrowdinCode(lang),
				CrowdinFile: file,
				LocalPath:   localFilename,
			})
		}
	}
	return options
}

// DownloadTranslations download translation files of the languages specified in the config file
func DownloadTranslations(config CrowdinSheetsConfig, translations []crowdin.ExportFileOptions) {
	api := crowdin.New(config.APIToken, config.ProjectID)

	p := mpb.New()
	var currentLang, currentFilename string

	nameFn := func(s *decor.Statistics) string {
		return currentLang + "/" + currentFilename
	}

	bar := p.AddBar(int64(len(translations)), mpb.PrependDecorators(
		decor.DynamicName(nameFn, 10, 0),
	), mpb.AppendDecorators(
		decor.Percentage(5, 0),
		decor.Counters(" | Files: %s/%s", 0, 0, 0),
	))

	for _, op := range translations {
		currentLang = path.Base(path.Dir(op.LocalPath))
		currentFilename = path.Base(op.LocalPath)

		folder, _ := filepath.Split(op.LocalPath)
		_ = os.MkdirAll(folder, 0755)
		err := api.ExportFile(&op)
		if err != nil {
			fmt.Printf("Failed to download file: %v, %v", op, err)
		}
		bar.Incr(1)
	}

	p.Stop()
}
