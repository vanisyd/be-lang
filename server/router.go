package server

import "strings"

type Route struct {
	Path     string
	Handler  Action
	Method   string
	Children []Route
}

type RouterMapping map[string]interface{}
type RouterMappingItem struct {
	Route    Route
	Children RouterMapping
}

var routerMap RouterMapping = RouterMapping{}

func GetRouterMap() RouterMapping {
	return routerMap
}

func BuildRouterMap(routes []Route) {
	prepareRouterMap(routes, routerMap)
}

func GetRouteByPath(path string) Route {
	var mappingItem interface{}
	mapping := GetRouterMap()
	pathParts := strings.Split(path, "/")

	for _, part := range pathParts {
		if part != "" {
			if part[:1] != "/" {
				part = "/" + part
			}

			mappingItem = prepareRouterPath(part, mapping, mappingItem)
		}
	}

	route, ok := mappingItem.(RouterMappingItem)
	if ok {
		if route.Route.Handler == nil {
			indexMapping := prepareRouterPath("/", mapping, mappingItem)
			indexRoute, ok := indexMapping.(RouterMappingItem)
			if ok {
				route = indexRoute
			}
		}
		return route.Route
	}

	return Route{}
}

func prepareRouterPath(pathPart string, mapping RouterMapping, mappingItem interface{}) interface{} {
	var curMapping RouterMappingItem
	var ok bool

	curItem, isChild := mappingItem.(RouterMappingItem)
	if !isChild {
		curMapping, ok = mapping[pathPart].(RouterMappingItem)
	} else {
		curMapping, ok = curItem.Children[pathPart].(RouterMappingItem)
	}

	if ok {
		return curMapping
	}

	return nil
}

func prepareRouterMap(routes []Route, mapping RouterMapping) {
	for _, route := range routes {
		mapping[route.Path] = RouterMappingItem{
			Route: Route{
				Path:    route.Path,
				Handler: route.Handler,
				Method:  route.Method,
			},
		}
		if route.Children != nil {
			nestedMap := make(RouterMapping)
			prepareRouterMap(route.Children, nestedMap)
			mappingItem, ok := mapping[route.Path].(RouterMappingItem)
			if ok {
				mappingItem.Children = nestedMap
				mapping[route.Path] = mappingItem
			}
		}
	}
}
