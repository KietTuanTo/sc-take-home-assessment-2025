package folder_test

import (
	"errors"
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
		{
			name:  "Moving to same folder",
			src:   "alpha",
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
			err:  errors.New("error: cannot move a folder to itself"),
		},
		{
			name:  "Source and Destination from different orgs",
			src:   "alpha",
			dst:   "beta",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "beta",
				},
			},

			want: []folder.Folder{},
			err:  errors.New("error: cannot move a folder to a different organization"),
		},
		{
			name:  "Destination is an immediate child of Source",
			src:   "alpha",
			dst:   "beta",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta",
				},
			},

			want: []folder.Folder{},
			err:  errors.New("error: cannot move a folder to a child of itself"),
		},
		{
			name:  "Destination is a child of Source",
			src:   "alpha",
			dst:   "charlie",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta",
				},
				{
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie",
				},
			},

			want: []folder.Folder{},
			err:  errors.New("error: cannot move a folder to a child of itself"),
		},
		{
			name:  "Basic valid use case",
			src:   "alpha",
			dst:   "beta",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "beta",
				},
				{
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie",
				},
			},

			want: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "beta.alpha",
				},
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "beta",
				},
				{
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "beta.alpha.charlie",
				},
			},
			err: nil,
		},
		{
			name:  "Move folder to another root folder",
			src:   "bravo",
			dst:   "golf",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.delta.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.delta.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "delta",
					Paths: "alpha.delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "echo",
					Paths: "alpha.delta.echo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "golf",
					Paths: "golf",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},

			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "golf.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "golf.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "delta",
					Paths: "alpha.delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "echo",
					Paths: "alpha.delta.echo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "golf",
					Paths: "golf",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			err: nil,
		},
		{
			name:  "Move folder to a non-root folder",
			src:   "bravo",
			dst:   "delta",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "delta",
					Paths: "alpha.delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "echo",
					Paths: "alpha.delta.echo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "golf",
					Paths: "golf",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},

			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.delta.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.delta.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "delta",
					Paths: "alpha.delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "echo",
					Paths: "alpha.delta.echo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "golf",
					Paths: "golf",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			err: nil,
		},
		{
			name:  "Move folder to subsequent parent folder",
			src:   "charlie",
			dst:   "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.charlie",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.charlie.delta",
				},
				{
					Name:  "gamma",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.charlie.gamma",
				},
				{
					Name:  "epsilon",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.charlie.delta.epsilon",
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie.delta",
				},
				{
					Name:  "gamma",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie.gamma",
				},
				{
					Name:  "epsilon",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie.delta.epsilon",
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.src, tt.dst)

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("MoveFolder() = %v, want %v for output", get, tt.want)
			}

			if tt.err != nil && err == nil {
				t.Errorf("MoveFolder() = nil, want %v for error", tt.err)
			} else if tt.err == nil && err != nil {
				t.Errorf("MoveFolder() = %v, want nil for error", tt.err)
			} else if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("MoveFolder() = %v\n want %v for error", err, tt.err)
			}
		})
	}

}
