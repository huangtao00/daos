//
// (C) Copyright 2018-2019 Intel Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
// The Government's rights to use, modify, reproduce, release, perform, display,
// or disclose this software are subject to the terms of the Apache License as
// provided in Contract No. 8F-30005.
// Any reproduction of computer software, computer software documentation, or
// portions thereof marked with this legend must also reproduce the markings.
//

package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// AssertTrue asserts b is true
func AssertTrue(t *testing.T, b bool, message string) {
	t.Helper()

	if !b {
		t.Fatal(message)
	}
}

// AssertFalse asserts b is false
func AssertFalse(t *testing.T, b bool, message string) {
	t.Helper()

	if b {
		t.Fatal(message)
	}
}

// AssertEqual asserts b is equal to a
//
// Whilst suitable in most situations, reflect.DeepEqual() may not be
// suitable for nontrivial struct element comparisons, go-cmp should
// then be used but will introduce a third party dep.
func AssertEqual(
	t *testing.T, a interface{}, b interface{}, message string) {
	t.Helper()

	if reflect.DeepEqual(a, b) {
		return
	}
	if len(message) > 0 {
		message += ", "
	}

	t.Fatalf(message+"%#v != %#v", a, b)
}

// AssertStringsEqual sorts string slices before comparing.
func AssertStringsEqual(
	t *testing.T, a []string, b []string, message string) {
	t.Helper()

	sort.Strings(a)
	sort.Strings(b)

	AssertEqual(t, a, b, message)
}

// ExpectError asserts error contains expected message
func ExpectError(t *testing.T, actualErr error, expectedMessage string, desc interface{}) {
	t.Helper()

	if actualErr == nil {
		if expectedMessage != "" {
			t.Fatalf("expected a non-nil error: %v", desc)
		}
	} else if diff := cmp.Diff(expectedMessage, actualErr.Error()); diff != "" {
		t.Fatalf("unexpected error (-want, +got):\n%s\n", diff)
	}
}

// CmpErrBool compares two errors and returns a boolean value indicating equality
// or at least close similarity between their messages.
func CmpErrBool(want, got error) bool {
	if want == got {
		return true
	}

	if want == nil || got == nil {
		return false
	}
	if !strings.Contains(got.Error(), want.Error()) {
		return false
	}

	return true
}

// CmpErr compares two errors for equality or at least close similarity in their messages.
func CmpErr(t *testing.T, want, got error) {
	t.Helper()

	if !CmpErrBool(want, got) {
		t.Fatalf("unexpected error\n(wanted: %v, got: %v)", want, got)
	}
}

// LoadTestFiles reads inputs and outputs from file and do basic sanity checks.
// Both files contain entries of multiple lines separated by blank line.
// Return inputs and outputs, both of which are slices of string slices.
func LoadTestFiles(inFile string, outFile string) (
	inputs [][]string, outputs [][]string, err error) {

	inputs, err = SplitFile(inFile)
	if err != nil {
		return
	}
	outputs, err = SplitFile(outFile)
	if err != nil {
		return
	}

	if len(inputs) < 1 {
		err = fmt.Errorf("no inputs read from file")
	} else if len(inputs) != len(outputs) {
		err = fmt.Errorf("number of inputs and outputs not equal")
	}

	return
}

// ShowBufferOnFailure displays captured output on test failure. Should be run
// via defer in the test function.
func ShowBufferOnFailure(t *testing.T, buf fmt.Stringer) {
	t.Helper()

	if t.Failed() {
		fmt.Printf("captured log output:\n%s", buf.String())
	}
}

// DefaultCmpOpts gets default go-cmp comparison options for tests.
func DefaultCmpOpts() []cmp.Option {
	// Avoid comparing the internal Protobuf fields
	isHiddenPBField := func(path cmp.Path) bool {
		if strings.HasPrefix(path.Last().String(), ".XXX_") {
			return true
		}
		return false
	}
	return []cmp.Option{
		cmp.FilterPath(isHiddenPBField, cmp.Ignore()),
	}
}

// CreateTestDir creates a temporary test directory.
// It returns the path to the directory and a cleanup function.
func CreateTestDir(t *testing.T) (string, func()) {
	t.Helper()

	name := strings.Replace(t.Name(), "/", "-", -1)
	tmpDir, err := ioutil.TempDir("", name)
	if err != nil {
		t.Fatalf("Couldn't create temporary directory: %v", err)
	}

	return tmpDir, func() {
		os.RemoveAll(tmpDir)
	}
}
