package opendota

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func (f *Fetcher) Teams(ctx context.Context) ([]models.Team, error) {

	reader, closeFn, err := f.fetchData(ctx, url+"teams")
	if err != nil {
		return nil, fmt.Errorf("error while fetching data: %w", err)
	}
	defer closeFn()

	var teams []models.Team

	if err := json.NewDecoder(reader).Decode(&teams); err != nil {
		return nil, fmt.Errorf("error while gecoding: %w", err)
	}

	return teams, nil
}
