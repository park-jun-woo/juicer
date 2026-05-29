//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 경로가 stat 가능한 파일로 존재하는지 보고한다
package express

import "os"

// statFile reports whether path exists as a stat-able file.
func statFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
