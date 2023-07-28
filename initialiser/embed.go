package initialiser

import "embed"

// Static is an embedded file server containing static HTTP assets.
//
//go:embed *
var Assets embed.FS
