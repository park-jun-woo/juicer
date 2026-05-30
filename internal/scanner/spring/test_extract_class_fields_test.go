//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractClassFields 테스트
package spring

import "testing"

func TestExtractClassFields(t *testing.T) {
	fi := sFileInfo(t, `class UserDto {
		@JsonProperty("user_name") private String name;
		private int age;
		private static String CONST;
	}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	fields := extractClassFields(cls, fi.src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 (static skipped), got %d", len(fields))
	}
}
