package arm

//go:generate go get github.com/go-bindata/go-bindata/go-bindata
//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix data data/...
//go:generate gofmt -s -l -w bindata.go
//go:generate go get github.com/golang/mock/gomock
//go:generate go install github.com/golang/mock/mockgen
//go:generate mockgen -destination=../util/mocks/mock_$GOPACKAGE/arm.go -package=mock_$GOPACKAGE -source arm.go
//go:generate gofmt -s -l -w ../util/mocks/mock_$GOPACKAGE/arm.go
//go:generate goimports -local=github.com/openshift/openshift-azure -e -w ../util/mocks/mock_$GOPACKAGE/arm.go

import (
	"context"
	"encoding/json"
	"text/template"

	"github.com/openshift/openshift-azure/pkg/api"
	util "github.com/openshift/openshift-azure/pkg/util/template"
)

type Generator interface {
	Generate(ctx context.Context, cs *api.OpenShiftManagedCluster, isUpdate bool, backupBlob string) (map[string]interface{}, error)
}

type simpleGenerator struct {
	pluginConfig api.PluginConfig
}

var _ Generator = &simpleGenerator{}

// NewSimpleGenerator create a new SimpleGenerator
func NewSimpleGenerator(pluginConfig *api.PluginConfig) Generator {
	return &simpleGenerator{
		pluginConfig: *pluginConfig,
	}
}

func (g *simpleGenerator) Generate(ctx context.Context, cs *api.OpenShiftManagedCluster, isUpdate bool, backupBlob string) (map[string]interface{}, error) {
	masterStartup, err := Asset("master-startup.sh")
	if err != nil {
		return nil, err
	}

	nodeStartup, err := Asset("node-startup.sh")
	if err != nil {
		return nil, err
	}

	tmpl, err := Asset("azuredeploy.json")
	if err != nil {
		return nil, err
	}
	azuredeploy, err := util.Template(string(tmpl), template.FuncMap{
		"Startup": func(role api.AgentPoolProfileRole) ([]byte, error) {
			if role == api.AgentPoolProfileRoleMaster {
				return util.Template(string(masterStartup), nil, cs, map[string]interface{}{
					"IsRecovery":     len(backupBlob) > 0,
					"BackupBlobName": backupBlob,
					"Role":           role,
					"TestConfig":     g.pluginConfig.TestConfig,
				})
			}
			return util.Template(string(nodeStartup), nil, cs, map[string]interface{}{
				"Role":       role,
				"TestConfig": g.pluginConfig.TestConfig,
			})
		},
		"IsUpgrade": func() bool {
			return isUpdate
		},
	}, cs, map[string]interface{}{
		"TestConfig": g.pluginConfig.TestConfig,
	})
	if err != nil {
		return nil, err
	}

	var azuretemplate map[string]interface{}
	err = json.Unmarshal(azuredeploy, &azuretemplate)
	if err != nil {
		return nil, err
	}
	return azuretemplate, nil
}
