// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	insect "github.com/tamalsaha/apiserver-builder-demo/pkg/apis/insect"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_Bee_To_insect_Bee,
		Convert_insect_Bee_To_v1beta1_Bee,
		Convert_v1beta1_BeeList_To_insect_BeeList,
		Convert_insect_BeeList_To_v1beta1_BeeList,
		Convert_v1beta1_BeeSpec_To_insect_BeeSpec,
		Convert_insect_BeeSpec_To_v1beta1_BeeSpec,
		Convert_v1beta1_BeeStatus_To_insect_BeeStatus,
		Convert_insect_BeeStatus_To_v1beta1_BeeStatus,
		Convert_v1beta1_BeeStatusStrategy_To_insect_BeeStatusStrategy,
		Convert_insect_BeeStatusStrategy_To_v1beta1_BeeStatusStrategy,
		Convert_v1beta1_BeeStrategy_To_insect_BeeStrategy,
		Convert_insect_BeeStrategy_To_v1beta1_BeeStrategy,
		Convert_v1beta1_Scale_To_insect_Scale,
		Convert_insect_Scale_To_v1beta1_Scale,
		Convert_v1beta1_ScaleList_To_insect_ScaleList,
		Convert_insect_ScaleList_To_v1beta1_ScaleList,
	)
}

func autoConvert_v1beta1_Bee_To_insect_Bee(in *Bee, out *insect.Bee, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_BeeSpec_To_insect_BeeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_BeeStatus_To_insect_BeeStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Bee_To_insect_Bee is an autogenerated conversion function.
func Convert_v1beta1_Bee_To_insect_Bee(in *Bee, out *insect.Bee, s conversion.Scope) error {
	return autoConvert_v1beta1_Bee_To_insect_Bee(in, out, s)
}

func autoConvert_insect_Bee_To_v1beta1_Bee(in *insect.Bee, out *Bee, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_insect_BeeSpec_To_v1beta1_BeeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_insect_BeeStatus_To_v1beta1_BeeStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_insect_Bee_To_v1beta1_Bee is an autogenerated conversion function.
func Convert_insect_Bee_To_v1beta1_Bee(in *insect.Bee, out *Bee, s conversion.Scope) error {
	return autoConvert_insect_Bee_To_v1beta1_Bee(in, out, s)
}

func autoConvert_v1beta1_BeeList_To_insect_BeeList(in *BeeList, out *insect.BeeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]insect.Bee)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_BeeList_To_insect_BeeList is an autogenerated conversion function.
func Convert_v1beta1_BeeList_To_insect_BeeList(in *BeeList, out *insect.BeeList, s conversion.Scope) error {
	return autoConvert_v1beta1_BeeList_To_insect_BeeList(in, out, s)
}

func autoConvert_insect_BeeList_To_v1beta1_BeeList(in *insect.BeeList, out *BeeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Bee)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_insect_BeeList_To_v1beta1_BeeList is an autogenerated conversion function.
func Convert_insect_BeeList_To_v1beta1_BeeList(in *insect.BeeList, out *BeeList, s conversion.Scope) error {
	return autoConvert_insect_BeeList_To_v1beta1_BeeList(in, out, s)
}

func autoConvert_v1beta1_BeeSpec_To_insect_BeeSpec(in *BeeSpec, out *insect.BeeSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_BeeSpec_To_insect_BeeSpec is an autogenerated conversion function.
func Convert_v1beta1_BeeSpec_To_insect_BeeSpec(in *BeeSpec, out *insect.BeeSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_BeeSpec_To_insect_BeeSpec(in, out, s)
}

func autoConvert_insect_BeeSpec_To_v1beta1_BeeSpec(in *insect.BeeSpec, out *BeeSpec, s conversion.Scope) error {
	return nil
}

// Convert_insect_BeeSpec_To_v1beta1_BeeSpec is an autogenerated conversion function.
func Convert_insect_BeeSpec_To_v1beta1_BeeSpec(in *insect.BeeSpec, out *BeeSpec, s conversion.Scope) error {
	return autoConvert_insect_BeeSpec_To_v1beta1_BeeSpec(in, out, s)
}

func autoConvert_v1beta1_BeeStatus_To_insect_BeeStatus(in *BeeStatus, out *insect.BeeStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_BeeStatus_To_insect_BeeStatus is an autogenerated conversion function.
func Convert_v1beta1_BeeStatus_To_insect_BeeStatus(in *BeeStatus, out *insect.BeeStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_BeeStatus_To_insect_BeeStatus(in, out, s)
}

func autoConvert_insect_BeeStatus_To_v1beta1_BeeStatus(in *insect.BeeStatus, out *BeeStatus, s conversion.Scope) error {
	return nil
}

// Convert_insect_BeeStatus_To_v1beta1_BeeStatus is an autogenerated conversion function.
func Convert_insect_BeeStatus_To_v1beta1_BeeStatus(in *insect.BeeStatus, out *BeeStatus, s conversion.Scope) error {
	return autoConvert_insect_BeeStatus_To_v1beta1_BeeStatus(in, out, s)
}

func autoConvert_v1beta1_BeeStatusStrategy_To_insect_BeeStatusStrategy(in *BeeStatusStrategy, out *insect.BeeStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1beta1_BeeStatusStrategy_To_insect_BeeStatusStrategy is an autogenerated conversion function.
func Convert_v1beta1_BeeStatusStrategy_To_insect_BeeStatusStrategy(in *BeeStatusStrategy, out *insect.BeeStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_BeeStatusStrategy_To_insect_BeeStatusStrategy(in, out, s)
}

func autoConvert_insect_BeeStatusStrategy_To_v1beta1_BeeStatusStrategy(in *insect.BeeStatusStrategy, out *BeeStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_insect_BeeStatusStrategy_To_v1beta1_BeeStatusStrategy is an autogenerated conversion function.
func Convert_insect_BeeStatusStrategy_To_v1beta1_BeeStatusStrategy(in *insect.BeeStatusStrategy, out *BeeStatusStrategy, s conversion.Scope) error {
	return autoConvert_insect_BeeStatusStrategy_To_v1beta1_BeeStatusStrategy(in, out, s)
}

func autoConvert_v1beta1_BeeStrategy_To_insect_BeeStrategy(in *BeeStrategy, out *insect.BeeStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1beta1_BeeStrategy_To_insect_BeeStrategy is an autogenerated conversion function.
func Convert_v1beta1_BeeStrategy_To_insect_BeeStrategy(in *BeeStrategy, out *insect.BeeStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_BeeStrategy_To_insect_BeeStrategy(in, out, s)
}

func autoConvert_insect_BeeStrategy_To_v1beta1_BeeStrategy(in *insect.BeeStrategy, out *BeeStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_insect_BeeStrategy_To_v1beta1_BeeStrategy is an autogenerated conversion function.
func Convert_insect_BeeStrategy_To_v1beta1_BeeStrategy(in *insect.BeeStrategy, out *BeeStrategy, s conversion.Scope) error {
	return autoConvert_insect_BeeStrategy_To_v1beta1_BeeStrategy(in, out, s)
}

func autoConvert_v1beta1_Scale_To_insect_Scale(in *Scale, out *insect.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	return nil
}

// Convert_v1beta1_Scale_To_insect_Scale is an autogenerated conversion function.
func Convert_v1beta1_Scale_To_insect_Scale(in *Scale, out *insect.Scale, s conversion.Scope) error {
	return autoConvert_v1beta1_Scale_To_insect_Scale(in, out, s)
}

func autoConvert_insect_Scale_To_v1beta1_Scale(in *insect.Scale, out *Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	return nil
}

// Convert_insect_Scale_To_v1beta1_Scale is an autogenerated conversion function.
func Convert_insect_Scale_To_v1beta1_Scale(in *insect.Scale, out *Scale, s conversion.Scope) error {
	return autoConvert_insect_Scale_To_v1beta1_Scale(in, out, s)
}

func autoConvert_v1beta1_ScaleList_To_insect_ScaleList(in *ScaleList, out *insect.ScaleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]insect.Scale)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ScaleList_To_insect_ScaleList is an autogenerated conversion function.
func Convert_v1beta1_ScaleList_To_insect_ScaleList(in *ScaleList, out *insect.ScaleList, s conversion.Scope) error {
	return autoConvert_v1beta1_ScaleList_To_insect_ScaleList(in, out, s)
}

func autoConvert_insect_ScaleList_To_v1beta1_ScaleList(in *insect.ScaleList, out *ScaleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Scale)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_insect_ScaleList_To_v1beta1_ScaleList is an autogenerated conversion function.
func Convert_insect_ScaleList_To_v1beta1_ScaleList(in *insect.ScaleList, out *ScaleList, s conversion.Scope) error {
	return autoConvert_insect_ScaleList_To_v1beta1_ScaleList(in, out, s)
}
