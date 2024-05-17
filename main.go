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
	"flag"

	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	_ "net/http/pprof"
	"os"

	"github.com/go-logr/logr"
	"github.com/openshift-pipelines/pipeline-service-exporter/collector"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/version"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
)

var (
	mainLog       logr.Logger
	promlogConfig *promlog.Config
)

func init() {
	promlogConfig = &promlog.Config{}
}

func main() {
	var listenAddress string
	var metricsPath string
	var probeAddr string
	var pprofAddr string

	flag.StringVar(&listenAddress, "telemetry.address", ":9117", "Address at which pipeline-service metrics are exported.")
	flag.StringVar(&metricsPath, "telemetry-path", "/metrics", "Path at which pipeline-service metrics are exported.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.StringVar(&pprofAddr, "pprof-address", "", "The address the pprof endpoint binds to.")

	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	klog.InitFlags(flag.CommandLine)
	flag.Parse()

	/*
			FYI tracing set set with this zap argument on the deployment (see https://sdk.operatorframework.io/docs/building-operators/golang/references/logging/)
		          args:
		            - -zap-log-level=6
	*/

	logger := zap.New(zap.UseFlagOptions(&opts))
	ctrl.SetLogger(logger)
	mainLog = ctrl.Log.WithName("main")

	mainLog.Info("Starting pipeline_service_exporter", "version", version.Info())
	mainLog.Info("Build context", "build", version.BuildContext())
	mainLog.Info("Starting Server: ", "listen_address", listenAddress)

	ctx := ctrl.SetupSignalHandler()
	restConfig := ctrl.GetConfigOrDie()
	restConfig.QPS = 50
	restConfig.Burst = 50
	var mgr ctrl.Manager
	var err error
	mopts := ctrl.Options{
		MetricsBindAddress:     listenAddress,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
	}

	mgr, err = collector.NewManager(restConfig, mopts, pprofAddr)
	if err != nil {
		mainLog.Error(err, "unable to start controller-runtime manager")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err = mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		mainLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err = mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		mainLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	mainLog.Info("Starting controller-runtime manager")

	if err = mgr.Start(ctx); err != nil {
		mainLog.Error(err, "problem running controller-runtime manager")
		os.Exit(1)
	}

}
