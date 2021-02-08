package access_control

import (
	"github.com/casbin/casbin/v2"
)

type AccessControl interface {
	Enforce(sub, domain, mod, field, act string) (bool, error)
	AddPolicy(sub, domain, mod, field, act string) (bool, error)
}

type accessControl struct {
	enforcer *casbin.Enforcer
}

func NewAccessControl(enforcer *casbin.Enforcer) AccessControl {
	return &accessControl{
		enforcer: enforcer,
	}
}

func (ac *accessControl) Enforce(sub, domain, mod, field, act string) (bool, error) {
	ok, err := ac.enforcer.Enforce(sub, domain, mod, field, act)
	if err != nil {
		return ok, err
	}

	return ok, nil
}

func (ac *accessControl) AddPolicy(sub, domain, mod, field, act string) (bool, error) {
	ok, err := ac.enforcer.AddPolicy(sub, domain, mod, field, act)
	if err != nil {
		return ok, err
	}

	return ok, nil
}
