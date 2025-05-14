# WebhookNotificationWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the webhook notification. | 
**EventId** | **string** | Unique identifier of the linked event. | 
**Data** | **map[string]interface{}** | The content of the webhook. | 
**Status** | **string** | Status of the webhook notification. | 
**SendAttempts** | [**[]WebhookNotificationWithIncludesSendAttemptsInner**](WebhookNotificationWithIncludesSendAttemptsInner.md) |  | 

## Methods

### NewWebhookNotificationWithIncludes

`func NewWebhookNotificationWithIncludes(id string, eventId string, data map[string]interface{}, status string, sendAttempts []WebhookNotificationWithIncludesSendAttemptsInner, ) *WebhookNotificationWithIncludes`

NewWebhookNotificationWithIncludes instantiates a new WebhookNotificationWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationWithIncludesWithDefaults

`func NewWebhookNotificationWithIncludesWithDefaults() *WebhookNotificationWithIncludes`

NewWebhookNotificationWithIncludesWithDefaults instantiates a new WebhookNotificationWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookNotificationWithIncludes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookNotificationWithIncludes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookNotificationWithIncludes) SetId(v string)`

SetId sets Id field to given value.


### GetEventId

`func (o *WebhookNotificationWithIncludes) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookNotificationWithIncludes) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookNotificationWithIncludes) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetData

`func (o *WebhookNotificationWithIncludes) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookNotificationWithIncludes) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookNotificationWithIncludes) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetStatus

`func (o *WebhookNotificationWithIncludes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookNotificationWithIncludes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookNotificationWithIncludes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSendAttempts

`func (o *WebhookNotificationWithIncludes) GetSendAttempts() []WebhookNotificationWithIncludesSendAttemptsInner`

GetSendAttempts returns the SendAttempts field if non-nil, zero value otherwise.

### GetSendAttemptsOk

`func (o *WebhookNotificationWithIncludes) GetSendAttemptsOk() (*[]WebhookNotificationWithIncludesSendAttemptsInner, bool)`

GetSendAttemptsOk returns a tuple with the SendAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAttempts

`func (o *WebhookNotificationWithIncludes) SetSendAttempts(v []WebhookNotificationWithIncludesSendAttemptsInner)`

SetSendAttempts sets SendAttempts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


