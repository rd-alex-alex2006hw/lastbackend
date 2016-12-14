package common

import (
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
	"time"
)

// List of all resource kinds supported by the UI.
const (
	ResourceKindConfigMap               = "configmap"
	ResourceKindDaemonSet               = "daemonset"
	ResourceKindDeployment              = "deployment"
	ResourceKindEvent                   = "event"
	ResourceKindHorizontalPodAutoscaler = "horizontalpodautoscaler"
	ResourceKindIngress                 = "ingress"
	ResourceKindJob                     = "job"
	ResourceKindLimitRange              = "limitrange"
	ResourceKindNamespace               = "namespace"
	ResourceKindNode                    = "node"
	ResourceKindPersistentVolumeClaim   = "persistentvolumeclaim"
	ResourceKindPersistentVolume        = "persistentvolume"
	ResourceKindPod                     = "pod"
	ResourceKindReplicaSet              = "replicaset"
	ResourceKindReplicationController   = "replicationcontroller"
	ResourceKindResourceQuota           = "resourcequota"
	ResourceKindSecret                  = "secret"
	ResourceKindService                 = "service"
	ResourceKindStatefulSet             = "statefulset"
)

// ObjectMeta is metadata about an instance of a resource.
type ObjectMeta struct {
	// Name is unique within a namespace. Name is primarily intended for creation
	// idempotence and configuration definition.
	Name string `json:"name,omitempty"`

	// Namespace defines the space within which name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	// Label keys are of the form:
	//     label-key ::= prefixed-name | name
	//     prefixed-name ::= prefix '/' name
	//     prefix ::= DNS_SUBDOMAIN
	//     name ::= DNS_LABEL
	// The prefix is optional.  If the prefix is not specified, the key is assumed to be private
	// to the user.  Other system components that wish to use labels must specify a prefix.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are unstructured key value data stored with a resource that may be set by
	// external tooling. They are not queryable and should be preserved when modifying
	// objects.  Annotation keys have the same formatting restrictions as Label keys. See the
	// comments on Labels for details.
	Annotations map[string]string `json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp unversioned.Time `json:"creationTimestamp,omitempty"`
}

// TypeMeta describes an individual object in an API response or request with strings representing
// the type of the object.
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// In smalllettercase.
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
	Kind ResourceKind `json:"kind,omitempty"`
}

// ResourceKind is an unique name for each resource. It can used for API discovery and generic
// code that does things based on the kind. For example, there may be a generic "deleter"
// that based on resource kind, name and namespace deletes it.
type ResourceKind string

// NewObjectMeta returns internal endpoint name for the given service properties, e.g.,
// NewObjectMeta creates a new instance of ObjectMeta struct based on K8s object meta.
func NewObjectMeta(k8SObjectMeta api.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name:              k8SObjectMeta.Name,
		Namespace:         k8SObjectMeta.Namespace,
		Labels:            k8SObjectMeta.Labels,
		CreationTimestamp: k8SObjectMeta.CreationTimestamp,
		Annotations:       k8SObjectMeta.Annotations,
	}
}

// ListMeta describes list of objects, i.e. holds information about pagination options set for the
// list.
type ListMeta struct {
	// Total number of items on the list. Used for pagination.
	TotalItems int `json:"totalItems"`
}

// NewTypeMeta creates new type mete for the resource kind.
func NewTypeMeta(kind ResourceKind) TypeMeta {
	return TypeMeta{
		Kind: kind,
	}
}

// IsSelectorMatching returns true when an object with the given
// selector targets the same Resources (or subset) that
// the tested object with the given selector.
func IsSelectorMatching(labelSelector map[string]string,
	testedObjectLabels map[string]string) bool {

	// If service has no selectors, then assume it targets different Resource.
	if len(labelSelector) == 0 {
		return false
	}
	for label, value := range labelSelector {
		if rsValue, ok := testedObjectLabels[label]; !ok || rsValue != value {
			return false
		}
	}
	return true
}

// MetricsByPod is a metrics map by pod name.
type MetricsByPod struct {
	// Metrics by namespace and name of a pod.
	MetricsMap map[string]map[string]PodMetrics `json:"metricsMap"`
}

// MetricResult is a some sample measurement of a non-negative, integer quantity (for example,
// memory usage in bytes observed at some moment)
type MetricResult struct {
	Timestamp time.Time `json:"timestamp"`
	Value     uint64    `json:"value"`
}

// PodMetrics is a structure representing pods metrics, contains information about CPU and memory
// usage.
type PodMetrics struct {
	// Most recent measure of CPU usage on all cores in nanoseconds.
	CPUUsage *uint64 `json:"cpuUsage"`
	// Pod memory usage in bytes.
	MemoryUsage *uint64 `json:"memoryUsage"`
	// Timestamped samples of CPUUsage over some short period of history
	CPUUsageHistory []MetricResult `json:"cpuUsageHistory"`
	// Timestamped samples of pod memory usage over some short period of history
	MemoryUsageHistory []MetricResult `json:"memoryUsageHistory"`
}
