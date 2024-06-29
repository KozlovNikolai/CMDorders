package client

import (
	"context"
)

type IRemoteStore interface {
	GetByID(ctx context.Context, id uint64) (interface{}, error)
	GetList(ctx context.Context) (interface{}, error)
}
