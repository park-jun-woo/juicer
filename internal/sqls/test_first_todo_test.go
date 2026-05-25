//ff:func feature=ratchet type=session control=iteration dimension=1
//ff:what TestFirstTODO 테스트
package sqls

import (
	"testing"
)

func TestFirstTODO(t *testing.T) {
	tests := []struct {
		name    string
		methods []MethodStatus
		want    int
	}{
		{"found", []MethodStatus{{Status: "DONE"}, {Status: "TODO"}}, 1},
		{"not found", []MethodStatus{{Status: "DONE"}}, -1},
		{"empty", []MethodStatus{}, -1},
		{"first", []MethodStatus{{Status: "TODO"}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sess := &Session{Methods: tt.methods}
			got := firstTODO(sess)
			if got != tt.want {
				t.Errorf("firstTODO() = %d, want %d", got, tt.want)
			}
		})
	}
}
