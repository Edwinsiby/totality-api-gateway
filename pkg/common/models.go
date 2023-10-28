package common

import "gateway/pkg/pb"

type Config struct {
	Port           string `mapstructure:"PORT"`
	UserClientPort string `mapstructure:"USERCLIENTPORT"`
}

type UserID struct {
	UserID int32 `json:"userID"`
}

type UserIDList struct {
	UserID []int32 `json:"userID"`
}

type UserData struct {
	userID    int32
	FirstName string
	City      string
	Height    float32
	Married   bool
}

type UserListData struct {
	UsersFound    []*pb.UserDetails
	UsersNotFound []int32
}
