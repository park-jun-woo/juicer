//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestLaravelRulesToField 테스트
package laravel

import "testing"

func TestLaravelRulesToField(t *testing.T) {
	f := laravelRulesToField("email", []string{"required", "email"})
	if f.Name != "email" || f.Type != "string" {
		t.Fatalf("got %+v", f)
	}
	f2 := laravelRulesToField("age", []string{"integer", "min:0"})
	if f2.Type != "integer" {
		t.Fatalf("got %+v", f2)
	}
}
