package helper

import (
	"context"

	"github.com/casbin/casbin/v2"
	firestoreadapter "github.com/reedom/casbin-firestore-adapter"
	"gitlab.innovationup.stream/innovation-upstream/iu-common-go/repo"
	"unknwon.dev/clog/v2"
)

func GetAccessControl(ctx context.Context) AccessControl {
	db := repo.NewFirestoreClient(ctx)
	a := firestoreadapter.NewAdapter(db)
	enforcer, err := casbin.NewEnforcer("rbac_model.conf", a)
	if err != nil {
		clog.Fatal("%+v", err)
	}

	ac := NewAccessControl(enforcer)
	return ac
}
