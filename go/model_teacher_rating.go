/*
 * ITU Super API
 *
 * Test server for ITU Information
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TeacherRating struct {

	// The student giving the rating
	StudentId int32 `json:"student_id,omitempty"`

	// The teacher getting rated
	TeacherId int32 `json:"teacher_id,omitempty"`

	// The rating the student is giving
	Rating int32 `json:"rating,omitempty"`
}
