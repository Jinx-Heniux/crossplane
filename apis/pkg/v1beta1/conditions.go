/*
Copyright 2020 The Crossplane Authors.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// Condition types.
const (
	// A TypeInstalled indicates whether a package has been installed.
	TypeInstalled runtimev1alpha1.ConditionType = "Installed"

	// A TypeHealthy indicates whether a package is healthy.
	TypeHealthy runtimev1alpha1.ConditionType = "Healthy"
)

// Reasons a package is or is not installed.
const (
	ReasonUnpacking     runtimev1alpha1.ConditionReason = "UnpackingPackage"
	ReasonInactive      runtimev1alpha1.ConditionReason = "InactivePackageRevision"
	ReasonActive        runtimev1alpha1.ConditionReason = "ActivePackageRevision"
	ReasonUnhealthy     runtimev1alpha1.ConditionReason = "UnhealthyPackageRevision"
	ReasonHealthy       runtimev1alpha1.ConditionReason = "HealthyPackageRevision"
	ReasonUnknownHealth runtimev1alpha1.ConditionReason = "UnknownPackageRevisionHealth"
)

// Unpacking indicates that the package manager is waiting for a package
// revision to be unpacked.
func Unpacking() runtimev1alpha1.Condition {
	return runtimev1alpha1.Condition{
		Type:               TypeInstalled,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonUnpacking,
	}
}

// Inactive indicates that the package manager is waiting for a package
// revision to be transitioned to an active state.
func Inactive() runtimev1alpha1.Condition {
	return runtimev1alpha1.Condition{
		Type:               TypeInstalled,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonInactive,
	}
}

// Active indicates that the package manager has installed and activated
// a package revision.
func Active() runtimev1alpha1.Condition {
	return runtimev1alpha1.Condition{
		Type:               TypeInstalled,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonActive,
	}
}

// Unhealthy indicates that the current revision is unhealthy.
func Unhealthy() runtimev1alpha1.Condition {
	return runtimev1alpha1.Condition{
		Type:               TypeHealthy,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonUnhealthy,
	}
}

// Healthy indicates that the current revision is healthy.
func Healthy() runtimev1alpha1.Condition {
	return runtimev1alpha1.Condition{
		Type:               TypeHealthy,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonHealthy,
	}
}

// UnknownHealth indicates that the health of the current revision is unknown.
func UnknownHealth() runtimev1alpha1.Condition {
	return runtimev1alpha1.Condition{
		Type:               TypeHealthy,
		Status:             corev1.ConditionUnknown,
		LastTransitionTime: metav1.Now(),
		Reason:             ReasonUnknownHealth,
	}
}
