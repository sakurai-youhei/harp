// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package patch

import (
	"testing"

	bundlev1 "github.com/elastic/harp/api/gen/go/harp/bundle/v1"
)

func TestApply_InvalidCreateRule(t *testing.T) {
	tests := []struct {
		name string
		rule *bundlev1.PatchRule
	}{
		{
			name: "create with matchSecret selector",
			rule: &bundlev1.PatchRule{
				Selector: &bundlev1.PatchSelector{
					MatchSecret: &bundlev1.PatchSelectorMatchSecret{Strict: "password"},
				},
				Package: &bundlev1.PatchPackage{Create: true},
			},
		},
		{
			name: "create without selector",
			rule: &bundlev1.PatchRule{
				Package: &bundlev1.PatchPackage{Create: true},
			},
		},
		{
			name: "create with matchPath glob selector",
			rule: &bundlev1.PatchRule{
				Selector: &bundlev1.PatchSelector{
					MatchPath: &bundlev1.PatchSelectorMatchPath{Glob: "app/*"},
				},
				Package: &bundlev1.PatchPackage{Create: true},
			},
		},
		{
			name: "rule without package",
			rule: &bundlev1.PatchRule{
				Selector: &bundlev1.PatchSelector{
					MatchPath: &bundlev1.PatchSelectorMatchPath{Strict: "app/target"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := &bundlev1.Patch{
				ApiVersion: "harp.elastic.co/v1",
				Kind:       "BundlePatch",
				Meta:       &bundlev1.PatchMeta{Name: "test"},
				Spec:       &bundlev1.PatchSpec{Rules: []*bundlev1.PatchRule{tt.rule}},
			}
			b := &bundlev1.Bundle{Packages: []*bundlev1.Package{{Name: "app/target"}}}

			// Must return an error instead of panicking.
			if _, err := Apply(spec, b, map[string]interface{}{}); err == nil {
				t.Fatal("Apply() expected an error, got nil")
			}
		})
	}
}
