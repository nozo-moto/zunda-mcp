# EngineManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManifestVersion** | **string** | マニフェストのバージョン | 
**Name** | **string** | エンジン名 | 
**BrandName** | **string** | ブランド名 | 
**Uuid** | **string** | エンジンのUUID | 
**Url** | **string** | エンジンのURL | 
**Icon** | **string** | エンジンのアイコンをBASE64エンコードしたもの | 
**DefaultSamplingRate** | **int32** | デフォルトのサンプリング周波数 | 
**FrameRate** | **float32** | エンジンのフレームレート | 
**TermsOfService** | **string** | エンジンの利用規約 | 
**UpdateInfos** | [**[]UpdateInfo**](UpdateInfo.md) | エンジンのアップデート情報 | 
**DependencyLicenses** | [**[]LicenseInfo**](LicenseInfo.md) | 依存関係のライセンス情報 | 
**SupportedVvlibManifestVersion** | Pointer to **string** | エンジンが対応するvvlibのバージョン | [optional] 
**SupportedFeatures** | [**SupportedFeatures**](SupportedFeatures.md) | エンジンが持つ機能 | 

## Methods

### NewEngineManifest

`func NewEngineManifest(manifestVersion string, name string, brandName string, uuid string, url string, icon string, defaultSamplingRate int32, frameRate float32, termsOfService string, updateInfos []UpdateInfo, dependencyLicenses []LicenseInfo, supportedFeatures SupportedFeatures, ) *EngineManifest`

NewEngineManifest instantiates a new EngineManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineManifestWithDefaults

`func NewEngineManifestWithDefaults() *EngineManifest`

NewEngineManifestWithDefaults instantiates a new EngineManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifestVersion

`func (o *EngineManifest) GetManifestVersion() string`

GetManifestVersion returns the ManifestVersion field if non-nil, zero value otherwise.

### GetManifestVersionOk

`func (o *EngineManifest) GetManifestVersionOk() (*string, bool)`

GetManifestVersionOk returns a tuple with the ManifestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestVersion

`func (o *EngineManifest) SetManifestVersion(v string)`

SetManifestVersion sets ManifestVersion field to given value.


### GetName

`func (o *EngineManifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineManifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineManifest) SetName(v string)`

SetName sets Name field to given value.


### GetBrandName

`func (o *EngineManifest) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *EngineManifest) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *EngineManifest) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


### GetUuid

`func (o *EngineManifest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EngineManifest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EngineManifest) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetUrl

`func (o *EngineManifest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EngineManifest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EngineManifest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIcon

`func (o *EngineManifest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EngineManifest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EngineManifest) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetDefaultSamplingRate

`func (o *EngineManifest) GetDefaultSamplingRate() int32`

GetDefaultSamplingRate returns the DefaultSamplingRate field if non-nil, zero value otherwise.

### GetDefaultSamplingRateOk

`func (o *EngineManifest) GetDefaultSamplingRateOk() (*int32, bool)`

GetDefaultSamplingRateOk returns a tuple with the DefaultSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSamplingRate

`func (o *EngineManifest) SetDefaultSamplingRate(v int32)`

SetDefaultSamplingRate sets DefaultSamplingRate field to given value.


### GetFrameRate

`func (o *EngineManifest) GetFrameRate() float32`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *EngineManifest) GetFrameRateOk() (*float32, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameRate

`func (o *EngineManifest) SetFrameRate(v float32)`

SetFrameRate sets FrameRate field to given value.


### GetTermsOfService

`func (o *EngineManifest) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *EngineManifest) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *EngineManifest) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.


### GetUpdateInfos

`func (o *EngineManifest) GetUpdateInfos() []UpdateInfo`

GetUpdateInfos returns the UpdateInfos field if non-nil, zero value otherwise.

### GetUpdateInfosOk

`func (o *EngineManifest) GetUpdateInfosOk() (*[]UpdateInfo, bool)`

GetUpdateInfosOk returns a tuple with the UpdateInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInfos

`func (o *EngineManifest) SetUpdateInfos(v []UpdateInfo)`

SetUpdateInfos sets UpdateInfos field to given value.


### GetDependencyLicenses

`func (o *EngineManifest) GetDependencyLicenses() []LicenseInfo`

GetDependencyLicenses returns the DependencyLicenses field if non-nil, zero value otherwise.

### GetDependencyLicensesOk

`func (o *EngineManifest) GetDependencyLicensesOk() (*[]LicenseInfo, bool)`

GetDependencyLicensesOk returns a tuple with the DependencyLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyLicenses

`func (o *EngineManifest) SetDependencyLicenses(v []LicenseInfo)`

SetDependencyLicenses sets DependencyLicenses field to given value.


### GetSupportedVvlibManifestVersion

`func (o *EngineManifest) GetSupportedVvlibManifestVersion() string`

GetSupportedVvlibManifestVersion returns the SupportedVvlibManifestVersion field if non-nil, zero value otherwise.

### GetSupportedVvlibManifestVersionOk

`func (o *EngineManifest) GetSupportedVvlibManifestVersionOk() (*string, bool)`

GetSupportedVvlibManifestVersionOk returns a tuple with the SupportedVvlibManifestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVvlibManifestVersion

`func (o *EngineManifest) SetSupportedVvlibManifestVersion(v string)`

SetSupportedVvlibManifestVersion sets SupportedVvlibManifestVersion field to given value.

### HasSupportedVvlibManifestVersion

`func (o *EngineManifest) HasSupportedVvlibManifestVersion() bool`

HasSupportedVvlibManifestVersion returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EngineManifest) GetSupportedFeatures() SupportedFeatures`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EngineManifest) GetSupportedFeaturesOk() (*SupportedFeatures, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EngineManifest) SetSupportedFeatures(v SupportedFeatures)`

SetSupportedFeatures sets SupportedFeatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


