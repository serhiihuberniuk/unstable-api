package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/serhiihuberniuk/unstable-api/handler"
	"github.com/serhiihuberniuk/unstable-api/models"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Leagues(t *testing.T) {
	t.Parallel()

	type mockBehavior func(f *Mockfetcher, ctx context.Context)

	jsonData := `[
{
"leagueid":0,"ticket":"string","banner":"string","tier":"string","name":"string"}]
`
	bytesData := []byte(jsonData)

	ctx := context.Background()

	leagues := []models.Leagues{
		{
			LeagueID: 0,
			Ticket:   "string",
			Banner:   "string",
			Tier:     "string",
			Name:     "string",
		},
	}

	testCases := []struct {
		name         string
		mockBehavior mockBehavior
		expectedCode int
		expectedBody []byte
	}{
		{
			name: "OK",
			mockBehavior: func(f *Mockfetcher, ctx context.Context) {
				f.EXPECT().Leagues(ctx).Return(leagues, nil)
			},
			expectedCode: 200,
			expectedBody: bytesData,
		},
		{
			name: "timeout",
			mockBehavior: func(f *Mockfetcher, ctx context.Context) {
				f.EXPECT().Leagues(ctx).Return(nil, models.ErrTimeout)
			},
			expectedCode: 408,
			expectedBody: []byte(http.StatusText(408) + "\n"),
		},
		{
			name: "internal",
			mockBehavior: func(f *Mockfetcher, ctx context.Context) {
				f.EXPECT().Leagues(ctx).Return(nil, models.ErrServer)
			},
			expectedCode: 500,
			expectedBody: []byte(http.StatusText(500) + "\n"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			fetchMock := NewMockfetcher(ctrl)

			h := handler.New(fetchMock)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "https://api.opendota.com/leagues", nil)

			tc.mockBehavior(fetchMock, ctx)

			h.Leagues(w, r)

			assert.Equal(t, tc.expectedCode, w.Code)
			assert.Equal(t, string(tc.expectedBody), string(w.Body.Bytes()))
		})
	}
}
