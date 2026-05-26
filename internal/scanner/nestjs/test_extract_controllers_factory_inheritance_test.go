//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractControllers_FactoryInheritance 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestExtractControllers_FactoryInheritance(t *testing.T) {
	dir := t.TempDir()

	baseSrc := `
export function BaseController(CreateDto, UpdateDto) {
  class GenericsController {
    @Get()
    findAll() {}
    @Get('paginate')
    paginate() {}
    @Get('find/:id')
    findOne() {}
    @Post()
    create() {}
    @Put(':id')
    update() {}
    @Patch('archive/:id')
    archive() {}
    @Patch('unarchive/:id')
    unarchive() {}
    @Delete(':id')
    delete() {}
    @Delete()
    clear() {}
    @Get('search')
    search() {}
  }
  return GenericsController;
}
`
	writeFile(t, dir, "src/base/base.controller.ts", baseSrc)

	childSrc := `
import { BaseController } from '../base/base.controller';

@Controller('categories')
export class CategoryController extends BaseController<Category, CategoryDto>(CreateCategoryDto, UpdateCategoryDto) {
}
`
	childFile := filepath.Join(dir, "src/category/category.controller.ts")
	writeFile(t, dir, "src/category/category.controller.ts", childSrc)

	root, err := parseTypeScript([]byte(childSrc))
	if err != nil {
		t.Fatal(err)
	}

	controllers := extractControllers(root, []byte(childSrc), "src/category/category.controller.ts", childFile)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	ci := controllers[0]
	if ci.prefix != "categories" {
		t.Fatalf("expected prefix 'categories', got %q", ci.prefix)
	}
	if len(ci.endpoints) != 10 {
		t.Fatalf("expected 10 endpoints, got %d", len(ci.endpoints))
	}
}
