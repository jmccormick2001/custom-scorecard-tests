// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/operator-framework/api/pkg/manifests"
	"github.com/operator-framework/operator-registry/pkg/registry"
	"github.com/sirupsen/logrus"

	"github.com/jmccormick2001/custom-scorecard-tests/internal/tests"
	scapiv1alpha2 "github.com/operator-framework/operator-sdk/pkg/apis/scorecard/v1alpha2"
)

// this is the custom scorecard test example binary
// As with the Redhat scorecard test image, the bundle that is under
// test is expected to be mounted so that tests can inspect the
// bundle contents as part of their test implementations.
// The actual test is to be run is named and that name is passed
// as an argument to this binary.  This argument mechanism allows
// this binary to run various tests all from within a single
// test image.

func main() {
	entrypoint := os.Args[1:]
	if len(entrypoint) == 0 {
		log.Fatal("test name argument is required")
	}

	// Read the pod's untar'd bundle from a well-known path.
	cfg, err := GetBundle("/bundle")
	if err != nil {
		log.Fatal(err.Error())
	}

	var result scapiv1alpha2.ScorecardTestResult

	switch entrypoint[0] {
	case tests.CustomTest1Name:
		result = tests.CustomTest1(*cfg)
	case tests.CustomTest2Name:
		result = tests.CustomTest2(*cfg)
	default:
		result = printValidTests()
	}

	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

}

// printValidTests will print out full list of test names to give a hint to the end user on what the valid tests are
func printValidTests() (result scapiv1alpha2.ScorecardTestResult) {
	result.State = scapiv1alpha2.FailState
	result.Errors = make([]string, 0)
	result.Suggestions = make([]string, 0)

	str := fmt.Sprintf("Valid tests for this image include: %s, %s",
		tests.CustomTest1Name,
		tests.CustomTest2Name)
	result.Errors = append(result.Errors, str)
	return result
}

// GetBundle parses a Bundle from a given on-disk path returning a bundle
func GetBundle(bundlePath string) (bundle *registry.Bundle, err error) {

	// validate the path
	if _, err := os.Stat(bundlePath); os.IsNotExist(err) {
		return nil, err
	}

	validationLogOutput := new(bytes.Buffer)
	origOutput := logrus.StandardLogger().Out
	logrus.SetOutput(validationLogOutput)
	defer logrus.SetOutput(origOutput)

	// TODO evaluate another API call that would support the new
	// bundle format
	var bundles []*registry.Bundle
	//var bundleErrors []errors.ManifestResult
	_, bundles, _ = manifests.GetManifestsDir(bundlePath)

	if len(bundles) == 0 {
		return nil, fmt.Errorf("bundle was not found")
	}
	if bundles[0] == nil {
		return nil, fmt.Errorf("bundle is invalid nil value")
	}
	bundle = bundles[0]
	_, err = bundle.ClusterServiceVersion()
	if err != nil {
		return nil, fmt.Errorf("error in csv retrieval %s", err.Error())
	}

	return bundle, nil
}
