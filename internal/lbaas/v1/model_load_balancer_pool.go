/*
Load Balancer Management API

Load Balancer Management API is an API for managing load balancers.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the LoadBalancerPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPool{}

// LoadBalancerPool struct for LoadBalancerPool
type LoadBalancerPool struct {
	// ID of the pool
	Id string `json:"id"`
	// Date and time of creation
	CreatedAt time.Time `json:"created_at"`
	// Date and time of last update
	UpdatedAt time.Time `json:"updated_at"`
	// A name for the pool
	Name     string                   `json:"name"`
	Protocol LoadBalancerPoolProtocol `json:"protocol"`
	// ID of project pool is assigned to
	ProjectId string `json:"project_id"`
	// A list of ports associated with the pool
	Ports []LoadBalancerPort `json:"ports,omitempty"`
	// A list of origins assigned to the pool
	Origins []LoadBalancerPoolOrigin `json:"origins,omitempty"`
	// A list of load balancers assigned to the pool
	Loadbalancers        []LoadBalancerShort `json:"loadbalancers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPool LoadBalancerPool

// NewLoadBalancerPool instantiates a new LoadBalancerPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPool(id string, createdAt time.Time, updatedAt time.Time, name string, protocol LoadBalancerPoolProtocol, projectId string) *LoadBalancerPool {
	this := LoadBalancerPool{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Protocol = protocol
	this.ProjectId = projectId
	return &this
}

// NewLoadBalancerPoolWithDefaults instantiates a new LoadBalancerPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPoolWithDefaults() *LoadBalancerPool {
	this := LoadBalancerPool{}
	return &this
}

// GetId returns the Id field value
func (o *LoadBalancerPool) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancerPool) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LoadBalancerPool) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LoadBalancerPool) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LoadBalancerPool) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LoadBalancerPool) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *LoadBalancerPool) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancerPool) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *LoadBalancerPool) GetProtocol() LoadBalancerPoolProtocol {
	if o == nil {
		var ret LoadBalancerPoolProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetProtocolOk() (*LoadBalancerPoolProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *LoadBalancerPool) SetProtocol(v LoadBalancerPoolProtocol) {
	o.Protocol = v
}

// GetProjectId returns the ProjectId field value
func (o *LoadBalancerPool) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *LoadBalancerPool) SetProjectId(v string) {
	o.ProjectId = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetPorts() []LoadBalancerPort {
	if o == nil || IsNil(o.Ports) {
		var ret []LoadBalancerPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetPortsOk() ([]LoadBalancerPort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []LoadBalancerPort and assigns it to the Ports field.
func (o *LoadBalancerPool) SetPorts(v []LoadBalancerPort) {
	o.Ports = v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetOrigins() []LoadBalancerPoolOrigin {
	if o == nil || IsNil(o.Origins) {
		var ret []LoadBalancerPoolOrigin
		return ret
	}
	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetOriginsOk() ([]LoadBalancerPoolOrigin, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given []LoadBalancerPoolOrigin and assigns it to the Origins field.
func (o *LoadBalancerPool) SetOrigins(v []LoadBalancerPoolOrigin) {
	o.Origins = v
}

// GetLoadbalancers returns the Loadbalancers field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetLoadbalancers() []LoadBalancerShort {
	if o == nil || IsNil(o.Loadbalancers) {
		var ret []LoadBalancerShort
		return ret
	}
	return o.Loadbalancers
}

// GetLoadbalancersOk returns a tuple with the Loadbalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetLoadbalancersOk() ([]LoadBalancerShort, bool) {
	if o == nil || IsNil(o.Loadbalancers) {
		return nil, false
	}
	return o.Loadbalancers, true
}

// HasLoadbalancers returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasLoadbalancers() bool {
	if o != nil && !IsNil(o.Loadbalancers) {
		return true
	}

	return false
}

// SetLoadbalancers gets a reference to the given []LoadBalancerShort and assigns it to the Loadbalancers field.
func (o *LoadBalancerPool) SetLoadbalancers(v []LoadBalancerShort) {
	o.Loadbalancers = v
}

func (o LoadBalancerPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["protocol"] = o.Protocol
	toSerialize["project_id"] = o.ProjectId
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Origins) {
		toSerialize["origins"] = o.Origins
	}
	if !IsNil(o.Loadbalancers) {
		toSerialize["loadbalancers"] = o.Loadbalancers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPool) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPool := _LoadBalancerPool{}

	err = json.Unmarshal(bytes, &varLoadBalancerPool)

	if err != nil {
		return err
	}

	*o = LoadBalancerPool(varLoadBalancerPool)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "origins")
		delete(additionalProperties, "loadbalancers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPool struct {
	value *LoadBalancerPool
	isSet bool
}

func (v NullableLoadBalancerPool) Get() *LoadBalancerPool {
	return v.value
}

func (v *NullableLoadBalancerPool) Set(val *LoadBalancerPool) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPool) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPool(val *LoadBalancerPool) *NullableLoadBalancerPool {
	return &NullableLoadBalancerPool{value: val, isSet: true}
}

func (v NullableLoadBalancerPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}