package service

// //                     CourseRepository
//                        INTERFACE
//                            ↑
//                            │
//                            │
//              ┌─────────────┴─────────────┐
//              │                           │
//              │                           │
//  InMemoryCourseRepository       PostgresCourseRepository
//              │
//              ↓
//         []model.Course

//                     CourseService
//                          │
//                          │ has
//                          ↓
//               repository.CourseRepository
//                          │
//                          ↓
//                 InMemory Repository
import (
	// "fmt"
	"context"
	"course-api/internal/model"
	"course-api/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type CourseService struct {
	repo repository.CourseRepository
}

// As course repository is an interface which have it's behaviour methods
// A CourseService needs a repository, but I don't care which concrete repository it is.

// It could be:

// InMemoryCourseRepository
// PostgresCourseRepository
// MongoCourseRepository

// fmt.Println("")

// func NewCourseService(repo repository.CourseRepository)*CourseService{
// 	return &CourseService{
// 		repo:repo,
// 	}
// }

func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{
		repo: repo,
	}
}
func (s *CourseService) Create(
	ctx context.Context,
	course model.Course,
) (model.Course, error) {

	if course.CourseName == "" {
		return model.Course{}, ErrCourseNameRequired
	}

	_, err := s.repo.FindByName(ctx, course.CourseName)

	if err == nil {
		return model.Course{}, ErrCourseAlreadyExists
	}

	if !errors.Is(err, repository.ErrCourseNotFound) {
		return model.Course{}, err
	}

	course.CourseID = uuid.New().String()

	return s.repo.Create(ctx, course)
}

func (s *CourseService) FindAll(ctx context.Context) ([]model.Course, error) {
	return s.repo.FindAll(ctx)
}

func (s *CourseService) GetByID(ctx context.Context, id string) (model.Course, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CourseService) UpdateCourse(ctx context.Context, course model.Course) (model.Course, error) {
	if(course.CourseName=="") {
		return model.Course{},ErrCourseNameRequired
	}
	return s.repo.Update(ctx, course)
}

func (s * CourseService) DeleteCourse(ctx context.Context, id string) error{
	return s.repo.Delete(ctx,id)
}
