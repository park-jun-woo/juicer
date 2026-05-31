//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveEnumPathArg_SameFile 테스트 (@Controller(Enum.Member) → 값)
package nestjs

import "testing"

func TestResolveEnumPathArg_SameFile(t *testing.T) {
	src := []byte(`
enum RouteKey { Asset = 'assets', User = 'users' }

@Controller(RouteKey.Asset)
export class AssetController {
  @Get(RouteKey.User)
  findOne() {}
}
`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	ci, ok := buildControllerInfo(cls, src, "asset.controller.ts", "/abs/asset.controller.ts", map[string]string{}, root, "/tmp")
	if !ok {
		t.Fatal("expected ok")
	}
	if ci.prefix != "assets" {
		t.Fatalf("prefix: want %q got %q", "assets", ci.prefix)
	}
	if len(ci.endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(ci.endpoints))
	}
	if ci.endpoints[0].path != "users" {
		t.Fatalf("method path: want %q got %q", "users", ci.endpoints[0].path)
	}
}
