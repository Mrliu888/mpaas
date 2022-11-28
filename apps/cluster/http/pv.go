package http

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s"

	corev1 "k8s.io/api/core/v1"
)

func (h *handler) registryPVHandler(ws *restful.WebService) {
	tags := []string{"卷管理"}

	ws.Route(ws.GET("/{id}/pv").To(h.QueryPersistentVolumes).
		Doc("查询卷列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(corev1.PersistentVolumeList{})).
		Returns(200, "OK", corev1.PersistentVolumeList{}))
}

func (h *handler) QueryPersistentVolumes(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewListRequestFromHttp(r.Request)
	ins, err := client.ListPersistentVolume(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
