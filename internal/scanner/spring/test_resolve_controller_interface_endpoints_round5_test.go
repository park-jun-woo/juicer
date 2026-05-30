//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveControllerInterfaceEndpoints_Round5 테스트
package spring

import "testing"

func TestResolveControllerInterfaceEndpoints_Round5(t *testing.T) {
	fi := sFileInfo(t, `
interface UserApi {
    @GetMapping("/u")
    UserDto get();
}
@RestController
class UserController implements UserApi {
    public UserDto get() { return null; }
}
`)
	ci := &controllerInfo{
		className:  "UserController",
		file:       fi.relPath,
		absFile:    fi.absPath,
		interfaces: []string{"UserApi"},
		imports:    map[string]string{},
	}
	resolveControllerInterfaceEndpoints(ci, fi)

}
