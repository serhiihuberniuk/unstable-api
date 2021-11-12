package opendota

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.opendota.com/api/"

type Fetcher struct {
	client *http.Client
}

func New() *Fetcher {
	return &Fetcher{
		client: &http.Client{},
	}
}

func (f *Fetcher) fetchData(ctx context.Context, url string) (io.Reader, func(), error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error while creating request: %w", err)
	}

	resp, err := f.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, fmt.Errorf("error while sending response: %w", err)
	}

	closeFn := func() {
		resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, closeFn, errors.New("error status: " + resp.Status)
	}

	return bufio.NewReader(resp.Body), closeFn, nil
}
