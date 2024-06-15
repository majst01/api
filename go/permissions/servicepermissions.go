// Code generated discover.go. DO NOT EDIT.
package permissions

func GetServices() []string {
	return []string{
		"admin.v1.TenantService",
		"admin.v1.TokenService",
		"api.v1.HealthService",
		"api.v1.IPService",
		"api.v1.MethodService",
		"api.v1.ProjectService",
		"api.v1.TenantService",
		"api.v1.TokenService",
		"api.v1.VersionService",
	}
}

func GetServicePermissions() *ServicePermissions {
	return &ServicePermissions{
		Roles: Roles{
			Admin: &Admin{
				Editor: []string{},
				Viewer: []string{},
			},
			Tenant: &Tenant{
				Owner:  []string{},
				Editor: []string{},
				Viewer: []string{},
			},
			Project: &Project{
				Owner:  []string{},
				Editor: []string{},
				Viewer: []string{},
			},
		},
		Methods: map[string]bool{
			"/admin.v1.TenantService/List":          true,
			"/admin.v1.TokenService/List":           true,
			"/admin.v1.TokenService/Revoke":         true,
			"/api.v1.HealthService/Get":             true,
			"/api.v1.IPService/Allocate":            true,
			"/api.v1.IPService/Delete":              true,
			"/api.v1.IPService/Get":                 true,
			"/api.v1.IPService/List":                true,
			"/api.v1.IPService/Update":              true,
			"/api.v1.MethodService/List":            true,
			"/api.v1.MethodService/TokenScopedList": true,
			"/api.v1.ProjectService/Create":         true,
			"/api.v1.ProjectService/Delete":         true,
			"/api.v1.ProjectService/Get":            true,
			"/api.v1.ProjectService/Invite":         true,
			"/api.v1.ProjectService/InviteAccept":   true,
			"/api.v1.ProjectService/InviteDelete":   true,
			"/api.v1.ProjectService/InviteGet":      true,
			"/api.v1.ProjectService/InvitesList":    true,
			"/api.v1.ProjectService/List":           true,
			"/api.v1.ProjectService/RemoveMember":   true,
			"/api.v1.ProjectService/Update":         true,
			"/api.v1.ProjectService/UpdateMember":   true,
			"/api.v1.TenantService/Create":          true,
			"/api.v1.TenantService/Delete":          true,
			"/api.v1.TenantService/Get":             true,
			"/api.v1.TenantService/Invite":          true,
			"/api.v1.TenantService/InviteAccept":    true,
			"/api.v1.TenantService/InviteDelete":    true,
			"/api.v1.TenantService/InviteGet":       true,
			"/api.v1.TenantService/InvitesList":     true,
			"/api.v1.TenantService/List":            true,
			"/api.v1.TenantService/RemoveMember":    true,
			"/api.v1.TenantService/Update":          true,
			"/api.v1.TenantService/UpdateMember":    true,
			"/api.v1.TokenService/Create":           true,
			"/api.v1.TokenService/List":             true,
			"/api.v1.TokenService/Revoke":           true,
			"/api.v1.VersionService/Get":            true,
		},
		Visibility: Visibility{
			Public: map[string]bool{
				"/api.v1.HealthService/Get":                                      true,
				"/api.v1.MethodService/List":                                     true,
				"/api.v1.VersionService/Get":                                     true,
				"/grpc.reflection.v1.ServerReflection/ServerReflectionInfo":      true,
				"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo": true,
			},
			Private: map[string]bool{},
			Self: map[string]bool{
				"/api.v1.MethodService/TokenScopedList": true,
				"/api.v1.ProjectService/InviteAccept":   true,
				"/api.v1.ProjectService/InviteGet":      true,
				"/api.v1.ProjectService/List":           true,
				"/api.v1.TenantService/Create":          true,
				"/api.v1.TenantService/InviteAccept":    true,
				"/api.v1.TenantService/InviteGet":       true,
				"/api.v1.TenantService/List":            true,
				"/api.v1.TokenService/Create":           true,
				"/api.v1.TokenService/List":             true,
				"/api.v1.TokenService/Revoke":           true,
			},
		},
		Auditable: map[string]bool{
			"/admin.v1.TenantService/List":          true,
			"/admin.v1.TokenService/List":           true,
			"/admin.v1.TokenService/Revoke":         true,
			"/api.v1.HealthService/Get":             false,
			"/api.v1.IPService/Allocate":            true,
			"/api.v1.IPService/Delete":              true,
			"/api.v1.IPService/Get":                 false,
			"/api.v1.IPService/List":                false,
			"/api.v1.IPService/Update":              true,
			"/api.v1.MethodService/List":            true,
			"/api.v1.MethodService/TokenScopedList": true,
			"/api.v1.ProjectService/Create":         true,
			"/api.v1.ProjectService/Delete":         true,
			"/api.v1.ProjectService/Get":            false,
			"/api.v1.ProjectService/Invite":         true,
			"/api.v1.ProjectService/InviteAccept":   true,
			"/api.v1.ProjectService/InviteDelete":   true,
			"/api.v1.ProjectService/InviteGet":      false,
			"/api.v1.ProjectService/InvitesList":    false,
			"/api.v1.ProjectService/List":           false,
			"/api.v1.ProjectService/RemoveMember":   true,
			"/api.v1.ProjectService/Update":         true,
			"/api.v1.ProjectService/UpdateMember":   true,
			"/api.v1.TenantService/Create":          true,
			"/api.v1.TenantService/Delete":          true,
			"/api.v1.TenantService/Get":             false,
			"/api.v1.TenantService/Invite":          true,
			"/api.v1.TenantService/InviteAccept":    true,
			"/api.v1.TenantService/InviteDelete":    true,
			"/api.v1.TenantService/InviteGet":       false,
			"/api.v1.TenantService/InvitesList":     false,
			"/api.v1.TenantService/List":            false,
			"/api.v1.TenantService/RemoveMember":    true,
			"/api.v1.TenantService/Update":          true,
			"/api.v1.TenantService/UpdateMember":    true,
			"/api.v1.TokenService/Create":           true,
			"/api.v1.TokenService/List":             true,
			"/api.v1.TokenService/Revoke":           true,
			"/api.v1.VersionService/Get":            false,
		},
	}
}
