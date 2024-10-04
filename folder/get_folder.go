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

func (f *driver) GetChildren(parent FileNode) []Folder {
	children := []Folder{}
	for _, fileNodePtr := range parent.children {
		children = append(children, fileNodePtr.file)
		children = append(children, f.GetChildren(*fileNodePtr)...)
	}

	return children
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	org, exists := f.orgs[orgID]

	if !exists {
		return []Folder{}
	}

	var parentNode *FileNode
	for _, fileNode := range org.folders {
		if fileNode.file.Name == name {
			parentNode = &fileNode
		}
	}

	if parentNode == nil {
		return []Folder{}
	}

	return f.GetChildren(*parentNode)
}
