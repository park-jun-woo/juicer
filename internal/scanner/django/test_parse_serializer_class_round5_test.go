//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseSerializerClass_Round5 테스트
package django

import "testing"

func TestParseSerializerClass_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "class UserSerializer(serializers.Serializer):\n    name = serializers.CharField(max_length=10)\n")
	cls := djFirst(t, fi.root, "class_definition")
	ser := parseSerializerClass(cls, fi)
	if ser == nil || ser.name != "UserSerializer" {
		t.Fatalf("serializer: %+v", ser)
	}
}
