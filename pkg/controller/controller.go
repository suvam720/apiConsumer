package controller

import (
	"context"

	"github.com/suvam720/api/pkg/handler"
)

func Controller(ctx context.Context, userId string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := handler.Handler(ctx, userId); err != nil {
			return err
		}
	}

	return nil
}
