package schedule

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Record struct {
	Timestamp string `json:"timestamp"`
	Subject   string `json:"subject"`
	Professor string `json:"professor"`
	Room      string `json:"room"`
	Kind      string `json:"kind"`
}

type Schedule = map[string]map[string]map[string]map[string]map[string][][]*Record

// expects subgroup as "х" or "хх" (1093)
func GetSchedule(faculty, year, group, subgroup, week string, is_perp bool) (error, [][]*Record) {
	path := "./assets/rebuilt.json"
	if !is_perp {
		path = "./assets/lectures-rebuilt.json"
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	defer jsonFile.Close()

	// Read file contents
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal JSON into struct
	var data Schedule
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	return nil, data[faculty][year][group][subgroup][week]
}
