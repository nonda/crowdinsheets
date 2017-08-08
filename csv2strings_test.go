package main

import "testing"

func TestCsv2Strings(t *testing.T) {
	type args struct {
		csvSource []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CSV with commas",
			args: args{
				csvSource: []byte("Source,English\nhello,world\n\"hel,lo\",\"wor,ld\"\n"),
			},
			want:    "\"hello\" = \"world\";\n\"hel,lo\" = \"wor,ld\";",
			wantErr: false,
		},
		{
			name: "Empty CSV",
			args: args{
				csvSource: []byte("Source,English"),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Csv2Strings(tt.args.csvSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("Csv2Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Csv2Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
