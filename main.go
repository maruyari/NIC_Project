package main

import (
	"NIC-Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("website/*.html")
	//	r.Static("../website/css", "../website/css")
	//	r.Static("website/css/admin.css", "../NIC_Project/website/css/admin.css")

	models.ConnectDataBase()
	r.GET("/student.html", func(c *gin.Context) {
		var students []models.Student
		var boards []string
		var exam []string
		models.DB.Find(&students).Group("board_name").Pluck("board_name", &boards)
		//var boards = models.DB.Raw("SELECT DISTINCT board_name FROM students")
		models.DB.Find(&students).Group("examination_name").Pluck("examination_name", &exam)
		c.HTML(http.StatusOK, "student.html", gin.H{
			"board": boards,
			"exam":  exam,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
	r.GET("/admin", func(c *gin.Context) {
		var students []models.Student
		var boards []string
		var exam []string
		var school []string
		var year []uint

		models.DB.Find(&students).Group("board_name").Pluck("board_name", &boards)
		//var boards = models.DB.Raw("SELECT DISTINCT board_name FROM students")
		models.DB.Find(&students).Group("examination_name").Pluck("examination_name", &exam)
		models.DB.Find(&students).Group("school_name").Pluck("school_name", &school)
		models.DB.Find(&students).Group("year_of_exam").Pluck("year_of_exam", &year)
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"board":  boards,
			"exam":   exam,
			"school": school,
			"year":   year,
		})
	})
	//http.HandleFunc("/displayTable", displayTable)

	r.POST("/add", func(c *gin.Context) {
		c.Request.ParseForm()
		var students []models.Student
		b := c.Request.FormValue("BOARD")
		e := c.Request.FormValue("EXAM")
		s := c.Request.FormValue("SCHOOL")
		y := c.Request.FormValue("YEAR")
		var names []string
		var rollno []string
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_name", &names)
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_roll_no", &rollno)
		type TableData struct {
			Name   string
			RollNo string
		}
		data := []TableData{}
		for i := 0; i < len(rollno); i++ {
			data = append(data, TableData{Name: names[i], RollNo: rollno[i]})
		}
		c.HTML(http.StatusOK, "add.html", gin.H{
			"board":  b,
			"school": s,
			"exam":   e,
			"year":   y,
			"name":   names,
			"rollno": rollno,
			"data":   data,
		})
		//fmt.Fprint(c.Writer,b,e,s,y)

	})
	r.POST("/modify", func(c *gin.Context) {
		c.Request.ParseForm()
		var students []models.Student
		b := c.Request.FormValue("BOARD")
		e := c.Request.FormValue("EXAM")
		s := c.Request.FormValue("SCHOOL")
		y := c.Request.FormValue("YEAR")
		var names []string
		var rollno []string
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_name", &names)
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_roll_no", &rollno)
		type TableData struct {
			Name   string
			RollNo string
		}
		data := []TableData{}
		for i := 0; i < len(rollno); i++ {
			data = append(data, TableData{Name: names[i], RollNo: rollno[i]})
		}
		c.HTML(http.StatusOK, "modify.html", gin.H{
			"board":  b,
			"school": s,
			"exam":   e,
			"year":   y,
			"name":   names,
			"rollno": rollno,
			"data":   data,
		})
		//fmt.Fprint(c.Writer,b,e,s,y)

	})
	r.POST("/view", func(c *gin.Context) {
		c.Request.ParseForm()
		var students []models.Student
		b := c.Request.FormValue("BOARD")
		e := c.Request.FormValue("EXAM")
		s := c.Request.FormValue("SCHOOL")
		y := c.Request.FormValue("YEAR")
		var names []string
		var rollno []string
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_name", &names)
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_roll_no", &rollno)
		type TableData struct {
			Name   string
			RollNo string
		}
		data := []TableData{}
		for i := 0; i < len(rollno); i++ {
			data = append(data, TableData{Name: names[i], RollNo: rollno[i]})
		}
		c.HTML(http.StatusOK, "view.html", gin.H{
			"board":  b,
			"school": s,
			"exam":   e,
			"year":   y,
			"name":   names,
			"rollno": rollno,
			"data":   data,
		})
	})

	r.GET("/find/", func(c *gin.Context) {

		rollno := c.Query("rollno")
		c.JSON(http.StatusOK, rollno)

	})
	r.Run()
}
