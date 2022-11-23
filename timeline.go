package mastodon

import (
	"errors"
	"log"
	"strings"
)

const (
	DirectTimeline = "direct"
	HomeTimeline   = "home"
	PublicTimeline = "public"
)

func (c *client) GetTimeline(timeline, minId, maxId string, local, onlyMedia bool, limit int) ([]Status, error) {

	if limit < 10 {
		limit = 10
	}

	endPoint := NewUrlBuilder().
		Path("/api/v1/timelines")

	switch {
	case timeline == "home", timeline == "public", timeline == "direct":
		endPoint = endPoint.Path(timeline)

	case strings.HasPrefix(timeline, ":"), strings.HasPrefix(timeline, "#"):
		hashtag := timeline[1:]
		if hashtag == "" {
			return nil, errors.New("timelines API: empty hashtag")
		}
		endPoint = endPoint.Path("tag").Path(hashtag)

	case len(timeline) > 1 && strings.HasPrefix(timeline, "!"):
		// Check the timeline is a number
		for _, n := range timeline[1:] {
			if n < '0' || n > '9' {
				return nil, errors.New("timelines API: invalid list ID")
			}
		}
		endPoint = endPoint.Path("list").Path(timeline[1:])

	default:
		return nil, errors.New("GetTimelines: bad timelines argument")
	}

	endPoint = endPoint.ParamBoolIfTrue("local", timeline == PublicTimeline && local).
		ParamBoolIfTrue("only_media", onlyMedia)

	var statuses []Status

	err := c.request(GET, endPoint.Build(), nil, nil, &statuses)
	if err != nil {
		return nil, err
	}

	if c.IsDebug() {
		log.Printf("Got %d statuses", len(statuses))
	}

	return statuses, nil
}
