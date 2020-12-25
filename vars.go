package esri

import (
	"fmt"
	"os"
)

var (
	host             = os.Getenv("ARCGIS_HOST")
	airspacelinkSlug = os.Getenv("ARCGIS_ASL_SLUG")
	baseURL          = fmt.Sprintf("%v/%v/arcgis/rest/services", host, airspacelinkSlug)
)
