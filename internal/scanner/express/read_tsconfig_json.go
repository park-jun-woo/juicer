//ff:func feature=scan type=parse control=iteration dimension=1 topic=express
//ff:what tsconfig.json 또는 tsconfig.app.json 파일을 읽어 바이트를 반환한다
package express

import (
	"os"
	"path/filepath"
)

func readTsconfigJSON(absRoot string) []byte {
	candidates := []string{"tsconfig.json", "tsconfig.app.json"}
	for _, name := range candidates {
		data, err := os.ReadFile(filepath.Join(absRoot, name))
		if err == nil {
			return data
		}
	}
	return nil
}
