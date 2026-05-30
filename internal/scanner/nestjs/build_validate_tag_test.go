//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what buildValidateTag 테스트
package nestjs

import "testing"

func TestBuildValidateTag(t *testing.T) {
	cases := []struct {
		optional   bool
		validators []string
		want       string
	}{
		{false, []string{"email"}, "required,email"},
		{false, nil, "required"},
		{true, []string{"email", "min:1"}, "email,min:1"},
		{true, nil, ""},
	}
	for _, c := range cases {
		if got := buildValidateTag(c.optional, c.validators); got != c.want {
			t.Errorf("buildValidateTag(%v,%v)=%q want %q", c.optional, c.validators, got, c.want)
		}
	}
}
