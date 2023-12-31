package upstream

import (
	"context"
	"fmt"

	"github.com/infraboard/mpaas/apps/gateway/provider/apisix/common"
)

// 创建Upstream
// /apisix/admin/upstreams
func (c *Client) CreateUpstream(ctx context.Context, in *CreateUpstreamRequeset) (
	*Upstream, error) {
	raw, err := c.c.
		Post("upstreams").
		Body(in.ToJSON()).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

// 查询Upstream详情
// /apisix/admin/upstreams/{id}
func (c *Client) DescribeUpstream(ctx context.Context, in *DescribeUpstreamRequest) (
	*Upstream, error) {
	resp := common.NewReponse()
	err := c.c.
		Get("upstreams").
		Suffix(in.UpStreamId).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	ins := NewUpstream()
	err = resp.GetValue(ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 在 Upstream 中添加一个节点
// /apisix/admin/upstreams/100
func (c *Client) AddNodeToUpstream(ctx context.Context, in *AddNodeToUpstreamRequest) (
	*Upstream, error) {
	raw, err := c.c.
		Patch("upstreams").
		Suffix(in.UpStreamId).
		Body(in.ToJSON()).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}

// 更新 Upstream 中单个节点
func (c *Client) UpdateUpstreamNode(ctx context.Context, in *UpdateUpstreamNodeRequeset) (
	*Upstream, error) {
	return nil, nil
}

// 删除 Upstream 中的一个节点
func (c *Client) RemoveNodeFromUpstream(ctx context.Context, in *RemoveNodeFromUpstreamRequest) (
	*Upstream, error) {
	return nil, nil
}

// 删除Upstream
// /apisix/admin/upstreams/{id}
func (c *Client) DeleteUpstream(ctx context.Context, in *DeleteUpstreamRequest) (
	*Upstream, error) {
	raw, err := c.c.
		Delete("upstreams").
		Suffix(in.UpStreamId).
		Do(ctx).
		Raw()
	fmt.Print(raw, err)
	return nil, nil
}
