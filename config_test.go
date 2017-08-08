package main

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    CrowdinSheetsConfig
		wantErr bool
	}{
		{
			name: "Parse example config",
			args: args{filename: "crowdin.yml.example"},
			want: CrowdinSheetsConfig{
				ProjectID: "PROJECT_NAME",
				APIToken:  "PROJECT_API_KEY",
				Languages: []string{
					"zh-CN", "zh-TW", "es-US", "it", "fr", "nl", "de", "ru",
				},
				Files: []string{
					"Localizable.strings",
					"android.xml",
				},
				OutputFolder: "./translations",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
