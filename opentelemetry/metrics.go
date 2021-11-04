package opentelemetry

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

type ChainbridgeMetrics struct {
	DepositEventCount metric.Int64Counter
}

func newChainbridgeMetrics(meter metric.Meter) *ChainbridgeMetrics {
	return &ChainbridgeMetrics{
		DepositEventCount: metric.Must(meter).NewInt64Counter(
			"chainbridge.DepositEventCount",
			metric.WithDescription("Number of deposit events across all chains"),
		),
	}
}

func initPrometheusMetrics(addr string) (*ChainbridgeMetrics, error) {
	config := prometheus.Config{}
	c := controller.New(
		processor.NewFactory(
			selector.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(config.DefaultHistogramBoundaries),
			),
			export.CumulativeExportKindSelector(),
			processor.WithMemory(true),
		),
	)
	exp, err := prometheus.New(config, c)
	if err != nil {
		return nil, err
	}

	global.SetMeterProvider(exp.MeterProvider())
	go func() {
		http.HandleFunc(addr, exp.ServeHTTP)
		_ = http.ListenAndServe(addr, nil)
	}()

	meter := c.Meter("")
	return newChainbridgeMetrics(meter), nil
}

func initOpenTelemetryMetrics(opts ...otlpmetrichttp.Option) (*ChainbridgeMetrics, error) {
	ctx := context.Background()

	client := otlpmetrichttp.NewClient(opts...)
	exp, err := otlpmetric.New(ctx, client)
	if err != nil {
		return nil, err
	}

	selector := simple.NewWithInexpensiveDistribution()
	proc := processor.NewFactory(selector, export.CumulativeExportKindSelector())
	cont := controller.New(proc, controller.WithExporter(exp))
	global.SetMeterProvider(cont)

	err = cont.Start(ctx)
	if err != nil {
		return nil, err
	}

	meter := cont.Meter("")
	metrics := newChainbridgeMetrics(meter)

	return metrics, nil
}