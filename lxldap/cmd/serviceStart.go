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
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start lxldap API server",
	Long:  `Start lxldap API server.`,
	Run:   runStartServer,
}

func init() {
	serviceCmd.AddCommand(serviceStartCmd)
}

func runStartServer(cmd *cobra.Command, args []string) {
	router := mux.NewRouter()
	router.HandleFunc("/", rootResponse).Methods("GET")
	router.HandleFunc("/user/free", userFree).Methods("GET")
	router.HandleFunc("/group/free", groupFree).Methods("GET")
	log.Info("Starting OpenLDAP info service in localhost:" + lib.Config.Apiserver.Port)
	fmt.Println("Starting OpenLDAP info service in localhost:" + lib.Config.Apiserver.Port)
	log.Fatal(http.ListenAndServeTLS(":"+lib.Config.Apiserver.Port, lib.Config.Apiserver.Cert, lib.Config.Apiserver.Key, router))
}
func rootResponse(w http.ResponseWriter, r *http.Request) {

	log := log.WithFields(
		log.Fields{
			"ip":  r.RemoteAddr,
			"URI": r.RequestURI,
		})

	log.Info("Request root to api service")

	w.Write([]byte("OpenLDAP Service"))
}
func userFree(w http.ResponseWriter, r *http.Request) {

	log := log.WithFields(
		log.Fields{
			"ip":  r.RemoteAddr,
			"URI": r.RequestURI,
		})

	log.Info("Request free user Id to api service")

	w.Write([]byte(strconv.Itoa(lib.NextFreeUid(false))))
}
func groupFree(w http.ResponseWriter, r *http.Request) {

	log := log.WithFields(
		log.Fields{
			"ip":  r.RemoteAddr,
			"URI": r.RequestURI,
		})

	log.Info("Request free group Id to api service")

	w.Write([]byte(strconv.Itoa(lib.NextFreeGid(false))))
}
