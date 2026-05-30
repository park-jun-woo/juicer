//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractPairStringValue: 키매칭+string값 / 키불일치 / string아님 분기
package express

import "testing"

func TestExtractPairStringValue_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: '/users', other: x };`))
	if got := extractPairStringValue(firstObject(t, fi), fi.Src, "path"); got != "/users" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractPairStringValue_KeyMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { foo: 'bar' };`))
	if got := extractPairStringValue(firstObject(t, fi), fi.Src, "path"); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractPairStringValue_ValueNotString(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: somevar };`))
	if got := extractPairStringValue(firstObject(t, fi), fi.Src, "path"); got != "" {
		t.Fatalf("got %q", got)
	}
}
