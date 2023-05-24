package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl cluster.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(cluster.AppName).(cluster.Service)
}