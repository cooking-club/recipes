package db

import "github.com/uptrace/bun"

type Department struct {
	bun.BaseModel `bun:"table:departments"`

	ID        uint8  `bun:",pk,autoincrement"`
	ShortName string `bun:"type:varchar(15)"`
	FullName  string
	Faculty   *string `bun:"type:enum('ИВТФ','ИФФ','ЭМФ')"` // todo: replace string with go native enum
}

type Room struct {
	bun.BaseModel `bun:"table:rooms"`

	ID       uint16 `bun:",pk,autoincrement"`
	Building string `bun:"type:enum('А','Б','В','Г','Д','С'),notnull"`
	Floor    uint8  `bun:",notnull"`
	Number   uint8  `bun:",notnull"`
	Label    string `bun:"type:varchar(31)"`
}

type Professor struct {
	bun.BaseModel `bun:"table:professors"`

	ID           uint16     `bun:",pk,autoincrement"`
	FirstName    string     `bun:"type:varchar(63)"`
	Patronymic   string     `bun:"type:varchar(63)"`
	Surname      string     `bun:"type:varchar(63)"`
	DepartmentID uint8      `bun:",notnull"`
	Department   Department `bun:"rel:belongs-to,join:department_id=id,notnull"`
}

type Group struct {
	bun.BaseModel `bun:"table:groups"`

	ID           uint16 `bun:",pk,autoincrement"`
	Label        string `bun:"type:varchar(15)"`
	StartYear    uint8  `bun:",notnull"`
	Number       uint8  `bun:",notnull"`
	Subgroup     uint8  `bun:",notnull"`
	DepartmentID uint8  `bun:",notnull"`

	Department Department `bun:"rel:belongs-to,join:department_id=id,notnull"`
	Records    []Record   `bun:"m2m:group_records,join:Group=Record"`
}

type Record struct {
	bun.BaseModel `bun:"table:records"`

	ID          uint32 `bun:"type:mediumint unsigned,pk,autoincrement"`
	Label       string `bun:",notnull"`
	Kind        string `bun:"type:enum('lab','sem','lec'),notnull"`
	Position    uint8  `bun:",notnull"`
	ProfessorID uint16 `bun:",notnull"`
	RoomID      uint16 `bun:",notnull"`

	Professor Professor `bun:"rel:belongs-to,join:professor_id=id,notnull"`
	Room      Room      `bun:"rel:belongs-to,join:room_id=id,notnull"`
	Groups    []Group   `bun:"m2m:group_records,join:Record=Group"`
}

type GroupRecord struct {
	bun.BaseModel `bun:"table:group_records"`

	GroupID  uint16 `bun:",pk"`
	RecordId uint32 `bun:",pk"`
	Group    Group  `bun:"rel:belongs-to,join:group_id=id"`
	Record   Record `bun:"rel:belongs-to,join:record_id=id"`
}
