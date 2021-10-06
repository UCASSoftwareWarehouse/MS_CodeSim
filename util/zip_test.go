package util

import "testing"

func TestUnzip(t *testing.T) {
	err := Unzip("/Users/purchaser/Desktop/GinAPI.zip", "/Users/purchaser/Desktop/new_dir")
	t.Logf("err=[%v]", err)
}

func TestZip(t *testing.T) {
	type args struct {
		srcFile string
		destZip string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Zip(tt.args.srcFile, tt.args.destZip); (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
