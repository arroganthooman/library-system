package usecase

import (
	"fmt"
	"reflect"
	"testing"

	repo "github.com/arroganthooman/library-system/internal/repository"
)

func TestGetAllBook(t *testing.T) {
	type fields struct {
		bookRepo MockBookRepository
		want     []repo.Book
		wantErr  bool
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success",
			fields: fields{
				bookRepo: MockBookRepository{
					GetAllBookRes: []repo.Book{
						{
							ID:     1,
							Title:  "test title",
							Author: "fikri.akmal",
						},
					},
				},
				want: []repo.Book{
					{
						ID:     1,
						Title:  "test title",
						Author: "fikri.akmal",
					},
				},
			},
		},
		{
			name: "Error",
			fields: fields{
				bookRepo: MockBookRepository{
					GetAllBookErr: fmt.Errorf("any"),
				},
				want: []repo.Book{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			usecase := NewBookUsecase(&tc.fields.bookRepo)
			res, err := usecase.GetAllBook()
			if err != nil && !tc.fields.wantErr {
				t.Logf("testcase %s wantErr: %+v, got: %+v", tc.name, tc.fields.wantErr, true)
			}

			isEqualGotWithWant := reflect.DeepEqual(res, tc.fields.want)
			if isEqualGotWithWant {
				t.Logf("testcase %s want: %+v, got: %+v", tc.name, tc.fields.want, res)
			}
		})
	}
}
