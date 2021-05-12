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
	"testing"
)

func TestIsCI(t *testing.T) {
	os.Unsetenv("CI")
	os.Unsetenv("GITHUB_ACTION")
	if IsCI() {
		t.Error("Detected CI, but shouldn't have")
	}

	os.Setenv("CI", "1")
	if !IsCI() {
		t.Error("Did not detect CI")
	}
	os.Unsetenv("CI")

	os.Setenv("GITHUB_ACTION", "1")
	if !IsCI() {
		t.Error("Did not detect CI")
	}
}

func TestWhichCI(t *testing.T) {
	// doesn't detect when noting set
	os.Unsetenv("CI")
	os.Unsetenv("GITHUB_ACTION")
	found, ci_name := WhichCI()
	if found {
		t.Error("Detected CI, but shouldn't have")
	}

	// detects generic CI
	os.Setenv("CI", "1")
	found, ci_name = WhichCI()
	if  !found {
		t.Error("Did not detect CI")
	} else {
		if ci_name != "unknown" {
			t.Errorf("Should have found 'unknown' CI, instead %v", ci_name)
		}
	}
	os.Unsetenv("CI")

	// detects github actions 
	os.Setenv("GITHUB_ACTION", "1")
	found, ci_name = WhichCI();
	if  !found {
		t.Error("Did not detect CI")
	} else {
		if ci_name != "github-actions" {
			t.Errorf("Should have found 'github-actions' CI, instead %v", ci_name)
		}
	}

	// detects github actions before an unknown CI
	os.Setenv("CI", "1")
	os.Setenv("GITHUB_ACTION", "1")
	found, ci_name = WhichCI();
	if  !found {
		t.Error("Did not detect CI")
	} else {
		if ci_name != "github-actions" {
			t.Errorf("Should have found 'github-actions' CI, instead %v", ci_name)
		}
	}
}
