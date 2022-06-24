module github.com/vorsprung/soundslike

go 1.13

require (
	github.com/dlclark/metaphone3 v0.0.0-20190903202417-5fe87fcdd547
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/loads v0.21.1
	github.com/go-openapi/runtime v0.24.1
	github.com/go-openapi/spec v0.20.4
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/jessevdk/go-flags v1.4.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
)

replace github.com/vorsprung/soundslike => ./

replace github.com/vorsprung/strfmt v0.21.2 => github.com/vorsprung/strfmt v0.21.3-0.20220624084400-6899f8befc19
