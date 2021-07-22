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

package user

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/elasticache-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ElastiCache{}
	_ = &svcapitypes.User{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeUsersOutput
	resp, err = rm.sdkapi.DescribeUsersWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeUsers", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UserNotFound" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.Users {
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.AccessString != nil {
			ko.Spec.AccessString = elem.AccessString
		} else {
			ko.Spec.AccessString = nil
		}
		if elem.Authentication != nil {
			f2 := &svcapitypes.Authentication{}
			if elem.Authentication.PasswordCount != nil {
				f2.PasswordCount = elem.Authentication.PasswordCount
			}
			if elem.Authentication.Type != nil {
				f2.Type = elem.Authentication.Type
			}
			ko.Status.Authentication = f2
		} else {
			ko.Status.Authentication = nil
		}
		if elem.Engine != nil {
			ko.Spec.Engine = elem.Engine
		} else {
			ko.Spec.Engine = nil
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		} else {
			ko.Status.Status = nil
		}
		if elem.UserGroupIds != nil {
			f5 := []*string{}
			for _, f5iter := range elem.UserGroupIds {
				var f5elem string
				f5elem = *f5iter
				f5 = append(f5, &f5elem)
			}
			ko.Status.UserGroupIDs = f5
		} else {
			ko.Status.UserGroupIDs = nil
		}
		if elem.UserId != nil {
			ko.Spec.UserID = elem.UserId
		} else {
			ko.Spec.UserID = nil
		}
		if elem.UserName != nil {
			ko.Spec.UserName = elem.UserName
		} else {
			ko.Spec.UserName = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	rm.setSyncedCondition(resp.Users[0].Status, &resource{ko})
	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeUsersInput, error) {
	res := &svcsdk.DescribeUsersInput{}

	if r.ko.Spec.UserID != nil {
		res.SetUserId(*r.ko.Spec.UserID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateUserOutput
	_ = resp
	resp, err = rm.sdkapi.CreateUserWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateUser", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Authentication != nil {
		f2 := &svcapitypes.Authentication{}
		if resp.Authentication.PasswordCount != nil {
			f2.PasswordCount = resp.Authentication.PasswordCount
		}
		if resp.Authentication.Type != nil {
			f2.Type = resp.Authentication.Type
		}
		ko.Status.Authentication = f2
	} else {
		ko.Status.Authentication = nil
	}
	if resp.Status != nil {
		ko.Status.Status = resp.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.UserGroupIds != nil {
		f5 := []*string{}
		for _, f5iter := range resp.UserGroupIds {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		ko.Status.UserGroupIDs = f5
	} else {
		ko.Status.UserGroupIDs = nil
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomCreateUserSetOutput(ctx, desired, resp, ko)
	if err != nil {
		return nil, err
	}
	rm.setSyncedCondition(resp.Status, &resource{ko})
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateUserInput, error) {
	res := &svcsdk.CreateUserInput{}

	if r.ko.Spec.AccessString != nil {
		res.SetAccessString(*r.ko.Spec.AccessString)
	}
	if r.ko.Spec.Engine != nil {
		res.SetEngine(*r.ko.Spec.Engine)
	}
	if r.ko.Spec.NoPasswordRequired != nil {
		res.SetNoPasswordRequired(*r.ko.Spec.NoPasswordRequired)
	}
	if r.ko.Spec.Passwords != nil {
		f3 := []*string{}
		for _, f3iter := range r.ko.Spec.Passwords {
			var f3elem string
			if f3iter != nil {
				tmpSecret, err := rm.rr.SecretValueFromReference(ctx, f3iter)
				if err != nil {
					return nil, err
				}
				if tmpSecret != "" {
					f3elem = tmpSecret
				}
			}
			f3 = append(f3, &f3elem)
		}
		res.SetPasswords(f3)
	}
	if r.ko.Spec.UserID != nil {
		res.SetUserId(*r.ko.Spec.UserID)
	}
	if r.ko.Spec.UserName != nil {
		res.SetUserName(*r.ko.Spec.UserName)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	updated, err = rm.CustomModifyUser(ctx, desired, latest, delta)
	if updated != nil || err != nil {
		return updated, err
	}
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}
	rm.populateUpdatePayload(input, desired, delta)

	var resp *svcsdk.ModifyUserOutput
	_ = resp
	resp, err = rm.sdkapi.ModifyUserWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyUser", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Authentication != nil {
		f2 := &svcapitypes.Authentication{}
		if resp.Authentication.PasswordCount != nil {
			f2.PasswordCount = resp.Authentication.PasswordCount
		}
		if resp.Authentication.Type != nil {
			f2.Type = resp.Authentication.Type
		}
		ko.Status.Authentication = f2
	} else {
		ko.Status.Authentication = nil
	}
	if resp.Status != nil {
		ko.Status.Status = resp.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.UserGroupIds != nil {
		f5 := []*string{}
		for _, f5iter := range resp.UserGroupIds {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		ko.Status.UserGroupIDs = f5
	} else {
		ko.Status.UserGroupIDs = nil
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomModifyUserSetOutput(ctx, desired, resp, ko)
	if err != nil {
		return nil, err
	}
	rm.setSyncedCondition(resp.Status, &resource{ko})
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.ModifyUserInput, error) {
	res := &svcsdk.ModifyUserInput{}

	if r.ko.Spec.UserID != nil {
		res.SetUserId(*r.ko.Spec.UserID)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteUserOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteUserWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteUser", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteUserInput, error) {
	res := &svcsdk.DeleteUserInput{}

	if r.ko.Spec.UserID != nil {
		res.SetUserId(*r.ko.Spec.UserID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.User,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Message()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "UserAlreadyExists",
		"UserQuotaExceeded",
		"DuplicateUserName",
		"InvalidParameterValue",
		"InvalidParameterCombination",
		"InvalidUserState",
		"UserNotFound",
		"DefaultUserAssociatedToUserGroup":
		return true
	default:
		return false
	}
}
