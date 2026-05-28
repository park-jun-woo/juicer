//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_DataAnnotations -- Data Annotation 변환 테스트
package dotnet

import "testing"

func TestScan_DataAnnotations(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "Controllers/UsersController.cs", basicControllerSource)
	writeFile(t, dir, "Models/CreateUserRequest.cs", basicCreateUserDtoSource)
	writeFile(t, dir, "Models/UserDto.cs", basicUserDtoSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	ep2 := result.Endpoints[2]
	if ep2.Request == nil || ep2.Request.Body == nil {
		t.Fatalf("expected body")
	}
	fields := ep2.Request.Body.Fields
	if len(fields) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(fields))
	}

	nameField := fields[0]
	if nameField.Name != "Name" {
		t.Errorf("field[0] name: want Name, got %s", nameField.Name)
	}
	if nameField.Validate != "required" {
		t.Errorf("field[0] validate: want required, got %s", nameField.Validate)
	}
	if nameField.MaxLength == nil || *nameField.MaxLength != 100 {
		t.Errorf("field[0] maxLength: want 100")
	}

	emailField := fields[1]
	if emailField.Validate != "email" {
		t.Errorf("field[1] validate: want email, got %s", emailField.Validate)
	}

	ageField := fields[2]
	if ageField.Minimum == nil || *ageField.Minimum != 0 {
		t.Errorf("field[2] minimum: want 0")
	}
	if ageField.Maximum == nil || *ageField.Maximum != 150 {
		t.Errorf("field[2] maximum: want 150")
	}
}
