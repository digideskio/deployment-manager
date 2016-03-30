/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package repo

import (
	"testing"
)

var (
	TestRepoBucket         = "kubernetes-charts-testing"
	TestRepoURL            = "gs://" + TestRepoBucket
	TestRepoType           = GCSRepoType
	TestRepoFormat         = GCSRepoFormat
	TestRepoCredentialName = "default"
)

func TestValidRepoURL(t *testing.T) {
	tr, err := NewRepo(TestRepoURL, TestRepoCredentialName, TestRepoBucket, string(TestRepoFormat), string(TestRepoType))
	if err != nil {
		t.Fatal(err)
	}

	if err := validateRepo(tr, TestRepoURL, TestRepoCredentialName, TestRepoFormat, TestRepoType); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidRepoURL(t *testing.T) {
	_, err := newRepo("%:invalid&url:%", TestRepoCredentialName, TestRepoBucket, TestRepoFormat, TestRepoType)
	if err == nil {
		t.Fatalf("expected error did not occur for invalid URL")
	}
}

func TestDefaultCredentialName(t *testing.T) {
	tr, err := newRepo(TestRepoURL, "", TestRepoBucket, TestRepoFormat, TestRepoType)
	if err != nil {
		t.Fatalf("cannot create repo using default credential name")
	}

	TestRepoCredentialName := "default"
	haveCredentialName := tr.GetCredentialName()
	if haveCredentialName != TestRepoCredentialName {
		t.Fatalf("unexpected credential name; want: %s, have %s.", TestRepoCredentialName, haveCredentialName)
	}
}

func TestInvalidRepoFormat(t *testing.T) {
	_, err := newRepo(TestRepoURL, TestRepoCredentialName, TestRepoBucket, "", TestRepoType)
	if err == nil {
		t.Fatalf("expected error did not occur for invalid format")
	}
}

func TestValidateRepoURL(t *testing.T) {
	validURLs := []string{
		"https://host/bucket",
		"http://host/bucket",
		"gs://bucket-name",
	}
	invalidURL := "charts"

	for _, url := range validURLs {
		err := ValidateRepoURL(url)
		if err != nil {
			t.Fatalf("Expected repo url: %v to be valid but threw an error", url)
		}
	}

	err := ValidateRepoURL(invalidURL)
	if err == nil {
		t.Fatalf("Expected repo url: %v to be invalid but did not throw an error", invalidURL)
	}
}
