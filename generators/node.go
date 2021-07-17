package generators

import (
	_ "embed"
)

//go:embed node/function.js
var function string

//go:embed node/package.json
var pkg string
