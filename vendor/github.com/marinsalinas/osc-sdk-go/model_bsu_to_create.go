/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// BsuToCreate Information about the BSU volume to create.
type BsuToCreate struct {
	// Set to `true` by default, which means that the volume is deleted when the VM is terminated. If set to `false`, the volume is not deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The number of I/O operations per second (IOPS). This parameter must be specified only if you create an `io1` volume. The maximum number of IOPS allowed for `io1` volumes is `13000`.
	Iops *int64 `json:"Iops,omitempty"`
	// The ID of the snapshot used to create the volume.
	SnapshotId *string `json:"SnapshotId,omitempty"`
	// The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.
	VolumeSize *int64 `json:"VolumeSize,omitempty"`
	// The type of the volume (`standard` \\| `io1` \\| `gp2`). If not specified in the request, a `standard` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).
	VolumeType *string `json:"VolumeType,omitempty"`
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *BsuToCreate) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetDeleteOnVmDeletionOk() (bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret, false
	}
	return *o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *BsuToCreate) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *BsuToCreate) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *BsuToCreate) GetIops() int64 {
	if o == nil || o.Iops == nil {
		var ret int64
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetIopsOk() (int64, bool) {
	if o == nil || o.Iops == nil {
		var ret int64
		return ret, false
	}
	return *o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *BsuToCreate) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given int64 and assigns it to the Iops field.
func (o *BsuToCreate) SetIops(v int64) {
	o.Iops = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *BsuToCreate) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetSnapshotIdOk() (string, bool) {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret, false
	}
	return *o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *BsuToCreate) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *BsuToCreate) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *BsuToCreate) GetVolumeSize() int64 {
	if o == nil || o.VolumeSize == nil {
		var ret int64
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetVolumeSizeOk() (int64, bool) {
	if o == nil || o.VolumeSize == nil {
		var ret int64
		return ret, false
	}
	return *o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *BsuToCreate) HasVolumeSize() bool {
	if o != nil && o.VolumeSize != nil {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int64 and assigns it to the VolumeSize field.
func (o *BsuToCreate) SetVolumeSize(v int64) {
	o.VolumeSize = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *BsuToCreate) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetVolumeTypeOk() (string, bool) {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret, false
	}
	return *o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *BsuToCreate) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *BsuToCreate) SetVolumeType(v string) {
	o.VolumeType = &v
}

type NullableBsuToCreate struct {
	Value        BsuToCreate
	ExplicitNull bool
}

func (v NullableBsuToCreate) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBsuToCreate) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}