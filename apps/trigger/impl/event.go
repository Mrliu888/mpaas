package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/trigger"
)

// 应用事件处理
func (i *impl) HandleEvent(ctx context.Context, in *trigger.Event) (
	*trigger.Record, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := trigger.NewRecord(in)

	switch in.Provider {
	case trigger.EVENT_PROVIDER_GITLAB:
		// 校验请求
		if err := in.GitlabEvent.Validate(); err != nil {
			return nil, exception.NewBadRequest(err.Error())
		}

		// 获取该服务对应事件的构建配置
		req := build.NewQueryBuildConfigRequest()
		req.AddService(in.GitlabEvent.EventToken)
		req.Event = in.GitlabEvent.EventName
		req.SetEnabled(true)
		set, err := i.build.QueryBuildConfig(ctx, req)
		if err != nil {
			return nil, err
		}

		matched := set.MatchBranch(in.GitlabEvent.GetBranch())
		for index := range matched.Items {
			// 执行构建配置匹配的流水线
			buildConf := matched.Items[index]
			bs := i.RunBuildConf(ctx, in, buildConf)
			ins.AddBuildStatus(bs)
		}
	}

	// 保存
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a deploy document error, %s", err)
	}
	return ins, nil
}

func (i *impl) RunBuildConf(ctx context.Context, in *trigger.Event, buildConf *build.BuildConfig) *trigger.BuildStatus {
	bs := trigger.NewBuildStatus(buildConf)

	pipelineId := buildConf.Spec.PipielineId()
	if pipelineId == "" {
		i.log.Debugf("构建配置: %s, 未配置流水线", buildConf.Spec.Name)
		return bs
	}

	runReq := pipeline.NewRunPipelineRequest(pipelineId)
	runReq.RunBy = "@" + in.UUID()
	runReq.TriggerMode = pipeline.TRIGGER_MODE_EVENT
	runReq.DryRun = in.SkipRunPipeline

	// 补充Build用户自定义变量
	runReq.AddRunParam(buildConf.BuildRunParams().Params...)

	// 补充构建时系统变量
	switch buildConf.Spec.TargetType {
	case build.TARGET_TYPE_IMAGE:
		ib := buildConf.Spec.ImageBuild
		// 注入Dockerfile位置信息
		runReq.AddRunParam(job.NewRunParam(
			build.SYSTEM_VARIABLE_APP_DOCKERFILE,
			ib.GetDockerFileWithDefault(build.DEFAULT_DOCKER_FILE_PATH),
		))
		// 注入推送代码仓库相关信息
		runReq.AddRunParam(job.NewRunParam(
			build.SYSTEM_VARIABLE_IMAGE_REPOSITORY,
			ib.GetImageRepositoryWithDefault(in.GitlabEvent.DefaultRepository()),
		))
	}

	// 补充Git信息
	runReq.AddRunParam(in.GitlabEvent.GitRunParams().Params...)
	// 补充版本信息
	switch buildConf.Spec.VersionNamedRule {
	case build.VERSION_NAMED_RULE_DATE_BRANCH_COMMIT:
		runReq.AddRunParam(in.GitlabEvent.DateCommitVersion(buildConf.Spec.VersionPrefix))
	case build.VERSION_NAMED_RULE_GIT_TAG:
		runReq.AddRunParam(in.GitlabEvent.TagVersion(buildConf.Spec.VersionPrefix))
	}

	i.log.Debugf("run pipeline req: %s", runReq.ToJson())
	pt, err := i.task.RunPipeline(ctx, runReq)
	if err != nil {
		bs.ErrorMessage = err.Error()
	} else {
		bs.PiplineTaskId = pt.Meta.Id
		bs.PiplineTask = pt
	}

	return bs
}

// 查询事件
func (i *impl) QueryRecord(ctx context.Context, in *trigger.QueryRecordRequest) (
	*trigger.RecordSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find event record error, error is %s", err)
	}

	set := trigger.NewRecordSet()
	// 循环
	for resp.Next(ctx) {
		ins := trigger.NewDefaultRecord()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode event record error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get event record count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
