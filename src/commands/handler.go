package commands

import (
	"context"

	"github.com/rinq/rinq-go/src/rinq"
)

// NewHandler returns a command handler that dispathes to the test handlers.
func NewHandler(p rinq.Peer) rinq.CommandHandler {
	return func(
		ctx context.Context,
		req rinq.Request,
		res rinq.Response,
	) {
		switch req.Command {

		case "command.success":
			Success(ctx, req, res, p)
		case "command.fail":
			Fail(ctx, req, res, p)
		case "command.fail-payload":
			FailWithPayload(ctx, req, res, p)
		case "command.error":
			Error(ctx, req, res, p)
		case "command.sleep":
			Sleep(ctx, req, res, p)

		case "notify.notify-me":
			NotifyMe(ctx, req, res, p)

		default:
			res.Fail("unknown-command", "no such command: %s", req.Command)
			req.Payload.Close()
		}
	}
}
