//ff:func feature=ddl type=test control=sequence
//ff:what topoSortTables 위상정렬/타이브레이크/순환 fallback 테스트
package ddl

import (
	"reflect"
	"testing"
)

func TestTopoSortTables(t *testing.T) {
	t.Run("linear dependency referenced first", func(t *testing.T) {
		tables := map[string]*Table{
			"image_embeddings": {Name: "image_embeddings", Constraints: []string{
				"FOREIGN KEY (entity_id) REFERENCES entities (id)"}},
			"entities": {Name: "entities", Constraints: []string{
				"FOREIGN KEY (org_id) REFERENCES orgs (id)"}},
			"orgs": {Name: "orgs"},
		}
		got := topoSortTables(tables)
		want := []string{"orgs", "entities", "image_embeddings"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("independent tables alphabetical", func(t *testing.T) {
		tables := map[string]*Table{
			"zebra": {Name: "zebra"},
			"apple": {Name: "apple"},
			"mango": {Name: "mango"},
		}
		got := topoSortTables(tables)
		want := []string{"apple", "mango", "zebra"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("quoted reference resolved (prisma keys)", func(t *testing.T) {
		tables := map[string]*Table{
			`"Token"`: {Name: `"Token"`, Constraints: []string{
				`FOREIGN KEY ("userId") REFERENCES "User" ("id")`}},
			`"User"`: {Name: `"User"`},
		}
		got := topoSortTables(tables)
		want := []string{`"User"`, `"Token"`}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("external reference ignored", func(t *testing.T) {
		tables := map[string]*Table{
			"a": {Name: "a", Constraints: []string{
				"FOREIGN KEY (x) REFERENCES nonexistent (id)"}},
		}
		got := topoSortTables(tables)
		want := []string{"a"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("cycle falls back alphabetically", func(t *testing.T) {
		tables := map[string]*Table{
			"b": {Name: "b", Constraints: []string{"FOREIGN KEY (a_id) REFERENCES a (id)"}},
			"a": {Name: "a", Constraints: []string{"FOREIGN KEY (b_id) REFERENCES b (id)"}},
		}
		got := topoSortTables(tables)
		want := []string{"a", "b"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
}
