//ff:func feature=scan type=test control=sequence topic=django
//ff:what round5 미커버 함수 직접 호출 테스트 (django)
package django

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func djFirst(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if found == nil && n.Type() == typ {
			found = n
		}
	})
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}

func TestUnquotePython_Round5(t *testing.T) {
	cases := map[string]string{
		`"x"`:       "x",
		`'y'`:       "y",
		`f"z"`:      "z",
		`r'p'`:      "p",
		`"""abc"""`: "abc",
		`a`:         "a",
		``:          "",
	}
	for in, want := range cases {
		if got := unquotePython(in); got != want {
			t.Errorf("unquotePython(%q)=%q want %q", in, got, want)
		}
	}
}

func TestWalkNodes_Round5(t *testing.T) {
	root, _ := parsePython([]byte("x = 1\ny = 2\n"))
	n := 0
	walkNodes(root, func(*sitter.Node) { n++ })
	if n < 3 {
		t.Fatalf("too few nodes: %d", n)
	}
}

func TestResolveViewName_Round5(t *testing.T) {
	// pure function: returns name unchanged or resolved
	got := resolveViewName("MyView")
	if got == "" {
		t.Fatal("expected non-empty resolution")
	}
}

func TestResolveViewSetMethods_Round5(t *testing.T) {
	methods := resolveViewSetMethods([]string{"ModelViewSet"})
	if len(methods) == 0 {
		t.Fatal("ModelViewSet should expand to CRUD methods")
	}
}

func TestResolveModelFields_Round5(t *testing.T) {
	// no-op / cache priming function; must not panic
	resolveModelFields()
}

func TestParsePathCall_Round5(t *testing.T) {
	src := []byte("urlpatterns = [path('users/', UserView.as_view())]")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := djFirst(t, root, "call")
	entry := parsePathCall(call, src)
	if entry == nil {
		t.Fatal("expected url entry")
	}
	if entry.pattern != "users/" {
		t.Errorf("pattern: %q", entry.pattern)
	}
	if entry.viewName != "UserView" {
		t.Errorf("viewName: %q", entry.viewName)
	}
}

func TestParsePathCallsInList_Round5(t *testing.T) {
	src := []byte("urlpatterns = [path('a/', AView.as_view()), path('b/', BView.as_view())]")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	list := djFirst(t, root, "list")
	entries := parsePathCallsInList(list, src)
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
}

func TestResolveCallArg_Round5(t *testing.T) {
	// include('app.urls') second arg
	src := []byte("path('api/', include('app.urls'))\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)
	var e urlEntry
	resolveCallArg(&e, pos[1], src)
	if !e.isInclude || e.includeModule != "app.urls" {
		t.Fatalf("include: %+v", e)
	}

	// MyView.as_view() second arg
	src2 := []byte("path('x/', MyView.as_view())\n")
	root2, _ := parsePython(src2)
	args2 := djFirst(t, root2, "argument_list")
	pos2 := positionalArgs(args2)
	var e2 urlEntry
	resolveCallArg(&e2, pos2[1], src2)
	if e2.viewName != "MyView" {
		t.Fatalf("viewName: %q", e2.viewName)
	}
}

func TestParseApiViewClass_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "class PingView(APIView):\n    def get(self, request):\n        return Response()\n")
	cls := djFirst(t, fi.root, "class_definition")
	av := parseAPIViewClass(cls, fi)
	if av == nil || av.name != "PingView" {
		t.Fatalf("apiview: %+v", av)
	}
}

func TestParseViewSetClass_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "class UserViewSet(ModelViewSet):\n    queryset = User.objects.all()\n")
	cls := djFirst(t, fi.root, "class_definition")
	vs := parseViewSetClass(cls, fi)
	if vs == nil || vs.name != "UserViewSet" {
		t.Fatalf("viewset: %+v", vs)
	}
}

func TestParseSerializerClass_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "class UserSerializer(serializers.Serializer):\n    name = serializers.CharField(max_length=10)\n")
	cls := djFirst(t, fi.root, "class_definition")
	ser := parseSerializerClass(cls, fi)
	if ser == nil || ser.name != "UserSerializer" {
		t.Fatalf("serializer: %+v", ser)
	}
}

func TestParseSerializerFieldAssignment_Round5(t *testing.T) {
	src := []byte("name = serializers.CharField(max_length=10)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := djFirst(t, root, "expression_statement")
	f := parseSerializerFieldAssignment(stmt, src)
	if f == nil || f.Name != "name" {
		t.Fatalf("field: %+v", f)
	}
}

func TestParseFuncView_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "@api_view(['GET'])\ndef ping(request):\n    return Response()\n")
	fn := djFirst(t, fi.root, "function_definition")
	fv := parseFuncView(fn, fi)
	if fv == nil {
		t.Fatal("expected func view")
	}
	if fv.name != "ping" {
		t.Errorf("name: %q", fv.name)
	}
}

func TestExtractApiViewDecorator_Round5(t *testing.T) {
	src := []byte("@api_view(['GET', 'POST'])\ndef v(request):\n    return Response()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := djFirst(t, root, "function_definition")
	methods := extractAPIViewDecorator(fn, src)
	if len(methods) != 2 {
		t.Fatalf("expected 2 methods, got %v", methods)
	}
}

func TestParseDecoratedAction_Round5(t *testing.T) {
	src := []byte("@action(detail=True, methods=['post'])\ndef activate(self, request, pk=None):\n    return Response()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	dec := djFirst(t, root, "decorated_definition")
	action := parseDecoratedAction(dec, src)
	if action == nil {
		t.Fatal("expected action")
	}
}

func TestParseRegisterCall_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "router.register('users', UserViewSet)\n")
	call := djFirst(t, fi.root, "call")
	reg := parseRegisterCall(call, fi)
	if reg == nil {
		t.Fatal("expected registration")
	}
}

func TestPositionalArgs_Round5(t *testing.T) {
	src := []byte("path('a/', AView.as_view(), name='a')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)
	// two positional ('a/', AView.as_view()), keyword excluded
	if len(pos) != 2 {
		t.Fatalf("expected 2 positional args, got %d", len(pos))
	}
}

func TestResolveSecondArg_And_CallArg_Round5(t *testing.T) {
	// include(...) second arg
	src := []byte("path('api/', include('app.urls'))\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)
	if len(pos) < 2 {
		t.Fatalf("expected 2 args, got %d", len(pos))
	}
	var e urlEntry
	resolveSecondArg(&e, pos[1], src)
	if !e.isInclude || e.includeModule != "app.urls" {
		t.Fatalf("include: %+v", e)
	}

	// as_view() second arg
	src2 := []byte("path('x/', MyView.as_view())\n")
	root2, _ := parsePython(src2)
	args2 := djFirst(t, root2, "argument_list")
	pos2 := positionalArgs(args2)
	var e2 urlEntry
	resolveSecondArg(&e2, pos2[1], src2)
	if e2.viewName != "MyView" {
		t.Fatalf("viewName: %q", e2.viewName)
	}
}

func TestApplyOneConstraint_Round5(t *testing.T) {
	src := []byte("name = serializers.CharField(max_length=12)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	// the keyword_argument node holds the integer child
	kw := djFirst(t, root, "keyword_argument")
	f := &scanner.Field{}
	applyOneConstraint(f, "max_length", kw, src)
	if f.MaxLength == nil || *f.MaxLength != 12 {
		t.Fatalf("max_length: %v", f.MaxLength)
	}
}

func TestBuildAPIViewEndpoints_Round5(t *testing.T) {
	entry := urlEntry{pattern: "ping/<int:pk>/", viewName: "PingView"}
	av := &apiviewInfo{name: "PingView", methods: []string{"GET", "POST"}, file: "views.py", line: 1}
	eps := buildAPIViewEndpoints(entry, av, map[string]serializerInfo{})
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	if eps[0].Handler != "PingView.get" {
		t.Errorf("handler: %q", eps[0].Handler)
	}
}

func TestParseFile_And_BuildAllEndpoints_Round5(t *testing.T) {
	dir := t.TempDir()
	urls := "from django.urls import path\nfrom .views import PingView\nurlpatterns = [path('ping/', PingView.as_view())]\n"
	views := "from rest_framework.views import APIView\nfrom rest_framework.response import Response\nclass PingView(APIView):\n    def get(self, request):\n        return Response()\n"
	if err := os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urls), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "views.py"), []byte(views), 0o644); err != nil {
		t.Fatal(err)
	}
	fiUrls, err := parseFile(dir, filepath.Join(dir, "urls.py"))
	if err != nil {
		t.Fatal(err)
	}
	if fiUrls.root == nil {
		t.Fatal("parseFile root nil")
	}
	fiViews, err := parseFile(dir, filepath.Join(dir, "views.py"))
	if err != nil {
		t.Fatal(err)
	}
	eps := buildAllEndpoints([]fileInfo{*fiUrls, *fiViews})
	if len(eps) == 0 {
		t.Fatalf("expected endpoints, got %d", len(eps))
	}
}
