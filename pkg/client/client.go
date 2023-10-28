package client

import (
	"context"
	"errors"
	"gateway/pkg/common"
	"gateway/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Client pb.UserClient
}

func InitUserClient(cfg common.Config) (UserClient, error) {
	clientConn, err := grpc.Dial(cfg.UserClientPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return UserClient{}, err
	}
	return NewUserClient(pb.NewUserClient(clientConn)), nil
}

func NewUserClient(client pb.UserClient) UserClient {
	return UserClient{
		Client: client,
	}
}

func (c UserClient) HealthCheck(ctx context.Context) (*pb.Response, error) {
	resp, err := c.Client.HealthCheck(ctx, &pb.Request{
		Data: "Hi user server",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c UserClient) UserDetails(ctx context.Context, req *pb.UserDetailsRequest) (*pb.UserDetailsResponse, error) {
	resp, err := c.Client.UserDetails(ctx, req)
	if err != nil {
		return nil, errors.New("Client error :" + err.Error())
	}
	return resp, nil
}

func (c UserClient) UserListDetails(ctx context.Context, req *pb.UserListDetailsRequest) (*pb.UserListDetailsResponse, error) {
	resp, err := c.Client.UserListDetails(ctx, req)
	if err != nil {
		return nil, errors.New("Client error :" + err.Error())
	}
	return resp, nil
}
