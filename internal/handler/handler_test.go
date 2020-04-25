package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"workshop/internal/api"
	"workshop/internal/api/mocks"
	"workshop/internal/handler"
)

func TestHandler_Hello(t *testing.T) {

	tests := []struct {
		name     string
		joke     *api.JokeResponse
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			joke:     &api.JokeResponse{Joke: "test joke"},
			err:      nil,
			codeWant: 200,
			bodyWant: "test joke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := &mocks.Client{}
			apiMock.On("GetJoke").
				Return(tt.joke, tt.err)

			h := handler.NewHandler(apiMock, "")

			req, _ := http.NewRequest("GET", "/hello", nil)
			rr := httptest.NewRecorder()

			h.Hello(rr, req)

			gotRaw, _ := ioutil.ReadAll(rr.Body)
			got := string(gotRaw)

			if got != tt.bodyWant {
				t.Errorf("wrong respode body %s want %s", got, tt.bodyWant)
			}

			if status := rr.Result().StatusCode; status != tt.codeWant {
				t.Errorf("wrong respode status %d want %d", status, tt.codeWant)
			}

		})
	}
}
