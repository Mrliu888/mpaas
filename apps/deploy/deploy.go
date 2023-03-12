package deploy

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mpaas/apps/job"
	meta "github.com/infraboard/mpaas/common/meta"
	"github.com/infraboard/mpaas/provider/k8s/network"
	"github.com/infraboard/mpaas/provider/k8s/workload"
	v1 "k8s.io/api/core/v1"
)

func NewDeploymentSet() *DeploymentSet {
	return &DeploymentSet{
		Items: []*Deployment{},
	}
}

func (s *DeploymentSet) Add(item *Deployment) {
	s.Items = append(s.Items, item)
}

func NewDefaultDeploy() *Deployment {
	return &Deployment{
		Spec:       NewCreateDeploymentRequest(),
		Credential: NewCredential(),
		Status:     NewStatus(),
	}
}

func (d *Deployment) SetDefault() {
	if d.Status == nil {
		d.Status = NewStatus()
	}
}

func (d *Deployment) ValidateToken(token string) error {
	if d.Spec == nil {
		return nil
	}

	if !d.Spec.AuthEnabled {
		return nil
	}

	if d.Credential.Token != token {
		return fmt.Errorf("集群访问Token不合法")
	}

	return nil
}

func (d *Deployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateDeploymentRequest
		Status *Status `json:"status"`
	}{d.Meta, d.Spec, d.Status})
}

func (d *Deployment) GetK8sClusterId() string {
	if d.Spec == nil {
		return ""
	}
	if d.Spec.K8STypeConfig == nil {
		return ""
	}
	return d.Spec.K8STypeConfig.ClusterId
}

// 部署时的系统变量, 在部署任务时注入
func (d *Deployment) SystemVariable() ([]v1.EnvVar, error) {
	items := []*job.RunParam{}
	switch d.Spec.Type {
	case TYPE_KUBERNETES:
		wc := d.Spec.K8STypeConfig
		// 与k8s部署相关的系统变量, 不要注入版本信息, 部署版本由用户自己传人
		wl, err := wc.GetWorkLoad()
		if err != nil {
			return nil, err
		}
		variables := wl.SystemVaraible(d.Spec.ServiceName)
		addr, _ := variables.ImageDetail()
		items = append(items,
			job.NewRunParam(
				job.SYSTEM_VARIABLE_WORKLOAD_KIND,
				strings.ToLower(d.Spec.K8STypeConfig.WorkloadKind),
			).SetReadOnly(true).SetSearchLabel(true),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_WORKLOAD_NAME,
				variables.WorkloadName,
			).SetReadOnly(true).SetSearchLabel(true),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_SERVICE_NAME,
				d.Spec.ServiceName,
			).SetReadOnly(true).SetSearchLabel(true),
			job.NewRunParam(
				job.SYSTEM_VARIABLE_IMAGE_REPOSITORY,
				addr,
			),
		)
	}

	return job.ParamsToEnvVar(items), nil
}

func (c *K8STypeConfig) GetWorkLoad() (*workload.WorkLoad, error) {
	return workload.ParseWorkloadFromYaml(c.WorkloadKind, c.WorkloadConfig)
}

func (c *K8STypeConfig) Merge(target *K8STypeConfig) error {
	if target == nil {
		return nil
	}

	// 不能更新集群的clusterId
	target.ClusterId = c.ClusterId

	return mergo.MergeWithOverwrite(c, target)
}

func (c *K8STypeConfig) GetServiceObj() (*v1.Service, error) {
	if c.Service == "" {
		return nil, nil
	}
	return network.ParseServiceFromYaml(c.Service)
}

func NewCredential() *Credential {
	return &Credential{}
}

func NewStatus() *Status {
	return &Status{}
}

func (s *Status) MarkCreating() {
	s.Stage = STAGE_CREATING
}

func (s *Status) Update(target *Status) error {
	if target == nil {
		return fmt.Errorf("Status为nil")
	}

	if target.Stage <= STAGE_CREATING {
		return fmt.Errorf("更新的状态不合法, 不能更新为PENDDING或者CREATING")
	}

	target.UpdateAt = time.Now().Unix()
	s = target

	return nil
}

func (s *Status) UpdateK8sWorkloadStatus(status *workload.WorkloadStatus) {
	if status == nil {
		return
	}

	switch status.Stage {
	case workload.WORKLOAD_STAGE_PROGERESS:
		s.Stage = STAGE_UPGRADING
	case workload.WORKLOAD_STAGE_ACTIVE:
		s.Stage = STAGE_ACTIVE
	case workload.WORKLOAD_STAGE_ERROR:
		s.Stage = STAGE_ERROR
	}
	s.Reason = status.Reason
	s.Message = status.Message
	s.UpdateAt = time.Now().Unix()
}

func NewAccessAddressFromK8sService(svc *v1.Service) *AccessAddress {
	address := NewAccessAddress()
	for i := range svc.Spec.Ports {
		p := svc.Spec.Ports[i]
		name := fmt.Sprintf(
			"%s_PORT_%d_%s",
			strings.ToUpper(svc.Name),
			p.Port,
			strings.ToUpper(string(p.Protocol)),
		)
		address.AddServiceEnv(name, "REDIS_SERVICE_NAME_PORT_6379_TCP")
	}

	return &AccessAddress{}
}

func NewServiceEnv(name, example string) *ServiceEnv {
	return &ServiceEnv{
		Name:    name,
		Example: example,
	}
}

func NewAccessAddress() *AccessAddress {
	return &AccessAddress{
		ServiceEnv: []*ServiceEnv{},
	}
}

func (a *AccessAddress) AddServiceEnv(name, example string) {
	a.ServiceEnv = append(a.ServiceEnv, NewServiceEnv(name, example))
}
