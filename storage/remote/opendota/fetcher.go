package opendota

import (
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

func (f *Fetcher) fetchData(ctx context.Context, url string, decode func(reader io.Reader) error) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error while creating request: %w", err)
	}

	resp, err := f.client.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("error while sending response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error status: " + resp.Status)
	}

	if err = decode(resp.Body); err != nil {
		return fmt.Errorf("error while decoding data: %w", err)
	}

	return nil
}
