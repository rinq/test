package commands

import (
	"context"

	"github.com/rinq/rinq-go/src/rinq"
)

// NotifyMe sends a notification back to the caller containing the request payload.
func NotifyMe(
	ctx context.Context,
	req rinq.Request,
	res rinq.Response,
	peer rinq.Peer,
) {
	defer req.Payload.Close()

	sess := peer.Session()

	if err := sess.Notify(
		ctx,
		req.Source.Ref().ID,
		"notify-me",
		req.Payload,
	); err != nil {
		res.Error(err)
	} else {
		res.Close()
	}
}
