package client

import (
	"context"
)

type IRemoteStore interface {
	GetByID(ctx context.Context, id int) (interface{}, error)
	GetList(ctx context.Context) (interface{}, error)
}
