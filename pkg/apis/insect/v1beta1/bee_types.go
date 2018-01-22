
/*
Copyright YEAR The Kubernetes Authors.

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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bee
// +k8s:openapi-gen=true
// +resource:path=bees,strategy=BeeStrategy
type Bee struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeeSpec   `json:"spec,omitempty"`
	Status BeeStatus `json:"status,omitempty"`
}

// BeeSpec defines the desired state of Bee
type BeeSpec struct {
}

// BeeStatus defines the observed state of Bee
type BeeStatus struct {
}

// Validate checks that an instance of Bee is well formed
func (BeeStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*insect.Bee)
	log.Printf("Validating fields for Bee %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Bee field values
func (BeeSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Bee)
	// set default field values here
	log.Printf("Defaulting fields for Bee %s\n", obj.Name)
}
