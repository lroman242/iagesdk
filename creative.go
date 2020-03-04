package iagesdk

//Template types
const TemplateNativeSquare = 100
const TemplateNativeHorizontal = 101
const TemplateNativeVertical = 102
const TemplateNativeImageAppInstall = 103
const TemplateNativeVideoAppInstall = 104
const TemplateNativeImageFeed = 105
const TemplateNativeVideoFeed = 106

const TemplateBannerFullscreen = 200
const TemplateBannerExpandable = 201
const TemplateBannerAdhesion = 202
const TemplateBannerAdhesionFullscreen = 203
const TemplateBanner240x400 = 204
const TemplateBanner300x50 = 205
const TemplateBanner300x250 = 206
const TemplateBanner320x50 = 207
const TemplateBanner468x60 = 208
const TemplateBanner640x100 = 209
const TemplateBanner728x90 = 210
const TemplateBannerCustom = 211

const TemplateVideo = 300

//Banner types
const VideoInline = 1
const VideoWrapper = 2
const BannerByURL = 3
const BannerHTMLByURL = 4
const BannerHTMLByCode = 5

const MRAID2 = 5
const MRAID3 = 6

// Creative describe iAGE`s base creative data structure
type Creative struct {
	ID           int  `json:"id"`
	LineItemId   int  `json:"LineItemId"`
	Disabled     bool `json:"disabled"`
	TemplateType int  `json:"templateType"`
}
