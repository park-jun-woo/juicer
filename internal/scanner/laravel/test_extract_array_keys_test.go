//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractArrayKeys 테스트
package laravel

import "testing"

func TestExtractArrayKeys(t *testing.T) {
	fi := mustParsePHP(t, `<?php
class R {
    public function toArray($request) {
        return [
            'id' => $this->id,
            'name' => $this->name,
        ];
    }
}`)
	methods := findAllByType(fi.root, "method_declaration")
	if len(methods) == 0 {
		t.Fatal("no method")
	}
	fields := extractArrayKeys(methods[0], fi.src)
	if len(fields) != 2 || fields[0].Name != "id" || fields[1].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
	if fields[0].Type != "string" {
		t.Fatalf("default type: %+v", fields[0])
	}
}
