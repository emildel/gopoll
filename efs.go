package ui

import "embed"

//go:embed "assets"
var Files embed.FS

//go:embed "tls"
var TLS embed.FS
