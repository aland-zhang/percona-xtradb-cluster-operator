// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the pxc v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=pxc.percona.com
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "pxc.percona.com", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

func init() {
	SchemeBuilder.Register(
		&PerconaXtraDBCluster{}, &PerconaXtraDBClusterList{},
		&PerconaXtraDBClusterBackup{}, &PerconaXtraDBClusterBackupList{},
	)
}