package http

import (
	"encoding/json"
	"fmt"
	"github.com/GrassBusinessLabs/eduprog-go-back/config"
	"github.com/GrassBusinessLabs/eduprog-go-back/config/container"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/http/controllers"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router(cont container.Container) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes, middleware.Logger, cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api", func(apiRouter chi.Router) {
		// Health
		apiRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())
			healthRouter.Handle("/*", NotFoundJSON())
		})

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			// Public routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Route("/auth", func(apiRouter chi.Router) {
					AuthRouter(apiRouter, cont.AuthController, cont.AuthMw)
				})
			})

			// Protected routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Use(cont.AuthMw)

				UserRouter(apiRouter, cont.UserController)
				PatientRouter(apiRouter, cont.PatientController)
				DoctorRouter(apiRouter, cont.DoctorController)
				AppointmentRouter(apiRouter, cont.AppointmentController)
				TreatmentRouter(apiRouter, cont.TreatmentController)

				apiRouter.Handle("/*", NotFoundJSON())
			})
		})
	})

	router.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, config.GetConfiguration().FileStorageLocation))
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	return router
}

func AuthRouter(r chi.Router, ac controllers.AuthController, amw func(http.Handler) http.Handler) {
	r.Route("/", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/register",
			ac.Register(),
		)
		apiRouter.Post(
			"/login",
			ac.Login(),
		)
		apiRouter.With(amw).Post(
			"/change-pwd",
			ac.ChangePassword(),
		)
		apiRouter.With(amw).Post(
			"/logout",
			ac.Logout(),
		)
	})
}

func UserRouter(r chi.Router, uc controllers.UserController) {
	r.Route("/users", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			uc.FindMe(),
		)
		apiRouter.Get(
			"/list",
			uc.ShowList(),
		)
		apiRouter.Put(
			"/",
			uc.Update(),
		)
		apiRouter.Delete(
			"/",
			uc.Delete(),
		)
	})
}

func PatientRouter(r chi.Router, pc controllers.PatientController) {
	r.Route("/patients", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			pc.SavePatient(),
		)
		apiRouter.Put(
			"/{pId}",
			pc.UpdatePatient(),
		)
		apiRouter.Get(
			"/",
			pc.ShowList(),
		)
		apiRouter.Get(
			"/{pId}",
			pc.FindById(),
		)
		apiRouter.Delete(
			"/{pId}",
			pc.DeletePatient(),
		)
	})
}

func DoctorRouter(r chi.Router, dc controllers.DoctorController) {
	r.Route("/doctors", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			dc.SaveDoctor(),
		)
		apiRouter.Put(
			"/{dId}",
			dc.UpdateDoctor(),
		)
		apiRouter.Get(
			"/",
			dc.ShowList(),
		)
		apiRouter.Get(
			"/{dId}",
			dc.FindById(),
		)
		apiRouter.Delete(
			"/{dId}",
			dc.DeleteDoctor(),
		)
	})
}

func AppointmentRouter(r chi.Router, ac controllers.AppointmentController) {
	r.Route("/appointments", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			ac.SaveAppointment(),
		)
		apiRouter.Put(
			"/{aId}",
			ac.UpdateAppointment(),
		)
		apiRouter.Get(
			"/",
			ac.ShowList(),
		)
		apiRouter.Get(
			"/{aId}",
			ac.FindById(),
		)
		apiRouter.Get(
			"/byPatient/{pId}",
			ac.FindByPatientId(),
		)
		apiRouter.Delete(
			"/{aId}",
			ac.DeleteAppointment(),
		)
	})
}

func TreatmentRouter(r chi.Router, tc controllers.TreatmentController) {
	r.Route("/treatments", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			tc.SaveTreatment(),
		)
		apiRouter.Put(
			"/{tId}",
			tc.UpdateTreatment(),
		)
		apiRouter.Get(
			"/",
			tc.ShowList(),
		)
		apiRouter.Get(
			"/{tId}",
			tc.FindById(),
		)
		apiRouter.Get(
			"/byPatient/{pId}",
			tc.FindByPatientId(),
		)
		apiRouter.Delete(
			"/{tId}",
			tc.DeleteTreatment(),
		)
	})
}

func NotFoundJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Resource Not Found")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("Ok")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}
