//ff:func feature=scan type=parse control=iteration dimension=1 topic=express
//ff:what tsconfig.json에서 compilerOptions.paths와 baseUrl을 읽어 alias 맵을 반환한다
package express

import (
	"encoding/json"
	"path/filepath"
	"strings"
)

func loadTsconfigPaths(absRoot string) map[string]string {
	aliases := make(map[string]string)
	raw := readTsconfigJSON(absRoot)
	if raw == nil {
		return aliases
	}
	var cfg struct {
		CompilerOptions struct {
			BaseUrl string              `json:"baseUrl"`
			Paths   map[string][]string `json:"paths"`
		} `json:"compilerOptions"`
	}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return aliases
	}
	baseUrl := cfg.CompilerOptions.BaseUrl
	for pattern, targets := range cfg.CompilerOptions.Paths {
		if len(targets) == 0 {
			continue
		}
		prefix := strings.TrimSuffix(pattern, "*")
		replacement := strings.TrimSuffix(targets[0], "*")
		if baseUrl != "" {
			replacement = filepath.Join(baseUrl, replacement) + string(filepath.Separator)
		}
		aliases[prefix] = replacement
	}
	return aliases
}
