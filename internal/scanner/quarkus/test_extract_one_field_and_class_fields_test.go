//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractOneFieldAndClassFields 테스트
package quarkus

import "testing"

func TestExtractOneFieldAndClassFields(t *testing.T) {
	fi := qFileInfo(t, `class UserDto {
		@JsonProperty("user_name") private String name;
		@NotNull private int age;
		private static String CONST;
	}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	fields := extractClassFields(cls, fi.src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields (static skipped), got %d: %+v", len(fields), fields)
	}
	if fields[0].Name != "name" || fields[0].JSON != "user_name" {
		t.Fatalf("field0: %+v", fields[0])
	}
}
