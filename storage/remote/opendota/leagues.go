package opendota

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func (f *Fetcher) Leagues(ctx context.Context) ([]models.Leagues, error) {
	var leagues []models.Leagues
	decode := func(reader io.Reader) error {
		if err := json.NewDecoder(reader).Decode(&leagues); err != nil {
			return fmt.Errorf("error while decoding from json: %w", err)
		}
		return nil
	}

	if err := f.fetchData(ctx, url+"leagues", decode); err != nil {
		return nil, fmt.Errorf("error while fetching data: %w", err)
	}

	return leagues, nil
}
