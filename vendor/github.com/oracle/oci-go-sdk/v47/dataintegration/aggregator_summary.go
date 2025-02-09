// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// AggregatorSummary A summary type containing information about the object's aggregator including its type, key, name and description.
type AggregatorSummary struct {

	// The type of the aggregator.
	Type *string `mandatory:"false" json:"type"`

	// The key of the aggregator object.
	Key *string `mandatory:"false" json:"key"`

	// The name of the aggregator.
	Name *string `mandatory:"false" json:"name"`

	// The identifier of the aggregator.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The description of the aggregator.
	Description *string `mandatory:"false" json:"description"`
}

func (m AggregatorSummary) String() string {
	return common.PointerString(m)
}
