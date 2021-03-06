/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package usage

import (
	"fmt"
	"os"
)

var usageStr = `
Usage: %s [options]
Server Options:
	--conf				Specify local configuration file path
Common Options:
	-h, --help			Show this message
`

// usage will print out the flag options for the server.
func HelpCallback() {
	fmt.Printf(usageStr, os.Args[0])
	os.Exit(0)
}
