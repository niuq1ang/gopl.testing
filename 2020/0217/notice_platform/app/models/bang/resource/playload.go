package resource

import (
	"encoding/json"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
)

const (
	ResourceTypeFile     = 1 // 文件资源，包括图片、语音、视频和一般文件
	ResourceTypeLocation = 2 // 位置资源

	ResourceTypeFileLabel     = "file"
	ResourceTypeLocationLabel = "location"
)

func NewResourcePayload(r *FileResource) (*ResourcePayload, error) {
	m := make(map[string]*ExifValType)
	if utils.NonEmptyString(r.Exif) {
		json.Unmarshal([]byte(r.Exif), &m)
	}
	p := &ResourcePayload{
		UUID:        r.UUID,
		CreateTime:  r.CreateTime,
		OwnerUUID:   r.OwnerUUID,
		ReferenceID: r.ReferenceID,
		Status:      r.Status,
		Name:        r.Name,
		Hash:        r.Hash,
		MIME:        r.MIME,
		Size:        r.Size,
		ImageWidth:  r.ImageWidth,
		ImageHeight: r.ImageHeight,
		Exif:        m,
	}
	var err error
	if p.Type, err = r.TypeLabel(); err != nil {
		return nil, errors.Trace(err)
	}

	if p.ReferenceType, err = utils.LabelForEntityType(r.ReferenceType); err != nil {
		return nil, errors.Trace(err)
	}
	return p, nil
}

func (r *Resource) TypeLabel() (label string, err error) {
	switch r.Type {
	case ResourceTypeFile:
		label = ResourceTypeFileLabel
	case ResourceTypeLocation:
		label = ResourceTypeLocationLabel
	default:
		err = errors.InvalidEnumError(r.Type, errors.Resource, errors.Type)
	}
	return
}
