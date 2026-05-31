//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestCollectParamRoutersFromUsage_PathGuard: path 형태 가드로 비-path 첫 인자 라우터 후보 탈락 검증
package express

import "testing"

func TestCollectParamRoutersFromUsage_PathGuard(t *testing.T) {
	// 첫 문자열 인자가 path 형태(`/` 시작 또는 `*`)일 때만 라우터로 등록한다.
	tests := []struct {
		src      string
		obj      string
		register bool
	}{
		{`req.get('user-agent');`, "req", false},
		{`config.get('urls');`, "config", false},
		{`model.get('status');`, "model", false},
		{`r.get('/x', h);`, "r", true},
		{`app.get('*', h);`, "app", true},
	}
	for _, tt := range tests {
		fi := mustParse(t, []byte(tt.src))
		routers := map[string]bool{}
		collectParamRoutersFromUsage(fi, routers)
		if routers[tt.obj] != tt.register {
			t.Errorf("%s: routers[%q]=%v, want %v", tt.src, tt.obj, routers[tt.obj], tt.register)
		}
	}
}
