//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what drfFieldToOpenAPI — 모든 DRF 필드 타입 변환 분기를 검증
package django

import "testing"

func TestDrfFieldToOpenAPI_AllBranches(t *testing.T) {
	cases := []struct {
		in         string
		wantType   string
		wantFormat string
	}{
		{"SlugField", "string", ""},
		{"SlugRelatedField", "string", ""},
		{"DecimalField", "number", ""},
		{"NullBooleanField", "boolean", ""},
		{"TimeField", "string", "time"},
		{"ImageField", "string", "binary"},
		{"ListSerializer", "array", ""},
		{"ManyRelatedField", "array", ""},
		{"DictField", "object", ""},
		{"HStoreField", "object", ""},
		{"ChoiceField", "string", ""},
		{"IPAddressField", "string", "ipv4"},
		{"GenericIPAddressField", "string", "ipv4"},
		{"DurationField", "string", ""},
		{"SerializerMethodField", "string", ""},
		{"UnknownField", "string", ""}, // default
	}
	for _, c := range cases {
		got := drfFieldToOpenAPI(c.in)
		if got.Type != c.wantType || got.Format != c.wantFormat {
			t.Errorf("%s -> {%q,%q}, want {%q,%q}", c.in, got.Type, got.Format, c.wantType, c.wantFormat)
		}
	}
}
