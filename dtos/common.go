package dtos

import (
	"os"
	"strings"
)

const (
	UNKNOWN_ERROR   = 10000
	NOT_FOUND       = 404
)

func FormatBody(m map[string]interface{}) map[string]interface{}{
	m["error_code"] = 0
	return m
}

func FormatError(code int,message string)map[string]interface{} {
	switch code {
	case UNKNOWN_ERROR:
		message = "message.error.unknown"
		break
	case NOT_FOUND:
		message = "message.error.404"
	}

	return map[string]interface{}{
		"error":true,
		"error_code":code,
		"error_desc":message,
	}
}

func FormatPhoto(img string,thumb string,domain string) map[string]interface{} {
	if len(img) == 0{
		return map[string]interface{}{}
	}
	if len(thumb) == 0{
		thumb = img
	}
	if len(domain) == 0{
		domain = os.Getenv("SHOP_URL")
	}
	if !strings.HasPrefix(thumb,"http") || !strings.HasPrefix(thumb,"https"){
		thumb = domain+"/"+thumb
	}
	if !strings.HasPrefix(img,"http") || !strings.HasPrefix(img,"https"){
		img = domain+"/"+img
	}
	return map[string]interface{}{
		"width": nil,
		"height": nil,
		"thumb": thumb,
		"large": img,
	}
}
