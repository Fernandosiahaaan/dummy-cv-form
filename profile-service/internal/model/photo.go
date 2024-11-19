package model

const (
	DirStaticUploadFolder = "../.appdata/upload/photo"
	DirStaticUploadFile   = "/app/upload/photo/"
)

type BodyUploadRequest struct {
	Base64Img string `json:"base64img"`
}

type BodyUploadResponse struct {
	ProfileCode int64  `json:"profileCode"`
	PhotoURL    string `json:"photoURL"`
}

type BodyDownloadResponse struct {
	Base64Img string `json:"data"`
}

type BodyDeleteResponse struct {
	ProfileCode int64 `json:"profileCode"`
}
