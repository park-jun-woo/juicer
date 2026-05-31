//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what splitAttrs 괄호/대괄호 균형 유지하며 @속성 분리 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestSplitAttrs(t *testing.T) {
	got := splitAttrs("@id @default(now()) @map([a, b])")
	want := []string{"@id", "@default(now())", "@map([a, b])"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := splitAttrs(""); len(got) != 0 {
		t.Errorf("empty: %v", got)
	}
	// @ inside parens must not split
	got = splitAttrs(`@default("a@b")`)
	if !reflect.DeepEqual(got, []string{`@default("a@b")`}) {
		t.Errorf("nested @: got %v", got)
	}
}
