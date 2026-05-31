//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveEnumPathArg_CrossFile 테스트 (import된 enum 멤버 경로 해석)
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestResolveEnumPathArg_CrossFile(t *testing.T) {
	dir := t.TempDir()

	enumSrc := `export enum RouteKey { Asset = 'assets', User = 'users' }`
	writeFile(t, dir, "src/enum.ts", enumSrc)

	ctrlSrc := `
import { RouteKey } from '../enum';

@Controller(RouteKey.Asset)
export class AssetController {
  @Get()
  findAll() {}
}
`
	ctrlFile := filepath.Join(dir, "src/asset/asset.controller.ts")
	writeFile(t, dir, "src/asset/asset.controller.ts", ctrlSrc)

	root, err := parseTypeScript([]byte(ctrlSrc))
	if err != nil {
		t.Fatal(err)
	}
	controllers := extractControllers(root, []byte(ctrlSrc), "src/asset/asset.controller.ts", ctrlFile, dir)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	if controllers[0].prefix != "assets" {
		t.Fatalf("cross-file prefix: want %q got %q", "assets", controllers[0].prefix)
	}
}
