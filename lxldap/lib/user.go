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

// LdapUser structure for manage ldap users
type LdapUser struct {
	Cn        string
	Uidnumber string
	Gidnumber string
	Sn        string
}

// GetMaxUidOL return max UID from OpenLDAP
func GetMaxUidOL() int {
	var ldapServer = Config.Openldap.Server
	tempPort, _ := strconv.Atoi(Config.Openldap.Port)
	var ldapTLSPort = uint16(tempPort)

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Search for higgest uidNumber in OUusers
	var baseDN = Config.Openldap.OUusers + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,                // The filter to apply
		[]string{"uidNumber"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	var uids []int

	for _, entry := range sr.Entries {
		uidnew, _ := strconv.Atoi(entry.GetAttributeValue("uidNumber"))
		uids = append(uids, uidnew)
	}

	slice.Sort(uids[:], func(i, j int) bool {
		return uids[i] < uids[j]
	})

	return uids[len(uids)-1]
}

// GetMaxUidAD return max UID from AD
func GetMaxUidAD() int {

	//Open AD connection
	var ldapServer = Config.Ad.Server
	tempPort, _ := strconv.Atoi(Config.Ad.Port)
	var ldapTLSPort = uint16(tempPort)

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Search for higgest uidNumber in OUusers
	var baseDN = Config.Ad.OUusers + "," + Config.Ad.BaseDN
	var filter = "(&(objectClass=user)(gidNumber>=8000)(uidNumber>=9000))"

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,                // The filter to apply
		[]string{"uidNumber"}, // A list attributes to retrieve
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

	var uids []int

	for _, entry := range sr.Entries {
		uidnew, _ := strconv.Atoi(entry.GetAttributeValue("uidNumber"))
		uids = append(uids, uidnew)
	}

	slice.Sort(uids[:], func(i, j int) bool {
		return uids[i] < uids[j]
	})

	return uids[len(uids)-1]
}

func GetUidCounterOL() int {
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
		[]string{"uidNumber"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	UidNumber, _ := strconv.Atoi(sr.Entries[0].GetAttributeValue("uidNumber"))

	return UidNumber
}

func UpdateUidCounterOL(uidNumber int) {
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

	//Modify uidNumber in OUconfig
	modify := ldap.NewModifyRequest("uid=counters,"+Config.Openldap.OUconfig+","+Config.Openldap.BaseDN, nil)
	modify.Replace("uidNumber", []string{strconv.Itoa(uidNumber)})

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("UID upadted to value " + strconv.Itoa(uidNumber))
}

func NextFreeUid(verbose bool) int {

	MaxUidAD := GetMaxUidAD()
	MaxUidOL := GetMaxUidOL()
	UidCounter := GetUidCounterOL()

	if verbose {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Max UID from AD", "Max UID from OL", "UID OL Counter"})
		table.Append([]string{strconv.Itoa(MaxUidAD), strconv.Itoa(MaxUidOL), strconv.Itoa(UidCounter)})
		table.Render()
	}

	if (MaxUidAD >= UidCounter) || (MaxUidOL >= UidCounter) {
		UpdateUidCounterOL(Max(MaxUidAD, MaxUidOL) + 1)
		return Max(MaxUidAD, MaxUidOL) + 1
	} else {
		return UidCounter
	}

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// ListUserOL list all users
func ListUserOL() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	log := log.WithFields(
		log.Fields{
			"username": user.Username,
			"uid":      user.Uid,
		})

	log.Info("Request user list from command line")

	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUusers + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // The filter to apply
		[]string{"dn", "cn", "uidNumber", "gidNumber", "sn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	var users []LdapUser
	for _, entry := range sr.Entries {
		users = append(users, LdapUser{Cn: entry.GetAttributeValue("cn"), Uidnumber: entry.GetAttributeValue("uidNumber"), Gidnumber: entry.GetAttributeValue("gidNumber"), Sn: entry.GetAttributeValue("sn")})
	}

	slice.Sort(users[:], func(i, j int) bool {
		return users[i].Uidnumber < users[j].Uidnumber
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cn", "Uid", "Gid", "Name"})

	for _, entry := range users {
		table.Append([]string{entry.Cn, entry.Uidnumber, entry.Gidnumber, entry.Sn})
	}

	table.Render()
}
