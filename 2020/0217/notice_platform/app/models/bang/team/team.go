package team

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

const (
	TeamStatusNormal   = 1
	TeamStatusDisabled = 2
)

type Team struct {
	UUID        string `db:"uuid" json:"uuid"`
	Status      int    `db:"status" json:"status"`
	Name        string `db:"name" json:"name"`
	Owner       string `db:"owner" json:"owner"`
	Logo        string `db:"logo" json:"logo"`
	CoverUrl    string `db:"cover_url" json:"cover_url"`
	Domain      string `db:"domain" json:"domain"`
	CreateTime  int64  `db:"create_time" json:"create_time"`
	ExpireTime  int64  `db:"expire_time" json:"expire_time"` // 团队有效截止时间(second)
	Scale       int    `db:"scale" json:"-"`
	Csm         string `db:"csm" json:"-"`
	Type        int    `db:"type" json:"-"`
	SprintAlias string `db:"sprint_alias" json:"-"`
	OrgUUID     string `db:"org_uuid" json:"org_uuid"`
	Workdays    string `db:"workdays" json:"-"`
	Workhours   int64  `db:"workhours" json:"workhours"`
}

func ListAllTeam(src gorp.SqlExecutor) ([]*Team, error) {
	var results []*Team
	sql := "SELECT uuid, create_time, status, name, owner, domain, logo, cover_url, "
	sql += "expire_time, scale, sprint_alias, type, csm, org_uuid, workdays, workhours FROM team WHERE status=?;"
	_, err := src.Select(&results, sql, TeamStatusNormal)
	if err != nil {
		return nil, errors.Sql(err)
	}

	return results, nil
}
