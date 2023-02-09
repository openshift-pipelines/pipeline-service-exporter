package main

import (
	"github.com/bnallapeta/pipeline-service-exporter/collector"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	"github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	listenAddress = kingpin.Flag("telemetry.address", "Address at which pipeline-service metrics are exported.").Default(".9117").String()
	metricsPath   = kingpin.Flag("telemetry-path", "Path at which pipeline-service metrics are exported.").Default("/metrics").String()
	toolkitFlags  = kingpinflag.AddFlags(kingpin.CommandLine, ":9117")
)

const (
	exporterName = "pipeline_service_exporter"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print(exporterName))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting pipeline_service_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build", version.BuildContext())
	level.Info(logger).Log("msg", "Starting Server: ", "listen_address", *listenAddress)

	prometheus.MustRegister(version.NewCollector(exporterName))

	psCollector, err := collector.NewCollector(logger)
	if err != nil {
		level.Error(logger).Log("msg", "Couldn't create collector", "error", err)
		os.Exit(1)
	}

	prometheus.MustRegister(psCollector)

	// Define a channel to watch out for any termination signals
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	// Listen for the termination signals from the OS
	go func() {
		level.Info(logger).Log("msg", "Listening and waiting for graceful stop")
		sig := <-gracefulStop
		level.Info(logger).Log("msg", "Caught signal: %+v. Waiting 2 seconds...", "signal", sig)
		time.Sleep(2 * time.Second)
		level.Info(logger).Log("msg", "Terminating pipeline_service_exporter on port: ", "listen_address", *listenAddress)
		os.Exit(0)
	}()

	// Define paths
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Pipeline Exporter</title></head>
			<body>
			<h1>Pipeline Exporter</h1>
			<p><a href='` + *metricsPath + `'>Metrics</a></p>
			</body>
			</html>`))
	})

	// Start the server
	srv := &http.Server{Addr: *listenAddress}
	if err := web.ListenAndServe(srv, toolkitFlags, logger); err != nil {
		level.Error(logger).Log("error", "Port Listen Address error", "reason", err)
		os.Exit(1)
	}
}
