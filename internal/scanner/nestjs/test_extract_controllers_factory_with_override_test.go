//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractControllers_FactoryWithOverride 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestExtractControllers_FactoryWithOverride(t *testing.T) {
	dir := t.TempDir()

	baseSrc := `
export function BaseController(CreateDto, UpdateDto) {
  class GenericsController {
    @Get()
    findAll() {}
    @Post()
    create() {}
  }
  return GenericsController;
}
`
	writeFile(t, dir, "src/base/base.controller.ts", baseSrc)

	childSrc := `
import { BaseController } from '../base/base.controller';

@Controller('custom')
export class CustomController extends BaseController<Item, ItemDto>(CreateItemDto, UpdateItemDto) {
  @Get()
  findAll() {}
}
`
	childFile := filepath.Join(dir, "src/custom/custom.controller.ts")
	writeFile(t, dir, "src/custom/custom.controller.ts", childSrc)

	root, err := parseTypeScript([]byte(childSrc))
	if err != nil {
		t.Fatal(err)
	}

	controllers := extractControllers(root, []byte(childSrc), "src/custom/custom.controller.ts", childFile)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	ci := controllers[0]
	if len(ci.endpoints) != 2 {
		t.Fatalf("expected 2 endpoints (override), got %d", len(ci.endpoints))
	}
}
