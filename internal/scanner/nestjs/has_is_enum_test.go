//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what hasIsEnum 테스트
package nestjs

import "testing"

func TestHasIsEnum(t *testing.T) {
	if !hasIsEnum([]string{"IsString", "IsEnum"}) {
		t.Fatal("expected true for IsEnum in list")
	}
	if hasIsEnum([]string{"IsString", "IsInt"}) {
		t.Fatal("expected false when IsEnum not in list")
	}
	if hasIsEnum(nil) {
		t.Fatal("expected false for nil")
	}
}
