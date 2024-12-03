// Copyright (c) 2023-2024 Tigera, Inc. All rights reserved.
/*

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

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

// ESGatewayDeployment is the configuration for the es-gateway Deployment.
type ESGatewayDeployment struct {

	// Spec is the specification of the es-gateway Deployment.
	// +optional
	Spec *ESGatewayDeploymentSpec `json:"spec,omitempty"`
}

// ESGatewayDeploymentSpec defines configuration for the es-gateway Deployment.
type ESGatewayDeploymentSpec struct {

	// Template describes the es-gateway Deployment pod that will be created.
	// +optional
	Template *ESGatewayDeploymentPodTemplateSpec `json:"template,omitempty"`
}

// ESGatewayDeploymentPodTemplateSpec is the es-gateway Deployment's PodTemplateSpec
type ESGatewayDeploymentPodTemplateSpec struct {

	// Spec is the es-gateway Deployment's PodSpec.
	// +optional
	Spec *ESGatewayDeploymentPodSpec `json:"spec,omitempty"`
}

// ESGatewayDeploymentPodSpec is the es-gateway Deployment's PodSpec.
type ESGatewayDeploymentPodSpec struct {
	// InitContainers is a list of es-gateway init containers.
	// If specified, this overrides the specified es-gateway Deployment init containers.
	// If omitted, the es-gateway Deployment will use its default values for its init containers.
	// +optional
	InitContainers []ESGatewayDeploymentInitContainer `json:"initContainers,omitempty"`

	// Containers is a list of es-gateway containers.
	// If specified, this overrides the specified es-gateway Deployment containers.
	// If omitted, the es-gateway Deployment will use its default values for its containers.
	// +optional
	Containers []ESGatewayDeploymentContainer `json:"containers,omitempty"`
}

// ESGatewayDeploymentContainer is a es-gateway Deployment container.
type ESGatewayDeploymentContainer struct {
	// Name is an enum which identifies the es-gateway Deployment container by name.
	// Supported values are: tigera-secure-es-gateway
	// +kubebuilder:validation:Enum=tigera-secure-es-gateway
	Name string `json:"name"`

	// Resources allows customization of limits and requests for compute resources such as cpu and memory.
	// If specified, this overrides the named es-gateway Deployment container's resources.
	// If omitted, the es-gateway Deployment will use its default value for this container's resources.
	// +optional
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
}

// ESGatewayDeploymentInitContainer is a es-gateway Deployment init container.
type ESGatewayDeploymentInitContainer struct {
	// Name is an enum which identifies the es-gateway Deployment init container by name.
	// Supported values are: tigera-secure-elasticsearch-cert-key-cert-provisioner
	// +kubebuilder:validation:Enum=tigera-secure-elasticsearch-cert-key-cert-provisioner
	Name string `json:"name"`

	// Resources allows customization of limits and requests for compute resources such as cpu and memory.
	// If specified, this overrides the named es-gateway Deployment init container's resources.
	// If omitted, the es-gateway Deployment will use its default value for this init container's resources.
	// +optional
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
}

func (c *ESGatewayDeployment) GetMetadata() *Metadata {
	return nil
}

func (c *ESGatewayDeployment) GetMinReadySeconds() *int32 {
	return nil
}

func (c *ESGatewayDeployment) GetPodTemplateMetadata() *Metadata {
	return nil
}

func (c *ESGatewayDeployment) GetInitContainers() []v1.Container {
	if c.Spec != nil {
		if c.Spec.Template != nil {
			if c.Spec.Template.Spec != nil {
				if c.Spec.Template.Spec.InitContainers != nil {
					cs := make([]v1.Container, len(c.Spec.Template.Spec.InitContainers))
					for i, v := range c.Spec.Template.Spec.InitContainers {
						// Only copy and return the init container if it has resources set.
						if v.Resources == nil {
							continue
						}
						c := v1.Container{Name: v.Name, Resources: *v.Resources}
						cs[i] = c
					}
					return cs
				}
			}
		}
	}
	return nil
}

func (c *ESGatewayDeployment) GetContainers() []v1.Container {
	if c.Spec != nil {
		if c.Spec.Template != nil {
			if c.Spec.Template.Spec != nil {
				if c.Spec.Template.Spec.Containers != nil {
					cs := make([]v1.Container, len(c.Spec.Template.Spec.Containers))
					for i, v := range c.Spec.Template.Spec.Containers {
						// Only copy and return the container if it has resources set.
						if v.Resources == nil {
							continue
						}
						c := v1.Container{Name: v.Name, Resources: *v.Resources}
						cs[i] = c
					}
					return cs
				}
			}
		}
	}
	return nil
}

func (c *ESGatewayDeployment) GetAffinity() *v1.Affinity {
	return nil
}

func (c *ESGatewayDeployment) GetTopologySpreadConstraints() []v1.TopologySpreadConstraint {
	return nil
}

func (c *ESGatewayDeployment) GetNodeSelector() map[string]string {
	return nil
}

func (c *ESGatewayDeployment) GetTolerations() []v1.Toleration {
	return nil
}

func (c *ESGatewayDeployment) GetTerminationGracePeriodSeconds() *int64 {
	return nil
}

func (c *ESGatewayDeployment) GetDeploymentStrategy() *appsv1.DeploymentStrategy {
	return nil
}

func (c *ESGatewayDeployment) GetPriorityClassName() string {
	return ""
}