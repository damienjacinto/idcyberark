package version

var (
	// buildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// release is a semantic version of current build
	Release = "unset"
)