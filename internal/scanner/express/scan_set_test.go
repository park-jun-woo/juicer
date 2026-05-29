//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what scanSet — 디렉터리를 스캔해 엔드포인트 키 집합을 만든다
package express

import "testing"

func scanSet(t *testing.T, dir string) map[string]bool {
	t.Helper()
	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	var eps []endpointLike
	for _, e := range result.Endpoints {
		eps = append(eps, endpointLike{e.Method, e.Path})
	}
	return epSet(eps)
}
