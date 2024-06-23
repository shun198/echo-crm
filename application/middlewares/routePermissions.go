package middlewares

import (
	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/route"
)

func RoutePermissions(method string, route string, role uint8) bool {
	return config.Contains(int(role), routePermissions[route][method])
}

var routePermissions = mergeMaps(
	route.GetUserPermission(),
)

func mergeMaps[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
