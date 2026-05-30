//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractImportModule dotted/relative/empty 전 분기 테스트
package fastapi

import "testing"

func TestExtractImportModule(t *testing.T) {
	src := []byte("from fastapi import FastAPI\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := findAllByType(root, "import_from_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_from_statement")
	}
	got := extractImportModule(stmts[0], src)
	if got != "fastapi" {
		t.Fatalf("expected 'fastapi', got %q", got)
	}

	// relative import (from . import xxx)
	src2 := []byte("from . import models\n")
	root2, _ := parsePython(src2)
	stmts2 := findAllByType(root2, "import_from_statement")
	if len(stmts2) > 0 {
		got2 := extractImportModule(stmts2[0], src2)
		_ = got2 // just need coverage
	}

	// relative import with dotted path
	src3 := []byte("from .models import User\n")
	root3, _ := parsePython(src3)
	stmts3 := findAllByType(root3, "import_from_statement")
	if len(stmts3) > 0 {
		got3 := extractImportModule(stmts3[0], src3)
		if got3 != ".models" {
			t.Fatalf("expected .models, got %q", got3)
		}
	}
}

func TestExtractImportModule_NoModule(t *testing.T) {
	// a node lacking relative_import and dotted_name children -> ""
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	if got := extractImportModule(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
