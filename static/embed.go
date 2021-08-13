package static

import "embed"

//go:embed *
// Static is an embedded file server containing static HTTP assets.
var Assets embed.FS

