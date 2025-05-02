# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **int32** | 音階 | [optional] 
**FrameLength** | **int32** | 音符のフレーム長 | 
**Lyric** | **string** | 音符の歌詞 | 

## Methods

### NewNote

`func NewNote(frameLength int32, lyric string, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Note) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Note) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Note) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Note) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Note) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Note) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *Note) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Note) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Note) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *Note) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetFrameLength

`func (o *Note) GetFrameLength() int32`

GetFrameLength returns the FrameLength field if non-nil, zero value otherwise.

### GetFrameLengthOk

`func (o *Note) GetFrameLengthOk() (*int32, bool)`

GetFrameLengthOk returns a tuple with the FrameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameLength

`func (o *Note) SetFrameLength(v int32)`

SetFrameLength sets FrameLength field to given value.


### GetLyric

`func (o *Note) GetLyric() string`

GetLyric returns the Lyric field if non-nil, zero value otherwise.

### GetLyricOk

`func (o *Note) GetLyricOk() (*string, bool)`

GetLyricOk returns a tuple with the Lyric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyric

`func (o *Note) SetLyric(v string)`

SetLyric sets Lyric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


