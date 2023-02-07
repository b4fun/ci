package ci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	setEnv := func(k string, v string) GetEnvFunc {
		return func(s string) string {
			if s == k {
				return v
			}
			return ""
		}
	}

	cases := []struct {
		getEnv   GetEnvFunc
		expected Name
	}{
		{
			getEnv:   setEnv("GERRIT_PROJECT", "1"),
			expected: Gerrit,
		},
		{
			getEnv:   setEnv("SYSTEM_TEAMFOUNDATIONCOLLECTIONURI", "1"),
			expected: AzurePipelines,
		},
		{
			getEnv:   setEnv("BITRISE_IO", "1"),
			expected: Bitrise,
		},
		{
			getEnv:   setEnv("BUDDY_WORKSPACE_ID", "1"),
			expected: Buddy,
		},
		{
			getEnv:   setEnv("BUILDKITE", "1"),
			expected: BuildKite,
		},
		{
			getEnv:   setEnv("CIRRUS_CI", "1"),
			expected: Cirrus,
		},
		{
			getEnv:   setEnv("GITLAB_CI", "1"),
			expected: Gitlab,
		},
		{
			getEnv:   setEnv("APPVEYOR", "1"),
			expected: Appveyor,
		},
		{
			getEnv:   setEnv("CIRCLECI", "1"),
			expected: CircleCI,
		},
		{
			getEnv:   setEnv("SEMAPHORE", "1"),
			expected: Semaphore,
		},
		{
			getEnv:   setEnv("DRONE", "1"),
			expected: Drone,
		},
		{
			getEnv:   setEnv("DSARI", "1"),
			expected: Dsari,
		},
		{
			getEnv:   setEnv("GITHUB_ACTIONS", "1"),
			expected: GithubActions,
		},
		{
			getEnv:   setEnv("TDDIUM", "1"),
			expected: Tddium,
		},
		{
			getEnv:   setEnv("SCREWDRIVER", "1"),
			expected: Screwdriver,
		},
		{
			getEnv:   setEnv("STRIDER", "1"),
			expected: Strider,
		},
		{
			getEnv:   setEnv("TASKCLUSTER_ROOT_URL", "1"),
			expected: TaskCluster,
		},
		{
			getEnv:   setEnv("JENKINS_URL", "1"),
			expected: Jenkins,
		},
		{
			getEnv:   setEnv("bamboo.buildKey", "1"),
			expected: Bamboo,
		},
		{
			getEnv:   setEnv("GO_PIPELINE_NAME", "1"),
			expected: GoCD,
		},
		{
			getEnv:   setEnv("HUDSON_URL", "1"),
			expected: Hudson,
		},
		{
			getEnv:   setEnv("WERCKER", "1"),
			expected: Wercker,
		},
		{
			getEnv:   setEnv("NETLIFY", "1"),
			expected: Netlify,
		},
		{
			getEnv:   setEnv("NOW_GITHUB_DEPLOYMENT", "1"),
			expected: NowGitHub,
		},
		{
			getEnv:   setEnv("GITLAB_DEPLOYMENT", "1"),
			expected: NowGitLab,
		},
		{
			getEnv:   setEnv("BITBUCKET_DEPLOYMENT", "1"),
			expected: NowBitbucket,
		},
		{
			getEnv:   setEnv("NOW_BUILDER", "1"),
			expected: Now,
		},
		{
			getEnv:   setEnv("VERCEL_GITHUB_DEPLOYMENT", "1"),
			expected: VercelGitHub,
		},
		{
			getEnv:   setEnv("VERCEL_GITLAB_DEPLOYMENT", "1"),
			expected: VercelGitLab,
		},
		{
			getEnv:   setEnv("VERCEL_BITBUCKET_DEPLOYMENT", "1"),
			expected: VercelBitbucket,
		},
		{
			getEnv:   setEnv("VERCEL_URL", "1"),
			expected: Vercel,
		},
		{
			getEnv:   setEnv("MAGNUM", "1"),
			expected: Magnum,
		},
		{
			getEnv:   setEnv("NEVERCODE", "1"),
			expected: Nevercode,
		},
		{
			getEnv:   setEnv("RENDER", "1"),
			expected: Render,
		},
		{
			getEnv:   setEnv("SAIL_CI", "1"),
			expected: Sail,
		},
		{
			getEnv:   setEnv("SHIPPABLE", "1"),
			expected: Shippable,
		},
		{
			getEnv:   setEnv("TEAMCITY_VERSION", "1"),
			expected: TeamCity,
		},
		{
			getEnv:   setEnv("CI_NAME", "foobar"),
			expected: Name("foobar"),
		},
		{
			getEnv:   setEnv("NODE", "/usr/local/.heroku/node/bin/node"),
			expected: Heroku,
		},
		{
			getEnv:   setEnv("TRAVIS", "1"),
			expected: TravisCI,
		},
		{
			getEnv:   setEnv("CODEBUILD_SRC_DIR", "1"),
			expected: AWSCodeBuild,
		},
		{
			getEnv:   setEnv("CI", "woodpecker"),
			expected: Woodpecker,
		},
		{
			getEnv:   setEnv("CI", "woodpecker"),
			expected: Woodpecker,
		},
		{
			getEnv:   setEnv("CI", "true"),
			expected: Custom,
		},
		{
			getEnv:   setEnv("CI", "1"),
			expected: Custom,
		},
		{
			getEnv:   setEnv("BUILDER_OUTPUT", "1"),
			expected: Builder,
		},
		{
			getEnv:   setEnv("", ""),
			expected: NoCIDetected,
		},
	}

	for idx := range cases {
		c := cases[idx]
		t.Run(fmt.Sprintf("case #%d - %s", idx, c.expected), func(t *testing.T) {
			v := Detect(DetectFromEnv(c.getEnv))
			assert.Equal(t, c.expected, v)
		})
	}
}
