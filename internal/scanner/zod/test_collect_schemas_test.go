//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectSchemas 테스트
package zod

import "testing"

func TestCollectSchemas(t *testing.T) {
	root, src := parseTS(t, `
const userSchema = z.object({ name: z.string() });
const x = 5;
`)
	schemas := CollectSchemas(root, src)
	if _, ok := schemas["userSchema"]; !ok {
		t.Fatalf("expected userSchema, got %v keys", len(schemas))
	}
}
