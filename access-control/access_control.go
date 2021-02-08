package access_control

import (
	"github.com/casbin/casbin/v2"
)

type AccessControl interface {
	Enforce(sub, domain, mod, field, act string) (bool, error)
}

type accessControl struct {
	enforcer *casbin.Enforcer
}

func NewAccessControl(enforcer *casbin.Enforcer) AccessControl {
	return &accessControl{
		enforcer: enforcer,
	}
}

func (r *accessControl) Enforce(sub, domain, mod, field, act string) (bool, error) {
	ok, err := r.enforcer.Enforce(sub, domain, mod, field, act)
	if err != nil {
		return ok, err
	}

	return ok, nil
}
