package app

import (
	"database/sql"
	"log/slog"

	"github.com/erwindrsno/PMVBD-CheckIn/internal/attendance"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/attendee"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/database"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/event"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/grade"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/school"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/subgrade"
	"github.com/erwindrsno/PMVBD-CheckIn/internal/user"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func New() *App {
	dbPath := "pmvbd-checkin-dev.db"
	db, err := database.InitDB(dbPath)
	if err != nil {
		slog.Error(err.Error())
	}

	router := gin.Default()

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173", "http://127.0.0.1:5173"}
	// router.Use(cors.New(config))
	router.Use(cors.Default())

	a := &App{
		DB:     db,
		Router: router,
	}

	a.setRoutes()
	return a
}

func (a *App) setRoutes() {
	a.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//school definition
	schoolStore := school.NewSQLiteStore(a.DB)
	schoolService := school.NewService(schoolStore)
	schoolHandler := school.NewHandler(schoolService)

	//event definition
	eventStore := event.NewSQLiteStore(a.DB)
	eventService := event.NewService(eventStore)
	eventHandler := event.NewHandler(eventService)

	//attendee definition
	attendeeStore := attendee.NewSQLiteStore(a.DB)
	attendeeService := attendee.NewService(attendeeStore)
	attendeeHandler := attendee.NewHandler(attendeeService)

	//grade definition
	gradeStore := grade.NewSQLiteStore(a.DB)
	gradeService := grade.NewService(gradeStore)
	gradeHandler := grade.NewHandler(gradeService)

	//subgrade definition
	subgradeStore := subgrade.NewSQLiteStore(a.DB)
	subgradeService := subgrade.NewService(subgradeStore)
	subgradeHandler := subgrade.NewHandler(subgradeService)

	//attendance definition
	attendanceStore := attendance.NewSQLiteStore(a.DB)
	attendanceService := attendance.NewService(attendanceStore, attendeeService)
	attendanceHandler := attendance.NewHandler(attendanceService)

	//user definition
	userStore := user.NewSQLiteStore(a.DB)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService)

	api := a.Router.Group("/api/v1")
	{
		schools := api.Group("/schools")
		{
			schools.GET("", schoolHandler.Read)
			schools.POST("", schoolHandler.Create)
		}

		events := api.Group("/events")
		{
			events.GET("", eventHandler.Read)
			events.POST("", eventHandler.Create)
			events.PATCH(":id/status", eventHandler.UpdateStatus)
			events.DELETE(":id", eventHandler.Delete)
		}

		attendees := api.Group("/attendees")
		{
			attendees.GET("", attendeeHandler.Read)
			attendees.POST("", attendeeHandler.Create)
			attendees.DELETE(":public_id", attendeeHandler.Delete)
		}

		attendances := api.Group("/attendances")
		{
			attendances.GET(":event_id", attendanceHandler.ReadByEventId)
			attendances.POST("", attendanceHandler.Capture)
			attendances.DELETE(":id", attendanceHandler.DeleteById)
		}

		grades := api.Group("/grades")
		{
			grades.GET("", gradeHandler.Read)
			grades.POST("", gradeHandler.Create)
		}

		subgrades := api.Group("/subgrades")
		{
			subgrades.GET("", subgradeHandler.Read)
			subgrades.POST("", subgradeHandler.Create)
		}

		users := api.Group("/users")
		{
			users.POST("", userHandler.Create)
		}
	}
}

func (a *App) Run(port string) {
	a.Router.Run(port)
}

func (a *App) CloseDB() {
	a.DB.Close()
}
