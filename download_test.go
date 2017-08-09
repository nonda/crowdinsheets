package main

import (
	"path"
	"reflect"
	"testing"

	crowdin "github.com/medisafe/go-crowdin"
)

func Test_configToFiles(t *testing.T) {
	type args struct {
		config CrowdinSheetsConfig
	}
	tests := []struct {
		name string
		args args
		want []crowdin.ExportFileOptions
	}{
		{
			name: "Must covert successfully",
			args: args{config: CrowdinSheetsConfig{
				ProjectID: "PROJECT_NAME",
				APIToken:  "PROJECT_API_KEY",
				Languages: []string{
					"zh-CN",
				},
				Files: []string{
					"Localizable.strings",
				},
				OutputFolder: "./translations",
			}},
			want: []crowdin.ExportFileOptions{
				crowdin.ExportFileOptions{
					CrowdinFile: "Localizable.strings",
					Language:    "zh_CN",
					LocalPath:   path.Join("./translations", "zh_CN", "Localizable.strings"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configToFiles(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configToFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertOsxCodeToCrowdinCode(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zh-CN",
			args: args{language: "zh-Hans"},
			want: "zh-CN",
		},
		{
			name: "zh-TW",
			args: args{language: "zh-Hant"},
			want: "zh-TW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertOsxCodeToCrowdinCode(tt.args.language); got != tt.want {
				t.Errorf("convertOsxCodeToCrowdinCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
