//ff:func feature=scan type=convert control=sequence
//ff:what TestPickBestResponse 테스트
package scanner

import (
	"testing"
)

func TestPickBestResponse(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		r := pickBestResponse([]Response{{Status: "200"}})
		if r.Status != "200" {
			t.Error("expected 200")
		}
	})

	t.Run("multiple - richer wins", func(t *testing.T) {
		r := pickBestResponse([]Response{
			{Status: "200"},
			{Status: "200", TypeName: "User", Fields: []Field{{Name: "id"}}},
		})
		if r.TypeName != "User" {
			t.Error("expected richer response to win")
		}
	})

	t.Run("first has type name", func(t *testing.T) {
		r := pickBestResponse([]Response{
			{Status: "200", TypeName: "User"},
			{Status: "200"},
		})
		if r.TypeName != "User" {
			t.Error("expected first with type name to win")
		}
	})
}
