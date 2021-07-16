package scanrfc

import (
	"reflect"
	"testing"
)

func TestFetchRFC(t *testing.T) {
	type args struct {
		rfc []string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchRFC(tt.args.rfc...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchRFC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchRFC() = %v, want %v", got, tt.want)
			}
		})
	}
}
