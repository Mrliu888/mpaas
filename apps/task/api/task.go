package api

import (
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/task"
)

// 用户自己手动管理任务状态相关接口
func (h *Handler) RegistryUserHandler(ws *restful.WebService) {
	tags := []string{"任务管理"}
	ws.Route(ws.POST("/{id}/status").To(h.UpdateJobTaskStatus).
		Doc("更新任务状态").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Reads(task.UpdateJobTaskStatusRequest{}).
		Writes(task.JobTask{}))

	ws.Route(ws.POST("/{id}/output").To(h.UpdateJobTaskOutput).
		Doc("更新任务输出").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Reads(task.UpdateJobTaskOutputRequest{}).
		Writes(task.JobTask{}))

	ws.Route(ws.GET("/{id}/log").To(h.WatchTaskLog).
		Doc("查询任务日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Reads(task.WatchJobTaskLogRequest{}).
		Writes(task.WatchJobTaskLogReponse{}))
}

func (h *Handler) UpdateJobTaskOutput(r *restful.Request, w *restful.Response) {
	req := task.NewUpdateJobTaskOutputRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateJobTaskOutput(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *Handler) UpdateJobTaskStatus(r *restful.Request, w *restful.Response) {
	req := task.NewUpdateJobTaskStatusRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateJobTaskStatus(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: 60 * time.Second,
		ReadBufferSize:   8192,
		WriteBufferSize:  8192,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (h *Handler) WatchTaskLog(r *restful.Request, w *restful.Response) {
	// websocket handshake
	ws, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		response.Failed(w, err)
		return
	}
	defer ws.Close()

	in := task.NewWatchJobTaskLogRequest(r.PathParameter("id"))

	// 读取请求
	term := task.NewTaskLogTerminal(ws)
	if err = term.ReadReq(in); err != nil {
		term.Failed(err)
		return
	}

	// 输出日志到Term中
	if err = h.service.WatchJobTaskLog(in, term); err != nil {
		term.Failed(err)
		return
	}

	term.Success("ok")
}
