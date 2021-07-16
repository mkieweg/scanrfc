package scanrfc

import (
	"context"
	"reflect"
	"testing"
)

func TestFetchRFC(t *testing.T) {
	type args struct {
		rfc []string
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchRFC(tt.args.rfc...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchRFC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchEntry(t *testing.T) {
	type args struct {
		ctx context.Context
		rfc string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchEntry(tt.args.ctx, tt.args.rfc)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetch(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetch(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchBibtex(t *testing.T) {
	type args struct {
		ctx context.Context
		rfc string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchBibtex(tt.args.ctx, tt.args.rfc)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchBibtex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchBibtex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchJson(t *testing.T) {
	type args struct {
		ctx context.Context
		rfc string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchJson(tt.args.ctx, tt.args.rfc)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
