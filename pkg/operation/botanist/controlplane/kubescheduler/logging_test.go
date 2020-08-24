// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubescheduler_test

import (
	. "github.com/gardener/gardener/pkg/operation/botanist/controlplane/kubescheduler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logging", func() {
	Describe("#LoggingConfiguration", func() {
		It("should return the expected logging parser and filter", func() {
			parser, filter, err := LoggingConfiguration()

			Expect(err).NotTo(HaveOccurred())
			Expect(parser).To(Equal(`[PARSER]
    Name        kubeSchedulerParser
    Format      regex
    Regex       ^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<log>.*)$
    Time_Key    time
    Time_Format %m%d %H:%M:%S.%L
`))

			Expect(filter).To(Equal(`[FILTER]
    Name                parser
    Match               kubernetes.*kube-scheduler*kube-scheduler*
    Key_Name            log
    Parser              kubeSchedulerParser
    Reserve_Data        True
`))
		})
	})
})
