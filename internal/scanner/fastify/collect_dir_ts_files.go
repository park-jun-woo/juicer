//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 주어진 디렉터리 하위의 .ts 파일 목록을 수집한다 (오류는 빈 목록으로 흡수)
package fastify

func collectDirTSFiles(dir string) []string {
	files, err := findTSFiles(dir)
	if err != nil {
		return nil
	}
	return files
}
