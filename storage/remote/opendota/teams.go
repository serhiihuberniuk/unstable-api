package opendota

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func (f *Fetcher) Teams(ctx context.Context) ([]models.Team, error) {
	var teams []models.Team
	decode := func(reader io.Reader) error {
		if err := json.NewDecoder(reader).Decode(&teams); err != nil {
			return fmt.Errorf("error while decoding from json: %w", err)
		}
		return nil
	}

	if err := f.fetchData(ctx, url+"teams", decode); err != nil {
		return nil, fmt.Errorf("error while fetching data: %w", err)
	}

	return teams, nil
}
