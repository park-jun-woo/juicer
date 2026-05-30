//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestParseSchema_Nil 테스트
package zod

import "testing"

func TestParseSchema_Nil(t *testing.T) {
	if got := ParseSchema(nil, nil); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
