package main

type Permissions struct {
	canSeeMessages      bool
	canDeleteMessages   bool
	canEditMessages     bool
	canKickMembers      bool
	canMakeMembersAdmin bool
	canAddMembers       bool
}

func SetUserPermissions(permissions *Permissions) int8 {
	flags := []bool{
		permissions.canSeeMessages,
		permissions.canDeleteMessages,
		permissions.canEditMessages,
		permissions.canKickMembers,
		permissions.canMakeMembersAdmin,
		permissions.canAddMembers,
	}

	var mask int8 = 0
	for i, flag := range flags {
		if flag {
			mask |= 1 << i
		}
	}
	return mask

}

func GetUserPermissions(permissionsMask int8) *Permissions {
	var flags [6]bool
	for i:=0 ;i<6;i++{
		flags[i] = permissionsMask & (1<<i)!=0
	}
	return &Permissions{
		canSeeMessages:      flags[0],
		canDeleteMessages:   flags[1],
		canEditMessages:     flags[2],
		canKickMembers:      flags[3],
		canMakeMembersAdmin: flags[4],
		canAddMembers:       flags[5],
	}
}
