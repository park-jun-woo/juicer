//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what applyArrayOption 테스트
package laravel

import "testing"

func TestApplyArrayOption_Prefix(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['prefix' => 'v1'];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	if len(elems) == 0 {
		t.Fatal("no array element")
	}
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if prefix != "/v1" && prefix != "v1" {
		t.Fatalf("prefix got %q", prefix)
	}
}

func TestApplyArrayOption_Middleware(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['middleware' => 'auth'];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if len(mw) == 0 {
		t.Fatalf("middleware not applied: %v", mw)
	}
}

func TestApplyArrayOption_UnknownKey(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['as' => 'name'];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if prefix != "" || len(mw) != 0 {
		t.Fatalf("unexpected change: prefix=%q mw=%v", prefix, mw)
	}
}

func TestApplyArrayOption_NoStringKey(t *testing.T) {
	// numeric-keyed / list element without leading string key
	fi := mustParsePHP(t, `<?php $x = [42];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	if len(elems) == 0 {
		t.Skip("no element")
	}
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if prefix != "" || len(mw) != 0 {
		t.Fatalf("unexpected change: prefix=%q mw=%v", prefix, mw)
	}
}
