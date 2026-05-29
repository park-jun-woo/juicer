//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what schema.prisma 파일/디렉터리를 파싱해 테이블맵+enum목록 반환 (엔트리포인트)
package prisma

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

// Parse reads a schema.prisma file (or a directory of *.prisma files),
// converts all models, and returns the resulting table state plus the enum
// types declared in the schema.
func Parse(path string) (map[string]*ddl.Table, []ddl.EnumType, error) {
	files, err := collectPrismaFiles(path)
	if err != nil {
		return nil, nil, err
	}
	var models []model
	var enums []ddl.EnumType
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, nil, fmt.Errorf("read %s: %w", f, err)
		}
		src := string(data)
		models = append(models, parseSource(src)...)
		enums = append(enums, parseEnums(src)...)
	}
	// convertModels consumes raw enum names for schema enum-set lookups; the
	// returned enum list quotes names so CREATE TYPE matches the column type.
	tables := convertModels(models, enums)
	return tables, quoteEnumNames(enums), nil
}
