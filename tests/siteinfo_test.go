package tests

import (
	"ecshopGoApi/services"
	"fmt"
	"testing"
)

func TestGetSiteInfo(t *testing.T) {
	siteInfo := services.GetSiteInfo()
	fmt.Println(siteInfo)
}
