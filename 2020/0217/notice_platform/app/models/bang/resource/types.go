package resource

type ExifValType struct {
	Val  string `json:"val"`
	Type int    `json:"type"`
}

type ResourcePayload struct {
	UUID          string                  `json:"uuid"`
	Type          string                  `json:"type"`
	ReferenceType string                  `json:"ref_type"`
	ReferenceID   string                  `json:"ref_id"`
	CreateTime    int64                   `json:"create_time"`
	OwnerUUID     string                  `json:"owner_uuid"`
	Status        int                     `json:"status"`
	Name          string                  `json:"name,omitempty"`
	Hash          string                  `json:"hash,omitempty"`
	MIME          string                  `json:"mime,omitempty"`
	Size          int64                   `json:"size,omitempty"`
	ImageWidth    int                     `json:"image_width,omitempty"`
	ImageHeight   int                     `json:"image_height,omitempty"`
	Exif          map[string]*ExifValType `json:"image_exif,omitempty"`
}

// 表示属于某个team的，由team中某个成员生成的，需要控制访问的资源
type Resource struct {
	UUID          string `db:"uuid"`
	ReferenceType int    `db:"reference_type"`
	ReferenceID   string `db:"reference_id"`
	TeamUUID      string `db:"team_uuid"`
	ProjectUUID   string `db:"project_uuid"`
	OwnerUUID     string `db:"owner_uuid"`
	Type          int    `db:"type"`
	Source        int    `db:"source"`
	ExtID         string `db:"ext_id"`
	Name          string `db:"name"`
	Status        int    `db:"status"`
	CreateTime    int64  `db:"create_time" json:"create_time"`
	Description   string `db:"description" json:"description"`
	ModifyTime    int64  `db:"modify_time" json:"-"`
}

type FileResource struct {
	Resource
	Hash        string `db:"hash"`
	MIME        string `db:"mime"`
	Size        int64  `db:"size"`
	ImageWidth  int    `db:"image_width"`
	ImageHeight int    `db:"image_height"`
	Exif        string `db:"exif"`
}
