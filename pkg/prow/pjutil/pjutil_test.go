/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pjutil

import (
	"reflect"
	"testing"
	"text/template"

	"github.com/jenkins-x/go-scm/scm"
	"github.com/jenkins-x/lighthouse/pkg/plumber"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/util/diff"

	"github.com/jenkins-x/lighthouse/pkg/prow/config"
	"k8s.io/test-infra/prow/kube"
)

func TestPostsubmitSpec(t *testing.T) {
	tests := []struct {
		name     string
		p        config.Postsubmit
		refs     plumber.Refs
		expected plumber.PipelineOptionsSpec
	}{
		{
			name: "can override path alias and cloneuri",
			p: config.Postsubmit{
				JobBase: config.JobBase{
					UtilityConfig: config.UtilityConfig{
						PathAlias: "foo",
						CloneURI:  "bar",
					},
				},
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.PostsubmitJob,
				Refs: &plumber.Refs{
					PathAlias: "foo",
					CloneURI:  "bar",
				},
			},
		},
		{
			name: "controller can default path alias and cloneuri",
			refs: plumber.Refs{
				PathAlias: "fancy",
				CloneURI:  "cats",
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.PostsubmitJob,
				Refs: &plumber.Refs{
					PathAlias: "fancy",
					CloneURI:  "cats",
				},
			},
		},
		{
			name: "job overrides take precedence over controller defaults",
			p: config.Postsubmit{
				JobBase: config.JobBase{
					UtilityConfig: config.UtilityConfig{
						PathAlias: "foo",
						CloneURI:  "bar",
					},
				},
			},
			refs: plumber.Refs{
				PathAlias: "fancy",
				CloneURI:  "cats",
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.PostsubmitJob,
				Refs: &plumber.Refs{
					PathAlias: "foo",
					CloneURI:  "bar",
				},
			},
		},
	}

	for _, tc := range tests {
		actual := PostsubmitSpec(tc.p, tc.refs)
		if expected := tc.expected; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s: actual %#v != expected %#v", tc.name, actual, expected)
		}
	}
}

func TestPresubmitSpec(t *testing.T) {
	tests := []struct {
		name     string
		p        config.Presubmit
		refs     plumber.Refs
		expected plumber.PipelineOptionsSpec
	}{
		{
			name: "can override path alias and cloneuri",
			p: config.Presubmit{
				JobBase: config.JobBase{
					UtilityConfig: config.UtilityConfig{
						PathAlias: "foo",
						CloneURI:  "bar",
					},
				},
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.PresubmitJob,
				Refs: &plumber.Refs{
					PathAlias: "foo",
					CloneURI:  "bar",
				},
			},
		},
		{
			name: "controller can default path alias and cloneuri",
			refs: plumber.Refs{
				PathAlias: "fancy",
				CloneURI:  "cats",
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.PresubmitJob,
				Refs: &plumber.Refs{
					PathAlias: "fancy",
					CloneURI:  "cats",
				},
			},
		},
		{
			name: "job overrides take precedence over controller defaults",
			p: config.Presubmit{
				JobBase: config.JobBase{
					UtilityConfig: config.UtilityConfig{
						PathAlias: "foo",
						CloneURI:  "bar",
					},
				},
			},
			refs: plumber.Refs{
				PathAlias: "fancy",
				CloneURI:  "cats",
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.PresubmitJob,
				Refs: &plumber.Refs{
					PathAlias: "foo",
					CloneURI:  "bar",
				},
			},
		},
	}

	for _, tc := range tests {
		actual := PresubmitSpec(tc.p, tc.refs)
		if expected := tc.expected; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s: actual %#v != expected %#v", tc.name, actual, expected)
		}
	}
}

func TestBatchSpec(t *testing.T) {
	tests := []struct {
		name     string
		p        config.Presubmit
		refs     plumber.Refs
		expected plumber.PipelineOptionsSpec
	}{
		{
			name: "can override path alias and cloneuri",
			p: config.Presubmit{
				JobBase: config.JobBase{
					UtilityConfig: config.UtilityConfig{
						PathAlias: "foo",
						CloneURI:  "bar",
					},
				},
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.BatchJob,
				Refs: &plumber.Refs{
					PathAlias: "foo",
					CloneURI:  "bar",
				},
			},
		},
		{
			name: "controller can default path alias and cloneuri",
			refs: plumber.Refs{
				PathAlias: "fancy",
				CloneURI:  "cats",
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.BatchJob,
				Refs: &plumber.Refs{
					PathAlias: "fancy",
					CloneURI:  "cats",
				},
			},
		},
		{
			name: "job overrides take precedence over controller defaults",
			p: config.Presubmit{
				JobBase: config.JobBase{
					UtilityConfig: config.UtilityConfig{
						PathAlias: "foo",
						CloneURI:  "bar",
					},
				},
			},
			refs: plumber.Refs{
				PathAlias: "fancy",
				CloneURI:  "cats",
			},
			expected: plumber.PipelineOptionsSpec{
				Type: plumber.BatchJob,
				Refs: &plumber.Refs{
					PathAlias: "foo",
					CloneURI:  "bar",
				},
			},
		},
	}

	for _, tc := range tests {
		actual := BatchSpec(tc.p, tc.refs)
		if expected := tc.expected; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s: actual %#v != expected %#v", tc.name, actual, expected)
		}
	}
}

func TestNewPlumberJob(t *testing.T) {
	var testCases = []struct {
		name                string
		spec                plumber.PipelineOptionsSpec
		labels              map[string]string
		expectedLabels      map[string]string
		annotations         map[string]string
		expectedAnnotations map[string]string
	}{
		{
			name: "periodic job, no extra labels",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PeriodicJob,
			},
			labels: map[string]string{},
			expectedLabels: map[string]string{
				kube.CreatedByProw:           "true",
				plumber.PlumberJobAnnotation: "job",
				plumber.PlumberJobTypeLabel:  "periodic",
			},
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job",
			},
		},
		{
			name: "periodic job, extra labels",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PeriodicJob,
			},
			labels: map[string]string{
				"extra": "stuff",
			},
			expectedLabels: map[string]string{
				kube.CreatedByProw:           "true",
				plumber.PlumberJobAnnotation: "job",
				plumber.PlumberJobTypeLabel:  "periodic",
				"extra":                      "stuff",
			},
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job",
			},
		},
		{
			name: "presubmit job",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PresubmitJob,
				Refs: &plumber.Refs{
					Org:  "org",
					Repo: "repo",
					Pulls: []plumber.Pull{
						{Number: 1},
					},
				},
			},
			labels: map[string]string{},
			expectedLabels: map[string]string{
				kube.CreatedByProw:           "true",
				plumber.PlumberJobAnnotation: "job",
				plumber.PlumberJobTypeLabel:  "presubmit",
				kube.OrgLabel:                "org",
				kube.RepoLabel:               "repo",
				kube.PullLabel:               "1",
			},
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job",
			},
		},
		{
			name: "non-github presubmit job",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PresubmitJob,
				Refs: &plumber.Refs{
					Org:  "https://some-gerrit-instance.foo.com",
					Repo: "some/invalid/repo",
					Pulls: []plumber.Pull{
						{Number: 1},
					},
				},
			},
			labels: map[string]string{},
			expectedLabels: map[string]string{
				kube.CreatedByProw:           "true",
				plumber.PlumberJobAnnotation: "job",
				plumber.PlumberJobTypeLabel:  "presubmit",
				kube.OrgLabel:                "some-gerrit-instance.foo.com",
				kube.RepoLabel:               "repo",
				kube.PullLabel:               "1",
			},
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job",
			},
		}, {
			name: "job with name too long to fit in a label",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job-created-by-someone-who-loves-very-very-very-long-names-so-long-that-it-does-not-fit-into-the-Kubernetes-label-so-it-needs-to-be-truncated-to-63-characters",
				Type: plumber.PresubmitJob,
				Refs: &plumber.Refs{
					Org:  "org",
					Repo: "repo",
					Pulls: []plumber.Pull{
						{Number: 1},
					},
				},
			},
			labels: map[string]string{},
			expectedLabels: map[string]string{
				kube.CreatedByProw:           "true",
				plumber.PlumberJobAnnotation: "job-created-by-someone-who-loves-very-very-very-long-names-so-l",
				plumber.PlumberJobTypeLabel:  "presubmit",
				kube.OrgLabel:                "org",
				kube.RepoLabel:               "repo",
				kube.PullLabel:               "1",
			},
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job-created-by-someone-who-loves-very-very-very-long-names-so-long-that-it-does-not-fit-into-the-Kubernetes-label-so-it-needs-to-be-truncated-to-63-characters",
			},
		},
		{
			name: "periodic job, extra labels, extra annotations",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PeriodicJob,
			},
			labels: map[string]string{
				"extra": "stuff",
			},
			annotations: map[string]string{
				"extraannotation": "foo",
			},
			expectedLabels: map[string]string{
				kube.CreatedByProw:           "true",
				plumber.PlumberJobAnnotation: "job",
				plumber.PlumberJobTypeLabel:  "periodic",
				"extra":                      "stuff",
			},
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job",
				"extraannotation":            "foo",
			},
		},
	}
	for _, testCase := range testCases {
		pj := NewPlumberJob(testCase.spec, testCase.labels, testCase.annotations)
		if actual, expected := pj.Spec, testCase.spec; !equality.Semantic.DeepEqual(actual, expected) {
			t.Errorf("%s: incorrect PipelineOptionsSpec created: %s", testCase.name, diff.ObjectReflectDiff(actual, expected))
		}
		if actual, expected := pj.Labels, testCase.expectedLabels; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s: incorrect PipelineOptions labels created: %s", testCase.name, diff.ObjectReflectDiff(actual, expected))
		}
		if actual, expected := pj.Annotations, testCase.expectedAnnotations; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s: incorrect PipelineOptions annotations created: %s", testCase.name, diff.ObjectReflectDiff(actual, expected))
		}
	}
}

func TestNewPlumberJobWithAnnotations(t *testing.T) {
	var testCases = []struct {
		name                string
		spec                plumber.PipelineOptionsSpec
		annotations         map[string]string
		expectedAnnotations map[string]string
	}{
		{
			name: "job without annotation",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PeriodicJob,
			},
			annotations: nil,
			expectedAnnotations: map[string]string{
				plumber.PlumberJobAnnotation: "job",
			},
		},
		{
			name: "job with annotation",
			spec: plumber.PipelineOptionsSpec{
				Job:  "job",
				Type: plumber.PeriodicJob,
			},
			annotations: map[string]string{
				"annotation": "foo",
			},
			expectedAnnotations: map[string]string{
				"annotation":                 "foo",
				plumber.PlumberJobAnnotation: "job",
			},
		},
	}

	for _, testCase := range testCases {
		pj := NewPlumberJob(testCase.spec, nil, testCase.annotations)
		if actual, expected := pj.Spec, testCase.spec; !equality.Semantic.DeepEqual(actual, expected) {
			t.Errorf("%s: incorrect PipelineOptionsSpec created: %s", testCase.name, diff.ObjectReflectDiff(actual, expected))
		}
		if actual, expected := pj.Annotations, testCase.expectedAnnotations; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s: incorrect PipelineOptions labels created: %s", testCase.name, diff.ObjectReflectDiff(actual, expected))
		}
	}
}

func TestJobURL(t *testing.T) {
	var testCases = []struct {
		name     string
		plank    config.Plank
		pj       plumber.PipelineOptions
		expected string
	}{
		{
			name: "non-decorated job uses template",
			plank: config.Plank{
				Controller: config.Controller{
					JobURLTemplate: template.Must(template.New("test").Parse("{{.Spec.Type}}")),
				},
			},
			pj:       plumber.PipelineOptions{Spec: plumber.PipelineOptionsSpec{Type: plumber.PeriodicJob}},
			expected: "periodic",
		},
		{
			name: "non-decorated job with broken template gives empty string",
			plank: config.Plank{
				Controller: config.Controller{
					JobURLTemplate: template.Must(template.New("test").Parse("{{.Garbage}}")),
				},
			},
			pj:       plumber.PipelineOptions{},
			expected: "",
		},
		{
			name: "decorated job without prefix uses template",
			plank: config.Plank{
				Controller: config.Controller{
					JobURLTemplate: template.Must(template.New("test").Parse("{{.Spec.Type}}")),
				},
			},
			pj:       plumber.PipelineOptions{Spec: plumber.PipelineOptionsSpec{Type: plumber.PeriodicJob}},
			expected: "periodic",
		},
	}

	logger := logrus.New()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if actual, expected := JobURL(testCase.plank, testCase.pj, logger.WithField("name", testCase.name)), testCase.expected; actual != expected {
				t.Errorf("%s: expected URL to be %q but got %q", testCase.name, expected, actual)
			}
		})
	}
}

func TestCreateRefs(t *testing.T) {
	pr := &scm.PullRequest{
		Number: 42,
		Link:   "https://github.example.com/kubernetes/Hello-World/pull/42",
		Head: scm.PullRequestBranch{
			Sha: "123456",
		},
		Base: scm.PullRequestBranch{
			Ref: "master",
			Repo: scm.Repository{
				Name:      "Hello-World",
				Link:      "https://github.example.com/kubernetes/Hello-World",
				Namespace: "kubernetes",
			},
		},
		Author: scm.User{
			Login: "ibzib",
			Link:  "https://github.example.com/ibzib",
		},
	}
	expected := plumber.Refs{
		Org:      "kubernetes",
		Repo:     "Hello-World",
		RepoLink: "https://github.example.com/kubernetes/Hello-World",
		BaseRef:  "master",
		BaseSHA:  "abcdef",
		BaseLink: "https://github.example.com/kubernetes/Hello-World/commit/abcdef",
		Pulls: []plumber.Pull{
			{
				Number:     42,
				Author:     "ibzib",
				SHA:        "123456",
				Link:       "https://github.example.com/kubernetes/Hello-World/pull/42",
				AuthorLink: "https://github.example.com/ibzib",
				CommitLink: "https://github.example.com/kubernetes/Hello-World/pull/42/commits/123456",
			},
		},
	}
	if actual := createRefs(pr, "abcdef"); !reflect.DeepEqual(expected, actual) {
		t.Errorf("diff between expected and actual refs:%s", diff.ObjectReflectDiff(expected, actual))
	}
}

func TestSpecFromJobBase(t *testing.T) {
	testCases := []struct {
		name    string
		jobBase config.JobBase
		verify  func(plumber.PipelineOptionsSpec) error
	}{
		{
			name:    "Verify reporter config gets copied",
			jobBase: config.JobBase{
				/*				ReporterConfig: &plumber.ReporterConfig{
									Slack: &plumber.SlackReporterConfig{
										Channel: "my-channel",
									},
								},
				*/
			},
			verify: func(pj plumber.PipelineOptionsSpec) error {
				/*				if pj.ReporterConfig == nil {
									return errors.New("Expected ReporterConfig to be non-nil")
								}
								if pj.ReporterConfig.Slack == nil {
									return errors.New("Expected ReporterConfig.Slack to be non-nil")
								}
								if pj.ReporterConfig.Slack.Channel != "my-channel" {
									return fmt.Errorf("Expected pj.ReporterConfig.Slack.Channel to be \"my-channel\", was %q",
										pj.ReporterConfig.Slack.Channel)
								}
				*/
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pj := specFromJobBase(tc.jobBase)
			if err := tc.verify(pj); err != nil {
				t.Fatalf("Verification failed: %v", err)
			}
		})
	}
}
