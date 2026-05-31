//ff:func feature=scan type=test topic=laravel control=iteration dimension=1
//ff:what appendAbortResponses 메서드 본문 abort() 호출에서 응답 추가(중복 제외) 테스트
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAppendAbortResponses(t *testing.T) {
	fi := mustParsePHP(t, `<?php
class C {
    public function show() {
        abort(404);
        abort_if($x, 403);
    }
}`)
	method := findAllByType(fi.root, "method_declaration")[0]
	cm := &controllerMethod{methodNode: method, src: fi.src}

	resp := appendAbortResponses(cm, nil)
	got := map[string]bool{}
	for _, r := range resp {
		got[r.Status] = true
	}
	if !got["404"] || !got["403"] {
		t.Errorf("abort responses: %+v", resp)
	}

	// pre-existing status not duplicated
	resp2 := appendAbortResponses(cm, []scanner.Response{{Status: "404"}})
	count404 := 0
	for _, r := range resp2 {
		if r.Status == "404" {
			count404++
		}
	}
	if count404 != 1 {
		t.Errorf("404 duplicated: %+v", resp2)
	}

	// nil methodNode -> unchanged
	if r := appendAbortResponses(&controllerMethod{}, nil); r != nil {
		t.Errorf("nil method: %+v", r)
	}
}
