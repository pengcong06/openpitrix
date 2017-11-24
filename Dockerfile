# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:alpine as builder

WORKDIR /go/src/openpitrix.io/openpitrix/
COPY . .
RUN go install ./cmd/...

RUN mv /go/bin/openpitrix-api     /go/bin/api
RUN mv /go/bin/openpitrix-app     /go/bin/app
RUN mv /go/bin/openpitrix-cluster /go/bin/cluster
RUN mv /go/bin/openpitrix-repo    /go/bin/repo
RUN mv /go/bin/openpitrix-runtime /go/bin/runtime

FROM alpine

COPY --from=builder /go/bin/api /usr/local/bin/
COPY --from=builder /go/bin/app /usr/local/bin/
COPY --from=builder /go/bin/cluster /usr/local/bin/
COPY --from=builder /go/bin/repo /usr/local/bin/
COPY --from=builder /go/bin/runtime /usr/local/bin/

# Glog
ENV OPENPITRIX_GLOG_LOGTOSTDERR=false
ENV OPENPITRIX_GLOG_ALSOLOGTOSTDERR=false
ENV OPENPITRIX_GLOG_STDERRTHRESHOLD=ERROR
ENV OPENPITRIX_GLOG_LOGDIR=

ENV OPENPITRIX_GLOG_LOGBACKTRACEAT=
ENV OPENPITRIX_GLOG_V=0
ENV OPENPITRIX_GLOG_VMODULE=

ENV OPENPITRIX_GLOG_COPYSTANDARDLOGTO=INFO

# database
ENV OPENPITRIX_DB_DBNAME=openpitrix
ENV OPENPITRIX_DB_ENCODING=utf8
ENV OPENPITRIX_DB_ENGINE=InnoDB
ENV OPENPITRIX_DB_HOST=openpitrix-mysql
ENV OPENPITRIX_DB_PORT=3306
ENV OPENPITRIX_DB_ROOTPASSWORD=password
ENV OPENPITRIX_DB_TYPE=mysql

# api service
ENV OPENPITRIX_API_HOST=openpitrix-api
ENV OPENPITRIX_API_PORT=8080

# app service
ENV OPENPITRIX_APP_HOST=openpitrix-app
ENV OPENPITRIX_APP_PORT=8081

# runtime service
ENV OPENPITRIX_RUNTIME_HOST=openpitrix-runtime
ENV OPENPITRIX_RUNTIME_PORT=8082

# cluster service
ENV OPENPITRIX_CLUSTER_HOST=openpitrix-cluster
ENV OPENPITRIX_CLUSTER_PORT=8083

# repo service
ENV OPENPITRIX_REPO_HOST=openpitrix-repo
ENV OPENPITRIX_REPO_PORT=8084

CMD ["sh"]