package app

import (
	apimachineryversion "k8s.io/apimachinery/pkg/version"
)

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

// GetOpenshiftVersion returns Overall Openshift codebase version
func GetOpenshiftVersion() apimachineryversion.Info {
	return apimachineryversion.Info{
		Major:      OpenshiftMajor,
		Minor:      OpenshiftMinor,
		GitCommit:  OpenshiftGitCommit,
		GitVersion: OpenshiftGitVersion,
		BuildDate:  OpenshiftBuildDate,
	}
}
