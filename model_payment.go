/*
Eduframe - API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eduframe

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Payment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payment{}

// Payment struct for Payment
type Payment struct {
	// Unique identifier of the payment.
	Id int32 `json:"id"`
	// Unique identifier of the linked invoice.
	InvoiceId int32 `json:"invoice_id"`
	// A number representing the total amount of the invoice.
	Amount string `json:"amount"`
	Currency Currency `json:"currency"`
	// Identifier of the payment method.
	PaymentMethodId int32 `json:"payment_method_id"`
	// The status of the payment.
	Status string `json:"status"`
	// The payment provider payment id.
	GatewayId NullableInt32 `json:"gateway_id"`
	// Date on which the payment was created.
	Date string `json:"date"`
}

type _Payment Payment

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment(id int32, invoiceId int32, amount string, currency Currency, paymentMethodId int32, status string, gatewayId NullableInt32, date string) *Payment {
	this := Payment{}
	this.Id = id
	this.InvoiceId = invoiceId
	this.Amount = amount
	this.Currency = currency
	this.PaymentMethodId = paymentMethodId
	this.Status = status
	this.GatewayId = gatewayId
	this.Date = date
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetId returns the Id field value
func (o *Payment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payment) SetId(v int32) {
	o.Id = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *Payment) GetInvoiceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *Payment) GetInvoiceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *Payment) SetInvoiceId(v int32) {
	o.InvoiceId = v
}

// GetAmount returns the Amount field value
func (o *Payment) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Payment) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Payment) SetAmount(v string) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *Payment) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Payment) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Payment) SetCurrency(v Currency) {
	o.Currency = v
}

// GetPaymentMethodId returns the PaymentMethodId field value
func (o *Payment) GetPaymentMethodId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value
// and a boolean to check if the value has been set.
func (o *Payment) GetPaymentMethodIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodId, true
}

// SetPaymentMethodId sets field value
func (o *Payment) SetPaymentMethodId(v int32) {
	o.PaymentMethodId = v
}

// GetStatus returns the Status field value
func (o *Payment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payment) SetStatus(v string) {
	o.Status = v
}

// GetGatewayId returns the GatewayId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Payment) GetGatewayId() int32 {
	if o == nil || o.GatewayId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.GatewayId.Get()
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetGatewayIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayId.Get(), o.GatewayId.IsSet()
}

// SetGatewayId sets field value
func (o *Payment) SetGatewayId(v int32) {
	o.GatewayId.Set(&v)
}

// GetDate returns the Date field value
func (o *Payment) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Payment) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Payment) SetDate(v string) {
	o.Date = v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["payment_method_id"] = o.PaymentMethodId
	toSerialize["status"] = o.Status
	toSerialize["gateway_id"] = o.GatewayId.Get()
	toSerialize["date"] = o.Date
	return toSerialize, nil
}

func (o *Payment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"invoice_id",
		"amount",
		"currency",
		"payment_method_id",
		"status",
		"gateway_id",
		"date",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPayment := _Payment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPayment)

	if err != nil {
		return err
	}

	*o = Payment(varPayment)

	return err
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


