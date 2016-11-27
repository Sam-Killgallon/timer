package timer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func Save_start_time(the_time time.Time) {
	data := loadData()

	t_day, s_hour := formatTime(the_time)
	e_hour := data[t_day]["end"]
	formatSave(t_day, s_hour, e_hour, data)
}

func Save_end_time(the_time time.Time) {
	data := loadData()

	t_day, e_hour := formatTime(the_time)
	s_hour := data[t_day]["start"]
	formatSave(t_day, s_hour, e_hour, data)
}

func Overtime() {
	const hour = "15:04"

	data := loadData()
	var total float64 = 0
	// one working day
	var work_day float64 = 60 * 8.5

	for _, val := range data {
		if val["start"] == "" || val["end"] == "" {
			continue
		}
		start, err := time.Parse(hour, val["start"])
		check(err)
		end, err := time.Parse(hour, val["end"])
		check(err)

		total = total + end.Sub(start).Minutes()
		total = total - work_day
	}

	formated := strconv.FormatFloat(total, 'f', 0, 64) + "m"
	parseable, err := time.ParseDuration(formated)
	check(err)
	fmt.Println(parseable)
}

func formatSave(day, start, end string, data map[string]map[string]string) {
	my_map := buildMap(start, end)

	data[day] = my_map
	date_json, err := json.MarshalIndent(data, "", "  ")
	check(err)

	saveData(date_json)
}

func saveData(json []byte) {
	check(ioutil.WriteFile("dates.json", json, 0644))
}

func loadData() map[string]map[string]string {
	data := readFile()

	var obj map[string]map[string]string
	err := json.Unmarshal(data, &obj)

	if err != nil {
		fmt.Println(err)
	}

	return obj
}

func readFile() []byte {
	data, err := ioutil.ReadFile("dates.json")
	check(err)
	return data
}

func formatTime(the_time time.Time) (string, string) {
	const day = "02-01-2006"
	const hour = "15:04"

	t_day := the_time.Format(day)
	t_hour := the_time.Format(hour)

	return t_day, t_hour
}

func buildMap(s_hour, e_hour string) map[string]string {
	return map[string]string{"start": s_hour, "end": e_hour}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
