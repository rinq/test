package commands

import (
	"context"
	"errors"
	"time"

	"github.com/rinq/rinq-go/src/rinq"
)

// Success sends a request payload back to the caller as the response payload.
func Success(
	ctx context.Context,
	req rinq.Request,
	res rinq.Response,
	peer rinq.Peer,
) {
	defer req.Payload.Close()

	res.Done(req.Payload)
}

// Fail sends a failure response.
func Fail(
	ctx context.Context,
	req rinq.Request,
	res rinq.Response,
	peer rinq.Peer,
) {
	defer req.Payload.Close()

	res.Fail("test-failure", "Failure requested by client.")
}

// FailWithPayload sends a failure response, using the response payload as the
// failure payload.
func FailWithPayload(
	ctx context.Context,
	req rinq.Request,
	res rinq.Response,
	peer rinq.Peer,
) {
	defer req.Payload.Close()

	res.Error(rinq.Failure{
		Type:    "test-failure",
		Message: "Failure requested by client.",
		Payload: req.Payload.Clone(),
	})
}

// Error sends a error response.
func Error(
	ctx context.Context,
	req rinq.Request,
	res rinq.Response,
	peer rinq.Peer,
) {
	defer req.Payload.Close()

	res.Error(errors.New("You done goofed."))
}

// Sleep closes the response after sleeping.
func Sleep(
	ctx context.Context,
	req rinq.Request,
	res rinq.Response,
	peer rinq.Peer,
) {
	defer req.Payload.Close()

	var millis time.Duration
	if err := req.Payload.Decode(&millis); err != nil {
		res.Error(err)
		return
	}

	time.Sleep(millis * time.Millisecond)
	res.Close()
}
