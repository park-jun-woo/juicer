//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 경로의 파일을 ctx.parsed에서 찾거나 신규 파싱하여 fileInfo를 반환한다 (캐시)
package express

func loadParsedFile(ctx *scanContext, path string) *fileInfo {
	if fi := ctx.parsed[path]; fi != nil {
		return fi
	}
	fi, err := parseFile(path)
	if err != nil {
		return nil
	}
	ctx.parsed[path] = fi
	return fi
}
