package resource

import "web/dataprovider/resource"

var resources []resource.Resource = []resource.Resource{
	UserResource{},
	AuthTokenResource{},
}

func Init() {
	for _, res := range resources {
		resource.PrepareResource(res)
	}
}
