package service 

import (
       _adminModel "github.com/onizukazaza/tar-ecom-api/pkg/admin/model"
)

type AdminService interface {
	SetRole(req *_adminModel.SetRoleReq) error
	
}