//ff:func feature=ddl type=command control=sequence
//ff:what 마이그레이션 디렉토리를 파싱하여 최종 DDL 문자열 반환
package ddl

// Run parses migrations in dir and returns the final DDL as a string.
func Run(dir string) (string, error) {
	tables, err := Parse(dir)
	if err != nil {
		return "", err
	}
	return Render(nil, tables), nil
}
