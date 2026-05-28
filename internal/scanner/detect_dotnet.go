//ff:func feature=scan type=extract control=sequence
//ff:what *.csproj에서 Microsoft.AspNetCore 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectDotnet(root string) bool {
	found := false
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || found {
			return err
		}
		if info.IsDir() {
			name := info.Name()
			if name == "bin" || name == "obj" || name == ".git" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".csproj") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(data)
		if strings.Contains(content, "Microsoft.AspNetCore") ||
			strings.Contains(content, "Microsoft.NET.Sdk.Web") {
			found = true
		}
		return nil
	})
	return found
}
