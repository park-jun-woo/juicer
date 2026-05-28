//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractBodyMemberAccess const body = await req.json() 후 body.field dot 접근 필드 추출 테스트
package supafunc

import (
	"testing"
)

func TestExtractBodyMemberAccess(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const body = await req.json()
  const result = await processTask(body.task)
  await updateStatus(body.status)
  return new Response(JSON.stringify(result))
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	// extractRequestJSON should fall back to extractBodyMemberAccess
	fields := extractRequestJSON(body, fi.Src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %v", len(fields), fields)
	}
	if fields[0] != "task" {
		t.Fatalf("expected first field 'task', got %q", fields[0])
	}
	if fields[1] != "status" {
		t.Fatalf("expected second field 'status', got %q", fields[1])
	}
}
