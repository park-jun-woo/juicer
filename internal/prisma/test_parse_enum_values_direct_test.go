//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what parseEnumValues 값 추출(빈줄/속성줄 스킵) 직접 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestParseEnumValuesDirect(t *testing.T) {
	body := []string{"ADMIN", "", `USER @map("user")`, "@@schema(\"x\")", "  GUEST  "}
	got := parseEnumValues(body)
	want := []string{"ADMIN", "USER", "GUEST"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
