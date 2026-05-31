//ff:type feature=scan type=model topic=django
//ff:what 클래스 이름→부모목록 인덱스 타입 (상속체인 전이 해소용)
package django

// classIndex maps a class name to its direct parent class names across all files.
// It is used to resolve inheritance chains transitively (e.g. a custom
// BaseViewSet that extends DRF's ModelViewSet).
type classIndex map[string][]string
