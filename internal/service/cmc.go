// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICmc interface {
		GetPrice(ctx context.Context, token string) float64
	}
)

var (
	localCmc ICmc
)

func Cmc() ICmc {
	if localCmc == nil {
		panic("implement not found for interface ICmc, forgot register?")
	}
	return localCmc
}

func RegisterCmc(i ICmc) {
	localCmc = i
}
