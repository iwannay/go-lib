// Copyright (c) 2018 Baidu, Inc.
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

package metrics

import (
	"testing"
)

func TestStateGet(t *testing.T) {
	var s State
	if s.Get() != "" {
		t.Errorf("init state expect \"\", but is:%s", s.Get())
	}
}

func TestStateSet(t *testing.T) {
	var s State
	s.Set("test")
	if s.Get() != "test" {
		t.Errorf("after set, expected to be test, but is:%s", s.Get())
	}
}
