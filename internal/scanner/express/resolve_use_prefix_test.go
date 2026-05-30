//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveUsePrefixes: 마운트 수집 / 비마운트 스킵
package express

import "testing"

func TestResolveUsePrefixes_Collects(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', userRouter); doStuff();`))
	mounts := resolveUsePrefixes(fi, map[string]bool{"app": true}, map[string]string{"userRouter": "./u.ts"})
	if len(mounts) != 1 || mounts[0].Prefix != "/api" || mounts[0].VarName != "userRouter" {
		t.Fatalf("got %+v", mounts)
	}
}

func TestResolveUsePrefixes_None(t *testing.T) {
	fi := mustParse(t, []byte(`doStuff();`))
	mounts := resolveUsePrefixes(fi, map[string]bool{"app": true}, nil)
	if len(mounts) != 0 {
		t.Fatalf("expected none, got %+v", mounts)
	}
}
