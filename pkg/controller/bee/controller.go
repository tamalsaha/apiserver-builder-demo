
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


package bee

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect/v1beta1"
	"github.com/tamalsaha/apiserver-builder-demo/pkg/controller/sharedinformers"
	listers "github.com/tamalsaha/apiserver-builder-demo/pkg/client/listers_generated/insect/v1beta1"
)

// +controller:group=insect,version=v1beta1,kind=Bee,resource=bees
type BeeControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Bee
	lister listers.BeeLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *BeeControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing bees labels
	c.lister = arguments.GetSharedInformers().Factory.Insect().V1beta1().Bees().Lister()
}

// Reconcile handles enqueued messages
func (c *BeeControllerImpl) Reconcile(u *v1beta1.Bee) error {
	// Implement controller logic here
	log.Printf("Running reconcile Bee for %s\n", u.Name)
	return nil
}

func (c *BeeControllerImpl) Get(namespace, name string) (*v1beta1.Bee, error) {
	return c.lister.Bees(namespace).Get(name)
}
