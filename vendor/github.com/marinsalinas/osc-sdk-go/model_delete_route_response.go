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

// DeleteRouteResponse struct for DeleteRouteResponse
type DeleteRouteResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	RouteTable      *RouteTable      `json:"RouteTable,omitempty"`
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteRouteResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRouteResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteRouteResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteRouteResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetRouteTable returns the RouteTable field value if set, zero value otherwise.
func (o *DeleteRouteResponse) GetRouteTable() RouteTable {
	if o == nil || o.RouteTable == nil {
		var ret RouteTable
		return ret
	}
	return *o.RouteTable
}

// GetRouteTableOk returns a tuple with the RouteTable field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRouteResponse) GetRouteTableOk() (RouteTable, bool) {
	if o == nil || o.RouteTable == nil {
		var ret RouteTable
		return ret, false
	}
	return *o.RouteTable, true
}

// HasRouteTable returns a boolean if a field has been set.
func (o *DeleteRouteResponse) HasRouteTable() bool {
	if o != nil && o.RouteTable != nil {
		return true
	}

	return false
}

// SetRouteTable gets a reference to the given RouteTable and assigns it to the RouteTable field.
func (o *DeleteRouteResponse) SetRouteTable(v RouteTable) {
	o.RouteTable = &v
}

type NullableDeleteRouteResponse struct {
	Value        DeleteRouteResponse
	ExplicitNull bool
}

func (v NullableDeleteRouteResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDeleteRouteResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
