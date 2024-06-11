package models

import "context"

type FriendRequest struct {
	Id          int `json:"id"`
	RequestorId int `json:"requestorId"`
	RecipientId int `json:"recipientId"`
}

type FriendRepository interface {
	CreateFriendRequest(c context.Context, req *FriendRequest) (int, error)
}
