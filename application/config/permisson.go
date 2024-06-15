package config

func SuperUserPermission() []int {
	return []int{0}
}

func AdminPermission() []int {
	return []int{0, 1}
}

func AuthenticatedPermission() []int {
	return []int{0, 1, 2}
}
