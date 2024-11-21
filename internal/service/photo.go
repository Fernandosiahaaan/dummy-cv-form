package service

import (
	"dummy-cv-form/internal/model"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func (s *Service) SavePhoto(profileCode int64, imageData []byte) (string, error) {
	filePath := fmt.Sprintf("%s/%d.png", model.DirStaticUploadFolder, profileCode)
	if err := os.MkdirAll(model.DirStaticUploadFolder, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directories. err = %v", err)
	}

	err := ioutil.WriteFile(filePath, imageData, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save image. err = %v", err)
	}

	photoURL := fmt.Sprintf("/app/upload/photo/%d.png", profileCode)
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return photoURL, err

	} else if profile == nil {
		return "", fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}
	profile.PhotoURL = photoURL

	_, err = s.UpdateProfile(profileCode, profile)
	if err != nil {
		return photoURL, nil
	}

	return photoURL, nil
}

func (s *Service) StorePhoto(profileCode int64) (*model.BodyDownloadResponse, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	if profile.PhotoURL == "" {
		return nil, fmt.Errorf("%s%d", model.PhotoErr01, profileCode)
	}

	filePath := fmt.Sprintf("%s/%d.png", model.DirStaticUploadFolder, profileCode)
	imgData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image '%s'", filePath)
	}

	// Convert image to base64 string
	base64Img := base64.StdEncoding.EncodeToString(imgData)
	dataResponse := "image/png;base64," + base64Img
	return &model.BodyDownloadResponse{Base64Img: dataResponse}, err
}

func (s *Service) DeletePhoto(profileCode int64) (*model.OnlyProfileCodeResponse, error) {
	// check profile code from db
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	if profile.PhotoURL == "" {
		return nil, fmt.Errorf("%s%d", model.PhotoErr01, profileCode)
	}

	filePath := fmt.Sprintf("%s/%d.png", model.DirStaticUploadFolder, profileCode)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found. err : %v", err)
	}

	err = os.Remove(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to delete image. err : %v", err)
	}

	profile.PhotoURL = ""
	_, err = s.UpdateProfile(profileCode, profile)
	if err != nil {
		return nil, err
	}

	return &model.OnlyProfileCodeResponse{ProfileCode: profileCode}, nil
}
