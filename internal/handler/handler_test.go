package handler

import (
	"net/http"
	"testing"
	"workshop/internal/api"
)

func TestHandler_Hello(t *testing.T) {
	type fields struct {
		jokeClient api.Client
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				jokeClient: tt.fields.jokeClient,
			}
			h.Hello(tt.args.w, tt.args.r)
		})
	}
}
