# SpeakerStyle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | スタイル名 | 
**Id** | **int32** | スタイルID | 
**Type** | Pointer to **string** | スタイルの種類。talk:音声合成クエリの作成と音声合成が可能。singing_teacher:歌唱音声合成用のクエリの作成が可能。frame_decode:歌唱音声合成が可能。sing:歌唱音声合成用のクエリの作成と歌唱音声合成が可能。 | [optional] [default to "talk"]

## Methods

### NewSpeakerStyle

`func NewSpeakerStyle(name string, id int32, ) *SpeakerStyle`

NewSpeakerStyle instantiates a new SpeakerStyle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeakerStyleWithDefaults

`func NewSpeakerStyleWithDefaults() *SpeakerStyle`

NewSpeakerStyleWithDefaults instantiates a new SpeakerStyle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpeakerStyle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpeakerStyle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpeakerStyle) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *SpeakerStyle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpeakerStyle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpeakerStyle) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *SpeakerStyle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpeakerStyle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpeakerStyle) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpeakerStyle) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


