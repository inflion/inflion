package store

import (
	"github.com/google/uuid"
	"github.com/inflion/inflion/internal/ops/flow"
)

type FlowData struct {
	Project string
	Id      uuid.UUID
	Body    string
}

type Project = string
type FlowId = uuid.UUID
type FlowBody = string

type FlowCreateRequest struct {
	Project Project
	Body    FlowBody
}

type FlowCreateResponse struct {
	Id FlowId
}

type FlowGetRequest struct {
	Project Project
	Id      FlowId
}

type FlowGetResponse struct {
	Body FlowBody
}

type FlowUpdateRequest struct {
	Project Project
	Id      FlowId
	Body    FlowBody
}

type FlowUpdateResponse struct {
	Id FlowId
}

type FlowDeleteRequest struct {
	Project Project
	Id      FlowId
}

type FlowDeleteResponse struct {
	Id FlowId
}

type Store interface {
	List(project string) ([]FlowData, error)
	Create(request FlowCreateRequest) (FlowCreateResponse, error)
	Get(request FlowGetRequest) (FlowGetResponse, error)
	Update(request FlowUpdateRequest) error
	Delete(request FlowDeleteRequest) error
}

type StoreFlowReader struct {
	project string
	id      FlowId
	store   Store
}

func (s StoreFlowReader) Read() (flow.Flow, error) {
	f, err := s.store.Get(FlowGetRequest{
		Project: s.project,
		Id:      s.id,
	})
	if err != nil {
		return flow.Flow{}, err
	}

	r, err := flow.Unmarshal([]byte(f.Body))
	if err != nil {
		return flow.Flow{}, err
	}

	return r, nil
}
