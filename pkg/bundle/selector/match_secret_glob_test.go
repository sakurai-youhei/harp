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

package selector

import (
	"testing"

	bundlev1 "github.com/elastic/harp/api/gen/go/harp/bundle/v1"
)

func TestMatchSecretGlob_MatchesSecretKey(t *testing.T) {
	p := &bundlev1.Package{
		Name: "app/target",
		Secrets: &bundlev1.SecretChain{
			Data: []*bundlev1.KV{{Key: "k"}},
		},
	}

	tests := []struct {
		name    string
		pattern string
		want    bool
	}{
		{name: "secret key", pattern: "k", want: true},
		{name: "secret key wildcard", pattern: "*", want: true},
		{name: "package path pattern", pattern: "app/*", want: false},
		{name: "package path", pattern: "app/target", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchSecretGlob(tt.pattern).IsSatisfiedBy(p); got != tt.want {
				t.Errorf("IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
