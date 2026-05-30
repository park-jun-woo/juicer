//ff:func feature=scan type=test control=sequence
//ff:what scanJoin 테스트 헬퍼
package gogin

import "path/filepath"

func scanJoin(dir, rel string) string { return filepath.Join(dir, rel) }
