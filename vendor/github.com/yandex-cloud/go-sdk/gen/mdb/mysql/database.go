// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mysql

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// DatabaseServiceClient is a mysql.DatabaseServiceClient with
// lazy GRPC connection initialization.
type DatabaseServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ mysql.DatabaseServiceClient = &DatabaseServiceClient{}

// Create implements mysql.DatabaseServiceClient
func (c *DatabaseServiceClient) Create(ctx context.Context, in *mysql.CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewDatabaseServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements mysql.DatabaseServiceClient
func (c *DatabaseServiceClient) Delete(ctx context.Context, in *mysql.DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewDatabaseServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements mysql.DatabaseServiceClient
func (c *DatabaseServiceClient) Get(ctx context.Context, in *mysql.GetDatabaseRequest, opts ...grpc.CallOption) (*mysql.Database, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewDatabaseServiceClient(conn).Get(ctx, in, opts...)
}

// List implements mysql.DatabaseServiceClient
func (c *DatabaseServiceClient) List(ctx context.Context, in *mysql.ListDatabasesRequest, opts ...grpc.CallOption) (*mysql.ListDatabasesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewDatabaseServiceClient(conn).List(ctx, in, opts...)
}