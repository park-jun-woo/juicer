//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestSubstituteTypeParams — 타입 파라미터 치환 테스트
package spring

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestSubstituteTypeParams(t *testing.T) {
	fields := []scanner.Field{
		{Name: "content", Type: "T"},
		{Name: "items", Type: "List<T>"},
		{Name: "total", Type: "integer"},
	}
	result := substituteTypeParams(fields, []string{"T"}, []string{"AlbumResponse"})
	if result[0].Type != "AlbumResponse" {
		t.Errorf("field[0].Type = %q, want AlbumResponse", result[0].Type)
	}
	if result[1].Type != "List<AlbumResponse>" {
		t.Errorf("field[1].Type = %q, want List<AlbumResponse>", result[1].Type)
	}
	if result[2].Type != "integer" {
		t.Errorf("field[2].Type = %q, want integer", result[2].Type)
	}
}
