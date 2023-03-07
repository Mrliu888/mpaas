package build

import (
	"encoding/json"
	"regexp"

	"github.com/infraboard/mpaas/common/meta"
)

// New 新建一个domain
func New(req *CreateBuildConfigRequest) (*BuildConfig, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &BuildConfig{
		Meta: meta.NewMeta(),
		Spec: req,
	}

	return d, nil
}

func NewBuildConfigSet() *BuildConfigSet {
	return &BuildConfigSet{
		Items: []*BuildConfig{},
	}
}

func (s *BuildConfigSet) Add(item *BuildConfig) {
	s.Items = append(s.Items, item)
}

// 比如: "foo.*"
func (s *BuildConfigSet) MatchBranch(branchRegExp string) *BuildConfigSet {
	set := NewBuildConfigSet()

	for i := range s.Items {
		item := s.Items[i]
		if item.Spec.Condition.MatchBranch(branchRegExp) {
			set.Add(item)
		}
	}

	return set
}

func NewDefaultBuildConfig() *BuildConfig {
	return &BuildConfig{
		Spec: NewCreateBuildConfigRequest(),
	}
}

func (b *BuildConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateBuildConfigRequest
	}{b.Meta, b.Spec})
}

func NewCreateBuildConfigRequest() *CreateBuildConfigRequest {
	return &CreateBuildConfigRequest{
		VersionPrefix: "v",
		Condition:     NewTrigger(),
		ImageBuild:    NewImageBuild(),
		PkgBuild:      NewPkgBuildConfig(),
		Labels:        make(map[string]string),
	}
}

func NewImageBuild() *ImageBuildConfig {
	return &ImageBuildConfig{
		BuildEnvVars: make(map[string]string),
		Extra:        make(map[string]string),
	}
}

func NewPkgBuildConfig() *PkgBuildConfig {
	return &PkgBuildConfig{
		Extra: make(map[string]string),
	}
}

func NewTrigger() *Trigger {
	return &Trigger{
		Events:   []string{},
		Branches: []string{},
	}
}

func (t *Trigger) AddEvent(event string) {
	t.Events = append(t.Events, event)
}

func (t *Trigger) AddBranche(branche string) {
	t.Branches = append(t.Branches, branche)
}

func (t *Trigger) MatchBranch(branchRegExp string) bool {
	for _, b := range t.Branches {
		ok, _ := regexp.MatchString(branchRegExp, b)
		if ok {
			return true
		}
	}

	return false
}
