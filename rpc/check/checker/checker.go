// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

//go:generate mockgen -destination ./checker_mock.go -package checker -source $GOFILE

package checker

import (
	"bookstore/rpc/check/check"
	"context"


	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CheckResp = check.CheckResp
	CheckReq  = check.CheckReq

	Checker interface {
		Check(ctx context.Context, in *CheckReq) (*CheckResp, error)
	}

	defaultChecker struct {
		cli zrpc.Client
	}
)

func NewChecker(cli zrpc.Client) Checker {
	return &defaultChecker{
		cli: cli,
	}
}

func (m *defaultChecker) Check(ctx context.Context, in *CheckReq) (*CheckResp, error) {
	client := check.NewCheckerClient(m.cli.Conn())
	return client.Check(ctx, in)
}