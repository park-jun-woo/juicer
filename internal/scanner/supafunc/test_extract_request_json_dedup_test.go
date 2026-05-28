//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractRequestJSON_Dedup switch-case 반복 req.json() 구조분해 시 필드 중복 제거 테스트
package supafunc

import (
	"testing"
)

func TestExtractRequestJSON_Dedup(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const method = req.method
  switch (method) {
    case "scan": {
      const { repo } = await req.json()
      break
    }
    case "status": {
      const { repo } = await req.json()
      break
    }
    case "report": {
      const { repo } = await req.json()
      break
    }
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	fields := extractRequestJSON(body, fi.Src)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d: %v", len(fields), fields)
	}
	if fields[0] != "repo" {
		t.Fatalf("expected field 'repo', got %q", fields[0])
	}
}
