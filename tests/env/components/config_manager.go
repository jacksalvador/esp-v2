// Copyright 2018 Google Cloud Platform Proxy Authors
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

package components

import (
	"os"
	"os/exec"

	"github.com/golang/glog"
)

const (
	xDSPath = "../../bin/configmanager"
)

type ConfigManagerServer struct {
	*Cmd
}

func NewConfigManagerServer(debugMode bool, args []string) (*ConfigManagerServer, error) {
	if debugMode {
		args = append(args, "--logtostderr", "--v=2")
	}
	glog.Infof("config manager args: %v", args)
	cmd := exec.Command(xDSPath, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return &ConfigManagerServer{
		Cmd: &Cmd{
			name: "ConfigManager",
			Cmd:  cmd,
		},
	}, nil
}
