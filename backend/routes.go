package main

import (
	"arbie/internal_new"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"}, // Allow your frontend origin
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	/*
		router.GET("/next_2_go", internal.Get_Next_2_go)
		router.GET("/get_day_races/:date", internal.Get_Day_Races_Internal)
		router.GET("/get_race_details/:race_id", internal.Get_Race_View)
		router.GET("/get_related_races/:race_id", internal.Get_Other_Meet_Races)
		router.GET("/get_race_entrants/:race_id", internal.Get_Race_Entrants)
		router.GET("/get_entrant_price_history/:entrant_id", internal.Get_Entrant_Price_History)
	*/

	router.GET("/get_day_races/:timestamp", internal_new.Get_Selected_Day_Races)
	router.GET("/next_2_go", internal_new.Get_Next_2_go)
	router.GET("/get_race_details", internal_new.Get_Race_View)
	router.GET("/get_entrant_price_history/:entrant_id", internal_new.Get_Entrant_Price_History)
	router.GET("/get_race_entrants/:race_id", internal_new.Get_Race_Entrants)

	router.Run("localhost:8080")
}
