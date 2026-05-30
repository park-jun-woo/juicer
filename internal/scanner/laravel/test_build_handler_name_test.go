//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestBuildHandlerName 테스트
package laravel

import "testing"

func TestBuildHandlerName(t *testing.T) {
	cases := []struct{ ctrl, action, want string }{
		{"", "", "closure"},
		{"", "show", "show"},
		{"UserController", "", "UserController"},
		{"UserController", "show", "UserController@show"},
	}
	for _, c := range cases {
		if got := buildHandlerName(c.ctrl, c.action); got != c.want {
			t.Errorf("buildHandlerName(%q,%q)=%q want %q", c.ctrl, c.action, got, c.want)
		}
	}
}
