package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

type FileNode struct {
	file     Folder
	parent   *FileNode
	children []*FileNode
}

type Organization struct {
	folders []*FileNode
}

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

func NewDriver(folders []Folder) IDriver {
	orgs := GenerateOrgs(folders)

	return &driver{
		orgs: orgs,
	}
}

// NewFileNode returns a pointer to a FileNode, containing
// a given folder, with default iniitialized values for parent
// and child
func NewFileNode(folder Folder) *FileNode {
	return &FileNode{
		file:     folder,
		parent:   nil,
		children: []*FileNode{},
	}
}

// NewOrg returns an Organization struct
func NewOrg() Organization {
	return Organization{
		folders: []*FileNode{},
	}
}

type driver struct {
	orgs map[uuid.UUID]Organization
}

// FindFileNode returns a pointer to the FileNode with
// a given name, stored inside the same Organization
func FindFileNode(folders []*FileNode, name string) *FileNode {
	for _, f := range folders {
		if f.file.Name == name {
			return f
		}
	}

	return nil
}

// GenerateFileNodes returns a map hashed by UUIDs, storing
// Organizations which contains a slice of pointers to
// their organization's respective FileNodes
func GenerateFileNodes(folders []Folder, orgs map[uuid.UUID]Organization) {
	for _, f := range folders {
		_, exists := orgs[f.OrgId]
		if !exists {
			orgs[f.OrgId] = NewOrg()
		}
		org := orgs[f.OrgId]
		org.folders = append(org.folders, NewFileNode(f))
		orgs[f.OrgId] = org
	}
}

// GenerateNodeParents changes the 'parent' field of
// each folder in 'folders' to the FileNode whose file
// name is the immediate parent of each file given in
// their path
func GenerateNodeParents(folders []*FileNode) {
	for i, fileNode := range folders {
		curr_path := fileNode.file.Paths
		path_sections := strings.Split(curr_path, ".")

		if len(path_sections) <= 1 {
			continue
		}

		// The name of the immediate parent FileNode is given by the second
		// last directory in the file's path, as the last is itself
		parent := FindFileNode(folders, path_sections[len(path_sections)-2])
		if parent == nil {
			continue
		}

		folders[i].parent = parent
		parent.children = append(parent.children, fileNode)
	}
}

// GenerateNodeChildren returns a slice of FileNode pointers, who are
// the immediate children of a given 'parentNode"
func GenerateNodeChildren(folders []*FileNode, parentNode *FileNode) []*FileNode {
	children := []*FileNode{}
	for _, childNode := range folders {
		if childNode.parent != nil && childNode.parent.file.Name == parentNode.file.Name {
			children = append(children, childNode)
		}
	}
	return children
}

// GenerateOrgs returns a map of Organizations, hashed
// by the Organization's OrgId and containing a slice of
// pointers to all FileNodes contained in that Organization
func GenerateOrgs(folders []Folder) map[uuid.UUID]Organization {
	orgs := map[uuid.UUID]Organization{}
	GenerateFileNodes(folders, orgs)

	for _, org := range orgs {
		GenerateNodeParents(org.folders)
		for i, fileNode := range org.folders {
			org.folders[i].children = GenerateNodeChildren(org.folders, fileNode)
		}
	}

	return orgs
}
