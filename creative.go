package iagesdk

//Template types

// TemplateNativeSquare - is a value that should be used with API calls to manage NativeSquareCreative
// like: CreateNativeSquareCreativeRequest, UpdateNativeSquareCreativeRequest, etc.
const TemplateNativeSquare = 100

// TemplateNativeHorizontal - is a value that should be used with API calls to manage NativeHorizontalCreative
// like: CreateNativeHorizontalCreativeRequest, UpdateNativeHorizontalCreativeRequest, etc.
const TemplateNativeHorizontal = 101

// TemplateNativeVertical - is a value that should be used with API calls to manage NativeVerticalCreative
// like: CreateNativeVerticalCreativeRequest, UpdateNativeVerticalCreativeRequest, etc.
const TemplateNativeVertical = 102

// TemplateNativeImageAppInstall - is a value that should be used with API calls to manage NativeImageAppInstallCreative
// like: CreateNativeImageAppInstallRequest, UpdateNativeImageAppInstallRequest, etc.
const TemplateNativeImageAppInstall = 103

// TemplateNativeVideoAppInstall - is a value that should be used with API calls to manage NativeVideoAppInstallCreative
// like: CreateNativeVideoAppInstallRequest, UpdateNativeVideoAppInstallRequest, etc.
const TemplateNativeVideoAppInstall = 104

// TemplateNativeImageFeed - is a value that should be used with API calls to manage NativeImageFeedCreative
// like: CreateNativeImageFeedRequest, UpdateNativeImageFeedRequest, etc.
const TemplateNativeImageFeed = 105

// TemplateNativeVideoFeed - is a value that should be used with API calls to manage NativeVideoFeedCreative
// like: CreateNativeVideoFeedRequest, UpdateNativeVideoFeedRequest, etc.
const TemplateNativeVideoFeed = 106

// TemplateBannerFullscreen - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have fullscreen size
const TemplateBannerFullscreen = 200

// TemplateBannerExpandable - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner can be expanded
const TemplateBannerExpandable = 201

// TemplateBannerAdhesion - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner would be adhesion to the top or bottom of the screen
const TemplateBannerAdhesion = 202

// TemplateBannerAdhesionFullscreen - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner would be adhesion to the top or bottom of the screen
const TemplateBannerAdhesionFullscreen = 203

// TemplateBanner240x400 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 240px and height = 400px
const TemplateBanner240x400 = 204

// TemplateBanner300x50 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 300px and height = 50px
const TemplateBanner300x50 = 205

// TemplateBanner300x250 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 300px and height = 250px
const TemplateBanner300x250 = 206

// TemplateBanner320x50 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 320px and height = 50px
const TemplateBanner320x50 = 207

// TemplateBanner468x60 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 468px and height = 60px
const TemplateBanner468x60 = 208

// TemplateBanner640x100 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 640px and height = 100px
const TemplateBanner640x100 = 209

// TemplateBanner728x90 - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner should have width = 728px and height = 90px
const TemplateBanner728x90 = 210

// TemplateBannerCustom - is a value for template type that should be used with API calls to manage any kind of
// banner. It will indicate that banner width and height should be provided in request with additional parameters (CustomWidth & CustomHeight)
const TemplateBannerCustom = 211

// TemplateVideo - is a value for template type that should be used with API calls to manage any kind of video.
const TemplateVideo = 300

//Banner types

// VideoInline - is a value for banner type that should be used with API calls to manage inline video banners
const VideoInline = 1

// VideoWrapper - is a value for banner type that should be used with API calls to manage wrapped video banners
const VideoWrapper = 2

// BannerByURL - is a value for banner type that should be used with API calls to manage banners created / downloaded from URL
const BannerByURL = 3

// BannerHTMLByURL - is a value for banner type that should be used with API calls to manage html banners created / downloaded from URL
const BannerHTMLByURL = 4

// BannerHTMLByCode - is a value for banner type that should be used with API calls to manage banners created from custom code
const BannerHTMLByCode = 5

// MRAID2 - value represent supported API framework MRAID2.0
const MRAID2 = 5

// MRAID3 - value represent supported API framework MRAID3.0
const MRAID3 = 6

// Creative describe iAGE`s base creative data structure
type Creative struct {
	ID           int  `json:"id"`
	LineItemID   int  `json:"lineItemId"`
	Disabled     bool `json:"disabled"`
	TemplateType int  `json:"templateType"`
}
