//ff:func feature=sql type=parse control=sequence
//ff:what cwd 기준 .codist 디렉토리 경로 반환
package sqls

// SessionDir returns the .codist directory path under cwd.
//
func SessionDir() string {
	return sessionDirName
}

