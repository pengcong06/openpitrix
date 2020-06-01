// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// openpitrix all in one
package main

import (
	"openpitrix.io/openpitrix/pkg/apigateway"
	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/service/app"
	"openpitrix.io/openpitrix/pkg/service/attachment"
	"openpitrix.io/openpitrix/pkg/service/category"
	"openpitrix.io/openpitrix/pkg/service/helm"
	"openpitrix.io/openpitrix/pkg/service/isv"
	"openpitrix.io/openpitrix/pkg/service/repo"
	"openpitrix.io/openpitrix/pkg/service/repo_indexer"
	"openpitrix.io/openpitrix/pkg/service/runtime"
)

func getConf(database string) *config.Config {
	cfg := config.GetConf()
	cfg.Mysql.Database = database
	return cfg
}

func main() {
	go category.Serve(getConf("app"))
	go isv.Serve(getConf("isv"))
	go repo_indexer.Serve(getConf("repo"))
	go repo.Serve(getConf("repo"))
	go runtime.Serve(getConf("runtime"))
	go app.Serve(getConf("app"))
	go attachment.Serve(getConf("attachment"))
	go helm.Serve(getConf(""))
	//todo replace helm.Serve
	go helm.HelmServe(getConf(""))

	apigateway.Serve(getConf(""))
}
