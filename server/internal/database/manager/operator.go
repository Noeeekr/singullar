// package OperationsManager provides reliable database OperationsManager for the api
package manager

import (
	"strings"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

type Operator struct {
	*borm.Transaction
}

func NewOperator(tx *borm.Transaction) *Operator {
	return &Operator{
		Transaction: tx,
	}
}

func addSelectClassesFilters(query *borm.Query, filters *[]*FilterClassesOptions) *borm.ConditionalQuery {
	conditions := make([]*borm.ConditionalQuery, len(*filters))
	for i, filter := range *filters {
		condition := []*borm.ConditionalQuery{}
		if filter.Id != nil {
			condition = append(condition, query.Field("c.id").IsEqual(*filter.Id))
		}
		if filter.ClassName == nil {
			condition = append(condition, query.Field("c.name").IsLike("%", false))
		} else {
			condition = append(condition, query.Field("c.name").IsLike("%"+*filter.ClassName+"%", false))
		}
		if filter.StudentName != nil {
			condition = append(condition, query.Compose(query.And(
				query.Field("u.name").IsLike("%"+strings.ToLower(*filter.StudentName)+"%", false),
				query.Field("u.role").IsEqual(models.STUDENT),
			)))
		}
		if filter.TeacherName != nil {
			condition = append(condition, query.Compose(query.And(
				query.Field("u.name").IsLike("%"+strings.ToLower(*filter.TeacherName)+"%", false),
				query.Field("u.role").IsEqual(models.TEACHER),
			)))
		}
		if filter.Series != nil {
			condition = append(condition, query.Field("c.series").IsEqual(*filter.Series))
		}
		if filter.Segment != nil {
			condition = append(condition, query.Field("c.segment").IsEqual(*filter.Segment))
		}
		if filter.CreationYear != nil {
			// query.And("c.created_at").After(filter.CreationYear) => Which is time.Time()
		}
		conditions[i] = query.Compose(query.And(condition...))
	}
	return query.Compose(query.Or(conditions...))
}
func addSelectStudentsFilters(query *borm.Query, filters *[]FilterStudentOptions) *borm.ConditionalQuery {
	conditions := make([]*borm.ConditionalQuery, len(*filters))
	for i, filter := range *filters {
		condition := []*borm.ConditionalQuery{}
		if filter.Segment == models.UNKNOWN_SEGMENT {
			condition = append(condition, query.Field("u.segment").IsEqual(nil))
		} else {
			condition = append(condition, query.Field("u.segment").IsEqual(filter.Segment))
		}
		if filter.Email != "" {
			condition = append(condition, query.Field("u.email").IsEqual(filter.Email))
		}
		if filter.Name != "" {
			condition = append(condition, query.Field("u.name").IsLike("%"+strings.ToLower(filter.Name)+"%", false))
		}
		if filter.ID != nil {
			condition = append(condition, query.Field("u.id").IsEqual(filter.ID))
		}
		if filter.ClassID != nil {
			condition = append(condition, query.Field("c.id").IsEqual(filter.ClassID))
		}
		conditions[i] = query.And(condition...)
	}
	return query.Compose(query.Or(conditions...))
}
