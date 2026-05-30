//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractUseMount: 정상마운트 + 각 nil 분기
package express

import "testing"

func TestExtractUseMount_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', userRouter);`))
	m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, map[string]string{"userRouter": "./user.ts"})
	if m == nil || m.Prefix != "/api" || m.VarName != "userRouter" || m.SourceRouter != "app" || m.FilePath != "./user.ts" {
		t.Fatalf("got %+v", m)
	}
}

func TestExtractUseMount_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`use('/api', r);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestExtractUseMount_UnregisteredRouter(t *testing.T) {
	fi := mustParse(t, []byte(`other.use('/api', r);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestExtractUseMount_NoIdentifierObject(t *testing.T) {
	// member object is itself a member_expression -> no direct identifier child
	fi := mustParse(t, []byte(`a.b.use('/x', r);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"a": true, "b": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestExtractUseMount_NotUse(t *testing.T) {
	fi := mustParse(t, []byte(`app.get('/api', r);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestExtractUseMount_ParseNil(t *testing.T) {
	// app.use with single arg -> parseUseMountArgs returns nil
	fi := mustParse(t, []byte(`app.use(mw);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestExtractUseMount_NoArgsNode(t *testing.T) {
	fi := mustParse(t, []byte("app.use`x`;"))
	if m := extractUseMount(firstCallExpr(t, fi), fi.Src, map[string]bool{"app": true}, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}
