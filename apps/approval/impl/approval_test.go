package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryApproval(t *testing.T) {
	req := approval.NewQueryApprovalRequest()
	set, err := impl.QueryApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestDescribeApproval(t *testing.T) {
	req := approval.NewDescribeApprovalRequest(conf.C.DEVCLOUD_DEPLOY_APPROVAL_ID)
	ins, err := impl.DescribeApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestEditApproval(t *testing.T) {
	req := approval.NewEditApprovalRequest(conf.C.DEVCLOUD_DEPLOY_APPROVAL_ID)
	ins, err := impl.EditApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCreateApproval(t *testing.T) {
	req := approval.NewCreateApprovalRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.CreateBy = "test"
	req.Title = "v1.0.0"
	req.Describe = "发布说明, 支持Markdown语法"
	req.AddOperator("test@default")
	req.AddAuditor("test@default")
	req.AutoRun = true
	tools.MustReadYamlFile("test/create.yml", req.PipelineSpec)
	set, err := impl.CreateApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestUpdateApprovalStatus(t *testing.T) {
	req := approval.NewUpdateApprovalStatusRequest(conf.C.DEVCLOUD_DEPLOY_APPROVAL_ID)
	req.Status.Stage = approval.STAGE_PASSED
	req.UpdateBy = "test"
	req.Status.AuditComment = "好好干，日子会越来越甜"
	ins, err := impl.UpdateApprovalStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestCloseApproval(t *testing.T) {
	req := approval.NewUpdateApprovalStatusRequest(conf.C.DEVCLOUD_DEPLOY_APPROVAL_ID)
	req.Status.Stage = approval.STAGE_CLOSED
	req.UpdateBy = "test"
	req.Status.CloseComment = "好好干，日子会越来越甜"
	ins, err := impl.UpdateApprovalStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteApproval(t *testing.T) {
	req := approval.NewDeleteApprovalRequest(conf.C.DEVCLOUD_DEPLOY_APPROVAL_ID)
	ins, err := impl.DeleteApproval(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
