package scanrfc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func FetchRFC(rfc ...string) [][]byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	n := len(rfc)
	out := make([][]byte, n)
	for i, r := range rfc {
		e, err := fetchEntry(ctx, r)
		if err != nil {
			log.Error(err)
			continue
		}
		out[i] = e
	}
	return out
}

func fetchEntry(ctx context.Context, rfc string) ([]byte, error) {
	resp, err := fetchBibtex(ctx, rfc)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			return fetchJson(ctx, rfc)
		} else {
			return nil, err
		}
	}
	return resp, nil
}

func fetch(ctx context.Context, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		return b, nil
	default:
		return nil, fmt.Errorf("could not fetch %v, http status code %v", url, resp.StatusCode)
	}
}

func fetchBibtex(ctx context.Context, rfc string) ([]byte, error) {
	return fetch(ctx, "https://datatracker.ietf.org/doc/"+rfc+"/bibtex/")
}

func fetchJson(ctx context.Context, rfc string) ([]byte, error) {
	return fetch(ctx, "https://datatracker.ietf.org/doc/"+rfc+"/doc.json")
}
