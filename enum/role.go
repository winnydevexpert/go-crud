package enum

type RoleType struct {
	Normal uint
	Admin  uint
}

var RoleEnum = RoleType{
	Normal: 1,
	Admin:  2,
}
