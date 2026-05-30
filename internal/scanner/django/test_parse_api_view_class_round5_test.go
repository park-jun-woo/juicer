//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseApiViewClass_Round5 테스트
package django

import "testing"

func TestParseApiViewClass_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "class PingView(APIView):\n    def get(self, request):\n        return Response()\n")
	cls := djFirst(t, fi.root, "class_definition")
	av := parseAPIViewClass(cls, fi)
	if av == nil || av.name != "PingView" {
		t.Fatalf("apiview: %+v", av)
	}
}
