package controller

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMusicCtrl_FetchMusic(t *testing.T) {
	type expected struct {
		result entity.MusicResult
		err    error
	}
	tests := []struct {
		name     string
		expected expected
		want     string
		wantErr  bool
	}{
		{
			name: "Success Fetch Music",
			expected: expected{
				result: entity.MusicResult{},
			},
			want: `{"data":{"list":null,"page":0,"row":0,"count":0},"message":"Success get data","status":"OK"}` + "\n",
			wantErr: false,
		},
		{
			name: "Error fetch data, return error response",
			expected: expected{
				err: errors.New("error"),
			},
			want:    `{"data":"","message":"error","status":"Bad Request"}` + "\n",
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := createCtx()
			mockService := service.NewMockMusicServiceItf(ctrl)
			mockService.EXPECT().FetchMusic(gomock.Any()).Return(tt.expected.result, tt.expected.err)

			d := &MusicCtrl{
				service: mockService,
			}
			if err := d.FetchMusic(ctx); (err != nil) != tt.wantErr {
				t.Errorf("FetchMusic() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestMusicCtrl_DetailMusic(t *testing.T) {
	type expected struct {
		result entity.Music
		err    error
	}
	tests := []struct {
		name     string
		expected expected
		want     string
		wantErr  bool
	}{
		{
			name: "Success Get Detail Music",
			expected: expected{
				result: entity.Music{},
				err:    nil,
			},
			want: `{"data":{"id":0,"title":"","singer":"","duration":"","album":"","release_year":""},"message":"Success get data","status":"OK"}` + "\n",
			wantErr: false,
		},
		{
			name: "Error get data, return error response",
			expected: expected{
				err: errors.New("error"),
			},
			want:    `{"data":"","message":"error","status":"Bad Request"}` + "\n",
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := createCtx()
			mockService := service.NewMockMusicServiceItf(ctrl)
			mockService.EXPECT().DetailMusic(gomock.Any()).Return(tt.expected.result, tt.expected.err)

			d := &MusicCtrl{
				service: mockService,
			}
			if err := d.DetailMusic(ctx); (err != nil) != tt.wantErr {
				t.Errorf("DetailMusic() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}
