module github.com/inflion/inflion/apiserver

go 1.15

require (
	cloud.google.com/go v0.51.0 // indirect
	github.com/blang/semver v3.5.0+incompatible // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/go-openapi/spec v0.19.9
	github.com/gogo/protobuf v1.3.1
	github.com/google/go-cmp v0.4.0
	github.com/google/gofuzz v1.1.0
	github.com/googleapis/gnostic v0.5.1
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/moby/term v0.0.0-20200312100748-672ec06f55cd // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/stretchr/testify v1.5.1
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200910180754-dd1b699fc489
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.27.0
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/apimachinery v0.19.2
	k8s.io/apiserver v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/component-base v0.19.2
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.3.0
	k8s.io/kube-openapi v0.0.0-20200923155610-8b5066479488
	k8s.io/utils v0.0.0-20201005171033-6301aaf42dc7
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20201003235837-18112a7b933b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20201003235655-10b38829b621
	k8s.io/client-go => k8s.io/client-go v0.0.0-20201004000108-758467711e07
	k8s.io/component-base => k8s.io/component-base v0.0.0-20201004000625-609bde980a40
)
