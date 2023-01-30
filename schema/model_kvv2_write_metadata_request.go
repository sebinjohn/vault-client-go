/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KVv2WriteMetadataRequest struct for KVv2WriteMetadataRequest
type KVv2WriteMetadataRequest struct {
	// If true the key will require the cas parameter to be set on all write requests. If false, the backend’s configuration will be used.
	CasRequired bool `json:"cas_required"`
	// User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret.
	CustomMetadata map[string]interface{} `json:"custom_metadata"`
	// The length of time before a version is deleted. If not set, the backend's configured delete_version_after is used. Cannot be greater than the backend's delete_version_after. A zero duration clears the current setting. A negative duration will cause an error.
	DeleteVersionAfter int32 `json:"delete_version_after"`
	// The number of versions to keep. If not set, the backend’s configured max version is used.
	MaxVersions int32 `json:"max_versions"`
}

// NewKVv2WriteMetadataRequestWithDefaults instantiates a new KVv2WriteMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2WriteMetadataRequestWithDefaults() *KVv2WriteMetadataRequest {
	var this KVv2WriteMetadataRequest

	return &this
}

func (o KVv2WriteMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cas_required"] = o.CasRequired
	toSerialize["custom_metadata"] = o.CustomMetadata
	toSerialize["delete_version_after"] = o.DeleteVersionAfter
	toSerialize["max_versions"] = o.MaxVersions

	return json.Marshal(toSerialize)
}