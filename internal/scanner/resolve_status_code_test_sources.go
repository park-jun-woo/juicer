package scanner

var resolveStatusCodeConstSrc = `package test

const StatusOK = 200

type T struct{}

func (t T) StatusOK() int { return StatusOK }

func f() int {
	return StatusOK
}
`

var resolveStatusCodeFallbackSrc = `package test

const code = 404

func f() int {
	return code + 0
}
`
