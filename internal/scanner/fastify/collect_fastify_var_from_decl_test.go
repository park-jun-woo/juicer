//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what collectFastifyVarFromDecl 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func declOfType(t *testing.T, src string) ([]*sitter.Node, *fileInfo) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	return findAllByType(fi.Root, "lexical_declaration"), fi
}

func TestCollectFastifyVarFromDecl_Match(t *testing.T) {
	decls, fi := declOfType(t, "const app = Fastify();\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if !instances["app"] {
		t.Fatalf("expected 'app' instance, got %v", instances)
	}
}

func TestCollectFastifyVarFromDecl_LowercaseMatch(t *testing.T) {
	decls, fi := declOfType(t, "const srv = fastify();\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if !instances["srv"] {
		t.Fatalf("expected 'srv' instance, got %v", instances)
	}
}

func TestCollectFastifyVarFromDecl_NonInstances(t *testing.T) {
	// no init call, member call, and non-Fastify call: none collected
	decls, fi := declOfType(t, "const a = 5;\nconst b = obj.create();\nconst c = Express();\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if len(instances) != 0 {
		t.Fatalf("expected no instances, got %v", instances)
	}
}

func TestCollectFastifyVarFromDecl_MultipleDeclarators(t *testing.T) {
	// one declaration with two declarators, only the Fastify one collected
	decls, fi := declOfType(t, "const app = Fastify(), other = 5;\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if !instances["app"] || instances["other"] {
		t.Fatalf("expected only app, got %v", instances)
	}
}
