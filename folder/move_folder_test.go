package folder_test

import (
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
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.src, tt.dst)

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("MoveFolder() = %v, want %v for output", get, tt.want)
			}

			if tt.err != err {
				t.Errorf("MoveFolder() = %v, want %v for error", err, tt.err)
			}
		})
	}

}
