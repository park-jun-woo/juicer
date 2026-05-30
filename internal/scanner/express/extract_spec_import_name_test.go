//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractSpecImportName: alias / name 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstImportSpec(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	specs := findAllByType(fi.Root, "import_specifier")
	if len(specs) == 0 {
		t.Fatal("no import_specifier")
	}
	return specs[0]
}

func TestExtractSpecImportName_Alias(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router as R } from 'express';`))
	if got := extractSpecImportName(firstImportSpec(t, fi), fi.Src); got != "R" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSpecImportName_Name(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router } from 'express';`))
	if got := extractSpecImportName(firstImportSpec(t, fi), fi.Src); got != "Router" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSpecImportName_IdentifierFallback(t *testing.T) {
	// node lacking name/alias fields but containing an identifier child:
	// an arguments node from a call falls back to the first identifier.
	fi := mustParse(t, []byte(`f(myArg);`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	if got := extractSpecImportName(args, fi.Src); got != "myArg" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSpecImportName_None(t *testing.T) {
	// node with neither name/alias fields nor identifier child -> ""
	fi := mustParse(t, []byte(`f("s");`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	if got := extractSpecImportName(args, fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
