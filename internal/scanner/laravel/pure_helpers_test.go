//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what applyValidationRule / buildHandlerName / buildResourcePath / childrenOfType / copyMiddleware / classMatches / classNameMatches / collectionResourceName 테스트
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyValidationRule(t *testing.T) {
	// type rule path
	f := &scanner.Field{}
	num := false
	applyValidationRule(f, "integer", &num)
	if f.Type != "integer" || !num {
		t.Fatalf("type: %+v num=%v", f, num)
	}
	// format rule path
	f2 := &scanner.Field{}
	num2 := false
	applyValidationRule(f2, "email", &num2)
	if f2.Type != "string" {
		t.Fatalf("format: %+v", f2)
	}
	// flag rule path
	f3 := &scanner.Field{}
	applyValidationRule(f3, "nullable", &num2)
	if !f3.Nullable {
		t.Fatalf("flag: %+v", f3)
	}
	// constraint path (fallthrough)
	f4 := &scanner.Field{}
	applyValidationRule(f4, "max:5", &num2)
	if f4.MaxLength == nil {
		t.Fatalf("constraint: %+v", f4)
	}
}

func TestBuildHandlerName(t *testing.T) {
	cases := []struct{ ctrl, action, want string }{
		{"", "", "closure"},
		{"", "show", "show"},
		{"UserController", "", "UserController"},
		{"UserController", "show", "UserController@show"},
	}
	for _, c := range cases {
		if got := buildHandlerName(c.ctrl, c.action); got != c.want {
			t.Errorf("buildHandlerName(%q,%q)=%q want %q", c.ctrl, c.action, got, c.want)
		}
	}
}

func TestBuildResourcePath(t *testing.T) {
	p, param := buildResourcePath("users")
	if p != "users" || param != "user" {
		t.Fatalf("single: %q %q", p, param)
	}
	p2, param2 := buildResourcePath("users.posts")
	if p2 != "users/{user}/posts" {
		t.Fatalf("nested path: %q", p2)
	}
	if param2 != "post" {
		t.Fatalf("nested param: %q", param2)
	}
}

func TestChildrenOfType(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['a' => 1, 'b' => 2];`)
	arrs := findAllByType(fi.root, "array_creation_expression")
	if len(arrs) == 0 {
		t.Fatal("no array")
	}
	elems := childrenOfType(arrs[0], "array_element_initializer")
	if len(elems) != 2 {
		t.Fatalf("expected 2, got %d", len(elems))
	}
}

func TestCopyMiddleware(t *testing.T) {
	orig := []string{"auth", "throttle"}
	cp := copyMiddleware(orig)
	if len(cp) != 2 || cp[0] != "auth" {
		t.Fatalf("got %v", cp)
	}
	cp[0] = "changed"
	if orig[0] != "auth" {
		t.Fatal("copy shares backing array")
	}
}

func TestCopyMiddleware_Empty(t *testing.T) {
	if cp := copyMiddleware(nil); len(cp) != 0 {
		t.Fatalf("got %v", cp)
	}
}

func TestClassMatches(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserController {}`)
	if !classMatches(&fi, "UserController") {
		t.Fatal("expected match")
	}
	if classMatches(&fi, "Other") {
		t.Fatal("unexpected match")
	}
}

func TestClassNameMatches(t *testing.T) {
	fi := mustParsePHP(t, `<?php class Foo {}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	if !classNameMatches(cls, fi.src, "Foo") {
		t.Fatal("expected match Foo")
	}
	if classNameMatches(cls, fi.src, "Bar") {
		t.Fatal("unexpected match Bar")
	}
}

func TestCollectionResourceName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = UserResource::collection($users);`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if got := collectionResourceName(scoped[0], fi.src); got != "UserResource" {
		t.Fatalf("got %q", got)
	}
}

func TestCollectionResourceName_NotCollection(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = UserResource::make($user);`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if got := collectionResourceName(scoped[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestCollectionResourceName_NotResourceSuffix(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = Foo::collection($users);`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if got := collectionResourceName(scoped[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
