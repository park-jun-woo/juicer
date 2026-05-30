//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractToArrayFields 테스트
package laravel

import "testing"

func TestExtractToArrayFields(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserResource {
		public function toArray($request) {
			return [ 'id' => $this->id, 'name' => $this->name ];
		}
	}`)
	fields := extractToArrayFields(&fi, "UserResource")
	if len(fields) != 2 {
		t.Fatalf("got %+v", fields)
	}
}
