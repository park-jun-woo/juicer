//ff:type feature=scan type=model topic=django
//ff:what APIView 정보 구조체
package django

// apiviewInfo holds information about a DRF APIView class.
type apiviewInfo struct {
	name            string
	methods         []string // HTTP methods defined (get, post, etc.)
	serializerClass string
	file            string
	line            int
}
