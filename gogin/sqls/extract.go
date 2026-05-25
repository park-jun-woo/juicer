//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 디렉토리에서 *_repo.go 파일을 파싱하여 SQL 스켈레톤 추출
package sqls

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Extract scans dir for *_repo.go files and returns SQL skeletons.
//
func Extract(dir string) (*SkeletonResult, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %w", dir, err)
	}

	var methods []MethodSkeleton

	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, "_repo.go") {
			continue
		}
		if strings.HasSuffix(name, "_test.go") || strings.HasSuffix(name, "_iface.go") {
			continue
		}

		path := filepath.Join(dir, name)
		ms, err := parseFile(path)
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", name, err)
		}
		methods = append(methods, ms...)
	}

	sort.Slice(methods, func(i, j int) bool {
		if methods[i].Repo != methods[j].Repo {
			return methods[i].Repo < methods[j].Repo
		}
		return methods[i].Method < methods[j].Method
	})

	return &SkeletonResult{Methods: methods}, nil
}

