//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectZodSchemas: z.object 스키마 수집 위임 검증
package express

import "testing"

func TestCollectZodSchemas(t *testing.T) {
	fi := mustParse(t, []byte(`const userSchema = z.object({ name: z.string() });`))
	schemas := collectZodSchemas(fi)
	if _, ok := schemas["userSchema"]; !ok {
		t.Fatalf("expected userSchema collected, got %v keys", len(schemas))
	}
}

func TestCollectZodSchemas_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	schemas := collectZodSchemas(fi)
	if len(schemas) != 0 {
		t.Fatalf("expected none, got %d", len(schemas))
	}
}
