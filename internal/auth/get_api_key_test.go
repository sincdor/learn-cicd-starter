package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}

	headers := make(http.Header)
	headers["Authorization"] = append(make([]string, 1), "1234")

	tests := []test{
		{
			name:    "Valid Authorization header",
			headers: http.Header{"Authorization": {"ApiKey 1234"}},
			want:    "1234",
			wantErr: false,
		},
		{
			name:    "Missing ApiKey",
			headers: http.Header{"Authorization": {"1234"}},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
