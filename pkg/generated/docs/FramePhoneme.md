# FramePhoneme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phoneme** | **string** | 音素 | 
**FrameLength** | **int32** | 音素のフレーム長 | 
**NoteId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFramePhoneme

`func NewFramePhoneme(phoneme string, frameLength int32, ) *FramePhoneme`

NewFramePhoneme instantiates a new FramePhoneme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFramePhonemeWithDefaults

`func NewFramePhonemeWithDefaults() *FramePhoneme`

NewFramePhonemeWithDefaults instantiates a new FramePhoneme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneme

`func (o *FramePhoneme) GetPhoneme() string`

GetPhoneme returns the Phoneme field if non-nil, zero value otherwise.

### GetPhonemeOk

`func (o *FramePhoneme) GetPhonemeOk() (*string, bool)`

GetPhonemeOk returns a tuple with the Phoneme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneme

`func (o *FramePhoneme) SetPhoneme(v string)`

SetPhoneme sets Phoneme field to given value.


### GetFrameLength

`func (o *FramePhoneme) GetFrameLength() int32`

GetFrameLength returns the FrameLength field if non-nil, zero value otherwise.

### GetFrameLengthOk

`func (o *FramePhoneme) GetFrameLengthOk() (*int32, bool)`

GetFrameLengthOk returns a tuple with the FrameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameLength

`func (o *FramePhoneme) SetFrameLength(v int32)`

SetFrameLength sets FrameLength field to given value.


### GetNoteId

`func (o *FramePhoneme) GetNoteId() string`

GetNoteId returns the NoteId field if non-nil, zero value otherwise.

### GetNoteIdOk

`func (o *FramePhoneme) GetNoteIdOk() (*string, bool)`

GetNoteIdOk returns a tuple with the NoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteId

`func (o *FramePhoneme) SetNoteId(v string)`

SetNoteId sets NoteId field to given value.

### HasNoteId

`func (o *FramePhoneme) HasNoteId() bool`

HasNoteId returns a boolean if a field has been set.

### SetNoteIdNil

`func (o *FramePhoneme) SetNoteIdNil(b bool)`

 SetNoteIdNil sets the value for NoteId to be an explicit nil

### UnsetNoteId
`func (o *FramePhoneme) UnsetNoteId()`

UnsetNoteId ensures that no value is present for NoteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


