//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractSerializersFromFile 테스트
package django

import "testing"

func TestExtractSerializersFromFile(t *testing.T) {
	src := `
class UserSerializer(serializers.ModelSerializer):
    name = serializers.CharField()

class Plain:
    pass
`
	fi := newTestFileInfo(t, src)
	ss := extractSerializersFromFile(fi)
	if len(ss) != 1 {
		t.Fatalf("expected 1 serializer (Plain skipped), got %d", len(ss))
	}
	if ss[0].name != "UserSerializer" {
		t.Errorf("name = %q, want UserSerializer", ss[0].name)
	}
}
