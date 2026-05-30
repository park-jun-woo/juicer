//ff:func feature=scan type=test control=sequence topic=actix
//ff:what receiverIsWebScopeOrResource — 수신자 헤드 판별을 검증
package actix

import "testing"

func TestReceiverIsWebScopeOrResource(t *testing.T) {
	cases := []struct {
		src  string
		suff string // method-call suffix of the receiver chain to test
		want bool
	}{
		{`fn f() { web::scope("/p").service(a); }`, ".service", true},
		{`fn f() { web::resource("/r").route(b); }`, ".route", true},
		{`fn f() { cfg.service(a); }`, ".service", false},
	}
	for _, c := range cases {
		root, err := parseRust([]byte(c.src))
		if err != nil {
			t.Fatal(err)
		}
		call := findCallByFuncSuffix(root, []byte(c.src), c.suff)
		if call == nil {
			t.Fatalf("call %q not found in %q", c.suff, c.src)
		}
		fe := findChildByType(call, "field_expression")
		recv := findFieldReceiver(fe)
		if got := receiverIsWebScopeOrResource(recv, []byte(c.src)); got != c.want {
			t.Errorf("src %q: got %v, want %v", c.src, got, c.want)
		}
	}
}
