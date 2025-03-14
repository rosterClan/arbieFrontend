package internal_new

import (
	"arbie/db_new"
	"arbie/new_models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Throw_Error(msg string, c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"err": msg})
}

func Get_Selected_Day_Races(c *gin.Context) {
	time := c.Param("timestamp")

	meets, err := db_new.Get_Races_On_Timestamp(time)
	if err != nil {
		Throw_Error("Couldn't get race_ids", c)
	}

	c.IndentedJSON(http.StatusOK, meets)

}

func Get_Next_2_go(c *gin.Context) {
	next_2_go_races, err := db_new.Get_Next_2_Go_Race()
	if err != nil {
		Throw_Error("Couldn't get race_ids", c)
	}

	c.IndentedJSON(http.StatusOK, next_2_go_races)
}

func Get_Race_Entrants(c *gin.Context) {
	entrant_ids, err := db_new.Get_EntrantIDs_By_RaceID(c.Param("race_id"))
	if err != nil {
		Throw_Error("Bad race_id", c)
	}

	entrants := make([]new_models.Entrant, 0)
	for _, entrant_id := range entrant_ids {
		entrant, err := db_new.Get_Entrant_By_ID(entrant_id)
		if err != nil {
			Throw_Error("Bad entrant_id", c)
		}
		entrants = append(entrants, entrant)
	}

	c.IndentedJSON(http.StatusOK, entrants)
}

func Get_Entrant_Price_History(c *gin.Context) {
	entrant, err := db_new.Get_Entrant_By_ID(c.Param("entrant_id"))
	if err != nil {
		Throw_Error("Couldn't identify entrant", c)
	}

	entrant_prices, err := db_new.Get_Price_By_EntrantID(c.Param("entrant_id"))
	if err != nil {
		Throw_Error("Couldn't identidy entrant prices", c)
	}

	entrant.Odds = entrant_prices
	c.IndentedJSON(http.StatusOK, entrant)
}

func Get_Race_View(c *gin.Context) {
	track := c.Query("track")
	round := c.Query("round")
	start_time := c.Query("start_time")

	race_ids, err := db_new.Get_Race_Ids(track, round, start_time)
	if err != nil {
		Throw_Error("Couldn't identidy entrant prices", c)
	}

	entrants, err := db_new.Get_Entrants_By_Raceids(race_ids)
	entrants, err = db_new.Fill_Entrant_Platform_Odds(race_ids, entrants)

	layout := "2006-01-02T15:04:05 -07:00"
	t, err := time.Parse(layout, start_time)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
	}

	var race new_models.Race

	race.Track_Name = track
	race.Round = race.Round
	race.Start_Time = t
	race.Entrants = entrants

	fmt.Println(track, round, start_time)
	c.IndentedJSON(http.StatusOK, race)
}
