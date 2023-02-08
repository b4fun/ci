package ci

import (
	"os"
	"regexp"
)

// GetEnvFunc provides environment variable value.
type GetEnvFunc func(string) string

type detectOpts struct {
	GetEnvFrom GetEnvFunc
}

func (o *detectOpts) getEnv(k string) string {
	f := o.GetEnvFrom
	if f == nil {
		f = os.Getenv
	}

	return f(k)
}

// DetectOpts configures Detect.
type DetectOpts interface {
	apply(*detectOpts)
}

type detectOptsFn func(*detectOpts)

func (f detectOptsFn) apply(o *detectOpts) {
	f(o)
}

// DetectFromEnv sets environment variable source for Detect.
func DetectFromEnv(f GetEnvFunc) DetectOpts {
	return detectOptsFn(func(do *detectOpts) {
		do.GetEnvFrom = f
	})
}

var herokuNodeRegex = regexp.MustCompile(`\/\.heroku\/node\/bin\/node$`)

// Detect detects CI environment name.
// ref: https://github.com/npm/ci-detect/blob/main/lib/index.js
func Detect(opts ...DetectOpts) Name {
	o := new(detectOpts)
	for _, f := range opts {
		f.apply(o)
	}

	notEmpty := func(k string) bool {
		return o.getEnv(k) != ""
	}

	// NOTE: we expect to allow overriding CI name via CI_NAME env
	//       and bypass all remaining checks. Therefore, keep this check on the top.
	if v := o.getEnv("CI_NAME"); v != "" {
		// codeship and a few others
		return Name(v)
	}
	if notEmpty("GERRIT_PROJECT") {
		return Gerrit
	}
	if notEmpty("SYSTEM_TEAMFOUNDATIONCOLLECTIONURI") {
		return AzurePipelines
	}
	if notEmpty("BITRISE_IO") {
		return Bitrise
	}
	if notEmpty("BUDDY_WORKSPACE_ID") {
		return Buddy
	}
	if notEmpty("BUILDKITE") {
		return BuildKite
	}
	if notEmpty("CIRRUS_CI") {
		return Cirrus
	}
	if notEmpty("GITLAB_CI") {
		return Gitlab
	}
	if notEmpty("APPVEYOR") {
		return Appveyor
	}
	if notEmpty("CIRCLECI") {
		return CircleCI
	}
	if notEmpty("SEMAPHORE") {
		return Semaphore
	}
	if notEmpty("DRONE") {
		return Drone
	}
	if notEmpty("DSARI") {
		return Dsari
	}
	if notEmpty("GITHUB_ACTIONS") {
		return GithubActions
	}
	if notEmpty("TDDIUM") {
		return Tddium
	}
	if notEmpty("SCREWDRIVER") {
		return Screwdriver
	}
	if notEmpty("STRIDER") {
		return Strider
	}
	if notEmpty("TASKCLUSTER_ROOT_URL") {
		return TaskCluster
	}
	if notEmpty("JENKINS_URL") {
		return Jenkins
	}
	if notEmpty("bamboo.buildKey") {
		return Bamboo
	}
	if notEmpty("GO_PIPELINE_NAME") {
		return GoCD
	}
	if notEmpty("HUDSON_URL") {
		return Hudson
	}
	if notEmpty("WERCKER") {
		return Wercker
	}
	if notEmpty("NETLIFY") {
		return Netlify
	}
	if notEmpty("NOW_GITHUB_DEPLOYMENT") {
		return NowGitHub
	}
	if notEmpty("GITLAB_DEPLOYMENT") {
		return NowGitLab
	}
	if notEmpty("BITBUCKET_DEPLOYMENT") {
		return NowBitbucket
	}
	if notEmpty("BITBUCKET_BUILD_NUMBER") {
		return BitbucketPipelines
	}
	if notEmpty("NOW_BUILDER") {
		return Now
	}
	if notEmpty("VERCEL_GITHUB_DEPLOYMENT") {
		return VercelGitHub
	}
	if notEmpty("VERCEL_GITLAB_DEPLOYMENT") {
		return VercelGitLab
	}
	if notEmpty("VERCEL_BITBUCKET_DEPLOYMENT") {
		return VercelBitbucket
	}
	if notEmpty("VERCEL_URL") {
		return Vercel
	}
	if notEmpty("MAGNUM") {
		return Magnum
	}
	if notEmpty("NEVERCODE") {
		return Nevercode
	}
	if notEmpty("RENDER") {
		return Render
	}
	if notEmpty("SAIL_CI") {
		return Sail
	}
	if notEmpty("SHIPPABLE") {
		return Shippable
	}
	if notEmpty("TEAMCITY_VERSION") {
		return TeamCity
	}
	if v := o.getEnv("NODE"); v != "" && herokuNodeRegex.MatchString(v) {
		return Heroku
	}
	if notEmpty("TRAVIS") {
		return TravisCI
	}
	if notEmpty("CODEBUILD_SRC_DIR") {
		return AWSCodeBuild
	}
	switch o.getEnv("CI") {
	case "woodpecker":
		return Woodpecker
	case "true", "1":
		return Custom
	}
	if notEmpty("BUILDER_OUTPUT") {
		return Builder
	}

	return NoCIDetected
}
