package v2

var authenticationAllowlistMethods = map[string]bool{
	"/lockets.api.v2.WorkspaceService/GetWorkspaceProfile":        true,
	"/lockets.api.v2.WorkspaceSettingService/GetWorkspaceSetting": true,
	"/lockets.api.v2.AuthService/GetAuthStatus":                   true,
	"/lockets.api.v2.AuthService/SignIn":                          true,
	"/lockets.api.v2.AuthService/SignInWithSSO":                   true,
	"/lockets.api.v2.AuthService/SignOut":                         true,
	"/lockets.api.v2.AuthService/SignUp":                          true,
	"/lockets.api.v2.UserService/GetUser":                         true,
	"/lockets.api.v2.UserService/SearchUsers":                     true,
	"/lockets.api.v2.LocketService/ListLockets":                       true,
	"/lockets.api.v2.LocketService/GetLocket":                         true,
	"/lockets.api.v2.LocketService/SearchLockets":                     true,
	"/lockets.api.v2.LocketService/ListLocketResources":               true,
	"/lockets.api.v2.LocketService/ListLocketRelations":               true,
	"/lockets.api.v2.LocketService/ListLocketComments":                true,
	"/lockets.api.v2.LinkService/GetLinkMetadata":                 true,
}

// isUnauthorizeAllowedMethod returns whether the method is exempted from authentication.
func isUnauthorizeAllowedMethod(fullMethodName string) bool {
	return authenticationAllowlistMethods[fullMethodName]
}

var allowedMethodsOnlyForAdmin = map[string]bool{
	"/lockets.api.v2.UserService/CreateUser": true,
}

// isOnlyForAdminAllowedMethod returns true if the method is allowed to be called only by admin.
func isOnlyForAdminAllowedMethod(methodName string) bool {
	return allowedMethodsOnlyForAdmin[methodName]
}
