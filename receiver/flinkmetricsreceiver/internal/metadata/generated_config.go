// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for flinkmetrics metrics.
type MetricsConfig struct {
	FlinkJobCheckpointCount           MetricConfig `mapstructure:"flink.job.checkpoint.count"`
	FlinkJobCheckpointInProgress      MetricConfig `mapstructure:"flink.job.checkpoint.in_progress"`
	FlinkJobLastCheckpointSize        MetricConfig `mapstructure:"flink.job.last_checkpoint.size"`
	FlinkJobLastCheckpointTime        MetricConfig `mapstructure:"flink.job.last_checkpoint.time"`
	FlinkJobRestartCount              MetricConfig `mapstructure:"flink.job.restart.count"`
	FlinkJvmClassLoaderClassesLoaded  MetricConfig `mapstructure:"flink.jvm.class_loader.classes_loaded"`
	FlinkJvmCPULoad                   MetricConfig `mapstructure:"flink.jvm.cpu.load"`
	FlinkJvmCPUTime                   MetricConfig `mapstructure:"flink.jvm.cpu.time"`
	FlinkJvmGcCollectionsCount        MetricConfig `mapstructure:"flink.jvm.gc.collections.count"`
	FlinkJvmGcCollectionsTime         MetricConfig `mapstructure:"flink.jvm.gc.collections.time"`
	FlinkJvmMemoryDirectTotalCapacity MetricConfig `mapstructure:"flink.jvm.memory.direct.total_capacity"`
	FlinkJvmMemoryDirectUsed          MetricConfig `mapstructure:"flink.jvm.memory.direct.used"`
	FlinkJvmMemoryHeapCommitted       MetricConfig `mapstructure:"flink.jvm.memory.heap.committed"`
	FlinkJvmMemoryHeapMax             MetricConfig `mapstructure:"flink.jvm.memory.heap.max"`
	FlinkJvmMemoryHeapUsed            MetricConfig `mapstructure:"flink.jvm.memory.heap.used"`
	FlinkJvmMemoryMappedTotalCapacity MetricConfig `mapstructure:"flink.jvm.memory.mapped.total_capacity"`
	FlinkJvmMemoryMappedUsed          MetricConfig `mapstructure:"flink.jvm.memory.mapped.used"`
	FlinkJvmMemoryMetaspaceCommitted  MetricConfig `mapstructure:"flink.jvm.memory.metaspace.committed"`
	FlinkJvmMemoryMetaspaceMax        MetricConfig `mapstructure:"flink.jvm.memory.metaspace.max"`
	FlinkJvmMemoryMetaspaceUsed       MetricConfig `mapstructure:"flink.jvm.memory.metaspace.used"`
	FlinkJvmMemoryNonheapCommitted    MetricConfig `mapstructure:"flink.jvm.memory.nonheap.committed"`
	FlinkJvmMemoryNonheapMax          MetricConfig `mapstructure:"flink.jvm.memory.nonheap.max"`
	FlinkJvmMemoryNonheapUsed         MetricConfig `mapstructure:"flink.jvm.memory.nonheap.used"`
	FlinkJvmThreadsCount              MetricConfig `mapstructure:"flink.jvm.threads.count"`
	FlinkMemoryManagedTotal           MetricConfig `mapstructure:"flink.memory.managed.total"`
	FlinkMemoryManagedUsed            MetricConfig `mapstructure:"flink.memory.managed.used"`
	FlinkOperatorRecordCount          MetricConfig `mapstructure:"flink.operator.record.count"`
	FlinkOperatorWatermarkOutput      MetricConfig `mapstructure:"flink.operator.watermark.output"`
	FlinkTaskRecordCount              MetricConfig `mapstructure:"flink.task.record.count"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		FlinkJobCheckpointCount: MetricConfig{
			Enabled: true,
		},
		FlinkJobCheckpointInProgress: MetricConfig{
			Enabled: true,
		},
		FlinkJobLastCheckpointSize: MetricConfig{
			Enabled: true,
		},
		FlinkJobLastCheckpointTime: MetricConfig{
			Enabled: true,
		},
		FlinkJobRestartCount: MetricConfig{
			Enabled: true,
		},
		FlinkJvmClassLoaderClassesLoaded: MetricConfig{
			Enabled: true,
		},
		FlinkJvmCPULoad: MetricConfig{
			Enabled: true,
		},
		FlinkJvmCPUTime: MetricConfig{
			Enabled: true,
		},
		FlinkJvmGcCollectionsCount: MetricConfig{
			Enabled: true,
		},
		FlinkJvmGcCollectionsTime: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryDirectTotalCapacity: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryDirectUsed: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryHeapCommitted: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryHeapMax: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryHeapUsed: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryMappedTotalCapacity: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryMappedUsed: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryMetaspaceCommitted: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryMetaspaceMax: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryMetaspaceUsed: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryNonheapCommitted: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryNonheapMax: MetricConfig{
			Enabled: true,
		},
		FlinkJvmMemoryNonheapUsed: MetricConfig{
			Enabled: true,
		},
		FlinkJvmThreadsCount: MetricConfig{
			Enabled: true,
		},
		FlinkMemoryManagedTotal: MetricConfig{
			Enabled: true,
		},
		FlinkMemoryManagedUsed: MetricConfig{
			Enabled: true,
		},
		FlinkOperatorRecordCount: MetricConfig{
			Enabled: true,
		},
		FlinkOperatorWatermarkOutput: MetricConfig{
			Enabled: true,
		},
		FlinkTaskRecordCount: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for flinkmetrics resource attributes.
type ResourceAttributesConfig struct {
	FlinkJobName       ResourceAttributeConfig `mapstructure:"flink.job.name"`
	FlinkResourceType  ResourceAttributeConfig `mapstructure:"flink.resource.type"`
	FlinkSubtaskIndex  ResourceAttributeConfig `mapstructure:"flink.subtask.index"`
	FlinkTaskName      ResourceAttributeConfig `mapstructure:"flink.task.name"`
	FlinkTaskmanagerID ResourceAttributeConfig `mapstructure:"flink.taskmanager.id"`
	HostName           ResourceAttributeConfig `mapstructure:"host.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		FlinkJobName: ResourceAttributeConfig{
			Enabled: true,
		},
		FlinkResourceType: ResourceAttributeConfig{
			Enabled: true,
		},
		FlinkSubtaskIndex: ResourceAttributeConfig{
			Enabled: true,
		},
		FlinkTaskName: ResourceAttributeConfig{
			Enabled: true,
		},
		FlinkTaskmanagerID: ResourceAttributeConfig{
			Enabled: true,
		},
		HostName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for flinkmetrics metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
