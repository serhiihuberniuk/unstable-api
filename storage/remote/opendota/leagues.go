package opendota

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func (f *Fetcher) Leagues(ctx context.Context) ([]models.Leagues, error) {
	reader, closeFn, err := f.fetchData(ctx, url+"leagues")
	if err != nil {
		return nil, fmt.Errorf("error while fetching data: %w", err)
	}
	defer closeFn()

	var leagues []models.Leagues
	if err = json.NewDecoder(reader).Decode(&leagues); err != nil {
		return nil, fmt.Errorf("error while decoding from json: %w", err)
	}

	return leagues, nil
}
