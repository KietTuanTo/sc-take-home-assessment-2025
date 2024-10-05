package folder_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		src     string
		dst     string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		err     error
	}{
		{
			name:  "Invalid Source Folder",
			src:   "beta",
			dst:   "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
			},

			want: []folder.Folder{},
			err:  errors.New("error: source folder does not exist"),
		},
		{
			name:  "Invalid Destination Folder",
			src:   "alpha",
			dst:   "beta",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
			},

			want: []folder.Folder{},
			err:  errors.New("error: destination folder does not exist"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.src, tt.dst)

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("MoveFolder() = %v, want %v for output", get, tt.want)
			}

			fmt.Println(tt.err)
			fmt.Println(err)
			if tt.err.Error() != err.Error() {
				t.Errorf("MoveFolder() = %v\n want %v for error", err, tt.err)
			}
		})
	}

}
