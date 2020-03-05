package iagesdk

// NativeImageFeedCreative describe iAGE`s data structure for image feed creative
type NativeImageFeedCreative struct {
	CreativeImage
	NativeIconFile ImageFile `json:"nativeIconFile"`
}
