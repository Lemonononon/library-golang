package auth

type RoleType int

const (
	RoleTypeInvalid RoleType = 0
	RoleTypeUser    RoleType = 1
	RoleTypeAdmin   RoleType = 2
)
