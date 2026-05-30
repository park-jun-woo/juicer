//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractSpecImportName_Name 테스트
package express

import "testing"

func TestExtractSpecImportName_Name(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router } from 'express';`))
	if got := extractSpecImportName(firstImportSpec(t, fi), fi.Src); got != "Router" {
		t.Fatalf("got %q", got)
	}
}
