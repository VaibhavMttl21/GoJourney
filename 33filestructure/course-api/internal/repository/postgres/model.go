package postgres

import "course-api/internal/model"

type CourseEntity struct {
    CourseID      string `gorm:"column:course_id;type:uuid;primaryKey"`
    CourseName    string `gorm:"column:course_name;not null"`
    CoursePrice   int    `gorm:"column:course_price;not null"`
    AuthorName    string `gorm:"column:author_name"`
    AuthorWebsite string `gorm:"column:author_website"`
}
// Whenever GORM works with CourseEntity, use the courses table.
func (CourseEntity) TableName() string {
    return "courses"
}
// this is because our api have model have two tables but postgres has one so we have to convert 
//This transformation belongs in the repository because it's specifically about persistence/database representation.
func toCourseEntity(course model.Course) CourseEntity {
    courseentity := CourseEntity{
        CourseID:    course.CourseID,
        CourseName:  course.CourseName,
        CoursePrice: course.CoursePrice,
    }

    if course.Author != nil {
        courseentity.AuthorName = course.Author.FullName
        courseentity.AuthorWebsite = course.Author.Website
    }

    return courseentity
}
// as the db will return as a courseentity we have to convert it inot our mode course so that 
// api doesnot get any issue ike for findall , find by id and all 
func toCourse(entity CourseEntity) model.Course {
	var author *model.Author

	if entity.AuthorName != "" || entity.AuthorWebsite != "" {
		author = &model.Author{
			FullName: entity.AuthorName,
			Website:  entity.AuthorWebsite,
		}
	}

	return model.Course{
		CourseID:    entity.CourseID,
		CourseName:  entity.CourseName,
		CoursePrice: entity.CoursePrice,
		Author:      author,
	}
}