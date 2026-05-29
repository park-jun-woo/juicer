//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what urlpatterns RHS가 list/래퍼 call 모두에서 path()를 수집하는지 검증한다
package django

import "testing"

func TestCollectFromURLPatternsRHS(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want int
	}{
		{"list", "urlpatterns = [path('a/', v1), path('b/', v2)]\n", 2},
		{"wrapper", "urlpatterns = i18n_patterns(path('a/', v1))\n", 1},
	}
	for _, tc := range cases {
		root, err := parsePython([]byte(tc.src))
		if err != nil {
			t.Fatalf("%s: parse: %v", tc.name, err)
		}
		nodes := findAllByType(root, "assignment")
		if len(nodes) == 0 {
			t.Fatalf("%s: no assignment node", tc.name)
		}
		got := collectFromURLPatternsRHS(nodes[0], []byte(tc.src))
		if len(got) != tc.want {
			t.Errorf("%s: got %d entries, want %d", tc.name, len(got), tc.want)
		}
	}
}
