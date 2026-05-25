//ff:func feature=sql type=parse control=sequence
//ff:what cwd 기준 .huma 디렉토리 경로 반환
package sqls

// SessionDir returns the .huma directory path under cwd.
//
func SessionDir() string {
	return sessionDirName
}

