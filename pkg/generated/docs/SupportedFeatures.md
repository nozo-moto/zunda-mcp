# SupportedFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustMoraPitch** | **bool** | モーラごとの音高の調整 | 
**AdjustPhonemeLength** | **bool** | 音素ごとの長さの調整 | 
**AdjustSpeedScale** | **bool** | 全体の話速の調整 | 
**AdjustPitchScale** | **bool** | 全体の音高の調整 | 
**AdjustIntonationScale** | **bool** | 全体の抑揚の調整 | 
**AdjustVolumeScale** | **bool** | 全体の音量の調整 | 
**AdjustPauseLength** | Pointer to **bool** | 句読点などの無音時間の調整 | [optional] 
**InterrogativeUpspeak** | **bool** | 疑問文の自動調整 | 
**SynthesisMorphing** | **bool** | 2種類のスタイルでモーフィングした音声を合成 | 
**Sing** | Pointer to **bool** | 歌唱音声合成 | [optional] 
**ManageLibrary** | Pointer to **bool** | 音声ライブラリのインストール・アンインストール | [optional] 
**ReturnResourceUrl** | Pointer to **bool** | キャラクター情報のリソースをURLで返送 | [optional] 

## Methods

### NewSupportedFeatures

`func NewSupportedFeatures(adjustMoraPitch bool, adjustPhonemeLength bool, adjustSpeedScale bool, adjustPitchScale bool, adjustIntonationScale bool, adjustVolumeScale bool, interrogativeUpspeak bool, synthesisMorphing bool, ) *SupportedFeatures`

NewSupportedFeatures instantiates a new SupportedFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedFeaturesWithDefaults

`func NewSupportedFeaturesWithDefaults() *SupportedFeatures`

NewSupportedFeaturesWithDefaults instantiates a new SupportedFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustMoraPitch

`func (o *SupportedFeatures) GetAdjustMoraPitch() bool`

GetAdjustMoraPitch returns the AdjustMoraPitch field if non-nil, zero value otherwise.

### GetAdjustMoraPitchOk

`func (o *SupportedFeatures) GetAdjustMoraPitchOk() (*bool, bool)`

GetAdjustMoraPitchOk returns a tuple with the AdjustMoraPitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustMoraPitch

`func (o *SupportedFeatures) SetAdjustMoraPitch(v bool)`

SetAdjustMoraPitch sets AdjustMoraPitch field to given value.


### GetAdjustPhonemeLength

`func (o *SupportedFeatures) GetAdjustPhonemeLength() bool`

GetAdjustPhonemeLength returns the AdjustPhonemeLength field if non-nil, zero value otherwise.

### GetAdjustPhonemeLengthOk

`func (o *SupportedFeatures) GetAdjustPhonemeLengthOk() (*bool, bool)`

GetAdjustPhonemeLengthOk returns a tuple with the AdjustPhonemeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustPhonemeLength

`func (o *SupportedFeatures) SetAdjustPhonemeLength(v bool)`

SetAdjustPhonemeLength sets AdjustPhonemeLength field to given value.


### GetAdjustSpeedScale

`func (o *SupportedFeatures) GetAdjustSpeedScale() bool`

GetAdjustSpeedScale returns the AdjustSpeedScale field if non-nil, zero value otherwise.

### GetAdjustSpeedScaleOk

`func (o *SupportedFeatures) GetAdjustSpeedScaleOk() (*bool, bool)`

GetAdjustSpeedScaleOk returns a tuple with the AdjustSpeedScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustSpeedScale

`func (o *SupportedFeatures) SetAdjustSpeedScale(v bool)`

SetAdjustSpeedScale sets AdjustSpeedScale field to given value.


### GetAdjustPitchScale

`func (o *SupportedFeatures) GetAdjustPitchScale() bool`

GetAdjustPitchScale returns the AdjustPitchScale field if non-nil, zero value otherwise.

### GetAdjustPitchScaleOk

`func (o *SupportedFeatures) GetAdjustPitchScaleOk() (*bool, bool)`

GetAdjustPitchScaleOk returns a tuple with the AdjustPitchScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustPitchScale

`func (o *SupportedFeatures) SetAdjustPitchScale(v bool)`

SetAdjustPitchScale sets AdjustPitchScale field to given value.


### GetAdjustIntonationScale

`func (o *SupportedFeatures) GetAdjustIntonationScale() bool`

GetAdjustIntonationScale returns the AdjustIntonationScale field if non-nil, zero value otherwise.

### GetAdjustIntonationScaleOk

`func (o *SupportedFeatures) GetAdjustIntonationScaleOk() (*bool, bool)`

GetAdjustIntonationScaleOk returns a tuple with the AdjustIntonationScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustIntonationScale

`func (o *SupportedFeatures) SetAdjustIntonationScale(v bool)`

SetAdjustIntonationScale sets AdjustIntonationScale field to given value.


### GetAdjustVolumeScale

`func (o *SupportedFeatures) GetAdjustVolumeScale() bool`

GetAdjustVolumeScale returns the AdjustVolumeScale field if non-nil, zero value otherwise.

### GetAdjustVolumeScaleOk

`func (o *SupportedFeatures) GetAdjustVolumeScaleOk() (*bool, bool)`

GetAdjustVolumeScaleOk returns a tuple with the AdjustVolumeScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustVolumeScale

`func (o *SupportedFeatures) SetAdjustVolumeScale(v bool)`

SetAdjustVolumeScale sets AdjustVolumeScale field to given value.


### GetAdjustPauseLength

`func (o *SupportedFeatures) GetAdjustPauseLength() bool`

GetAdjustPauseLength returns the AdjustPauseLength field if non-nil, zero value otherwise.

### GetAdjustPauseLengthOk

`func (o *SupportedFeatures) GetAdjustPauseLengthOk() (*bool, bool)`

GetAdjustPauseLengthOk returns a tuple with the AdjustPauseLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustPauseLength

`func (o *SupportedFeatures) SetAdjustPauseLength(v bool)`

SetAdjustPauseLength sets AdjustPauseLength field to given value.

### HasAdjustPauseLength

`func (o *SupportedFeatures) HasAdjustPauseLength() bool`

HasAdjustPauseLength returns a boolean if a field has been set.

### GetInterrogativeUpspeak

`func (o *SupportedFeatures) GetInterrogativeUpspeak() bool`

GetInterrogativeUpspeak returns the InterrogativeUpspeak field if non-nil, zero value otherwise.

### GetInterrogativeUpspeakOk

`func (o *SupportedFeatures) GetInterrogativeUpspeakOk() (*bool, bool)`

GetInterrogativeUpspeakOk returns a tuple with the InterrogativeUpspeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrogativeUpspeak

`func (o *SupportedFeatures) SetInterrogativeUpspeak(v bool)`

SetInterrogativeUpspeak sets InterrogativeUpspeak field to given value.


### GetSynthesisMorphing

`func (o *SupportedFeatures) GetSynthesisMorphing() bool`

GetSynthesisMorphing returns the SynthesisMorphing field if non-nil, zero value otherwise.

### GetSynthesisMorphingOk

`func (o *SupportedFeatures) GetSynthesisMorphingOk() (*bool, bool)`

GetSynthesisMorphingOk returns a tuple with the SynthesisMorphing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesisMorphing

`func (o *SupportedFeatures) SetSynthesisMorphing(v bool)`

SetSynthesisMorphing sets SynthesisMorphing field to given value.


### GetSing

`func (o *SupportedFeatures) GetSing() bool`

GetSing returns the Sing field if non-nil, zero value otherwise.

### GetSingOk

`func (o *SupportedFeatures) GetSingOk() (*bool, bool)`

GetSingOk returns a tuple with the Sing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSing

`func (o *SupportedFeatures) SetSing(v bool)`

SetSing sets Sing field to given value.

### HasSing

`func (o *SupportedFeatures) HasSing() bool`

HasSing returns a boolean if a field has been set.

### GetManageLibrary

`func (o *SupportedFeatures) GetManageLibrary() bool`

GetManageLibrary returns the ManageLibrary field if non-nil, zero value otherwise.

### GetManageLibraryOk

`func (o *SupportedFeatures) GetManageLibraryOk() (*bool, bool)`

GetManageLibraryOk returns a tuple with the ManageLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageLibrary

`func (o *SupportedFeatures) SetManageLibrary(v bool)`

SetManageLibrary sets ManageLibrary field to given value.

### HasManageLibrary

`func (o *SupportedFeatures) HasManageLibrary() bool`

HasManageLibrary returns a boolean if a field has been set.

### GetReturnResourceUrl

`func (o *SupportedFeatures) GetReturnResourceUrl() bool`

GetReturnResourceUrl returns the ReturnResourceUrl field if non-nil, zero value otherwise.

### GetReturnResourceUrlOk

`func (o *SupportedFeatures) GetReturnResourceUrlOk() (*bool, bool)`

GetReturnResourceUrlOk returns a tuple with the ReturnResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnResourceUrl

`func (o *SupportedFeatures) SetReturnResourceUrl(v bool)`

SetReturnResourceUrl sets ReturnResourceUrl field to given value.

### HasReturnResourceUrl

`func (o *SupportedFeatures) HasReturnResourceUrl() bool`

HasReturnResourceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


