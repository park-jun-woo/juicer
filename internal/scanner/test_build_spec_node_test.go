//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildSpecNode 테스트
package scanner

import (
	"testing"
)

func TestBuildSpecNode(t *testing.T) {
	t.Run("with paths", func(t *testing.T) {
		result := &ScanResult{
			Endpoints: []Endpoint{
				{Method: "GET", Path: "/users", Handler: "h.GetUsers"},
			},
		}
		node := buildSpecNode(result)
		if node == nil {
			t.Fatal("expected non-nil node")
		}
	})

	t.Run("with schemas", func(t *testing.T) {
		result := &ScanResult{
			Endpoints: []Endpoint{
				{
					Method:  "POST",
					Path:    "/users",
					Handler: "h.CreateUser",
					Request: &Request{
						Body: &Body{
							TypeName: "CreateUserRequest",
							Fields:   []Field{{Name: "name", Type: "string"}},
						},
					},
				},
			},
		}
		node := buildSpecNode(result)
		if node == nil {
			t.Fatal("expected non-nil node")
		}
	})

	t.Run("Any method", func(t *testing.T) {
		result := &ScanResult{
			Endpoints: []Endpoint{
				{Method: "Any", Path: "/health", Handler: "h.Health"},
			},
		}
		node := buildSpecNode(result)
		if node == nil {
			t.Fatal("expected non-nil node")
		}
	})
}
