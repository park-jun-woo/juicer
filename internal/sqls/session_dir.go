//ff:func feature=sql type=parse control=sequence
//ff:what cwd 기준 .juicer 디렉토리 경로 반환
package sqls

// SessionDir returns the .juicer directory path under cwd.
//
func SessionDir() string {
	return sessionDirName
}

