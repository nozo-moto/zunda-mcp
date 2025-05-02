# Score

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | [**[]Note**](Note.md) | 音符のリスト | 

## Methods

### NewScore

`func NewScore(notes []Note, ) *Score`

NewScore instantiates a new Score object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreWithDefaults

`func NewScoreWithDefaults() *Score`

NewScoreWithDefaults instantiates a new Score object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *Score) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Score) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Score) SetNotes(v []Note)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


