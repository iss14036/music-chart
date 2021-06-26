package controller

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createCtx() *mockEchoContext {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e := echo.New()

	c := e.NewContext(req, rec)
	c.Set(constant.EchoUserID, "123")
	return &mockEchoContext{c}
}

type mockEchoContext struct {
	echo.Context
}

func (m *mockEchoContext) getResponseBody() []byte {
	res := m.Response().Writer.(*httptest.ResponseRecorder)
	b, _ := ioutil.ReadAll(res.Body)
	return b
}

func TestUserCtrl_GetUser(t *testing.T) {
	type expected struct {
		result entity.User
		err    error
	}
	tests := []struct {
		name     string
		expected expected
		want     string
		wantErr  bool
	}{
		{
			name: "Success Get User",
			expected: expected{
				result: entity.User{
					ID:       1,
					Username: "test",
					FullName: "test",
					Hobby:    "test",
				},
			},
			want: `{"data":{"id":1,"username":"test","password":"","full_name":"test","hobby":"test","gender":"",` +
				`"address":""},"message":"Success get data","status":"OK"}` + "\n",
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
			mockService := service.NewMockUserServiceItf(ctrl)
			mockService.EXPECT().GetUser(gomock.Any()).Return(tt.expected.result, tt.expected.err)

			d := &UserCtrl{
				service: mockService,
			}
			if err := d.GetUser(ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestUserCtrl_InsertUser(t *testing.T) {
	type expected struct {
		err error
	}
	tests := []struct {
		name     string
		expected expected
		body     string
		want     string
		wantErr  bool
	}{
		{
			name: "Success insert User",
			want: `{"data":"","message":"Success create data","status":"Created"}` + "\n",
			body: `{
					"username": "test",
					"password": "test",
					"full_name": "testing",
					"hobby": "coding",
					"gender": "pria",
					"address": "tangsel"
				}`,
			wantErr: false,
		}, {
			name: "Error insert user, return error response",
			expected: expected{
				err: errors.New("error"),
			},
			want:    `{"data":"","message":"There is field that empty. Please check again","status":"Bad Request"}` + "\n",
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
			mockService := service.NewMockUserServiceItf(ctrl)
			if tt.expected.err == nil {
				mockService.EXPECT().InsertUser(gomock.Any()).Return(tt.expected.err)
			}

			d := &UserCtrl{
				service: mockService,
			}

			if err := d.InsertUser(ctx); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestUserCtrl_Login(t *testing.T) {
	type expected struct {
		token string
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
			name: "Success Login",
			want: `{"data":{"token":"test"},"message":"Success Login","status":"OK"}` + "\n",
			body: `{
					"username": "test",
					"password": "test"
				}`,
			expected: expected{
				token: "test",
			},
			wantErr: false,
		}, {
			name: "Error Login, return error response",
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
			mockService := service.NewMockUserServiceItf(ctrl)
			mockService.EXPECT().Login(gomock.Any()).Return(tt.expected.token, tt.expected.err)

			d := &UserCtrl{
				service: mockService,
			}

			if err := d.Login(ctx); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ctx.getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}
