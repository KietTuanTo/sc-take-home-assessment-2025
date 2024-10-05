package folder

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

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

func (f *driver) GetChildren(parent *FileNode) []Folder {
	nodeChildren := []Folder{}
	for _, fileNodePtr := range parent.children {
		nodeChildren = append(nodeChildren, fileNodePtr.file)
		nodeChildren = append(nodeChildren, f.GetChildren(fileNodePtr)...)
	}

	return nodeChildren
}

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

	return f.GetChildren(parentNode)
}
