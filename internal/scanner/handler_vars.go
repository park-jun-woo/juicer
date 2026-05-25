package scanner

// gin.Context 요청 메서드 분류
var bindMethods = map[string]bool{
	"ShouldBindJSON": true, "BindJSON": true,
	"ShouldBind": true, "Bind": true,
	"ShouldBindQuery": true, "BindQuery": true,
	"ShouldBindUri": true, "BindUri": true,
}

var queryMethods = map[string]bool{
	"Query": true, "DefaultQuery": true, "GetQuery": true,
}

var paramMethods = map[string]bool{
	"Param": true,
}

var formMethods = map[string]bool{
	"PostForm": true, "DefaultPostForm": true, "GetPostForm": true,
}

var fileMethods = map[string]bool{
	"FormFile": true,
}

var rawBodyMethods = map[string]bool{
	"GetRawData": true,
}

// 응답 메서드와 인자 수 패턴
var responseMethods = map[string]string{
	"JSON":                "json",
	"AbortWithStatusJSON": "json",
	"String":              "string",
	"Data":                "data",
	"DataFromReader":      "data",
	"File":                "file",
	"Redirect":            "redirect",
	"Status":              "status",
	"AbortWithStatus":     "status",
}
