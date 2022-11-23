package mastodon

import "errors"

var ErrNotFound = errors.New("404 not-found")

var ErrInvalidParameter = errors.New("invalid parameter")

var ErrTooManyMedia = errors.New("too many media (max 4)")

var ErrInvalidID = errors.New("invalid ID")
