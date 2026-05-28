//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestFieldTypeToScannerType_ArrayObject — List<Photo>가 []Photo를 반환한다
package spring

import "testing"

func TestFieldTypeToScannerType_ArrayObject(t *testing.T) {
	tests := []struct {
		jtype string
		want  string
	}{
		{"List<Photo>", "[]Photo"},
		{"List<String>", "array:string"},
		{"List<Integer>", "array:integer"},
		{"Set<Tag>", "[]Tag"},
		{"ArrayList<Comment>", "[]Comment"},
	}
	for _, tt := range tests {
		t.Run(tt.jtype, func(t *testing.T) {
			got := fieldTypeToScannerType(tt.jtype)
			if got != tt.want {
				t.Errorf("fieldTypeToScannerType(%q) = %q, want %q", tt.jtype, got, tt.want)
			}
		})
	}
}
