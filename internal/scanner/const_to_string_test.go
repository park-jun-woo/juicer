package scanner

import (
	"go/constant"
	"testing"
)

func TestConstToString_Int(t *testing.T) {
	v := constant.MakeInt64(200)
	got := constToString(v)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}

func TestConstToString_String(t *testing.T) {
	v := constant.MakeString("hello")
	got := constToString(v)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
