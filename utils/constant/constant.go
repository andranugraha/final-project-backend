package constant

const (
	// Role
	AdminRoleId = 1
	UserRoleId  = 2

	// Level
	NewbieLevelId = 1
	JuniorLevelId = 2
	SeniorLevelId = 3
	MasterLevelId = 4

	// Status
	DraftStatus   = "draft"
	PublishStatus = "publish"
)

var (
	// Role
	MapRoleId = map[int]string{
		AdminRoleId: "admin",
		UserRoleId:  "user",
	}
)
