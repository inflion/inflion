package context

import "strings"

type ExecutionContext struct {
	ExecutionFields ExecutionFields
}

func NewExecutionContext() ExecutionContext {
	return ExecutionContext{
		ExecutionFields: ExecutionFields{
			Fields: map[string]ExecutionFields{},
			Values: map[string]interface{}{},
		},
	}
}

func NewExecutionContextWithFields(fields map[string]ExecutionFields) ExecutionContext {
	return ExecutionContext{
		ExecutionFields: ExecutionFields{
			Fields: fields,
			Values: map[string]interface{}{},
		},
	}
}

func (e ExecutionContext) AddFields(key string, fields ExecutionFields) ExecutionContext {
	e.ExecutionFields.Fields[key] = fields
	return e
}

func (e ExecutionContext) GetValueByPath(path Path) interface{} {
	return e.ExecutionFields.getRecursively(path.toArray())
}

type ExecutionFields struct {
	Fields map[string]ExecutionFields
	Values map[string]interface{}
}

func (e ExecutionFields) getFields(key string) ExecutionFields {
	if v, ok := e.Fields[key]; ok {
		return v
	} else {
		return ExecutionFields{}
	}
}

func (e ExecutionFields) getRecursively(path []string) string {
	if len(path) == 1 {
		if v, ok := e.Values[path[0]]; ok {
			return v.(string)
		} else {
			return ""
		}
	} else {
		first, tail := path[0], path[1:]
		return e.getFields(first).getRecursively(tail)
	}
}

type Path struct {
	Path string
}

func NewPath(path string) Path {
	return Path{Path: path}
}

func (p Path) toArray() []string {
	return strings.Split(p.Path, ".")
}

func (p Path) getFirst() string {
	return p.toArray()[0]
}

func (p Path) getTail() []string {
	return p.toArray()[1:]
}

func (p Path) getFirstAndTail() (string, []string) {
	path := p.toArray()
	return path[0], path[1:]
}
