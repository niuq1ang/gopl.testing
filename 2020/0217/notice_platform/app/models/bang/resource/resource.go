package resource

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

func GetResource(src gorp.SqlExecutor, uuid string) (*Resource, error) {
	sql := "SELECT * FROM resource WHERE uuid=?;"
	resource := new(Resource)
	rows, err := src.Select(resource, sql, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		resource = rows[0].(*Resource)
		return resource, nil
	}
	return nil, nil
}

func GetFileResource(src gorp.SqlExecutor, uuid string) (*FileResource, error) {
	result, err := MapFileResourcesByUUID(src, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}

	fileResource, found := result[uuid]
	if found {
		return fileResource, nil
	} else {
		return nil, nil
	}
}

func MapFileResourcesByUUID(src gorp.SqlExecutor, uuid string) (map[string]*FileResource, error) {
	resources := make(map[string]*FileResource)
	result := make([]*FileResource, 0)
	sql := "SELECT * FROM file_resource_view WHERE uuid=?;"
	_, err := src.Select(&result, sql, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}

	for _, res := range result {
		res.ExtID = res.Hash
		resources[res.UUID] = res
	}
	return resources, nil
}
