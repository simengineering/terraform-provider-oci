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

// RegistryMetadata Information about the object and its parent.
type RegistryMetadata struct {

	// The owning object's key for this object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	// The registry version.
	RegistryVersion *int `mandatory:"false" json:"registryVersion"`

	// The identifying key for the object.
	Key *string `mandatory:"false" json:"key"`

	// Specifies whether this object is a favorite or not.
	IsFavorite *bool `mandatory:"false" json:"isFavorite"`
}

func (m RegistryMetadata) String() string {
	return common.PointerString(m)
}
