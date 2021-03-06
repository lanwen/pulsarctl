// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package functions

import (
	"errors"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/spf13/cobra"
)

var checkPutStateArgs = func(args []string) error {
	if len(args) < 2 {
		return errors.New("need to specified the state key and state value")
	}
	return nil
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"functions",
		"Interface for managing Pulsar Functions "+
			"(lightweight, Lambda-style compute processes that work with Pulsar)",
		"",
		"pf",
	)

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, stopFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, startFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, restartFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, statusFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, statsFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, querystateFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, putstateFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, triggerFunctionsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, downloadFunctionsCmd)

	return resourceCmd
}
