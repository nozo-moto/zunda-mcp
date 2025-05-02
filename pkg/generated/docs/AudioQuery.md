# AudioQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccentPhrases** | [**[]AccentPhrase**](AccentPhrase.md) | アクセント句のリスト | 
**SpeedScale** | **float32** | 全体の話速 | 
**PitchScale** | **float32** | 全体の音高 | 
**IntonationScale** | **float32** | 全体の抑揚 | 
**VolumeScale** | **float32** | 全体の音量 | 
**PrePhonemeLength** | **float32** | 音声の前の無音時間 | 
**PostPhonemeLength** | **float32** | 音声の後の無音時間 | 
**PauseLength** | Pointer to **NullableFloat32** |  | [optional] 
**PauseLengthScale** | Pointer to **float32** | 句読点などの無音時間（倍率）。デフォルト値は1 | [optional] [default to 1]
**OutputSamplingRate** | **int32** | 音声データの出力サンプリングレート | 
**OutputStereo** | **bool** | 音声データをステレオ出力するか否か | 
**Kana** | Pointer to **string** | [読み取り専用]AquesTalk 風記法によるテキスト。音声合成用のクエリとしては無視される | [optional] 

## Methods

### NewAudioQuery

`func NewAudioQuery(accentPhrases []AccentPhrase, speedScale float32, pitchScale float32, intonationScale float32, volumeScale float32, prePhonemeLength float32, postPhonemeLength float32, outputSamplingRate int32, outputStereo bool, ) *AudioQuery`

NewAudioQuery instantiates a new AudioQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioQueryWithDefaults

`func NewAudioQueryWithDefaults() *AudioQuery`

NewAudioQueryWithDefaults instantiates a new AudioQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccentPhrases

`func (o *AudioQuery) GetAccentPhrases() []AccentPhrase`

GetAccentPhrases returns the AccentPhrases field if non-nil, zero value otherwise.

### GetAccentPhrasesOk

`func (o *AudioQuery) GetAccentPhrasesOk() (*[]AccentPhrase, bool)`

GetAccentPhrasesOk returns a tuple with the AccentPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentPhrases

`func (o *AudioQuery) SetAccentPhrases(v []AccentPhrase)`

SetAccentPhrases sets AccentPhrases field to given value.


### GetSpeedScale

`func (o *AudioQuery) GetSpeedScale() float32`

GetSpeedScale returns the SpeedScale field if non-nil, zero value otherwise.

### GetSpeedScaleOk

`func (o *AudioQuery) GetSpeedScaleOk() (*float32, bool)`

GetSpeedScaleOk returns a tuple with the SpeedScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedScale

`func (o *AudioQuery) SetSpeedScale(v float32)`

SetSpeedScale sets SpeedScale field to given value.


### GetPitchScale

`func (o *AudioQuery) GetPitchScale() float32`

GetPitchScale returns the PitchScale field if non-nil, zero value otherwise.

### GetPitchScaleOk

`func (o *AudioQuery) GetPitchScaleOk() (*float32, bool)`

GetPitchScaleOk returns a tuple with the PitchScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitchScale

`func (o *AudioQuery) SetPitchScale(v float32)`

SetPitchScale sets PitchScale field to given value.


### GetIntonationScale

`func (o *AudioQuery) GetIntonationScale() float32`

GetIntonationScale returns the IntonationScale field if non-nil, zero value otherwise.

### GetIntonationScaleOk

`func (o *AudioQuery) GetIntonationScaleOk() (*float32, bool)`

GetIntonationScaleOk returns a tuple with the IntonationScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntonationScale

`func (o *AudioQuery) SetIntonationScale(v float32)`

SetIntonationScale sets IntonationScale field to given value.


### GetVolumeScale

`func (o *AudioQuery) GetVolumeScale() float32`

GetVolumeScale returns the VolumeScale field if non-nil, zero value otherwise.

### GetVolumeScaleOk

`func (o *AudioQuery) GetVolumeScaleOk() (*float32, bool)`

GetVolumeScaleOk returns a tuple with the VolumeScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeScale

`func (o *AudioQuery) SetVolumeScale(v float32)`

SetVolumeScale sets VolumeScale field to given value.


### GetPrePhonemeLength

`func (o *AudioQuery) GetPrePhonemeLength() float32`

GetPrePhonemeLength returns the PrePhonemeLength field if non-nil, zero value otherwise.

### GetPrePhonemeLengthOk

`func (o *AudioQuery) GetPrePhonemeLengthOk() (*float32, bool)`

GetPrePhonemeLengthOk returns a tuple with the PrePhonemeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePhonemeLength

`func (o *AudioQuery) SetPrePhonemeLength(v float32)`

SetPrePhonemeLength sets PrePhonemeLength field to given value.


### GetPostPhonemeLength

`func (o *AudioQuery) GetPostPhonemeLength() float32`

GetPostPhonemeLength returns the PostPhonemeLength field if non-nil, zero value otherwise.

### GetPostPhonemeLengthOk

`func (o *AudioQuery) GetPostPhonemeLengthOk() (*float32, bool)`

GetPostPhonemeLengthOk returns a tuple with the PostPhonemeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPhonemeLength

`func (o *AudioQuery) SetPostPhonemeLength(v float32)`

SetPostPhonemeLength sets PostPhonemeLength field to given value.


### GetPauseLength

`func (o *AudioQuery) GetPauseLength() float32`

GetPauseLength returns the PauseLength field if non-nil, zero value otherwise.

### GetPauseLengthOk

`func (o *AudioQuery) GetPauseLengthOk() (*float32, bool)`

GetPauseLengthOk returns a tuple with the PauseLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseLength

`func (o *AudioQuery) SetPauseLength(v float32)`

SetPauseLength sets PauseLength field to given value.

### HasPauseLength

`func (o *AudioQuery) HasPauseLength() bool`

HasPauseLength returns a boolean if a field has been set.

### SetPauseLengthNil

`func (o *AudioQuery) SetPauseLengthNil(b bool)`

 SetPauseLengthNil sets the value for PauseLength to be an explicit nil

### UnsetPauseLength
`func (o *AudioQuery) UnsetPauseLength()`

UnsetPauseLength ensures that no value is present for PauseLength, not even an explicit nil
### GetPauseLengthScale

`func (o *AudioQuery) GetPauseLengthScale() float32`

GetPauseLengthScale returns the PauseLengthScale field if non-nil, zero value otherwise.

### GetPauseLengthScaleOk

`func (o *AudioQuery) GetPauseLengthScaleOk() (*float32, bool)`

GetPauseLengthScaleOk returns a tuple with the PauseLengthScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseLengthScale

`func (o *AudioQuery) SetPauseLengthScale(v float32)`

SetPauseLengthScale sets PauseLengthScale field to given value.

### HasPauseLengthScale

`func (o *AudioQuery) HasPauseLengthScale() bool`

HasPauseLengthScale returns a boolean if a field has been set.

### GetOutputSamplingRate

`func (o *AudioQuery) GetOutputSamplingRate() int32`

GetOutputSamplingRate returns the OutputSamplingRate field if non-nil, zero value otherwise.

### GetOutputSamplingRateOk

`func (o *AudioQuery) GetOutputSamplingRateOk() (*int32, bool)`

GetOutputSamplingRateOk returns a tuple with the OutputSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSamplingRate

`func (o *AudioQuery) SetOutputSamplingRate(v int32)`

SetOutputSamplingRate sets OutputSamplingRate field to given value.


### GetOutputStereo

`func (o *AudioQuery) GetOutputStereo() bool`

GetOutputStereo returns the OutputStereo field if non-nil, zero value otherwise.

### GetOutputStereoOk

`func (o *AudioQuery) GetOutputStereoOk() (*bool, bool)`

GetOutputStereoOk returns a tuple with the OutputStereo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStereo

`func (o *AudioQuery) SetOutputStereo(v bool)`

SetOutputStereo sets OutputStereo field to given value.


### GetKana

`func (o *AudioQuery) GetKana() string`

GetKana returns the Kana field if non-nil, zero value otherwise.

### GetKanaOk

`func (o *AudioQuery) GetKanaOk() (*string, bool)`

GetKanaOk returns a tuple with the Kana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKana

`func (o *AudioQuery) SetKana(v string)`

SetKana sets Kana field to given value.

### HasKana

`func (o *AudioQuery) HasKana() bool`

HasKana returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


