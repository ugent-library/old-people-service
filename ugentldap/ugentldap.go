package ugentldap

import (
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/ugent-library/people-service/models"
)

type Client struct {
	url      string
	username string
	password string
}

type ClientConn struct {
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
	"departmentNumber",
}

func NewClient(config Config) *Client {
	return &Client{
		url:      config.Url,
		username: config.Username,
		password: config.Password,
	}
}

func (cli *Client) NewConn() (*ClientConn, error) {
	conn, err := ldap.DialURL(cli.url)
	if err != nil {
		return nil, err
	}

	if err = conn.Bind(cli.username, cli.password); err != nil {
		defer conn.Close()
		return nil, err
	}

	return &ClientConn{conn}, nil
}

func (conn *ClientConn) Close() error {
	return conn.conn.Close()
}

func (conn *ClientConn) SearchPeople(filter string, cb func(*models.Person) error) error {
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
		sr, err := conn.conn.Search(searchReq)
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
			if err := cb(mapToDummyPerson(entry)); err != nil {
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
		if _, err := conn.conn.Search(searchReq); err != nil {
			return err
		}
	}

	return nil
}

func (cli *Client) SearchPeople(filter string, cb func(*models.Person) error) error {
	uc, err := cli.NewConn()
	if err != nil {
		return err
	}
	defer uc.Close()
	return uc.SearchPeople(filter, cb)
}

func (cli *Client) SearchStudents(cb func(*models.Person) error) error {
	return cli.SearchPeople("(objectClass=ugentStudent)", cb)
}

func mapToDummyPerson(entry *ldap.Entry) *models.Person {
	np := models.NewPerson()
	np.Active = true

	for _, attr := range entry.Attributes {
		for _, val := range attr.Values {
			switch attr.Name {
			case "uid":
				np.OtherId = append(np.OtherId, &models.IdRef{
					Type: "ugent_username",
					Id:   val,
				})
			// contains current active ugentID
			case "ugentID":
				np.OtherId = append(np.OtherId, &models.IdRef{
					Type: "ugent_id",
					Id:   val,
				})
			// contains ugentID also (at the end)
			case "ugentHistoricIDs":
				np.OtherId = append(np.OtherId, &models.IdRef{
					Type: "historic_ugent_id",
					Id:   val,
				})
			case "ugentBarcode":
				np.OtherId = append(np.OtherId, &models.IdRef{
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
			case "departmentNumber":
				//OrganizationRef#Id is here an ugent id (e.g. CA20)
				np.Organization = append(np.Organization, models.NewOrganizationRef(val))
			}
		}
	}
	return np
}
