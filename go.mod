module openpitrix.io/openpitrix

go 1.13

require k8s.io/helm v2.14.3+incompatible

require k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible

require k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/semver v1.5.0
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/aws/aws-sdk-go v1.25.21
	github.com/bitly/go-simplejson v0.5.0
	github.com/chai2010/jsonmap v1.0.0
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/disintegration/imaging v1.6.1
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/structs v1.1.0
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/runtime v0.19.7
	github.com/go-openapi/spec v0.19.4
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.4
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocraft/dbr v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.2
	github.com/google/gops v0.3.6
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/robfig/cron v1.2.0
	github.com/sony/sonyflake v1.0.0
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli v1.22.1
	github.com/xeipuuv/gojsonschema v1.2.0
	go.etcd.io/etcd v3.3.17+incompatible
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/tools v0.0.0-20191029041327-9cc4af7d6b2c
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc v1.24.0
	gopkg.in/square/go-jose.v2 v2.4.0
	gopkg.in/yaml.v2 v2.2.5
	kubesphere.io/im v0.1.0
	openpitrix.io/iam v0.1.0
	openpitrix.io/notification v0.2.2
)

replace github.com/ugorji/go => github.com/ugorji/go v0.0.0-20190128213124-ee1426cffec0

require k8s.io/kubernetes v1.14.1

require k8s.io/cli-runtime v0.0.0-20190409023024-d644b00f3b79 // indirect

require (
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/chai2010/gettext-go v0.0.0-20191225085308-6b9f4b1008e1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/huandu/xstrings v1.3.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/rubenv/sql-migrate v0.0.0-20200212082348-64f95ea68aa3 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-00010101000000-000000000000
	k8s.io/apiextensions-apiserver v0.0.0-00010101000000-000000000000
	k8s.io/apiserver v0.0.0-00010101000000-000000000000 // indirect
	k8s.io/cloud-provider v0.0.0-20190409023720-1bc0c81fa51d // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200204173128-addea2498afe // indirect
	k8s.io/utils v0.0.0-20200229041039-0a110f9eb7ab // indirect
	sigs.k8s.io/kustomize v2.0.3+incompatible // indirect
	vbom.ml/util v0.0.0-20180919145318-efcd4e0f9787 // indirect
)

replace github.com/gocraft/dbr => github.com/gocraft/dbr v0.0.0-20180507214907-a0fd650918f6

replace github.com/docker/docker => github.com/docker/engine v0.0.0-20190423201726-d2cfbce3f3b0

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190409021813-1ec86e4da56c

replace k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
