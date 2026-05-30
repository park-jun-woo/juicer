//ff:type feature=scan type=test topic=quarkus
//ff:what sampleResource 테스트 보조 선언
package quarkus

const sampleResource = `
@Path("/users")
@RolesAllowed({"admin"})
public class UserResource {
    @GET
    @Path("/{id}")
    public UserDto get(@PathParam("id") Long id) { return null; }

    @POST
    public Response create(UserDto dto) { return Response.status(201).build(); }
}
`
