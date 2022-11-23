package mastodon

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	// VisibilityPublic Visible to everyone, shown in public timelines.
	VisibilityPublic = "public"
	// VisibilityUnlisted Visible to public, but not included in public timelines.
	VisibilityUnlisted = "unlisted"
	// VisibilityPrivate Visible to followers only, and to any mentioned users.
	VisibilityPrivate = "private"
	// VisibilityDirect Visible only to mentioned users. Deprecated.
	VisibilityDirect = "direct"
)

type PostStatus struct {
	IdempotencyKey string  // Optional, used to prevent duplicate posts, valid for 1 hour
	Text           string  // Text of the post
	InReplyTo      int64   // Id of the post this is a reply to
	MediaIds       []int64 // Media that is attached to this status.
	Sensitive      bool    // Is this status marked as sensitive content?
	SpoilerText    string  // Subject or summary line, below which status content is collapsed until expanded.
	Visibility     string  // Visibility of this status.
}

func (c *client) Post(status PostStatus) (*Status, error) {
	if status.Text == "" {
		return nil, ErrInvalidParameter
	}

	switch status.Visibility {
	case "", VisibilityPublic, VisibilityUnlisted, VisibilityPrivate, VisibilityDirect:
		// Okay
	default:
		return nil, ErrInvalidParameter
	}

	if len(status.MediaIds) > 4 {
		return nil, ErrTooManyMedia
	}

	params := make(RequestParams)
	params["status"] = status.Text

	if status.InReplyTo > 0 {
		params["in_reply_to_id"] = strconv.FormatInt(status.InReplyTo, 10)
	}

	resp := &Status{}

	err := c.request("POST", "/api/v1/statuses",
		func(req *http.Request, _ interface{}) error {
			if status.IdempotencyKey != "" {
				req.Header.Set("Idempotency-Key", status.IdempotencyKey)
			}

			for i, id := range status.MediaIds {
				if id < 1 {
					return ErrInvalidID
				}
				qID := fmt.Sprintf("media_ids[%d]", i)
				params[qID] = strconv.FormatInt(id, 10)
			}

			if status.Sensitive {
				params["sensitive"] = "true"
			}

			if status.SpoilerText != "" {
				params["spoiler_text"] = status.SpoilerText
			}

			if status.Visibility != "" {
				params["visibility"] = status.Visibility
			}

			return nil
		}, params, resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
