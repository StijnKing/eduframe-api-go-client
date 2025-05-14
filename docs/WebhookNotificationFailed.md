# WebhookNotificationFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the webhook notification. | 
**EventId** | **string** | Unique identifier of the linked event. | 
**WebhookId** | **string** | Unique identifier of the linked webhook. | 
**Data** | **map[string]interface{}** | The content of the webhook. | 
**Status** | **string** | Status of the webhook notification. | 
**CreatedAt** | **time.Time** | Timestamp of creation. | 

## Methods

### NewWebhookNotificationFailed

`func NewWebhookNotificationFailed(id string, eventId string, webhookId string, data map[string]interface{}, status string, createdAt time.Time, ) *WebhookNotificationFailed`

NewWebhookNotificationFailed instantiates a new WebhookNotificationFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationFailedWithDefaults

`func NewWebhookNotificationFailedWithDefaults() *WebhookNotificationFailed`

NewWebhookNotificationFailedWithDefaults instantiates a new WebhookNotificationFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookNotificationFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookNotificationFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookNotificationFailed) SetId(v string)`

SetId sets Id field to given value.


### GetEventId

`func (o *WebhookNotificationFailed) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookNotificationFailed) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookNotificationFailed) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetWebhookId

`func (o *WebhookNotificationFailed) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookNotificationFailed) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookNotificationFailed) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetData

`func (o *WebhookNotificationFailed) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookNotificationFailed) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookNotificationFailed) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetStatus

`func (o *WebhookNotificationFailed) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookNotificationFailed) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookNotificationFailed) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *WebhookNotificationFailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookNotificationFailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookNotificationFailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


