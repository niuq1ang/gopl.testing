package message

// message
const (
	MessageTypeDiscussion = 1
	MessageTypeSystem     = 2

	MessageTypeDiscussionLabel = "discussion"
	MessageTypeSystemLabel     = "system"

	taskManhoursRefMessage = "uuid, team_uuid, send_time, action, object_type, object_id, object_attr, old_value, new_value, ext"
)

type Message struct {
	UUID          string `db:"uuid" json:"uuid"`
	TeamUUID      string `db:"team_uuid" json:"team_uuid"`
	ReferenceType int    `db:"reference_type" json:"reference_type"`
	ReferenceID   string `db:"reference_id" json:"reference_id"`
	FromUUID      string `db:"from_uuid" json:"from_uuid"`
	ToUUID        string `db:"to_uuid" json:"to_uuid"`
	SendTime      int64  `db:"send_time" json:"send_time"`
	Type          int    `db:"type" json:"type"`
	Text          string `db:"message" json:"text"`
	ResourceUUID  string `db:"resource_uuid" json:"resource_uuid"`
	SubjectType   int    `db:"subject_type" json:"subject_type"`
	SubjectID     string `db:"subject_id" json:"subject_id"`
	Action        int    `db:"action" json:"action"`
	ObjectType    int    `db:"object_type" json:"object_type"`
	ObjectID      string `db:"object_id" json:"object_id"`
	ObjectName    string `db:"object_name" json:"object_name"`
	ObjectAttr    int    `db:"object_attr" json:"object_attr"`
	OldValue      string `db:"old_value" json:"old_value"`
	NewValue      string `db:"new_value" json:"new_value"`
	Ext           string `db:"ext" json:"ext"`
}
