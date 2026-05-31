//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractControllers_FactoryWithDirectMethods 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestExtractControllers_FactoryWithDirectMethods(t *testing.T) {
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

@Controller('products')
export class ProductController extends BaseController<Product, ProductDto>(CreateProductDto, UpdateProductDto) {
  @Get('expensive')
  findExpensive() {}
  @Get('search/name')
  searchByName() {}
}
`
	childFile := filepath.Join(dir, "src/product/product.controller.ts")
	writeFile(t, dir, "src/product/product.controller.ts", childSrc)

	root, err := parseTypeScript([]byte(childSrc))
	if err != nil {
		t.Fatal(err)
	}

	controllers := extractControllers(root, []byte(childSrc), "src/product/product.controller.ts", childFile, "/tmp")
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	ci := controllers[0]
	if len(ci.endpoints) != 4 {
		t.Fatalf("expected 4 endpoints (2 inherited + 2 direct), got %d", len(ci.endpoints))
	}
}
