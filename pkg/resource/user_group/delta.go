// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package user_group

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.UserGroupID, b.ko.Spec.UserGroupID) {
		delta.Add("Spec.UserGroupID", a.ko.Spec.UserGroupID, b.ko.Spec.UserGroupID)
	} else if a.ko.Spec.UserGroupID != nil && b.ko.Spec.UserGroupID != nil {
		if *a.ko.Spec.UserGroupID != *b.ko.Spec.UserGroupID {
			delta.Add("Spec.UserGroupID", a.ko.Spec.UserGroupID, b.ko.Spec.UserGroupID)
		}
	}

	if !ackcompare.SliceStringPEqual(a.ko.Spec.UserIDs, b.ko.Spec.UserIDs) {
		delta.Add("Spec.UserIDs", a.ko.Spec.UserIDs, b.ko.Spec.UserIDs)
	}

	return delta
}