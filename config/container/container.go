package container

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/config"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/app"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/database"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/http/controllers"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/http/middlewares"
	"github.com/go-chi/jwtauth/v5"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"net/http"
)

type Container struct {
	Middlewares
	Services
	Controllers
}

type Middlewares struct {
	AuthMw func(http.Handler) http.Handler
}

type Services struct {
	app.AuthService
	app.UserService
	app.PatientService
	app.DoctorService
	app.AppointmentService
	app.TreatmentService
}

type Controllers struct {
	controllers.AuthController
	controllers.UserController
	controllers.PatientController
	controllers.DoctorController
	controllers.AppointmentController
	controllers.TreatmentController
}

func New(conf config.Configuration) Container {
	tknAuth := jwtauth.New("HS256", []byte(conf.JwtSecret), nil)
	sess := getDbSess(conf)

	userRepository := database.NewUserRepository(sess)
	sessionRepository := database.NewSessRepository(sess)
	patientRepository := database.NewPatientRepository(sess)
	doctorRepository := database.NewDoctorRepository(sess)
	appointmentRepository := database.NewAppointmentRepository(sess)
	treatmentRepository := database.NewTreatmentRepository(sess)

	userService := app.NewUserService(userRepository)
	authService := app.NewAuthService(sessionRepository, userService, conf, tknAuth)
	patientService := app.NewPatientService(patientRepository)
	doctorService := app.NewDoctorService(doctorRepository)
	appointmentService := app.NewAppointmentService(appointmentRepository)
	treatmentService := app.NewTreatmentService(treatmentRepository)

	authController := controllers.NewAuthController(authService, userService)
	userController := controllers.NewUserController(userService)
	patientController := controllers.NewPatientController(patientService)
	doctorController := controllers.NewDoctorController(doctorService)
	appointmentController := controllers.NewAppointmentController(appointmentService)
	treatmentController := controllers.NewTreatmentController(treatmentService)

	authMiddleware := middlewares.AuthMiddleware(tknAuth, authService, userService)

	return Container{
		Middlewares: Middlewares{
			AuthMw: authMiddleware,
		},
		Services: Services{
			authService,
			userService,
			patientService,
			doctorService,
			appointmentService,
			treatmentService,
		},
		Controllers: Controllers{
			authController,
			userController,
			patientController,
			doctorController,
			appointmentController,
			treatmentController,
		},
	}
}

func getDbSess(conf config.Configuration) db.Session {
	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	//sess, err := sqlite.Open(
	//	sqlite.ConnectionURL{
	//		Database: conf.DatabasePath,
	//	})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}
	return sess
}
