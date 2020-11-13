package api

// AuthenticationModes represents auth modes supported by dashboard.
type AuthenticationModes map[AuthenticationMode]bool

// ProtectedResource represents basic information about resource that should be filtered out from Dashboard UI.
type ProtectedResource struct {
	// ResourceName is a name of the protected resource.
	ResourceName string
	// ResourceNamespace is a namespace of the protected resource. Should be empty if resource is non-namespaced.
	ResourceNamespace string
}

// IsEnabled returns true if given auth mode is supported, false otherwise.
func (self AuthenticationModes) IsEnabled(mode AuthenticationMode) bool {
	_, exists := self[mode]
	return exists
}

// Array returns array of auth modes supported by dashboard.
func (self AuthenticationModes) Array() []AuthenticationMode {
	modes := []AuthenticationMode{}
	for mode := range self {
		modes = append(modes, mode)
	}

	return modes
}

// Add adds given auth mode to AuthenticationModes map
func (self AuthenticationModes) Add(mode AuthenticationMode) {
	self[mode] = true
}

// AuthenticationMode represents auth mode supported by dashboard, i.e. basic.
type AuthenticationMode string

// String returns string representation of auth mode.
func (self AuthenticationMode) String() string {
	return string(self)
}

// Authentication modes supported by dashboard should be defined below.
const (
	Token AuthenticationMode = "token"
	Basic AuthenticationMode = "basic"
)