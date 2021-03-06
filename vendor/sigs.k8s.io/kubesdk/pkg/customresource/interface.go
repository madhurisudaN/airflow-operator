/*
Copyright 2018 Google LLC
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

package customresource

import (
	"sigs.k8s.io/kubesdk/pkg/component"
)

// Handle is an interface for operating on CRs
type Handle interface {
	HandleError(error)
	Components() []component.Component
}

// ValidateInterface - optional interface
type ValidateInterface interface {
	Validate() error
}

// ApplyDefaultsInterface - optional interface
type ApplyDefaultsInterface interface {
	ApplyDefaults()
}
