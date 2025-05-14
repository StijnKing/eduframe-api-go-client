# UpdateWebhookByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The callback url for Eduframe to send a HTTP POST payload to. | [optional] 
**Active** | Pointer to **bool** | State of webhook. | [optional] 
**Events** | Pointer to [**[]WebhookEvent**](WebhookEvent.md) | Array of events. | [optional] 

## Methods

### NewUpdateWebhookByIdRequest

`func NewUpdateWebhookByIdRequest() *UpdateWebhookByIdRequest`

NewUpdateWebhookByIdRequest instantiates a new UpdateWebhookByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookByIdRequestWithDefaults

`func NewUpdateWebhookByIdRequestWithDefaults() *UpdateWebhookByIdRequest`

NewUpdateWebhookByIdRequestWithDefaults instantiates a new UpdateWebhookByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UpdateWebhookByIdRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhookByIdRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhookByIdRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhookByIdRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetActive

`func (o *UpdateWebhookByIdRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateWebhookByIdRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateWebhookByIdRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateWebhookByIdRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEvents

`func (o *UpdateWebhookByIdRequest) GetEvents() []WebhookEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UpdateWebhookByIdRequest) GetEventsOk() (*[]WebhookEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UpdateWebhookByIdRequest) SetEvents(v []WebhookEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UpdateWebhookByIdRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


