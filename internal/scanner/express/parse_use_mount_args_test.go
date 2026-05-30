//ff:func feature=scan type=test control=sequence topic=express
//ff:what parseUseMountArgs: 정상 / 인자부족 / prefix비문자열 / router비식별자
package express

import "testing"

func TestParseUseMountArgs_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', userRouter);`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	m := parseUseMountArgs(args, fi.Src, map[string]string{"userRouter": "./u.ts"})
	if m == nil || m.Prefix != "/api" || m.VarName != "userRouter" || m.FilePath != "./u.ts" {
		t.Fatalf("got %+v", m)
	}
}

func TestParseUseMountArgs_TooFew(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api');`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if m := parseUseMountArgs(args, fi.Src, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestParseUseMountArgs_PrefixNotString(t *testing.T) {
	fi := mustParse(t, []byte(`app.use(mw, userRouter);`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if m := parseUseMountArgs(args, fi.Src, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}

func TestParseUseMountArgs_RouterNotIdent(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', 'literal');`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if m := parseUseMountArgs(args, fi.Src, nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}
