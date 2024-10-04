package folder

import (
	"github.com/gofrs/uuid"
)

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type FileNode struct {
	file     Folder
	parent   *FileNode
	children []*FileNode
}

func NewFileNode(folder Folder) FileNode {
	return FileNode{
		file:     folder,
		parent:   nil,
		children: []*FileNode{},
	}
}

type Organization struct {
	folders []FileNode
}

func NewOrg() Organization {
	return Organization{
		folders: []FileNode{},
	}
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data
	orgs map[uuid.UUID]Organization

	// example: feel free to change the data structure, if slice is not what you want
	// folders []Folder
}

func NewDriver(folders []Folder) IDriver {
	orgs := map[uuid.UUID]Organization{}

	for _, f := range folders {
		_, exists := orgs[f.OrgId]
		if !exists {
			orgs[f.OrgId] = NewOrg()
		}
		org := orgs[f.OrgId]
		org.folders = append(org.folders, NewFileNode(f))
		orgs[f.OrgId] = org
	}

	return &driver{
		orgs: orgs,
	}
}
