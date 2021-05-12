// Copyright (C) 2020 by Russell Smith <https://github.com/ukd1>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package detectci allows you to detect if, and what kind of CI service
// your code is running on, allowing you to behave differently when needed

package detectci

import (
	"os"
)

type ciIndicator struct {
	name string
}

var ciIndicatiors = map[string]ciIndicator{
	// This list is based on the list found at:
	// https://github.com/npm/ci-detect/blob/master/test/index.js
	"GERRIT_PROJECT":                     ciIndicator{name: "gerrit"},
	"GITLAB_CI":                          ciIndicator{name: "gitlab"},
	"CIRCLECI":                           ciIndicator{name: "circle-ci"},
	"TEAMCITY_VERSION":                   ciIndicator{name: "teamcity"},
	"SEMAPHORE":                          ciIndicator{name: "semaphore"},
	"DRONE":                              ciIndicator{name: "drone"},
	"GITHUB_ACTION":                      ciIndicator{name: "github-actions"},
	"TDDIUM":                             ciIndicator{name: "tddium"},
	"JENKINS_URL":                        ciIndicator{name: "jenkins"},
	"WERCKER":                            ciIndicator{name: "wercker"},
	"NETLIFY":                            ciIndicator{name: "netlify"},
	"NOW_GITHUB_DEPLOYMENT":              ciIndicator{name: "now-for-github"},
	"GITLAB_DEPLOYMENT":                  ciIndicator{name: "now-for-gitlab"},
	"BITBUCKET_DEPLOYMENT":               ciIndicator{name: "now-for-bitbucket"},
	"bamboo.buildKey":                    ciIndicator{name: "bamboo"},
	"GO_PIPELINE_NAME":                   ciIndicator{name: "gocd"},
	"TRAVIS":                             ciIndicator{name: "travis-ci"},
	"APPVEYOR":                           ciIndicator{name: "appveyor"},
	"CODEBUILD_BUILD_ID":                 ciIndicator{name: "aws-codebuild"},
	"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI": ciIndicator{name: "azure-pipelines"},
	"BITRISE_IO":                         ciIndicator{name: "bitrise"},
	"BUDDY_WORKSPACE_ID":                 ciIndicator{name: "buddy"},
	"BUILDKITE":                          ciIndicator{name: "buildkite"},
	"CIRRUS_CI":                          ciIndicator{name: "cirrus"},
	"DSARI":                              ciIndicator{name: "dsari"},
	"STRIDER":                            ciIndicator{name: "strider"},
	"TASKCLUSTER_ROOT_URL":               ciIndicator{name: "taskcluster"},
	"HUDSON_URL":                         ciIndicator{name: "hudson"},
	"NOW_BUILDER":                        ciIndicator{name: "now"},
	"MAGNUM":                             ciIndicator{name: "magnum"},
	"NEVERCODE":                          ciIndicator{name: "nevercode"},
	"RENDER":                             ciIndicator{name: "render"},
	"SAIL_CI":                            ciIndicator{name: "sail"},
	"SHIPPABLE":                          ciIndicator{name: "shippable"},
}

func IsCI() bool {
	found_ci := false

	if _, ok := os.LookupEnv("CI"); ok {
		found_ci = true
	} else {
		for env, _ := range ciIndicatiors {
			if _, ok := os.LookupEnv(env); ok {
				found_ci = true
				break
			}
		}
	}

	return found_ci
}

func WhichCI() (bool, string) {
	found_ci := false
	var ci ciIndicator

	for env, info := range ciIndicatiors {
		if _, ok := os.LookupEnv(env); ok {
			found_ci = true
			ci = info
			break
		}
	}

	if !found_ci {
		if _, ok := os.LookupEnv("CI"); ok {
			return true, "unknown"
		} else {
			return false, ""
		}
	} else {
		return true, ci.name
	}
}
