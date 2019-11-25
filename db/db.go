package db

import (
	"context"

	"github.com/matkinhig/echo-fw/types"
)

func GetAllStudent() (*[]types.Student, error) {

	var students []types.Student
	filter := map[string]interface{}{}

	cursor, err := Client.Database(DBName).Collection(StudentCols).Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	cursor.All(context.TODO(), &students)

	return &students, nil
}
