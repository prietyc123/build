// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMutateImageCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mutate Image Command Suite")
}
