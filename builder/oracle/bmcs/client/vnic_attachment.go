// Copyright (c) 2017 Oracle America, Inc.
// The contents of this file are subject to the Mozilla Public License Version
// 2.0 (the "License"); you may not use this file except in compliance with the
// License. If a copy of the MPL was not distributed with this file, You can
// obtain one at http://mozilla.org/MPL/2.0/

package bmcs

import (
	"time"
)

// VNICAttachmentService enables communicating with the BMCS compute API's VNIC
// attachment endpoint.
type VNICAttachmentService struct {
	client *baseClient
}

// NewVNICAttachmentService creates a new VNICAttachmentService for communicating with the
// BMCS compute API's instance related endpoints.
func NewVNICAttachmentService(s *baseClient) *VNICAttachmentService {
	return &VNICAttachmentService{
		client: s.New().Path("vnicAttachments/"),
	}
}

// VNICAttachment details the attachment of a VNIC to a BMCS instance.
type VNICAttachment struct {
	AvailabilityDomain string    `json:"availabilityDomain"`
	CompartmentID      string    `json:"compartmentId"`
	DisplayName        string    `json:"displayName,omitempty"`
	ID                 string    `json:"id"`
	InstanceID         string    `json:"instanceId"`
	LifecycleState     string    `json:"lifecycleState"`
	SubnetID           string    `json:"subnetId"`
	TimeCreated        time.Time `json:"timeCreated"`
	VNICID             string    `json:"vnicId"`
}

// ListVnicAttachmentsParams are the paramaters available when communicating
// with the ListVnicAttachments API endpoint.
type ListVnicAttachmentsParams struct {
	AvailabilityDomain string `url:"availabilityDomain,omitempty"`
	CompartmentID      string `url:"compartmentId"`
	InstanceID         string `url:"instanceId,omitempty"`
	VNICID             string `url:"vnicId,omitempty"`
}

// List returns an array of VNICAttachments.
func (s *VNICAttachmentService) List(params *ListVnicAttachmentsParams) ([]VNICAttachment, error) {
	vnicAttachments := new([]VNICAttachment)
	e := new(APIError)

	_, err := s.client.New().Get("").QueryStruct(params).Receive(vnicAttachments, e)
	err = firstError(err, e)

	return *vnicAttachments, err
}
