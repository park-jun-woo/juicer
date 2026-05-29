//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what schema.prisma 파일/디렉터리를 파싱해 테이블명->*ddl.Table 맵 반환 (엔트리포인트)
package prisma

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

// Parse reads a schema.prisma file (or a directory of *.prisma files),
// converts all models, and returns the resulting table state.
func Parse(path string) (map[string]*ddl.Table, error) {
	files, err := collectPrismaFiles(path)
	if err != nil {
		return nil, err
	}
	var models []model
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", f, err)
		}
		models = append(models, parseSource(string(data))...)
	}
	return convertModels(models), nil
}
