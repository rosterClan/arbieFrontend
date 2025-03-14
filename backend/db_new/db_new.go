package db_new

import (
	"arbie/new_models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "arbie_4.2"
)

var GLOBAL_DB_OBJ *sql.DB

func init() {
	open_connection()
}

func open_connection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	GLOBAL_DB_OBJ = db
}

func exec_query(query string) *sql.Rows {
	rows, err := GLOBAL_DB_OBJ.Query(query)
	if err != nil {
		return nil
	}
	return rows
}

func Get_Platform_Price_Fluctuations_By_EntrantID(entrant_id string, platform_name string) ([]new_models.Price, error) {

	//Updated
	rows := exec_query("SELECT odds, record_time FROM datrum WHERE horse='" + entrant_id + "' AND platform_name='" + platform_name + "' ORDER BY record_time DESC; ")

	price_fluctuations := make([]new_models.Price, 0)
	for rows.Next() {
		var curr_price_instance new_models.Price
		if err := rows.Scan(&curr_price_instance.Odds, &curr_price_instance.Record_Time); err != nil {
			rows.Close()
			return price_fluctuations, errors.New("scan error")
		}
		price_fluctuations = append(price_fluctuations, curr_price_instance)
	}
	rows.Close()

	return price_fluctuations, nil
}

func Get_Price_By_EntrantID(entrant_id string) ([]new_models.Prices, error) {

	//Updated
	rows := exec_query("SELECT DISTINCT platform_name FROM datrum WHERE horse='" + entrant_id + "';")

	all_prices := make([]new_models.Prices, 0)
	for rows.Next() {
		var platform_name string
		if err := rows.Scan(&platform_name); err != nil {
			rows.Close()
			return all_prices, errors.New("scan error")
		}

		platform_prices, err := Get_Platform_Price_Fluctuations_By_EntrantID(entrant_id, platform_name)
		if err != nil {
			rows.Close()
			return all_prices, errors.New("scan error")
		}

		var platform_price new_models.Prices
		if len(platform_prices) > 0 {
			platform_price.Current_Price = platform_prices[0]
		}
		platform_price.Price_Fluctuations = platform_prices
		platform_price.Platform_Name = platform_name

		all_prices = append(all_prices, platform_price)
	}
	rows.Close()

	return all_prices, nil
}

func Get_Entrant_By_ID(entrant_id string) (new_models.Entrant, error) {

	//Updated
	rows := exec_query("SELECT horse, horse, is_scratched FROM datrum WHERE horse='" + entrant_id + "';")

	var entrant new_models.Entrant
	for rows.Next() {
		if err := rows.Scan(&entrant.Entrant_Name, &entrant.Is_Scratched); err != nil {
			return entrant, errors.New("scan error")
		}
		break
	}

	return entrant, nil
}

func Get_EntrantIDs_By_RaceID(race_id string) ([]string, error) {

	//Updated
	rows := exec_query("SELECT horse FROM datrum WHERE race_id='" + race_id + "';")

	entrant_ids := make([]string, 0)
	for rows.Next() {
		var entrant_id string
		if err := rows.Scan(&entrant_id); err != nil {
			rows.Close()
			return entrant_ids, errors.New("scan error")
		}
		entrant_ids = append(entrant_ids, entrant_id)
	}

	return entrant_ids, nil
}

func Get_Next_2_Go_Race() ([]new_models.Race, error) {
	rows := exec_query("SELECT track, round, to_timestamp(avg(extract(epoch from start_time))) as time_stamps from race GROUP BY track, round ORDER BY time_stamps ASC;")
	races := make([]new_models.Race, 0)
	for rows.Next() {
		var curr_race new_models.Race
		if err := rows.Scan(&curr_race.Track_Name, &curr_race.Round, &curr_race.Start_Time); err != nil {
			rows.Close()
			return races, errors.New("scan error")
		}
		races = append(races, curr_race)
	}
	rows.Close()
	return races, nil
}

func Get_Related_RaceIds(race_id string) ([]string, error) {

	//Updated
	rows := exec_query("SELECT race.race_id FROM race JOIN track ON race.track = track.track_name WHERE DATE(race.start_time) in (SELECT DATE(start_time) FROM race WHERE race_id = '" + race_id + "') AND race.track in (SELECT race.track FROM race WHERE race_id = '" + race_id + "') ORDER BY race.round;")

	race_ids := make([]string, 0)
	for rows.Next() {
		var cur_race_id string
		if err := rows.Scan(&cur_race_id); err != nil {
			return race_ids, errors.New("scan error")
		}
		race_ids = append(race_ids, cur_race_id)
	}
	rows.Close()

	return race_ids, nil
}

func Get_Races_By_Timestamp_And_Track(time_stamp string, track_name string) ([]new_models.Race, error) {
	rows := exec_query("SELECT track, round, to_timestamp(avg(extract(epoch from start_time))) as time_stamps from race WHERE track='" + track_name + "' AND DATE(start_time)=DATE(to_timestamp(" + time_stamp + "/1000)) GROUP BY track, round ORDER BY round ASC;")

	races := make([]new_models.Race, 0)
	for rows.Next() {
		var curr_race new_models.Race
		if err := rows.Scan(&curr_race.Track_Name, &curr_race.Round, &curr_race.Start_Time); err != nil {
			rows.Close()
			return races, errors.New("scan error")
		}
		races = append(races, curr_race)
	}
	rows.Close()

	return races, nil
}

func Get_Races_On_Timestamp(time_stamp string) ([]new_models.Meet, error) {
	rows := exec_query("select DISTINCT track from race WHERE DATE(start_time)=DATE(to_timestamp(" + time_stamp + " / 1000));")

	meets := make([]new_models.Meet, 0)
	for rows.Next() {
		var track_name string
		if err := rows.Scan(&track_name); err != nil {
			rows.Close()
			return meets, errors.New("scan error")
		}

		races, err := Get_Races_By_Timestamp_And_Track(time_stamp, track_name)
		if err != nil {
			rows.Close()
			return meets, errors.New("scan error")
		}

		var meet new_models.Meet

		meet.Track_Name = track_name
		meet.Races = races

		meets = append(meets, meet)
	}
	rows.Close()

	return meets, nil
}

func Get_Race_Ids(track_name string, round string, start_time string) ([]string, error) {
	descriptionTemplate := `SELECT race_id FROM race WHERE track='%s' AND round=%s AND DATE(to_timestamp('%s', 'YYYY-MM-DD"T"HH24:MI:SS.US TZH:TZM') AT TIME ZONE 'UTC')=DATE(start_time AT TIME ZONE 'UTC');`
	newString := fmt.Sprintf(descriptionTemplate, track_name, round, start_time)

	rows := exec_query(newString)

	race_ids := make([]string, 0)
	for rows.Next() {
		var race_id string
		if err := rows.Scan(&race_id); err != nil {
			rows.Close()
			return race_ids, errors.New("scan error")
		}
		race_ids = append(race_ids, race_id)
	}
	rows.Close()
	return race_ids, nil
}

func Get_Entrants_By_Raceids(race_ids []string) ([]new_models.Entrant, error) {
	query_sub_str := "("
	for _, race_id := range race_ids {
		query_sub_str += ("'" + race_id + "'" + ",")
	}
	query_sub_str += "'0')"

	rows := exec_query("SELECT horse, SUM(is_scratched) > 0 FROM datrum WHERE race_id IN " + query_sub_str + " GROUP BY horse;")

	entrants := make([]new_models.Entrant, 0)
	for rows.Next() {
		var entrant new_models.Entrant
		if err := rows.Scan(&entrant.Entrant_Name, &entrant.Is_Scratched); err != nil {
			rows.Close()
			return entrants, errors.New("scan error")
		}
		entrants = append(entrants, entrant)
	}
	rows.Close()

	return entrants, nil
}

func Get_Individual_Entrant_Current_Odds(entrant string, race_id string) ([]new_models.Price, error) {
	rows := exec_query("SELECT odds, record_time FROM datrum WHERE race_id='" + race_id + "' AND horse='" + entrant + "' ORDER BY record_time DESC;")

	prices := make([]new_models.Price, 0)

	for rows.Next() {
		var currnet_price new_models.Price
		if err := rows.Scan(&currnet_price.Odds, &currnet_price.Record_Time); err != nil {
			rows.Close()
			return prices, errors.New("scan error")
		}
		prices = append(prices, currnet_price)
	}
	rows.Close()

	return prices, nil
}

func Get_Platform_Name_By_ID(race_id string) (string, error) {
	rows := exec_query("SELECT DISTINCT platform_name FROM datrum WHERE race_id='" + race_id + "';")

	platform_name := ""
	for rows.Next() {
		if err := rows.Scan(&platform_name); err != nil {
			rows.Close()
			return platform_name, errors.New("scan error")
		}
		break
	}
	rows.Close()

	return platform_name, nil
}

func Fill_Entrant_Platform_Odds(raceIDs []string, entrants []new_models.Entrant) ([]new_models.Entrant, error) {
	for i, entrant := range entrants {
		var platformPrices []new_models.Prices
		for _, raceID := range raceIDs {
			var currentPrices new_models.Prices
			individualPrices, err := Get_Individual_Entrant_Current_Odds(entrant.Entrant_Name, raceID)
			if err != nil {
				return nil, fmt.Errorf("error fetching odds for entrant %s in race %s: %w", entrant.Entrant_Name, raceID, err)
			}
			currentPrices.Platform_Name, err = Get_Platform_Name_By_ID(raceID)
			if err != nil {
				continue
			}
			if len(individualPrices) > 0 {
				currentPrices.Current_Price = individualPrices[0]
			}
			currentPrices.Price_Fluctuations = individualPrices
			platformPrices = append(platformPrices, currentPrices)
		}
		entrants[i].Odds = platformPrices
	}

	return entrants, nil
}
