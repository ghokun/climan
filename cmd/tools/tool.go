package tools

type Tool interface {
	Name() string
	Type() string
	Latest() string
	Current() string
	Description() string
	List() []string
	Install() error
	Remove() error
}

type ToolType string

const (
	APIClient    ToolType = "Kubernetes API Client"
	Serverless   ToolType = "Serverless"
	Pipeline     ToolType = "Pipeline"
	Vendor       ToolType = "Vendor"
	Distribution ToolType = "Distribution"
)

var Tools = map[string]Tool{
	"climan": Climan{},
}
