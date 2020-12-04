package configstore

type Config struct {
	Key   string
	Value string
}

type Project = string
type ConfigKey = string
type ConfigValue = string

type ConfigCreateOrUpdateRequest struct {
	Project Project
	Config  Config
}

type ConfigCreateOrUpdateResponse struct {
	Config Config
}

type ConfigListRequest struct {
	Project Project
	Key     string
}

type ConfigListResponse struct {
	Configs []Config
}

type ConfigGetRequest struct {
	Project Project
	Key     ConfigKey
}

type ConfigGetResponse struct {
	Config Config
}

type ConfigDeleteRequest struct {
	Project Project
	Key     ConfigKey
}

type FlowDeleteResponse struct {
	Config Config
}

type ConfigStore interface {
	CreateOrUpdate(request ConfigCreateOrUpdateRequest) (ConfigCreateOrUpdateResponse, error)
	Get(request ConfigGetRequest) (ConfigGetResponse, error)
	List(request ConfigListRequest) (ConfigListResponse, error)
	Delete(request ConfigDeleteRequest) error
}
