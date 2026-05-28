package echo

// echo.Context 요청 메서드 분류
var bindMethods = map[string]bool{
	"Bind": true,
}

var queryMethods = map[string]bool{
	"QueryParam":  true,
	"QueryParams": true,
}

var paramMethods = map[string]bool{
	"Param": true,
}

var formMethods = map[string]bool{
	"FormValue": true,
}

var fileMethods = map[string]bool{
	"FormFile": true,
}

// 응답 메서드와 인자 수 패턴
var responseMethods = map[string]string{
	"JSON":      "json",
	"JSONPretty": "json",
	"String":    "string",
	"HTML":      "html",
	"Blob":      "data",
	"Stream":    "data",
	"File":      "file",
	"Redirect":  "redirect",
	"NoContent": "status",
}
