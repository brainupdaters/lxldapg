// Copyright © 2019 Pau Roura <pau@brainupdaters.net>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"lxldap/lib"
	"os"
	"os/user"
	"strconv"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userFreeCmd = &cobra.Command{
	Use:   "free",
	Short: "show next free UID",
	Long:  `Returns the next User Identifier {uidNumber} to assing in a new group.`,
	Run:   runFreeUser,
}

func init() {
	userCmd.AddCommand(userFreeCmd)
}

func runFreeUser(cmd *cobra.Command, args []string) {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	log := log.WithFields(
		log.Fields{
			"username": user.Username,
			"uid":      user.Uid,
		})

	log.Info("Request free user Id from command line")

	NextFree := lib.NextFreeUid(Verbose)

	if Verbose {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Next Free UID"})
		table.Append([]string{strconv.Itoa(NextFree)})
		table.Render()
	} else {
		fmt.Println(NextFree)
	}
}
