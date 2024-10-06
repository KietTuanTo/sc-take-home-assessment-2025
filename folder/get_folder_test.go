package folder_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Test with Valid UUID",
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
					Name:  "gamma",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "gamma",
				},
			},

			want: []folder.Folder{
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
			},
		},
		{
			name:  "Test with Invalid UUID",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID + "."),
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
					Name:  "gamma",
					OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
					Paths: "gamma",
				},
			},

			want: []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", get, tt.want)
			}
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		src     string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Test with Invalid UUID",
			src:   "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID + "."),
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
		},
		{
			name:  "Test with Non Existent FileName",
			src:   "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta",
				},
			},

			want: []folder.Folder{},
		},
		{
			name:  "Test with File Name in Different Org",
			src:   "alpha",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
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
		},
		{
			name:  "Test only inside organisation",
			src:   "alpha",
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
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "alpha",
				},
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "alpha.beta",
				},
			},

			want: []folder.Folder{
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta",
				},
			},
		},
		{
			name:  "Test with Valid UUID",
			src:   "alpha",
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
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.delta",
				},
				{
					Name:  "gamma",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta.gamma",
				},
			},

			want: []folder.Folder{
				{
					Name:  "beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta",
				},
				{
					Name:  "gamma",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.beta.gamma",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.delta",
				},
			},
		},
		{
			name:  "Test with no children",
			src:   "alpha",
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
			},

			want: []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.src)

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", get, tt.want)
			}
		})
	}
}
