// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/mpcsdk/mpcCommon/riskAdminService/riskAdminServiceNats"
)

type (
	IRiskAdmin interface {
		RiskAdminCfg() *riskAdminServiceNats.RiskAdminCfg
	}
)

var (
	localRiskAdmin IRiskAdmin
)

func RiskAdmin() IRiskAdmin {
	if localRiskAdmin == nil {
		panic("implement not found for interface IRiskAdmin, forgot register?")
	}
	return localRiskAdmin
}

func RegisterRiskAdmin(i IRiskAdmin) {
	localRiskAdmin = i
}
