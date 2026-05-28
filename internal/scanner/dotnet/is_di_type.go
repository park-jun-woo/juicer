//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 타입이 DI 서비스 타입인지 확인한다
package dotnet

import "strings"

func isDIType(t string) bool {
	if diTypes[t] {
		return true
	}
	if strings.HasPrefix(t, "ILogger<") {
		return true
	}
	if strings.HasSuffix(t, "DbContext") {
		return true
	}
	if strings.HasPrefix(t, "I") && len(t) > 1 && t[1] >= 'A' && t[1] <= 'Z' {
		if strings.HasSuffix(t, "Service") || strings.HasSuffix(t, "Repository") {
			return true
		}
	}
	return false
}
