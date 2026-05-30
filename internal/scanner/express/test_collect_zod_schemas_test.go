//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectZodSchemas 테스트
package express

import "testing"

func TestCollectZodSchemas(t *testing.T) {
	fi := mustParse(t, []byte(`const userSchema = z.object({ name: z.string() });`))
	schemas := collectZodSchemas(fi)
	if _, ok := schemas["userSchema"]; !ok {
		t.Fatalf("expected userSchema collected, got %v keys", len(schemas))
	}
}
