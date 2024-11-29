// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package internal

import (
	"testing"
)

func TestSupportedVersions(t *testing.T) {
	supportedVersions := SupportedVersions()

	contains := func(arr []string, find string) bool {
		for _, i := range arr {
			if i == find {
				return true
			}
		}
		return false
	}

	for _, v := range []string{"0.1.0", "0.2.0"} {
		if contains(supportedVersions, v) {
			t.Errorf("expected %s to not be a supported version", v)
		}
	}

	for _, v := range []string{"0.3.0", "0.3.1", "0.4.0", "1.0.0", "1.1.0"} {
		if !contains(supportedVersions, v) {
			t.Errorf("expected %s to be a supported version", v)
		}
	}
}
