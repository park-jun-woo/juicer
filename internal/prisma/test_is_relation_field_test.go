//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what isRelationField 판정 테이블 테스트 (관계 필드 제외 + FK 스칼라 컬럼 유지)
package prisma

import "testing"

func TestIsRelationField(t *testing.T) {
	models := map[string]bool{"User": true, "Post": true}

	cases := []struct {
		name string
		in   field
		want bool
	}{
		{"navigation field to declared model", field{name: "author", baseType: "User"}, true},
		{"nullable navigation field", field{name: "author", baseType: "User", nullable: true}, true},
		{"array navigation field", field{name: "posts", baseType: "Post", array: true}, true},
		{"fk scalar column kept", field{name: "authorId", baseType: "Int"}, false},
		{"plain scalar string", field{name: "title", baseType: "String"}, false},
		{"unsupported type not a relation", field{name: "embedding", baseType: `Unsupported("vector(768)")`}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isRelationField(c.in, models); got != c.want {
				t.Errorf("isRelationField(%+v) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}
