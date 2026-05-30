//ff:func feature=scan type=test control=sequence topic=flask
//ff:what tryParseBlueprintAssignment 테스트
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstAssignment(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	assigns := findAllByType(root, "assignment")
	if len(assigns) == 0 {
		t.Fatal("no assignment")
	}
	return assigns[0], b
}

func TestTryParseBlueprintAssignment_Match(t *testing.T) {
	assign, b := firstAssignment(t, `api = Blueprint("api", __name__, url_prefix="/api")`+"\n")
	bp := tryParseBlueprintAssignment(assign, b)
	if bp == nil || bp.varName != "api" || bp.name != "api" || bp.urlPrefix != "/api" {
		t.Fatalf("blueprint = %+v", bp)
	}
}

func TestTryParseBlueprintAssignment_NotCall(t *testing.T) {
	assign, b := firstAssignment(t, "x = 5\n")
	if bp := tryParseBlueprintAssignment(assign, b); bp != nil {
		t.Fatalf("non-call should be nil, got %+v", bp)
	}
}

func TestTryParseBlueprintAssignment_NoLeftIdentifier(t *testing.T) {
	// attribute-target assignment has no identifier as a direct LHS child
	assign, b := firstAssignment(t, `obj.attr = Blueprint("x", __name__)`+"\n")
	if bp := tryParseBlueprintAssignment(assign, b); bp != nil {
		t.Fatalf("attribute LHS should be nil, got %+v", bp)
	}
}

func TestTryParseBlueprintAssignment_NoUrlPrefix(t *testing.T) {
	assign, b := firstAssignment(t, `bp = Blueprint("name", __name__)`+"\n")
	got := tryParseBlueprintAssignment(assign, b)
	if got == nil || got.urlPrefix != "" {
		t.Fatalf("expected empty url_prefix, got %+v", got)
	}
}

func TestTryParseBlueprintAssignment_NotBlueprint(t *testing.T) {
	assign, b := firstAssignment(t, `app = Flask(__name__)`+"\n")
	if bp := tryParseBlueprintAssignment(assign, b); bp != nil {
		t.Fatalf("Flask() should be nil, got %+v", bp)
	}
}
