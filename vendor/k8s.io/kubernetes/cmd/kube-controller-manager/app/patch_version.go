package app

var (
	// OpenshiftMajor return openshift major version
	OpenshiftMajor string
	// OpenshiftMinor returns openshift minor version
	OpenshiftMinor string
	// OpenshiftGitVersion return git version of openshift
	OpenshiftGitVersion string
	// OpenshiftGitCommit returns git commit
	OpenshiftGitCommit string
	// OpenshiftBuildDate returns build date of openshift
	OpenshiftBuildDate string
)

// OpenshiftVersionInfo stores version information about Openshift accessible to k8s packages
type OpenshiftVersionInfo struct {
	Major      string `json:"major"`
	Minor      string `json:"minor"`
	GitCommit  string `json:"gitCommit"`
	GitVersion string `json:"gitVersion"`
	BuildDate  string `json:"buildDate"`
}

// GetOpenshiftVersion returns Overall Openshift codebase version
func GetOpenshiftVersion() OpenshiftVersionInfo {
	return OpenshiftVersionInfo{
		Major:      OpenshiftMajor,
		Minor:      OpenshiftMinor,
		GitCommit:  OpenshiftGitCommit,
		GitVersion: OpenshiftGitVersion,
		BuildDate:  OpenshiftBuildDate,
	}
}
