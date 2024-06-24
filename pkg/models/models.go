package models

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/database"
	"gorm.io/gorm"
)

// School represents the School table
type School struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Departments []Department `gorm:"foreignKey:SchoolID" json:"departments,omitempty"`
}

func (s *School) Save() {
	DB.Create(&s)
}


// Department represents the Department table
type Department struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SchoolID    uint    `json:"school_id"`
	Courses     []Course `gorm:"foreignKey:DepartmentID" json:"courses,omitempty"`
}

func (d *Department) Save() {
	DB.Create(&d)
}

// Course represents the Course table
type Course struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
	Books        []Books `gorm:"foreignKey:CourseID" json:"books,omitempty"`
}

func (c *Course) Save() {
	DB.Create(&c)
}
// Book represents the Book table
type Books struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	File         string `json:"file"`
	Image        string `json:"image"`
	Author			 string `json:"author"`
	CourseID     uint   `json:"course_id"`
}

func (b *Books) Save() {
	DB.Create(&b)
}


var DB *gorm.DB

func init() {
	database.Connect()
	DB = database.DB
	err := DB.AutoMigrate(&School{}, &Department{}, &Course{}, &Books{})
	if err != nil {
		log.Fatal("fail to connect with database")
	}
}

// start school
func CreateSchools(c *fiber.Ctx) error {
	school := new(School)
	err := c.BodyParser(school)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	school.Save()
	return c.JSON(school)
}

func GetSchool(c *fiber.Ctx) error {
	var school School;
	id := c.Params("school_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).SendString("school id is not valid")
	}
	if err := DB.First(&school, Id).Error; err != nil {
		return c.Status(404).SendString("school not exist!");
	}
	return c.JSON(school);
}

func GetAllSchools(c *fiber.Ctx) error {
	var schools []School
	DB.Find(&schools)
	return c.JSON(schools)
}

func DeleteSchool(c *fiber.Ctx) error {
	SId := c.Params("school_id")
	s_id, err := strconv.ParseUint(SId, 10, 64)
	if err != nil {
		return c.Status(400).SendString("school id not found")
	}
	var school School
	DB.Where("id = ?", s_id).Delete(&school)
	return c.JSON(school)
}
func UpdateSchool(c *fiber.Ctx) error {
	var school School

	// Get school_id from path parameters
	id := c.Params("school_id")

	// Parse school_id to uint64
	Id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid school id")
	}
	
	// Check if school with given Id exists
	if err := DB.First(&school, Id).Error; err != nil {
		return c.Status(404).SendString("School not found")
	}

	// Parse request body into new_school
	var new_school School
	if err := c.BodyParser(&new_school); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	// Update fields if they are empty
	if new_school.Name == "" {
		new_school.Name = school.Name
	}
	if new_school.Description == "" {
		new_school.Description = school.Description
	}

	// Save updated school to database
	if err := DB.Model(&school).Updates(new_school).Error; err != nil {
		return c.Status(500).SendString("Failed to update school")
	}

	// Return updated school as JSON response
	return c.JSON(new_school)
}



// end school

// start department

func CreateDepartment(c *fiber.Ctx) error {
	department := new(Department)
	err := c.BodyParser(&department)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	School_ID := c.Params("school_id")
	SchoolID, err := strconv.ParseUint(School_ID, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid school ID")
	}
	department.SchoolID = uint(SchoolID)
	department.Save()
	return c.JSON(department)
}

func GetAllDepartment(c *fiber.Ctx) error {
	var department []Department
	DB.Find(&department)
	return c.JSON(department)
}

func GetDepartments(c *fiber.Ctx) error {
	Did := c.Params("school_id")
	d_id, err := strconv.ParseUint(Did, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid school ID")
	}
	var departments []Department
	DB.Where("school_id = ?", d_id).Find(&departments)
	return c.JSON(departments)
}

func GetDepartment(c *fiber.Ctx) error {
	id := c.Params("department_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid department id")
	}
	var department Department
	if err := DB.First(&department, Id).Error; err != nil {
		return c.Status(404).SendString("Department not found")
	}
	return c.JSON(department)
}

func UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("department_id")
	Id, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		return c.Status(400).SendString("Invalid department id")
	}
	var department Department
	if err := DB.First(&department, Id).Error; err != nil {
		return c.Status(404).SendString("Department not found")
	}

	var new_department Department
	if err := c.BodyParser(&new_department); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	if new_department.Name != "" {
		department.Name = new_department.Name
	}
	if new_department.Description != "" {
		department.Description = new_department.Description
	}

	if err := DB.Model(&department).Updates(new_department).Error; err != nil {
		return c.Status(500).SendString("Failed to update department")
	}
	return c.JSON(department)

}

func DeleteDepartment(c *fiber.Ctx) error {
	DId := c.Params("department_id")
	d_id, err := strconv.ParseUint(DId, 10, 64)
	if err != nil {
		return c.Status(400).SendString("department id not found")
	}
	var department Department
	DB.Where("id = ?", d_id).Delete(&department)
	return c.JSON(department)
}

// end department

// start course
func CreateCourse(c *fiber.Ctx) error {
	course := new(Course)
	err := c.BodyParser(&course)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	Did := c.Params("department_id")
	d_id, err := strconv.ParseUint(Did, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid department")
	}
	course.DepartmentID = uint(d_id)
	course.Save()
	return c.JSON(course)

}

func GetCourses(c *fiber.Ctx) error {
	var courses []Course
	result := DB.Find(&courses)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve courses",
	})
	}
	return c.JSON(courses)
}

func GetAllCourse(c *fiber.Ctx) error {
	Cid := c.Params("department_id")
	c_id, err := strconv.ParseUint(Cid, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid department")
	}
	var courses []Course
	DB.Where("department_id = ?", c_id).Find(&courses)
	return c.JSON(courses)
}

func GetCourse(c *fiber.Ctx) error {
	id := c.Params("course_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid course id")
	}
	var course Course
	if err := DB.First(&course, Id).Error; err != nil {
		return c.Status(404).SendString("Course not found")
	}
	return c.JSON(course)
}

func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("course_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid course id")
	}
	var course Course
	if err := DB.First(&course, Id).Error; err != nil {
		return c.Status(404).SendString("Course not found")
	}
	var new_course Course
	if err := c.BodyParser(&new_course); err != nil {
		return c.Status(400).SendString("Invalid input")
	}
	if new_course.Name != "" {
		course.Name = new_course.Name
	}
	if new_course.Description != "" {
		course.Description = new_course.Description
	}	
	if new_course.Code != "" {
		course.Code = new_course.Code
	}
	if err := DB.Model(&course).Updates(new_course).Error; err != nil {
		return c.Status(500).SendString("Failed to update course")
	}
	return c.JSON(course)
}

func DeleteCourse(c *fiber.Ctx) error {
	CId := c.Params("course_id")
	c_id, err := strconv.ParseUint(CId, 10, 64)
	if err != nil {
		return c.Status(400).SendString("course id not found")
	}
	var course Course
	DB.Where("id = ?", c_id).Delete(&course)
	return c.JSON(course)
}

// end course



// books

func CreateBook(c *fiber.Ctx) error {
	book := new(Books)
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	CId := c.Params("course_id")
	c_id, err := strconv.ParseUint(CId, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid course")
	}
	book.CourseID = uint(c_id)
	book.Save()
	return c.JSON(book)
}

func GetAllBooks(c *fiber.Ctx) error {
	Cid := c.Params("course_id")
	c_id, err := strconv.ParseUint(Cid, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid course")
	}
	var books []Books
	DB.Where("course_id = ?", c_id).Find(&books)
	return c.JSON(books)
}


// func GetBooks(c *fiber.Ctx) error {
// 	var books []Books
// 	// DB.Order("RAND()").Limit(5).Find(&books)
// 	// return c.JSON(books)
// 	result := DB.Order("RAND()").Limit(15).Find(&books)
//     if result.Error != nil {
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to retrieve books",
//         })
//     }
//     return c.JSON(books)
// }

func GetBooks(c *fiber.Ctx) error {
	// Parse query parameters directly from URL
	schoolID := c.Query("school_id")
	departmentID := c.Query("department_id")
	courseID := c.Query("course_id")
	author := c.Query("author")
	search := c.Query("search") // new search parameter

	var books []Books
	query := DB.Model(&Books{}).Order("RAND()").Limit(15)

	// Apply filters based on presence of query parameters
	if courseID != "" {
		query = query.Where("course_id = ?", courseID)
	} else if departmentID != "" {
		query = query.Joins("JOIN courses ON courses.id = books.course_id").
			Where("courses.department_id = ?", departmentID)
	} else if schoolID != "" {
		query = query.Joins("JOIN courses ON courses.id = books.course_id").
			Joins("JOIN departments ON departments.id = courses.department_id").
			Where("departments.school_id = ?", schoolID)
	}

	// Apply additional filters
	if author != "" {
		query = query.Where("author LIKE ?", "%" + author + "%")
	}

	// Apply search functionality
	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ? OR author LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve books",
		})
	}

	return c.JSON(books)
}




func DeleteBook(c *fiber.Ctx) error {
	BId := c.Params("book_id")
	b_id, err := strconv.ParseUint(BId, 10, 64)
	if err != nil {
		return c.Status(400).SendString("book id not found")
	}
	var book Books
	DB.Where("id = ?", b_id).Delete(&book)
	return c.JSON(book)
}

// end books
