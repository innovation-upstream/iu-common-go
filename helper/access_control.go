package helper

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/casbin/casbin/v2"
	firestoreadapter "github.com/reedom/casbin-firestore-adapter"
	access_control "gitlab.innovationup.stream/innovation-upstream/iu-common-go/access-control"
	"unknwon.dev/clog/v2"
)

func GetAccessControl(ctx context.Context, db *firestore.Client) access_control.AccessControl {
	a := firestoreadapter.NewAdapter(db)
	enforcer, err := casbin.NewEnforcer("rbac_model.conf", a)
	if err != nil {
		clog.Fatal("%+v", err)
	}

	ac := access_control.NewAccessControl(enforcer)
	return ac
}
