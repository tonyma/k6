module github.com/loadimpact/k6

go 1.14

require (
	github.com/Azure/go-ntlmssp v0.0.0-20180810175552-4a21cbd618b4
	github.com/DataDog/datadog-go v0.0.0-20180330214955-e67964b4021a
	github.com/GeertJohan/go.rice v0.0.0-20170420135705-c02ca9a983da
	github.com/PuerkitoBio/goquery v1.3.0
	github.com/Soontao/goHttpDigestClient v0.0.0-20170320082612-6d28bb1415c5
	github.com/andybalholm/brotli v0.0.0-20190704151324-71eb68cc467c
	github.com/dop251/goja v0.0.0-20210216182323-60bc6ebb9fc1
	github.com/dustin/go-humanize v0.0.0-20171111073723-bb3d318650d4
	github.com/fatih/color v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/websocket v1.4.2
	github.com/influxdata/influxdb1-client v0.0.0-20190402204710-8ff2fc3824fc
	github.com/jhump/protoreflect v1.7.0
	github.com/julienschmidt/httprouter v1.1.1-0.20180222160526-d18983907793
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/klauspost/compress v1.7.2
	github.com/kubernetes/helm v2.9.0+incompatible
	github.com/mailru/easyjson v0.7.4-0.20200812114229-8ab5ff9cd8e4
	github.com/manyminds/api2go v0.0.0-20180125085803-95be7bd0455e
	github.com/mattn/go-colorable v0.0.9
	github.com/mattn/go-isatty v0.0.4
	github.com/mccutchen/go-httpbin v1.1.2-0.20190116014521-c5cb2f4802fa
	github.com/mitchellh/mapstructure v0.0.0-20180220230111-00c29f56e238
	github.com/mstoykov/xk6-kafka-output v0.0.0-20210318151243-3cda7bd77727
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/oxtoacart/bpool v0.0.0-20150712133111-4e1c5567d7c2
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.1.1
	github.com/spf13/cobra v0.0.4-0.20180629152535-a114f312e075
	github.com/spf13/pflag v1.0.1
	github.com/stretchr/testify v1.2.2
	github.com/tidwall/gjson v1.6.1
	github.com/tidwall/pretty v1.0.2
	github.com/urfave/negroni v0.3.1-0.20180130044549-22c5532ea862
	github.com/zyedidia/highlight v0.0.0-20170330143449-201131ce5cf5
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20201008223702-a5fa9d4b7c91
	golang.org/x/text v0.3.3
	golang.org/x/time v0.0.0-20170927054726-6dc17368e09b
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/guregu/null.v3 v3.3.0
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.0
	github.com/stretchr/testify => github.com/stretchr/testify v1.2.1
	github.com/ugorji/go => github.com/ugorji/go v0.0.0-20180112141927-9831f2c3ac10
	golang.org/x/text => golang.org/x/text v0.3.0
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.1.1
)
