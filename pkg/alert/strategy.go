package alert

import "context"

type Strategy interface {
	EmitAlert(ctx context.Context, Alert Alert) error
}
