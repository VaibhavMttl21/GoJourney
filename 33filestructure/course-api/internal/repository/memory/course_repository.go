// repository/course_repository.go
//         │
//         └── interface
//             CourseRepository
//                   ↑
//                   │ satisfies
//                   │
// repository/memory/course_repository.go
//         │
//         └── struct
//             InMemoryCourseRepository

// Interface
//     ↓
// defines WHAT

// Struct implementation
//     ↓
// defines HOW
package memory

import (
	"context"

	"course-api/internal/model"
	"course-api/internal/repository"
)

type InMemoryCourseRepository struct {
	courses []model.Course
}

// Compile-time check:
// Make sure InMemoryCourseRepository satisfies repository.CourseRepository.
var _ repository.CourseRepository = (*InMemoryCourseRepository)(nil)
// construcor or we can say initalizer
// func NewInMemoryCourseRepository() *InMemoryCourseRepository {
// 	return &InMemoryCourseRepository{
// 		courses: make([]model.Course, 0),
// 	}
// }

func NewInMemoryCourseRepository() *InMemoryCourseRepository{
	return &InMemoryCourseRepository{
		courses: make([]model.Course,0),
	}
}


func (r *InMemoryCourseRepository) Create(
	ctx context.Context,
	course model.Course,
) (model.Course, error) {

	select {
	case <-ctx.Done():
		return model.Course{}, ctx.Err()
	default:
	}

	r.courses = append(r.courses, course)

	return course, nil
}

func (r *InMemoryCourseRepository) FindByID(
	ctx context.Context,
	id string,
) (model.Course, error) {

	for _, course := range r.courses {
		if course.CourseID == id {
			return course, nil
		}
	}

	return model.Course{}, repository.ErrCourseNotFound
}

func (r *InMemoryCourseRepository) FindByName(
	ctx context.Context,
	name string,
) (model.Course, error) {

	for _, course := range r.courses {
		if course.CourseName == name {
			return course, nil
		}
	}

	return model.Course{}, repository.ErrCourseNotFound
}

func (r *InMemoryCourseRepository) Update(
	ctx context.Context,
	course model.Course,
) (model.Course, error) {

	for i := range r.courses {
		if r.courses[i].CourseID == course.CourseID {
			r.courses[i] = course
			return course, nil
		}
	}

	return model.Course{}, repository.ErrCourseNotFound
}

func (r *InMemoryCourseRepository) Delete(
	ctx context.Context,
	id string,
) error {

	for i := range r.courses {
		if r.courses[i].CourseID == id {

			r.courses = append(
				r.courses[:i],
				r.courses[i+1:]...,
			)

			return nil
		}
	}

	return repository.ErrCourseNotFound
}

func (r * InMemoryCourseRepository) FindAll(ctx context.Context)([]model.Course,error){
	return r.courses,nil
}