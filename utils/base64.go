package utils

import "encoding/base64"

var encoding = base64.URLEncoding.WithPadding(base64.NoPadding)
