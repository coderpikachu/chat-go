// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

//go:generate mockgen -self_package=test/00Chat1/1internal/apiserver/service/v1 -destination mock_service.go -package v1 test/00Chat1/1internal/apiserver/service/v1 Service,UserSrv,SecretSrv,PolicySrv

import "test/00Chat1/1internal/apiserver/store"

// Service defines functions used to return resource interface.
type Service interface {
	Users() UserSrv
	Secrets() SecretSrv
	Policies() PolicySrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Secrets() SecretSrv {
	return newSecrets(s)
}

func (s *service) Policies() PolicySrv {
	return newPolicies(s)
}
