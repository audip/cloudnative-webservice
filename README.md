## Web service based on cloud native tools

## User Stories

### Required Stories
- [ ] webservice should be written in [Golang](https://golang.org/).
- [ ] webservice should emit metrics for [prometheus](https://prometheus.io/) to scrape.
- [ ] webservice should emit logs that are shipped by [fluentd](https://www.fluentd.org/) to [elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html).
- [ ] webservice should have tracing (opentracing API) built in, which can be viewed in [Jaeger](http://jaeger.readthedocs.io/en/latest/).
- [ ] webservice should have REST endpoints for other services to connect to.
- [ ] webservice should have CI (Travis CI) and CD (docker hub).
- [ ] webservice should be dockerized and use multi-stage builds to copy go binary.

### Optional Stories
- [ ] webservice should have good code coverage (tests), [goreport](https://goreportcard.com/) of A or above.
- [ ] webservice should store key/values in [etcd](https://coreos.com/etcd/docs/latest/).
- [ ] webservice should have a [GRPC](https://grpc.io/) endpoint.
- [ ] webservice should have a [Vitess](http://vitess.io/) or [CockroachDB](https://github.com/cockroachdb/cockroach) as a database.
- [ ] webservice should integrate the update framework using [notary](https://github.com/theupdateframework/notary)
- [ ] webservice should use [bazel](https://bazel.build/) as a code build system, instead of make
