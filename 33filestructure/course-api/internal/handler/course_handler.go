package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"course-api/internal/model"
	"course-api/internal/repository"
	"course-api/internal/response"
	"course-api/internal/service"

	"github.com/gorilla/mux"
)

type CourseHandler struct {
	service *service.CourseService
}

func NewCourseHandler(
	courseService *service.CourseService,
) *CourseHandler {
	return &CourseHandler{
		service: courseService,
	}
}

func (h *CourseHandler) CreateCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	var course model.Course

	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid JSON",
		)
		return
	}

	createdCourse, err := h.service.Create(
		r.Context(),
		course,
	)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrCourseNameRequired):
			response.Error(
				w,
				http.StatusBadRequest,
				"course name is required",
			)

		case errors.Is(err, service.ErrCourseAlreadyExists):
			response.Error(
				w,
				http.StatusConflict,
				"course already exists",
			)

		default:
			response.Error(
				w,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		return
	}

	response.JSON(
		w,
		http.StatusCreated,
		createdCourse,
	)
}

func (h *CourseHandler) GetAllCourses(
    w http.ResponseWriter,
    r *http.Request,
) {
    courses, err := h.service.FindAll(r.Context())

    if err != nil {
        response.Error(
            w,
            http.StatusInternalServerError,
            "internal server error",
        )
        return
    }

    response.JSON(
        w,
        http.StatusOK,
        courses,
    )
}

func (h * CourseHandler) GetByID(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	course,err := h.service.GetByID(r.Context(),id)

	if err != nil{
		if errors.Is(err, repository.ErrCourseNotFound){
			 response.Error(
				w,
				http.StatusNotFound,
				"Course Not Found",
			)
			return
		}
		response.Error(
			w,
			http.StatusInternalServerError,
			"internal server error",
		)
	}
	response.JSON(
		w,
		http.StatusOK,
		course,
	)
}

func (h * CourseHandler) ModifyCourse(w http.ResponseWriter,r * http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var course model.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	 if err != nil {
        response.Error(
            w,
            http.StatusBadRequest,
            "invalid JSON",
        )
        return
    }

	course.CourseID = id
	updatedcourse,err := h.service.UpdateCourse(r.Context(),course)
	if err != nil{
		if errors.Is(err, service.ErrCourseNameRequired){
			response.Error(
				w,
				http.StatusBadRequest,
				"course name is required",
			)
			return
		}else if errors.Is(err,repository.ErrCourseNotFound){
			response.Error(
				w,
				http.StatusNoContent,
				"course not found",
			)
			return
		}
		response.Error(
			w,
			http.StatusInternalServerError,
			"interna server error",
		)
	}
	response.JSON(
		w,
		http.StatusOK,
		updatedcourse,
	)
}

func (h * CourseHandler) DelCourse(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteCourse(r.Context(),id)
	if err!=nil{
		if errors.Is(err,repository.ErrCourseNotFound){
			response.Error(
				w,
				http.StatusNotFound,
				"Course Not Found",
			)
			return
		}
		response.Error(
			w,
			http.StatusInternalServerError,
			"Internal Server Error",
		)
	}
	w.WriteHeader(http.StatusNoContent)
}

//                         HTTP
            //              │
            //              ↓
            //   ┌────────────────────┐
            //   │      HANDLER       │
            //   │                    │
            //   │ Decode JSON        │
            //   │                    │
            //   └─────────┬──────────┘
            //             │
            //             │ Course
            //             ↓
            //   ┌────────────────────┐
            //   │      SERVICE       │
            //   │                    │
            //   │ Validate           │
            //   │ Check duplicate    │
            //   │ Generate ID        │
            //   │                    │
            //   └─────────┬──────────┘
            //             │
            //             │ Course
            //             ↓
            //   ┌────────────────────┐
            //   │    REPOSITORY      │
            //   │    INTERFACE       │
            //   └─────────┬──────────┘
            //             │
            //             ↓
            //   ┌────────────────────┐
            //   │   MEMORY REPO      │
            //   │                    │
            //   │ []Course           │
            //   └────────────────────┘