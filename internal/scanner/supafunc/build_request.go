//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what 추출된 bodyFields와 queryParams로 Request 구조체를 구성한다
package supafunc

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildRequest(bodyFields []string, queryParams []string) *scanner.Request {
	if len(bodyFields) == 0 && len(queryParams) == 0 {
		return nil
	}
	req := &scanner.Request{}
	if len(bodyFields) > 0 {
		req.Body = &scanner.Body{
			VarName: "body",
			Method:  "json",
			Fields:  buildBodyFields(bodyFields),
		}
	}
	for _, q := range queryParams {
		req.Query = append(req.Query, scanner.Param{Name: q, Type: "string"})
	}
	return req
}
