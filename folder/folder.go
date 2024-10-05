package folder

import (
	"strings"

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

func NewFileNode(folder Folder) *FileNode {
	return &FileNode{
		file:     folder,
		parent:   nil,
		children: []*FileNode{},
	}
}

type Organization struct {
	folders []*FileNode
}

func NewOrg() Organization {
	return Organization{
		folders: []*FileNode{},
	}
}

type driver struct {
	orgs map[uuid.UUID]Organization
}

func findFileNode(folders []*FileNode, name string) *FileNode {
	for _, f := range folders {
		if f.file.Name == name {
			return f
		}
	}

	return nil
}

func generateFileNodes(folders []Folder, orgs map[uuid.UUID]Organization) {
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

func generateNodeParents(folders []*FileNode) {
	for i, fileNode := range folders {
		curr_path := fileNode.file.Paths
		path_sections := strings.Split(curr_path, ".")

		if len(path_sections) <= 1 {
			continue
		}

		parent := findFileNode(folders, path_sections[len(path_sections)-2])
		if parent == nil {
			continue
		}

		folders[i].parent = parent
		parent.children = append(parent.children, fileNode)
	}
}

func generateNodeChildren(folders []*FileNode, parentNode *FileNode) []*FileNode {
	children := []*FileNode{}
	for _, childNode := range folders {
		if childNode.parent != nil && childNode.parent.file.Name == parentNode.file.Name {
			children = append(children, childNode)
		}
	}
	return children
}

func generateOrgs(folders []Folder) map[uuid.UUID]Organization {
	orgs := map[uuid.UUID]Organization{}
	generateFileNodes(folders, orgs)

	for _, org := range orgs {
		generateNodeParents(org.folders)
		for i, fileNode := range org.folders {
			org.folders[i].children = generateNodeChildren(org.folders, fileNode)
		}
	}

	return orgs
}

func NewDriver(folders []Folder) IDriver {
	orgs := generateOrgs(folders)

	return &driver{
		orgs: orgs,
	}
}
