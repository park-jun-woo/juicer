//ff:func feature=scan type=extract control=sequence dimension=1
//ff:what 디렉터리명이 언어 공통 테스트/픽스처 디렉터리(test, tests, __tests__, spec, fixtures, testdata, e2e 등)인지 판정한다
package scanner

import "strings"

// testDirs holds language-agnostic directory names that hold test code or
// fixtures and must be excluded from endpoint scanning to avoid false positives.
var testDirs = map[string]bool{
	"test":      true,
	"tests":     true,
	"__tests__": true,
	"spec":      true,
	"specs":     true,
	"__mocks__": true,
	"fixtures":  true,
	"testdata":  true,
	"e2e":       true,
}

// IsTestDir reports whether a directory name is a common test/fixture directory.
// The match is case-insensitive so "Tests" (C#/.NET) and "Spec" also match.
func IsTestDir(name string) bool {
	return testDirs[strings.ToLower(name)]
}
