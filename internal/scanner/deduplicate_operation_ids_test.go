//ff:func feature=scan type=test control=sequence
//ff:what TestDeduplicateOperationIDs 테스트
package scanner

import "testing"

func TestDeduplicateOperationIDs(t *testing.T) {
	t.Run("no duplicates", func(t *testing.T) {
		eps := []Endpoint{
			{Method: "GET", Path: "/api/v1/users", Handler: "h.ListUsers"},
			{Method: "POST", Path: "/api/v1/users", Handler: "h.CreateUser"},
		}
		got := deduplicateOperationIDs(eps)
		if got[0] != "listUsers" {
			t.Errorf("got[0] = %q, want %q", got[0], "listUsers")
		}
		if got[1] != "createUser" {
			t.Errorf("got[1] = %q, want %q", got[1], "createUser")
		}
	})

	t.Run("duplicate findAll from inheritance", func(t *testing.T) {
		eps := []Endpoint{
			{Method: "GET", Path: "/api/v1/categories", Handler: "CategoriesController.findAll"},
			{Method: "GET", Path: "/api/v1/products", Handler: "ProductsController.findAll"},
			{Method: "GET", Path: "/api/v1/users", Handler: "UsersController.findAll"},
		}
		got := deduplicateOperationIDs(eps)
		// 각각 다른 prefix가 붙어야 한다
		seen := map[string]bool{}
		for i := 0; i < 3; i++ {
			if seen[got[i]] {
				t.Fatalf("duplicate operationId: %s", got[i])
			}
			seen[got[i]] = true
		}
		// prefix가 categories, products, users여야 함
		if got[0] != "categoriesFindAll" {
			t.Errorf("got[0] = %q, want %q", got[0], "categoriesFindAll")
		}
		if got[1] != "productsFindAll" {
			t.Errorf("got[1] = %q, want %q", got[1], "productsFindAll")
		}
		if got[2] != "usersFindAll" {
			t.Errorf("got[2] = %q, want %q", got[2], "usersFindAll")
		}
	})

	t.Run("secondary duplicate adds suffix", func(t *testing.T) {
		// 같은 경로 prefix를 가진 중복 핸들러 (예: /api/v1/items와 /items)
		eps := []Endpoint{
			{Method: "GET", Path: "/api/v1/items", Handler: "A.findAll"},
			{Method: "GET", Path: "/items", Handler: "B.findAll"},
		}
		got := deduplicateOperationIDs(eps)
		// 둘 다 prefix가 items이므로 2차 중복 발생
		if got[0] == got[1] {
			t.Fatalf("expected different operationIds, both are %q", got[0])
		}
		if got[0] != "itemsFindAll" {
			t.Errorf("got[0] = %q, want %q", got[0], "itemsFindAll")
		}
		if got[1] != "itemsFindAll2" {
			t.Errorf("got[1] = %q, want %q", got[1], "itemsFindAll2")
		}
	})

	t.Run("same path GET POST same handler uses method prefix", func(t *testing.T) {
		// 동일 path의 GET/POST가 같은 핸들러를 공유 → path 세그먼트가 login으로 동일하여
		// doubling(loginLogin/loginLogin2)이 발생하던 케이스. method 접두로 구분한다.
		eps := []Endpoint{
			{Method: "GET", Path: "/login", Handler: "h.login"},
			{Method: "POST", Path: "/login", Handler: "h.login"},
		}
		got := deduplicateOperationIDs(eps)
		if got[0] != "getLogin" {
			t.Errorf("got[0] = %q, want %q", got[0], "getLogin")
		}
		if got[1] != "postLogin" {
			t.Errorf("got[1] = %q, want %q", got[1], "postLogin")
		}
		// doubling/무의미 접미가 발생하지 않아야 한다
		for i := 0; i < 2; i++ {
			if got[i] == "loginLogin" || got[i] == "loginLogin2" {
				t.Errorf("got[%d] = %q, doubling/suffix must not occur", i, got[i])
			}
		}
	})

	t.Run("empty input", func(t *testing.T) {
		got := deduplicateOperationIDs(nil)
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %d entries", len(got))
		}
	})
}
