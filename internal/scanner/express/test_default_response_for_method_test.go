//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestDefaultResponseForMethod: res 호출 없는 핸들러의 메서드별 기본값 (POST→201, DELETE→204, 그 외 200) (Phase140)
package express

import "testing"

func TestDefaultResponseForMethod(t *testing.T) {
	tests := []struct {
		method     string
		wantStatus string
	}{
		{"POST", "201"},
		{"post", "201"},
		{"DELETE", "204"},
		{"delete", "204"},
		{"GET", "200"},
		{"PUT", "200"},
		{"", "200"},
	}
	for _, tt := range tests {
		resps := defaultResponseForMethod(tt.method)
		if len(resps) != 1 || resps[0].Status != tt.wantStatus || resps[0].Kind != "json" {
			t.Errorf("method %q: got %+v, want %s/json", tt.method, resps, tt.wantStatus)
		}
	}
}
