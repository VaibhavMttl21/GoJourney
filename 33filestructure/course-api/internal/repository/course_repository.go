// Interface
//     ↓
// defines WHAT

// Struct implementation
//     ↓
// defines HOW

package repository

import (
	"context"
	"course-api/internal/model"
)

// Any Course Repository must know how to perform these five operations.
type CourseRepository interface{
	Create(ctx context.Context,course model.Course)(model.Course,error)
	FindByID(ctx context.Context, id string)(model.Course,error)
	FindByName(ctx context.Context, name string) (model.Course,error)
	Update(ctx context.Context, course model.Course) (model.Course, error)
	Delete(ctx context.Context, id string)(error)
	FindAll(ctx context.Context) ([]model.Course, error)
}