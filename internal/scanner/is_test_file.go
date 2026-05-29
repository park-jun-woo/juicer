//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 파일명이 언어별 테스트 네이밍(*.test.*, *.spec.*, *_test.*, *Test.java, *Tests.cs 등)에 매칭되는지 판정한다
package scanner

import "strings"

// testFileSuffixes holds full filename suffixes (lowercased) that mark a source
// file as a test, covering JS/TS (*.test.*, *.spec.*), Python/Rust (*_test.*),
// Java (*test.java, *tests.java) and C# (*test.cs, *tests.cs).
var testFileSuffixes = []string{
	"test.java", "tests.java", "test.cs", "tests.cs",
}

// testFileInfixes holds substrings that mark a source file as a test
// (TS/JS *.test.ts / *.spec.ts, Python/Rust *_test.py / *_spec.rs, etc.).
var testFileInfixes = []string{".test.", ".spec.", "_test.", "_spec."}

// IsTestFile reports whether a source file name is a test file by its naming
// convention.
func IsTestFile(name string) bool {
	lower := strings.ToLower(name)
	for _, infix := range testFileInfixes {
		if strings.Contains(lower, infix) {
			return true
		}
	}
	for _, suffix := range testFileSuffixes {
		if strings.HasSuffix(lower, suffix) {
			return true
		}
	}
	return false
}
