package ugentldap

import (
	"strings"

	"github.com/go-ldap/ldap/v3"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/models"
)

type UgentLdap struct {
	url      string
	username string
	password string
}

type UgentLdapConn struct {
	conn *ldap.Conn
}

type Config struct {
	Url      string
	Username string
	Password string
}

var ldapAttributes = []string{
	"objectClass",
	"uid",
	"ugentPreferredSn",
	"ugentPreferredGivenName",
	"ugentID",
	"ugentHistoricIDs",
	"ugentBirthDate",
	"mail",
	"ugentBarcode",
	"ugentJobCategory",
	"ugentAddressingTitle",
	"displayName",
	"ugentExpirationDate",
}

func NewClient(config Config) *UgentLdap {
	return &UgentLdap{
		url:      config.Url,
		username: config.Username,
		password: config.Password,
	}
}

func (cli *UgentLdap) NewConn() (*UgentLdapConn, error) {
	conn, err := ldap.DialURL(cli.url)
	if err != nil {
		return nil, err
	}

	if err = conn.Bind(cli.username, cli.password); err != nil {
		defer conn.Close()
		return nil, err
	}

	return &UgentLdapConn{conn}, nil
}

func (uc *UgentLdapConn) Close() error {
	return uc.conn.Close()
}

func (uc *UgentLdapConn) SearchPeople(filter string, cb func(*models.Person) error) error {
	searchReq := ldap.NewSearchRequest(
		"ou=people,dc=ugent,dc=be",
		ldap.ScopeSingleLevel,
		ldap.NeverDerefAliases,
		0, 0, false,
		filter,
		ldapAttributes,
		[]ldap.Control{},
	)

	/*
		Search with paging control, or SearchWithPaging(size)
		buffer all results into memory before returning it,
		using a lot of memory (250M). Now uses around 25K of memory.

		This is partly stolen from method SearchWithPaging
	*/
	pagingControl := ldap.NewControlPaging(2000)
	searchReq.Controls = append(searchReq.Controls, pagingControl)
	var cbErr error

	for {
		sr, err := uc.conn.Search(searchReq)
		if err != nil {
			return err
		}

		// pagingResult is hardly ever nil
		pagingResult := ldap.FindControl(sr.Controls, ldap.ControlTypePaging)
		if pagingResult == nil {
			pagingControl = nil
			break
		}

		for _, entry := range sr.Entries {
			if err := cb(mapToPerson(entry)); err != nil {
				cbErr = err
				break
			}
		}
		if cbErr != nil {
			break
		}

		// cookie is a cursor to the next page
		cookie := pagingResult.(*ldap.ControlPaging).Cookie
		if len(cookie) == 0 {
			// cookie is empty: server resources for paging are cleared automatically by the server
			pagingControl = nil
			break
		}
		pagingControl.SetCookie(cookie)
	}

	/*
		abandon paging: clear server side resources for paging.
		When callback returns an error, all server side resources
		for paging should be cleared/invalidated explicitly

		cf. https://www.ietf.org/rfc/rfc2696.txt:

		"A sequence of paged search requests is abandoned by the client
		sending a search request containing a pagedResultsControl with the
		size set to zero (0) and the cookie set to the last cookie returned
		by the server."
	*/
	if cbErr != nil && pagingControl != nil {
		pagingControl.PagingSize = 0
		if _, err := uc.conn.Search(searchReq); err != nil {
			return err
		}
	}

	return nil
}

func (cli *UgentLdap) SearchPeople(filter string, cb func(*models.Person) error) error {
	uc, err := cli.NewConn()
	if err != nil {
		return err
	}
	defer uc.Close()
	return uc.SearchPeople(filter, cb)
}

func mapToPerson(entry *ldap.Entry) *models.Person {
	np := models.NewPerson()
	np.Active = true

	for _, attr := range entry.Attributes {
		for _, val := range attr.Values {
			switch attr.Name {
			case "uid":
				np.OtherId = append(np.OtherId, &v1.IdRef{
					Type: "ugent_username",
					Id:   val,
				})
			// contains current active ugentID
			case "ugentID":
				np.OtherId = append(np.OtherId, &v1.IdRef{
					Type: "ugent_id",
					Id:   val,
				})
			// contains ugentID also (at the end)
			case "ugentHistoricIDs":
				np.OtherId = append(np.OtherId, &v1.IdRef{
					Type: "historic_ugent_id",
					Id:   val,
				})
			case "ugentBarcode":
				np.OtherId = append(np.OtherId, &v1.IdRef{
					Type: "ugent_barcode",
					Id:   val,
				})
			case "ugentPreferredGivenName":
				np.FirstName = val
			case "ugentPreferredSn":
				np.LastName = val
			case "displayName":
				np.FullName = val
			case "ugentBirthDate":
				np.BirthDate = val
			case "mail":
				np.Email = strings.ToLower(val)
			case "ugentJobCategory":
				np.JobCategory = append(np.JobCategory, val)
			case "ugentAddressingTitle":
				np.Title = val
			case "objectClass":
				np.ObjectClass = append(np.ObjectClass, val)
			case "ugentExpirationDate":
				np.ExpirationDate = val
			}
		}
	}
	return np
}
