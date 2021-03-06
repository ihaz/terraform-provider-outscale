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

// ListenerForCreation Information about the listener to create.
type ListenerForCreation struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort int64 `json:"BackendPort"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	BackendProtocol *string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included).
	LoadBalancerPort int64 `json:"LoadBalancerPort"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	LoadBalancerProtocol string `json:"LoadBalancerProtocol"`
	// The ID of the server certificate.
	ServerCertificateId *string `json:"ServerCertificateId,omitempty"`
}

// GetBackendPort returns the BackendPort field value
func (o *ListenerForCreation) GetBackendPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BackendPort
}

// SetBackendPort sets field value
func (o *ListenerForCreation) SetBackendPort(v int64) {
	o.BackendPort = v
}

// GetBackendProtocol returns the BackendProtocol field value if set, zero value otherwise.
func (o *ListenerForCreation) GetBackendProtocol() string {
	if o == nil || o.BackendProtocol == nil {
		var ret string
		return ret
	}
	return *o.BackendProtocol
}

// GetBackendProtocolOk returns a tuple with the BackendProtocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetBackendProtocolOk() (string, bool) {
	if o == nil || o.BackendProtocol == nil {
		var ret string
		return ret, false
	}
	return *o.BackendProtocol, true
}

// HasBackendProtocol returns a boolean if a field has been set.
func (o *ListenerForCreation) HasBackendProtocol() bool {
	if o != nil && o.BackendProtocol != nil {
		return true
	}

	return false
}

// SetBackendProtocol gets a reference to the given string and assigns it to the BackendProtocol field.
func (o *ListenerForCreation) SetBackendProtocol(v string) {
	o.BackendProtocol = &v
}

// GetLoadBalancerPort returns the LoadBalancerPort field value
func (o *ListenerForCreation) GetLoadBalancerPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LoadBalancerPort
}

// SetLoadBalancerPort sets field value
func (o *ListenerForCreation) SetLoadBalancerPort(v int64) {
	o.LoadBalancerPort = v
}

// GetLoadBalancerProtocol returns the LoadBalancerProtocol field value
func (o *ListenerForCreation) GetLoadBalancerProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerProtocol
}

// SetLoadBalancerProtocol sets field value
func (o *ListenerForCreation) SetLoadBalancerProtocol(v string) {
	o.LoadBalancerProtocol = v
}

// GetServerCertificateId returns the ServerCertificateId field value if set, zero value otherwise.
func (o *ListenerForCreation) GetServerCertificateId() string {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificateId
}

// GetServerCertificateIdOk returns a tuple with the ServerCertificateId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetServerCertificateIdOk() (string, bool) {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret, false
	}
	return *o.ServerCertificateId, true
}

// HasServerCertificateId returns a boolean if a field has been set.
func (o *ListenerForCreation) HasServerCertificateId() bool {
	if o != nil && o.ServerCertificateId != nil {
		return true
	}

	return false
}

// SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.
func (o *ListenerForCreation) SetServerCertificateId(v string) {
	o.ServerCertificateId = &v
}

type NullableListenerForCreation struct {
	Value        ListenerForCreation
	ExplicitNull bool
}

func (v NullableListenerForCreation) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableListenerForCreation) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
