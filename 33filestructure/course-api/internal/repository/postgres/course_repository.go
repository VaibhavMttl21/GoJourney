//	      main.go
//	         │
//	database.Connect()
//	         │
//	         ▼
//	     *gorm.DB
//	         │
//	         ▼
//
// NewPostgresCourseRepository()
//
//	            │
//	            ▼
//	PostgresCourseRepository
//	            │
//	            │ satisfies
//	            ▼
//	   CourseRepository
//	            │
//	            ▼
//	      CourseService
//	            │
//	            ▼
//	      CourseHandler
package postgres

import (
	"context"
	"course-api/internal/model"
	"course-api/internal/repository"
	"errors"

	// "course-api/internal/repository"

	"gorm.io/gorm"
)

type PostgresCourseRepository struct{
	db *gorm.DB
}

// contrustor
func NewPostgresCourseRepository (db *gorm.DB) *PostgresCourseRepository{
	return &PostgresCourseRepository{
		db:db,
	}
}

// var _ repository.CourseRepository = (*PostgresCourseRepository)(nil)

func (r * PostgresCourseRepository) Create(ctx context.Context, course model.Course)(model.Course,error){
	entity := toCourseEntity(course)
	err := r.db.WithContext(ctx).Create(&entity).Error

	if err != nil{
		return model.Course{},err
	}
	return course,nil

}

func (r *PostgresCourseRepository) FindAll(ctx context.Context)([]model.Course,error){
	var entities []CourseEntity
	err := r.db.WithContext(ctx).Find(&entities).Error
	
	if err != nil{
		return nil,err
	}
	courses := make([]model.Course,0,len(entities))

	for _, entity := range entities{
		courses = append(courses, toCourse(entity))
	}
	return courses,nil
}

func (r *PostgresCourseRepository) FindByID(ctx context.Context,id string,) (model.Course, error) {
	var entity CourseEntity

	err := r.db.WithContext(ctx).Where("course_id = ?", id).First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Course{}, repository.ErrCourseNotFound
		}

		return model.Course{}, err
	}

	return toCourse(entity), nil
}

func (r *PostgresCourseRepository) FindByName(ctx context.Context,name string,) (model.Course, error) {

	var entity CourseEntity

	err := r.db.WithContext(ctx).Where("course_name = ?", name).First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Course{}, repository.ErrCourseNotFound
		}

		return model.Course{}, err
	}

	return toCourse(entity), nil
}

func (r *PostgresCourseRepository) Update(ctx context.Context,course model.Course,) (model.Course, error) {

	entity := toCourseEntity(course)

	result := r.db.WithContext(ctx).Model(&CourseEntity{}).Where("course_id = ?", course.CourseID).Updates(&entity)

	if result.Error != nil {
		return model.Course{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Course{}, repository.ErrCourseNotFound
	}

	return course, nil
}

func (r *PostgresCourseRepository) Delete(ctx context.Context,id string,) error {

	result := r.db.WithContext(ctx).Where("course_id = ?", id).Delete(&CourseEntity{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return repository.ErrCourseNotFound
	}

	return nil
}