//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectRouterAliasesFromStmt: Router/별칭 import 추출 + 미해당 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstImportStmt(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	stmts := findAllByType(fi.Root, "import_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_statement")
	}
	return stmts[0]
}

func TestCollectRouterAliasesFromStmt_Plain(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router } from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if !aliases["Router"] {
		t.Fatalf("expected Router, got %v", aliases)
	}
}

func TestCollectRouterAliasesFromStmt_Alias(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router as R } from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if !aliases["R"] {
		t.Fatalf("expected alias R, got %v", aliases)
	}
}

func TestCollectRouterAliasesFromStmt_NonRouter(t *testing.T) {
	fi := mustParse(t, []byte(`import { Foo } from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if len(aliases) != 0 {
		t.Fatalf("expected none, got %v", aliases)
	}
}

func TestCollectRouterAliasesFromStmt_NoNamedImports(t *testing.T) {
	// default import only -> import_clause present but no named_imports
	fi := mustParse(t, []byte(`import express from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if len(aliases) != 0 {
		t.Fatalf("expected none, got %v", aliases)
	}
}

func TestCollectRouterAliasesFromStmt_NoClause(t *testing.T) {
	// side-effect import -> no import_clause
	fi := mustParse(t, []byte(`import 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if len(aliases) != 0 {
		t.Fatalf("expected none, got %v", aliases)
	}
}
