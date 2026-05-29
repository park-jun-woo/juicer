//ff:func feature=scan type=convert control=sequence
//ff:what TestDeduplicateEndpoints 테스트
package scanner

import (
	"testing"
)

func TestDeduplicateEndpoints(t *testing.T) {
	t.Run("no duplicates", func(t *testing.T) {
		eps := []Endpoint{
			{Method: "GET", Path: "/a"},
			{Method: "POST", Path: "/b"},
		}
		got := DeduplicateEndpoints(eps)
		if len(got) != 2 {
			t.Errorf("expected 2, got %d", len(got))
		}
	})

	t.Run("with duplicates - richer wins", func(t *testing.T) {
		eps := []Endpoint{
			{Method: "GET", Path: "/a"},
			{Method: "GET", Path: "/a", Responses: []Response{{TypeName: "User", Fields: []Field{{Name: "id"}}}}},
		}
		got := DeduplicateEndpoints(eps)
		if len(got) != 1 {
			t.Errorf("expected 1, got %d", len(got))
		}
		if len(got[0].Responses) != 1 {
			t.Error("expected richer endpoint to win")
		}
	})
}
