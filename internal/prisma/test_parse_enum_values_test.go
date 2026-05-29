//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what parseEnumValues가 본문 라인에서 값 이름만 추출(@map/속성 라인 무시)하는지 검증
package prisma

import (
	"reflect"
	"testing"
)

func TestParseEnumValues(t *testing.T) {
	cases := []struct {
		name string
		body []string
		want []string
	}{
		{"plain values", []string{"USER", "ADMIN"}, []string{"USER", "ADMIN"}},
		{"value with @map ignored on attr", []string{`USER @map("user")`, "ADMIN"}, []string{"USER", "ADMIN"}},
		{"attribute-only line skipped", []string{"USER", `@@map("roles")`, "ADMIN"}, []string{"USER", "ADMIN"}},
		{"empty", nil, []string{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := parseEnumValues(c.body)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("parseEnumValues(%v) = %v, want %v", c.body, got, c.want)
			}
		})
	}
}
