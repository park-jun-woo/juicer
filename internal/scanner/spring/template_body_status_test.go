//ff:func feature=scan type=test control=sequence topic=spring
//ff:what extractBodyStatus 테스트용 Java 소스 템플릿
package spring

var bodyStatusTests = []struct {
	name string
	src  string
	want string
}{
	{
		name: "ResponseEntity.ok",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public ResponseEntity<?> create() {
        return ResponseEntity.ok(new Item());
    }
}`,
		want: "200",
	},
	{
		name: "ResponseEntity.created",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public ResponseEntity<?> create() {
        return ResponseEntity.created(uri).body(item);
    }
}`,
		want: "201",
	},
	{
		name: "ResponseEntity.noContent",
		src: `
package com.example;
@RestController
public class C {
    @DeleteMapping("/a")
    public ResponseEntity<Void> delete() {
        return ResponseEntity.noContent().build();
    }
}`,
		want: "204",
	},
	{
		name: "ResponseEntity.badRequest",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public ResponseEntity<?> create() {
        return ResponseEntity.badRequest().body("error");
    }
}`,
		want: "400",
	},
	{
		name: "ResponseEntity.notFound",
		src: `
package com.example;
@RestController
public class C {
    @GetMapping("/a")
    public ResponseEntity<?> get() {
        return ResponseEntity.notFound().build();
    }
}`,
		want: "404",
	},
	{
		name: "ResponseEntity.status with HttpStatus",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public ResponseEntity<?> create() {
        return ResponseEntity.status(HttpStatus.CONFLICT).body("dup");
    }
}`,
		want: "409",
	},
	{
		name: "ResponseEntity.status with integer literal",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public ResponseEntity<?> create() {
        return ResponseEntity.status(202).body("accepted");
    }
}`,
		want: "202",
	},
	{
		name: "new ResponseEntity with HttpStatus",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public ResponseEntity<?> create() {
        return new ResponseEntity<>(item, HttpStatus.CREATED);
    }
}`,
		want: "201",
	},
	{
		name: "no ResponseEntity pattern",
		src: `
package com.example;
@RestController
public class C {
    @PostMapping("/a")
    public String create() {
        return "done";
    }
}`,
		want: "",
	},
}
