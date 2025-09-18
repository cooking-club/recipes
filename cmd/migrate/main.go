package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cooking-club/recipes/internal/db"
)

func init() {
	db.Init()
}

func main() {
	defer db.Close()

	fillFrom("./assets/rebuilt.json")
}

func fillFrom(file string) error {
	ctx := context.Background()

	jsonFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer jsonFile.Close()

	// Read file contents
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal JSON into struct
	var data map[string]map[string]map[string]map[string]map[string][][]*RecordInfo
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	f := "ИВТФ"
	// for f := range data {
	for y := range data[f] {
		for g := range data[f][y] {
			var srtyr uint8 = uint8(time.Now().Year() % 100)

			if strings.ContainsRune(g, 'м') {
				g = g[:2]
				srtyr -= 4
			}

			switch y {
			case "2":
				srtyr -= 1
			case "3":
				srtyr -= 2
			case "4":
				srtyr -= 3
			case "5":
				srtyr -= 4
			}

			num, err := strconv.Atoi(g)
			if err != nil {
				panic(err)
			}

			deptid := getDepartment(ctx, fmt.Sprint(num))

			for s := range data[f][y][g] {
				sg := uint8(strings.Count(s, "х") - 1)

				groupid := getGroup(ctx, srtyr, uint8(num), sg, deptid)

				for w := range data[f][y][g][s] {
					for d := range data[f][y][g][s][w] {
						for _, r := range data[f][y][g][s][w][d] {
							if r == nil {
								continue
							}

							roomid := getRoom(ctx, r.Room)
							profid := getProf(ctx, r.Professor)

							pos := 0
							if w == "1" {
								pos = 42
							}
							pos += 7 * d
							switch r.Timestamp {
							case "09:50":
								pos += 1
							case "11:40":
								pos += 2
							case "14:00":
								pos += 3
							case "15:50":
								pos += 4
							case "17:40":
								pos += 5
							case "19:25":
								pos += 6
							}

							recid := getRecord(ctx, profid, roomid, r.Subject, r.Kind, uint8(pos))
							log.Println("new record", recid)

							recordTo(ctx, recid, groupid)
						}
					}
				}
			}
		}
	}

	// }

	return nil
}

func getDepartment(ctx context.Context, code string) (deptid uint8) {
	err := db.Select((*db.Department)(nil)).Column("id").Where("short_name = ?", code).Scan(ctx, &deptid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newdept := db.Department{
				ShortName: code,
			}

			// inserts Department for every group
			res, err := db.Insert(&newdept).Exec(ctx)
			if err != nil {
				panic(err)
			}

			id, err := res.LastInsertId()
			if err != nil {
				panic(err)
			}

			return uint8(id)
		}
		panic(err)
	}
	return deptid
}

func getGroup(ctx context.Context, year, num, subg, dept uint8) uint16 {
	newgroup := db.Group{
		StartYear:    year,
		Number:       num,
		Subgroup:     subg,
		DepartmentID: dept,
	}

	res, err := db.Insert(&newgroup).Exec(ctx)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return uint16(id)
}

func getRoom(ctx context.Context, room string) (roomid uint16) {
	err := db.Select((*db.Room)(nil)).Column("id").Where("label = ?", room).Scan(ctx, &roomid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			runes := []rune(room)
			num := string(runes[1:4])
			building := string(runes[0:1])

			fullnum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			newroom := db.Room{
				Label:    room,
				Building: building,
				Floor:    uint8(fullnum / 100),
				Number:   uint8(fullnum % 100),
			}

			// inserts Department for every group
			res, err := db.Insert(&newroom).Exec(ctx)
			if err != nil {
				panic(err)
			}

			id, err := res.LastInsertId()
			if err != nil {
				panic(err)
			}

			return uint16(id)
		}
		panic(err)
	}
	return roomid
}

func getProf(ctx context.Context, fullname string) (profid uint16) {
	err := db.Select((*db.Professor)(nil)).Column("id").Where("surname = ?", fullname).Scan(ctx, &profid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newprof := db.Professor{
				Surname:      fullname,
				DepartmentID: 0,
			}

			// inserts Department for every group
			res, err := db.Insert(&newprof).Exec(ctx)
			if err != nil {
				panic(err)
			}

			id, err := res.LastInsertId()
			if err != nil {
				panic(err)
			}

			return uint16(id)
		}
		panic(err)
	}
	return profid
}

func getRecord(ctx context.Context, profid, roomid uint16, label, kind string, pos uint8) (recid uint32) {
	newrec := db.Record{
		ProfessorID: profid,
		RoomID:      roomid,
		Label:       label,
		Kind:        kind[:3],
		Position:    pos,
	}

	res, err := db.Insert(&newrec).Exec(ctx)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return uint32(id)
}

func recordTo(ctx context.Context, rid uint32, gid uint16) {
	newpair := db.GroupRecord{
		RecordId: rid,
		GroupID:  gid,
	}

	_, err := db.Insert(&newpair).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

type RecordInfo struct {
	Kind      string `json:"kind"`
	Professor string `json:"professor"`
	Room      string `json:"room"`
	Subject   string `json:"subject"`
	Timestamp string `json:"timestamp"`
}
