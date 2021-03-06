/*
 * ITU Super API
 *
 * Test server for ITU Information
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CourseCourseIdDelete",
		strings.ToUpper("Delete"),
		"/course/{course_id}",
		CourseCourseIdDelete,
	},

	Route{
		"CourseCourseIdGet",
		strings.ToUpper("Get"),
		"/course/{course_id}",
		CourseCourseIdGet,
	},

	Route{
		"CoursePost",
		strings.ToUpper("Post"),
		"/course",
		CoursePost,
	},

	Route{
		"CoursePut",
		strings.ToUpper("Put"),
		"/course",
		CoursePut,
	},

	Route{
		"AddStudent",
		strings.ToUpper("Post"),
		"/student",
		AddStudent,
	},

	Route{
		"DeleteStudent",
		strings.ToUpper("Delete"),
		"/student/{student_id}",
		DeleteStudent,
	},

	Route{
		"GetCourseRating",
		strings.ToUpper("Get"),
		"/student/{student_id}/{course_id}",
		GetCourseRating,
	},

	Route{
		"GetStudentByID",
		strings.ToUpper("Get"),
		"/student/{student_id}",
		GetStudentByID,
	},

	Route{
		"GetWorkload",
		strings.ToUpper("Get"),
		"/student/{student_id}/workload",
		GetWorkload,
	},

	Route{
		"UpdateRating",
		strings.ToUpper("Put"),
		"/student/{student_id}/{course_id}",
		UpdateRating,
	},

	Route{
		"TeacherPost",
		strings.ToUpper("Post"),
		"/teacher",
		TeacherPost,
	},

	Route{
		"TeacherTeacherIdGet",
		strings.ToUpper("Get"),
		"/teacher/{teacher_id}",
		TeacherTeacherIdGet,
	},
}
