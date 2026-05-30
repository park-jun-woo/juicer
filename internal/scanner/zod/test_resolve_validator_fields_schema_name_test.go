//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestResolveValidatorFields_SchemaName 테스트
package zod

import "testing"

func TestResolveValidatorFields_SchemaName(t *testing.T) {
	root, src := parseTS(t, `const userSchema = z.object({ name: z.string() });`)
	schemas := CollectSchemas(root, src)
	v := ValidatorInfo{Target: "json", SchemaName: "userSchema"}
	fields := ResolveValidatorFields(v, schemas, src, map[string][]byte{"userSchema": src})
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}
