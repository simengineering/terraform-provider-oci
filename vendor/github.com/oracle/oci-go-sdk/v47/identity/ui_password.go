// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// UiPassword A text password that enables a user to sign in to the Console, the user interface for interacting with Oracle
// Cloud Infrastructure.
// For more information about user credentials, see User Credentials (https://docs.cloud.oracle.com/Content/Identity/Concepts/usercredentials.htm).
type UiPassword struct {

	// The user's password for the Console.
	Password *string `mandatory:"false" json:"password"`

	// The OCID of the user.
	UserId *string `mandatory:"false" json:"userId"`

	// Date and time the password was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState UiPasswordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m UiPassword) String() string {
	return common.PointerString(m)
}

// UiPasswordLifecycleStateEnum Enum with underlying type: string
type UiPasswordLifecycleStateEnum string

// Set of constants representing the allowable values for UiPasswordLifecycleStateEnum
const (
	UiPasswordLifecycleStateCreating UiPasswordLifecycleStateEnum = "CREATING"
	UiPasswordLifecycleStateActive   UiPasswordLifecycleStateEnum = "ACTIVE"
	UiPasswordLifecycleStateInactive UiPasswordLifecycleStateEnum = "INACTIVE"
	UiPasswordLifecycleStateDeleting UiPasswordLifecycleStateEnum = "DELETING"
	UiPasswordLifecycleStateDeleted  UiPasswordLifecycleStateEnum = "DELETED"
)

var mappingUiPasswordLifecycleState = map[string]UiPasswordLifecycleStateEnum{
	"CREATING": UiPasswordLifecycleStateCreating,
	"ACTIVE":   UiPasswordLifecycleStateActive,
	"INACTIVE": UiPasswordLifecycleStateInactive,
	"DELETING": UiPasswordLifecycleStateDeleting,
	"DELETED":  UiPasswordLifecycleStateDeleted,
}

// GetUiPasswordLifecycleStateEnumValues Enumerates the set of values for UiPasswordLifecycleStateEnum
func GetUiPasswordLifecycleStateEnumValues() []UiPasswordLifecycleStateEnum {
	values := make([]UiPasswordLifecycleStateEnum, 0)
	for _, v := range mappingUiPasswordLifecycleState {
		values = append(values, v)
	}
	return values
}
