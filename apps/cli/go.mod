module nx-go-playground/cli

go 1.21

replace nx-go-playground/logger => ../../libs/logger
replace nx-go-playground/math => ../../libs/math
replace nx-go-playground/geometry => ../../libs/geometry

require (
	nx-go-playground/geometry v0.0.0-00010101000000-000000000000
	nx-go-playground/logger v0.0.0-00010101000000-000000000000
)

require nx-go-playground/math v0.0.0-00010101000000-000000000000 // indirect

