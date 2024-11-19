package util

var (
	// DATABASE ERRORS
	DBErrors = map[string]string{
		"UserNotFound": "[DB001] User not found",
		"MustBeUnique": "[DB002] Username or email must be unique",
	}
)
