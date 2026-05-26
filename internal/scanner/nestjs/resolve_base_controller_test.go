//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveBaseController_FactoryInheritance 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestResolveBaseController_FactoryInheritance(t *testing.T) {
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
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes found")
	}
	imports := extractImports(root, []byte(childSrc))
	eps := resolveBaseController(classes[0], []byte(childSrc), childFile, imports, "src/category/category.controller.ts")
	if len(eps) != 10 {
		t.Fatalf("expected 10 inherited endpoints, got %d", len(eps))
	}
}
