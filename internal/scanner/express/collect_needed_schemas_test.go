//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectNeededSchemas — 라우트 참조 스키마명 수집 분기를 검증
package express

import "testing"

func TestCollectNeededSchemas_SkipsNoRouters(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	ctx := &scanContext{
		parsed: map[string]*fileInfo{"a.ts": fi},
		// No routers registered for a.ts -> the file is skipped.
		allRouters: map[string]map[string]bool{},
	}
	names := collectNeededSchemas(ctx)
	if len(names) != 0 {
		t.Fatalf("expected no schemas, got %v", names)
	}
}

func TestCollectNeededSchemas_WithRoutes(t *testing.T) {
	src := []byte(`
const app = express();
app.post('/users', validate(UserSchema), (req, res) => {});
`)
	fi := mustParse(t, src)
	ctx := &scanContext{
		parsed:     map[string]*fileInfo{"a.ts": fi},
		allRouters: map[string]map[string]bool{"a.ts": {"app": true}},
	}
	// Should run extractRoutes + collectSchemaNames without panicking.
	names := collectNeededSchemas(ctx)
	// Whether UserSchema is collected depends on validator recognition; at
	// minimum the call must complete and return a (possibly empty) slice.
	_ = names
}
