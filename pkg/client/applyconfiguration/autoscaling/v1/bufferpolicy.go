// Copyright 2023 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// BufferPolicyApplyConfiguration represents an declarative configuration of the BufferPolicy type for use
// with apply.
type BufferPolicyApplyConfiguration struct {
	MaxReplicas *int32              `json:"maxReplicas,omitempty"`
	MinReplicas *int32              `json:"minReplicas,omitempty"`
	BufferSize  *intstr.IntOrString `json:"bufferSize,omitempty"`
}

// BufferPolicyApplyConfiguration constructs an declarative configuration of the BufferPolicy type for use with
// apply.
func BufferPolicy() *BufferPolicyApplyConfiguration {
	return &BufferPolicyApplyConfiguration{}
}

// WithMaxReplicas sets the MaxReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxReplicas field is set to the value of the last call.
func (b *BufferPolicyApplyConfiguration) WithMaxReplicas(value int32) *BufferPolicyApplyConfiguration {
	b.MaxReplicas = &value
	return b
}

// WithMinReplicas sets the MinReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReplicas field is set to the value of the last call.
func (b *BufferPolicyApplyConfiguration) WithMinReplicas(value int32) *BufferPolicyApplyConfiguration {
	b.MinReplicas = &value
	return b
}

// WithBufferSize sets the BufferSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BufferSize field is set to the value of the last call.
func (b *BufferPolicyApplyConfiguration) WithBufferSize(value intstr.IntOrString) *BufferPolicyApplyConfiguration {
	b.BufferSize = &value
	return b
}