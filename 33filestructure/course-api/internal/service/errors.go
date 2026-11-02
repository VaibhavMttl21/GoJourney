package service

import "errors"

var (
	ErrCourseNameRequired = errors.New("course name is required")
	ErrCourseAlreadyExists = errors.New("course already exists")
)