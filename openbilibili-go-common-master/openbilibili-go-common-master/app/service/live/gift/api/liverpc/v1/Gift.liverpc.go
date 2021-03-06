// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Gift.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/Gift.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ==============
// Gift Interface
// ==============

type Gift interface {
	// * 增加包裹道具
	//
	AddFreeGift(context.Context, *GiftAddFreeGiftReq) (*GiftAddFreeGiftResp, error)
}

// ====================
// Gift Live Rpc Client
// ====================

type giftRpcClient struct {
	client *liverpc.Client
}

// NewGiftRpcClient creates a Rpc client that implements the Gift interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewGiftRpcClient(client *liverpc.Client) Gift {
	return &giftRpcClient{
		client: client,
	}
}

func (c *giftRpcClient) AddFreeGift(ctx context.Context, in *GiftAddFreeGiftReq) (*GiftAddFreeGiftResp, error) {
	out := new(GiftAddFreeGiftResp)
	err := doRpcRequest(ctx, c.client, 1, "Gift.addFreeGift", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRpcRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message) (err error) {
	err = client.Call(ctx, version, method, in, out)
	return
}
