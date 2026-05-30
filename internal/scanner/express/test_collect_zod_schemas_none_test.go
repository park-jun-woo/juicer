//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectZodSchemas_None 테스트
package express

import "testing"

func TestCollectZodSchemas_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	schemas := collectZodSchemas(fi)
	if len(schemas) != 0 {
		t.Fatalf("expected none, got %d", len(schemas))
	}
}
