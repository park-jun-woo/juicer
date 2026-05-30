//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractSpecImportName_Alias 테스트
package express

import "testing"

func TestExtractSpecImportName_Alias(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router as R } from 'express';`))
	if got := extractSpecImportName(firstImportSpec(t, fi), fi.Src); got != "R" {
		t.Fatalf("got %q", got)
	}
}
