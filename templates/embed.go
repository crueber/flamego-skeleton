package templates

import "embed"

//go:embed *.tmpl include user
var Templates embed.FS
