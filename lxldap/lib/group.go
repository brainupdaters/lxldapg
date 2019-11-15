package lib

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/user"
	"strconv"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/bradfitz/slice"
	"gopkg.in/ldap.v3"
)

// LdapGroup structure for manage ldap groups
type LdapGroup struct {
	Cn        string
	Gidnumber string
}

// GetMaxGidOL return max GID from OpenLDAP
func GetMaxGidOL() int {
	var ldapServer = Config.Openldap.Server
	tempPort, _ := strconv.Atoi(Config.Openldap.Port)
	var ldapTLSPort = uint16(tempPort)

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Search for higgest uidNumber in OUusers
	var baseDN = Config.Openldap.OUgroups + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,                // The filter to apply
		[]string{"gidNumber"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	var gids []int

	for _, entry := range sr.Entries {
		gidnew, _ := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
		gids = append(gids, gidnew)
	}

	slice.Sort(gids[:], func(i, j int) bool {
		return gids[i] < gids[j]
	})

	return gids[len(gids)-1]
}

// GetMaxGidAD return max GID from AD
func GetMaxGidAD() int {

	//Open AD connection
	var ldapServer = Config.Ad.Server
	tempPort, _ := strconv.Atoi(Config.Ad.Port)
	var ldapTLSPort = uint16(tempPort)

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Search for higgest gidNumber in OUgroups
	var baseDN = Config.Ad.OUgroups + "," + Config.Ad.BaseDN
	var filter = "(&(objectClass=group)(gidNumber>=8000))"

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,                // The filter to apply
		[]string{"gidNumber"}, // A list attributes to retrieve
		nil,
	)

	l.Bind(Config.Ad.DN, Config.Ad.PS)
	if err != nil {
		log.Fatal(err)
	}

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	var gids []int

	for _, entry := range sr.Entries {
		gidnew, _ := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
		gids = append(gids, gidnew)
	}

	slice.Sort(gids[:], func(i, j int) bool {
		return gids[i] < gids[j]
	})

	return gids[len(gids)-1]
}

func GetGidCounterOL() int {
	var ldapServer = Config.Openldap.Server
	tempPort, _ := strconv.Atoi(Config.Openldap.Port)
	var ldapTLSPort = uint16(tempPort)

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Search for uidNumber in OUconfig
	baseDN := "uid=counters," + Config.Openldap.OUconfig + "," + Config.Openldap.BaseDN
	filter := "(&(objectClass=*))"

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,                // The filter to apply
		[]string{"gidNumber"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	GidNumber, _ := strconv.Atoi(sr.Entries[0].GetAttributeValue("gidNumber"))

	return GidNumber
}

func UpdateGidCounterOL(gidNumber int) {
	var ldapServer = Config.Openldap.Server
	tempPort, _ := strconv.Atoi(Config.Openldap.Port)
	var ldapTLSPort = uint16(tempPort)

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	l.Bind(Config.Openldap.DN, Config.Openldap.PS)
	if err != nil {
		log.Fatal(err)
	}

	//Modify gidNumber in OUconfig
	modify := ldap.NewModifyRequest("uid=counters,"+Config.Openldap.OUconfig+","+Config.Openldap.BaseDN, nil)
	modify.Replace("gidNumber", []string{strconv.Itoa(gidNumber)})

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("GID upadted to value " + strconv.Itoa(gidNumber))
}

func NextFreeGid(verbose bool) int {

	MaxGidAD := GetMaxGidAD()
	MaxGidOL := GetMaxGidOL()
	GidCounter := GetGidCounterOL()

	if verbose {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Max GID from AD", "Max GID from OL", "GID OL Counter"})
		table.Append([]string{strconv.Itoa(MaxGidAD), strconv.Itoa(MaxGidOL), strconv.Itoa(GidCounter)})
		table.Render()
	}

	if (MaxGidAD >= GidCounter) || (MaxGidOL >= GidCounter) {
		UpdateGidCounterOL(Max(MaxGidAD, MaxGidOL) + 1)
		return Max(MaxGidAD, MaxGidOL) + 1
	} else {
		return GidCounter
	}

}

func ListGroupOL() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	log := log.WithFields(
		log.Fields{
			"username": user.Username,
			"uid":      user.Uid,
		})

	log.Info("Request group list from command line")

	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUgroups + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,                            // The filter to apply
		[]string{"dn", "cn", "gidNumber"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	var groups []LdapGroup
	for _, entry := range sr.Entries {
		groups = append(groups, LdapGroup{Cn: entry.GetAttributeValue("cn"), Gidnumber: entry.GetAttributeValue("gidNumber")})
	}

	slice.Sort(groups[:], func(i, j int) bool {
		return groups[i].Gidnumber < groups[j].Gidnumber
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cn", "Gid"})

	for _, entry := range groups {
		table.Append([]string{entry.Cn, entry.Gidnumber})
	}

	table.Render()
}
