//ff:type feature=scan type=model topic=fastify
//ff:what 플러그인 마운트 정보 구조체
package fastify

type pluginMount struct {
	PluginRef    string
	Prefix       string
	FilePath     string
	SourceFile   string
	Inline       bool
	WrapperStart uint32
	WrapperEnd   uint32
}
