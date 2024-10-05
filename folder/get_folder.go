package folder

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders returns a slice of Folders
// corresponding to the generated sample data
func GetAllFolders() []Folder {
	return GetSampleData()
}

// GetFoldersByOrgID returns a slice of Folders
// which have a certain orgID
func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	value, exists := f.orgs[orgID]
	if !exists {
		return []Folder{}
	}

	res := []Folder{}
	for _, f := range value.folders {
		res = append(res, f.file)
	}
	return res
}

// GetChildren returns a slice of Folders containing
// all the children of a FileNode 'parent'
func GetChildren(parent *FileNode) []Folder {
	nodeChildren := []Folder{}
	for _, fileNodePtr := range parent.children {
		nodeChildren = append(nodeChildren, fileNodePtr.file)
		nodeChildren = append(nodeChildren, GetChildren(fileNodePtr)...)
	}

	return nodeChildren
}

// GetAllChildFolders returns the slice of Folders generated using
// GetChildren, but ensures that the orgID is valid, and the name of
// the file exists in the organization
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	org, exists := f.orgs[orgID]

	if !exists {
		return []Folder{}
	}

	var parentNode *FileNode = nil
	for _, fileNode := range org.folders {
		if fileNode.file.Name == name {
			parentNode = fileNode
		}
	}

	if parentNode == nil {
		return []Folder{}
	}

	return GetChildren(parentNode)
}
