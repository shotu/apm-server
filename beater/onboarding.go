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

package beater

import (
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

func notifyListening(config *Config, pubFct func(beat.Event)) {

	var isServerUp = func() bool {
		secure := config.SSL.isEnabled()
		return isServerUp(secure, config.Host, 10, time.Second)
	}

	if isServerUp() {
		logp.NewLogger("onboarding").Info("Publishing onboarding document")

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"processor": common.MapStr{"name": "onboarding", "event": "onboarding"},
				"observer":  common.MapStr{"listening": config.Host},
			},
		}
		pubFct(event)
	}
}
