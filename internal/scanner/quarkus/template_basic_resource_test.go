//ff:type feature=scan type=test topic=quarkus
//ff:what 기본 JAX-RS 리소스 테스트 템플릿
package quarkus

var basicUserResourceSource = `
package com.example.resource;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import com.example.dto.CreateUserRequest;
import com.example.dto.UserResponse;
import java.util.List;

@Path("/api/users")
@Produces(MediaType.APPLICATION_JSON)
public class UserResource {

    @GET
    public List<UserResponse> listUsers() {
        return null;
    }

    @GET
    @Path("/{id}")
    public UserResponse getUser(@PathParam("id") Long id) {
        return null;
    }

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    public UserResponse createUser(CreateUserRequest req) {
        return null;
    }
}
`

var basicCreateUserRequestSource = `
package com.example.dto;

import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;

public class CreateUserRequest {
    @NotNull
    private String name;

    @NotNull
    private String email;

    @Size(min = 8, max = 100)
    private String password;
}
`

var basicUserResponseSource = `
package com.example.dto;

public class UserResponse {
    private Long id;
    private String name;
    private String email;
}
`
