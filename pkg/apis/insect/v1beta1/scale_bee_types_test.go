
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


package v1beta1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect/v1beta1"
	. "github.com/tamalsaha/apiserver-builder-demo/pkg/client/clientset_generated/clientset/typed/insect/v1beta1"
)

var _ = Describe("Bee", func() {
	var instance Bee
	var expected Bee
	var client BeeInterface

	BeforeEach(func() {
		instance = Bee{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	Describe("when sending a scale request", func() {
		It("should return success", func() {
			client = cs.InsectV1beta1Client.Bees("bee-test-scale")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			scale := &Scale{}
			scale.Name = instance.Name
			restClient := cs.InsectV1beta1Client.RESTClient()
			err = restClient.Post().Namespace("bee-test-scale").
				Name(instance.Name).
				Resource("bees").
				SubResource("scale").
				Body(scale).Do().Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
