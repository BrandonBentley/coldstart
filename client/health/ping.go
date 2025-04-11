package health

import (
	"context"
	"log/slog"

	"github.com/BrandonBentley/slogctx"
)

func (c *Client) Ping(ctx context.Context) error {
	ctx = slogctx.WithAttrs(
		ctx,
		slog.String("function", "health.Client.Ping"),
	)

	slog.InfoContext(
		ctx,
		"Ping called but not yet implemented",
	)

	return nil
}
