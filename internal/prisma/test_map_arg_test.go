//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what mapArg 접두 일치 및 인용 인자 추출 테스트
package prisma

import "testing"

func TestMapArg(t *testing.T) {
	v, ok := mapArg(`@@map("accounts")`, "@@map")
	if !ok || v != "accounts" {
		t.Errorf("got (%q,%v), want (accounts,true)", v, ok)
	}
	if _, ok := mapArg(`@id`, "@@map"); ok {
		t.Error("prefix mismatch must be false")
	}
	if _, ok := mapArg(`@@map(noquote)`, "@@map"); ok {
		t.Error("no opening quote must be false")
	}
	if _, ok := mapArg(`@@map("unterminated`, "@@map"); ok {
		t.Error("no closing quote must be false")
	}
}
