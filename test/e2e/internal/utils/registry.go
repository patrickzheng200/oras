/*
Copyright The ORAS Authors.
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

package utils

import (
	"strings"

	"github.com/onsi/gomega"
	"oras.land/oras-go/v2/registry"
)

// Reference generates the reference string from given parameters.
func Reference(reg string, repo string, tagOrDigest string) string {
	ref := registry.Reference{
		Registry:   reg,
		Repository: repo,
		Reference:  strings.TrimSpace(tagOrDigest),
	}
	gomega.Expect(ref.Validate()).ShouldNot(gomega.HaveOccurred())
	return ref.String()
}
