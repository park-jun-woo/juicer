//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestPathMethodToOperationID 테스트
package scanner

import (
	"testing"
)

func TestPathMethodToOperationID(t *testing.T) {
	tests := []struct {
		method string
		path   string
		want   string
	}{
		{"GET", "/api/v1/users", "get_users"},
		{"POST", "/api/v1/admin/buildings/:buildingId", "post_admin_buildings_buildingId"},
		{"DELETE", "/users/:id", "delete_users_id"},
	}

	for _, tt := range tests {
		t.Run(tt.method+"_"+tt.path, func(t *testing.T) {
			got := pathMethodToOperationID(tt.method, tt.path)
			if got != tt.want {
				t.Errorf("pathMethodToOperationID() = %q, want %q", got, tt.want)
			}
		})
	}
}
