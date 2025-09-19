package db

import "github.com/uptrace/bun"

type Department struct {
	bun.BaseModel `bun:"table:departments"`

	ID        uint8   `bun:",pk,autoincrement" json:"id"`
	ShortName string  `bun:"type:varchar(15)" json:"shortName"`
	FullName  string  `json:"fullName"`
	Faculty   *string `bun:"type:enum('ИВТФ','ИФФ','ЭМФ', 'ФЭУ')" json:"faculty"` // todo: replace string with go native enum
}

type Room struct {
	bun.BaseModel `bun:"table:rooms"`

	ID       uint16 `bun:",pk,autoincrement" json:"id"`
	Building string `bun:"type:enum('А','Б','В','Г','Д','С'),notnull" json:"building"`
	Floor    uint8  `bun:",notnull" json:"floor"`
	Number   uint8  `bun:",notnull" json:"number"`
	Label    string `bun:"type:varchar(31)" json:"label"`
}

type Professor struct {
	bun.BaseModel `bun:"table:professors"`

	ID           uint16     `bun:",pk,autoincrement" json:"id"`
	FirstName    string     `bun:"type:varchar(63)" json:"firstName"`
	Patronymic   string     `bun:"type:varchar(63)" json:"patronymic"`
	Surname      string     `bun:"type:varchar(63)" json:"surname"`
	DepartmentID uint8      `bun:",notnull" json:"departmentId"`
	Department   Department `bun:"rel:belongs-to,join:department_id=id,notnull" json:"department"`
}

type Group struct {
	bun.BaseModel `bun:"table:groups"`

	ID           uint16 `bun:",pk,autoincrement" json:"id"`
	Label        string `bun:"type:varchar(15)" json:"label"`
	StartYear    uint8  `bun:",notnull" json:"startYear"`
	Number       uint8  `bun:",notnull" json:"number"`
	Subgroup     uint8  `bun:",notnull" json:"subgroup"`
	DepartmentID uint8  `bun:",notnull" json:"departmentId"`

	Department Department `bun:"rel:belongs-to,join:department_id=id,notnull" json:"department"`
	Records    []Record   `bun:"m2m:group_records,join:Group=Record" json:"records"`
}

type Record struct {
	bun.BaseModel `bun:"table:records"`

	ID          uint32 `bun:"type:mediumint unsigned,pk,autoincrement" json:"id"`
	Label       string `bun:",notnull" json:"label"`
	Kind        string `bun:"type:enum('lab','sem','lec'),notnull" json:"kind"`
	Position    uint8  `bun:",notnull" json:"position"`
	ProfessorID uint16 `bun:",notnull" json:"professorId"`
	RoomID      uint16 `bun:",notnull" json:"roomId"`

	Professor Professor `bun:"rel:belongs-to,join:professor_id=id,notnull" json:"professor"`
	Room      Room      `bun:"rel:belongs-to,join:room_id=id,notnull" json:"room"`
	Groups    []Group   `bun:"m2m:group_records,join:Record=Group" json:"groups"`
}

type GroupRecord struct {
	bun.BaseModel `bun:"table:group_records"`

	GroupID  uint16 `bun:",pk"`
	RecordId uint32 `bun:",pk"`
	Group    Group  `bun:"rel:belongs-to,join:group_id=id"`
	Record   Record `bun:"rel:belongs-to,join:record_id=id"`
}
