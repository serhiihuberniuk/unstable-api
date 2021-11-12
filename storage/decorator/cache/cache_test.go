package cache_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/serhiihuberniuk/unstable-api/models"
	"github.com/serhiihuberniuk/unstable-api/storage/decorator/cache"
	"github.com/stretchr/testify/assert"
)

func TestCacheDecorator_Leagues(t *testing.T) {
	t.Parallel()

	type decoratedMockBehavior func(d *Mockdecorated, ctx context.Context, leagues []models.Leagues)
	type cacheMockBehavior func(c *Mockcache, ctx context.Context, key string)

	inCtx := context.Background()
	leagues := []models.Leagues{
		{
			LeagueID: 0,
			Ticket:   "string",
			Banner:   "string",
			Tier:     "string",
			Name:     "string",
		},
	}
	var cacheData interface{}
	cacheData = leagues

	testCases := []struct {
		name                  string
		decoratedMockBehavior decoratedMockBehavior
		cacheMockBehavior     cacheMockBehavior
		errMessageExpected    string
		leaguesExpected       []models.Leagues
	}{
		{
			name: "OK",
			decoratedMockBehavior: func(d *Mockdecorated, ctx context.Context, leagues []models.Leagues) {
				d.EXPECT().Leagues(context.Background()).Return(leagues, nil)
			},
			cacheMockBehavior: func(c *Mockcache, ctx context.Context, key string) {
				c.EXPECT().Get(key).Return(nil, false)
				c.EXPECT().Put(context.Background(), key, leagues)
			},
			errMessageExpected: "",
			leaguesExpected:    leagues,
		},
		{
			name: "OK from cache",
			decoratedMockBehavior: func(d *Mockdecorated, ctx context.Context, leagues []models.Leagues) {
			},
			cacheMockBehavior: func(c *Mockcache, ctx context.Context, key string) {
				c.EXPECT().Get(key).Return(cacheData, true)
			},
			errMessageExpected: "",
			leaguesExpected:    leagues,
		},
		{
			name: "decorated error",
			decoratedMockBehavior: func(d *Mockdecorated, ctx context.Context, leagues []models.Leagues) {
				d.EXPECT().Leagues(ctx).Return(nil, errors.New("error"))
			},
			cacheMockBehavior: func(c *Mockcache, ctx context.Context, key string) {
				c.EXPECT().Get(key).Return(nil, false)
			},
			errMessageExpected: "error",
			leaguesExpected:    nil,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			decoratedMock := NewMockdecorated(ctrl)
			cacheMock := NewMockcache(ctrl)

			c := cache.NewCacheDecorator(decoratedMock, cacheMock)

			tc.decoratedMockBehavior(decoratedMock, inCtx, leagues)
			tc.cacheMockBehavior(cacheMock, inCtx, "/leagues")

			l, err := c.Leagues(inCtx)
			if tc.errMessageExpected == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.leaguesExpected, l)

				return
			}

			assert.Equal(t, tc.leaguesExpected, l)
			assert.Contains(t, err.Error(), tc.errMessageExpected)

		})
	}
}
