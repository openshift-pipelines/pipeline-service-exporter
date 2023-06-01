/*
 Copyright 2023 The Pipeline Service Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/openshift-pipelines/pipeline-service-exporter/collector"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
	_ "net/http/pprof"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
)

var (
	listenAddress = kingpin.Flag("telemetry.address", "Address at which pipeline-service metrics are exported.").Default(":9117").String()
	metricsPath   = kingpin.Flag("telemetry-path", "Path at which pipeline-service metrics are exported.").Default("/metrics").String()
	probeAddr     = kingpin.Flag("health-probe-bind-address", "The address the probe endpoint binds to.").Default(":8081").String()
	toolkitFlags  = kingpinflag.AddFlags(kingpin.CommandLine, ":9117")
	logger        log.Logger
	promlogConfig *promlog.Config
)

const (
	exporterName = "pipeline_service_exporter"
)

func init() {
	promlogConfig = &promlog.Config{}
	logger = promlog.New(promlogConfig)
}

func main() {

	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print(exporterName))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	level.Info(logger).Log("msg", "Starting pipeline_service_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build", version.BuildContext())
	level.Info(logger).Log("msg", "Starting Server: ", "listen_address", *listenAddress)

	ctx := ctrl.SetupSignalHandler()
	restConfig := ctrl.GetConfigOrDie()
	var mgr ctrl.Manager
	var err error
	mopts := ctrl.Options{
		MetricsBindAddress:     *listenAddress,
		Port:                   9443,
		HealthProbeBindAddress: *probeAddr,
	}

	mgr, err = collector.NewManager(restConfig, mopts, logger)
	if err != nil {
		level.Error(logger).Log("msg", "unable to start manager", "error", err)
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		level.Error(logger).Log("msg", "unable to set up health check", "error", err)
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		level.Error(logger).Log("msg", "unable ot set up ready check", "error", err)
		os.Exit(1)
	}

	level.Info(logger).Log("msg", "starting manager")

	if err := mgr.Start(ctx); err != nil {
		level.Error(logger).Log("msg", "problem running manager", "error", err)
		os.Exit(1)
	}

}
