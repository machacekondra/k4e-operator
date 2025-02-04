package indexer

import (
	"strings"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/project-flotta/flotta-operator/api/v1alpha1"
	"github.com/project-flotta/flotta-operator/internal/labels"
	flottalabels "github.com/project-flotta/flotta-operator/internal/labels"
)

const (
	// WorkloadByDeviceIndexKey is the name of the indexer for workloads by device
	WorkloadByDeviceIndexKey = "workload-by-device"

	// DeviceByWorkloadIndexKey is the key of the indexer for devices by workload
	DeviceByWorkloadIndexKey = "device-by-workload"
)

func WorkloadByDeviceIndexFunc(obj ctrlruntimeclient.Object) []string {
	workload, ok := obj.(*v1alpha1.EdgeWorkload)
	if !ok {
		return []string{}
	}
	var keys []string
	for name, value := range workload.Labels {
		if flottalabels.IsSelectorLabel(name) {
			keys = append(keys, CreateWorkloadIndexKey(name, value))
		}
	}
	return keys
}

func DeviceByWorkloadIndexFunc(obj ctrlruntimeclient.Object) []string {
	device, ok := obj.(*v1alpha1.EdgeDevice)
	if !ok {
		return []string{}
	}
	var keys []string
	// iterate over labels map and return a list of workloads as keys
	for name := range device.Labels {
		if flottalabels.IsWorkloadLabel(name) {
			keys = append(keys, CreateDeviceIndexKey(name))
		}
	}
	return keys
}

// CreateDeviceIndexKey creates a key for the device index which is basically the workload name
func CreateDeviceIndexKey(label string) string {
	return strings.TrimPrefix(label, labels.WorkloadLabelPrefix)
}

// CreateWorkloadIndexKey creates a key for the workload index
// The key is of the form: device or label
func CreateWorkloadIndexKey(label, value string) string {
	suffix := strings.TrimPrefix(label, labels.SelectorLabelPrefix)
	if suffix == labels.DeviceNameLabel {
		return value
	}
	return suffix
}
