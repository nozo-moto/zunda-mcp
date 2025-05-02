# FrameAudioQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**F0** | **[]float32** | フレームごとの基本周波数 | 
**Volume** | **[]float32** | フレームごとの音量 | 
**Phonemes** | [**[]FramePhoneme**](FramePhoneme.md) | 音素のリスト | 
**VolumeScale** | **float32** | 全体の音量 | 
**OutputSamplingRate** | **int32** | 音声データの出力サンプリングレート | 
**OutputStereo** | **bool** | 音声データをステレオ出力するか否か | 

## Methods

### NewFrameAudioQuery

`func NewFrameAudioQuery(f0 []float32, volume []float32, phonemes []FramePhoneme, volumeScale float32, outputSamplingRate int32, outputStereo bool, ) *FrameAudioQuery`

NewFrameAudioQuery instantiates a new FrameAudioQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameAudioQueryWithDefaults

`func NewFrameAudioQueryWithDefaults() *FrameAudioQuery`

NewFrameAudioQueryWithDefaults instantiates a new FrameAudioQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetF0

`func (o *FrameAudioQuery) GetF0() []float32`

GetF0 returns the F0 field if non-nil, zero value otherwise.

### GetF0Ok

`func (o *FrameAudioQuery) GetF0Ok() (*[]float32, bool)`

GetF0Ok returns a tuple with the F0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF0

`func (o *FrameAudioQuery) SetF0(v []float32)`

SetF0 sets F0 field to given value.


### GetVolume

`func (o *FrameAudioQuery) GetVolume() []float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *FrameAudioQuery) GetVolumeOk() (*[]float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *FrameAudioQuery) SetVolume(v []float32)`

SetVolume sets Volume field to given value.


### GetPhonemes

`func (o *FrameAudioQuery) GetPhonemes() []FramePhoneme`

GetPhonemes returns the Phonemes field if non-nil, zero value otherwise.

### GetPhonemesOk

`func (o *FrameAudioQuery) GetPhonemesOk() (*[]FramePhoneme, bool)`

GetPhonemesOk returns a tuple with the Phonemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonemes

`func (o *FrameAudioQuery) SetPhonemes(v []FramePhoneme)`

SetPhonemes sets Phonemes field to given value.


### GetVolumeScale

`func (o *FrameAudioQuery) GetVolumeScale() float32`

GetVolumeScale returns the VolumeScale field if non-nil, zero value otherwise.

### GetVolumeScaleOk

`func (o *FrameAudioQuery) GetVolumeScaleOk() (*float32, bool)`

GetVolumeScaleOk returns a tuple with the VolumeScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeScale

`func (o *FrameAudioQuery) SetVolumeScale(v float32)`

SetVolumeScale sets VolumeScale field to given value.


### GetOutputSamplingRate

`func (o *FrameAudioQuery) GetOutputSamplingRate() int32`

GetOutputSamplingRate returns the OutputSamplingRate field if non-nil, zero value otherwise.

### GetOutputSamplingRateOk

`func (o *FrameAudioQuery) GetOutputSamplingRateOk() (*int32, bool)`

GetOutputSamplingRateOk returns a tuple with the OutputSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSamplingRate

`func (o *FrameAudioQuery) SetOutputSamplingRate(v int32)`

SetOutputSamplingRate sets OutputSamplingRate field to given value.


### GetOutputStereo

`func (o *FrameAudioQuery) GetOutputStereo() bool`

GetOutputStereo returns the OutputStereo field if non-nil, zero value otherwise.

### GetOutputStereoOk

`func (o *FrameAudioQuery) GetOutputStereoOk() (*bool, bool)`

GetOutputStereoOk returns a tuple with the OutputStereo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStereo

`func (o *FrameAudioQuery) SetOutputStereo(v bool)`

SetOutputStereo sets OutputStereo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


