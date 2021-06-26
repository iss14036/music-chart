package controller

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFavoriteCtrl_FetchFavorite(t *testing.T) {
	type expected struct {
		result entity.FavoriteResult
		err    error
	}
	tests := []struct {
		name     string
		expected expected
		want     string
		wantErr  bool
	}{
		{
			name: "Success Fetch Favorite",
			expected: expected{
				result: entity.FavoriteResult{},
			},
			want: `{"data":{"list":null,"page":0,"row":0,"count":0},"message":"Success get data","status":"OK"}` + "\n",
			wantErr: false,
		},
		{
			name: "Error fetch data, return error response",
			expected: expected{
				err: errors.New("error"),
			},
			want:    `{"data":"","message":"error","status":"Internal Server Error"}` + "\n",
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := createCtx()
			mockService := service.NewMockFavoriteServiceItf(ctrl)
			mockService.EXPECT().GetFavorite(gomock.Any()).Return(tt.expected.result, tt.expected.err)

			d := &FavoriteCtrl{
				service: mockService,
			}
			if err := d.FetchFavorite(ctx); (err != nil) != tt.wantErr {
				t.Errorf("FetchFavorite() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestFavoriteCtrl_SetFavorite(t *testing.T) {
	type expected struct {
		err   error
	}
	tests := []struct {
		name     string
		expected expected
		body     string
		want     string
		wantErr  bool
	}{
		{
			name: "Success Set Favorite",
			want: `{"data":"","message":"Success create data","status":"Created"}` + "\n",
			body: `{
					"music_id": 1
				}`,
			wantErr: false,
		}, {
			name: "Error Set Favorite, return error response",
			expected: expected{
				err: errors.New("error login"),
			},
			want:    `{"data":"","message":"error login","status":"Internal Server Error"}` + "\n",
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx := createCtx()
			ctx.Context = c
			ctx.Set("user_id", float64(123))
			mockService := service.NewMockFavoriteServiceItf(ctrl)
			mockService.EXPECT().SetFavorite(gomock.Any()).Return(tt.expected.err)

			d := &FavoriteCtrl{
				service: mockService,
			}

			if err := d.SetFavorite(ctx); (err != nil) != tt.wantErr {
				t.Errorf("SetFavorite() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}
