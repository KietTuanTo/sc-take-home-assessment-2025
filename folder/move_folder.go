package folder

import (
	"errors"

	"github.com/gofrs/uuid"
)

// FindFolder returns a pointer to the FileNode and a string
// representation of its UUID, given the name of the Folder it
// contains
func FindFolder(name string, orgs map[uuid.UUID]Organization) (*FileNode, string) {
	for orgID, org := range orgs {
		for _, fileNode := range org.folders {
			if fileNode.file.Name == name {
				return fileNode, orgID.String()
			}
		}
	}

	return nil, ""
}

// CheckIsChild returns a boolean stating whether the
// given 'dst' FileNode is a child of 'src'
func CheckIsChild(src *FileNode, dst *FileNode) bool {
	for _, childNode := range src.children {
		if childNode.file.Name == dst.file.Name {
			return true
		} else if CheckIsChild(childNode, dst) {
			return true
		}
	}

	return false
}

// RemoveChild returns the children of 'parentNode'
// after removing the FileNode containing the
// Folder with name 'fileToRemove'
func RemoveChild(parentNode *FileNode, fileToRemove string) []*FileNode {
	for i, fileNode := range parentNode.children {
		if fileNode.file.Name == fileToRemove {
			return append(parentNode.children[:i], parentNode.children[:i+1]...)
		}
	}

	return parentNode.children
}

// ChangeChildPaths changes all paths of the Folders
// contained within child FileNodes of 'parentNode'
func ChangeChildPaths(parentNode *FileNode) {
	for _, childNode := range parentNode.children {
		childNode.file.Paths = parentNode.file.Paths + "." + childNode.file.Name
		ChangeChildPaths(childNode)
	}
}

// CreateFolderSlice returns a slice containing all the folders
// stored within the drive
func CreateFolderSlice(orgs map[uuid.UUID]Organization) []Folder {
	folders := []Folder{}
	for _, org := range orgs {
		for _, fileNode := range org.folders {
			folders = append(folders, fileNode.file)
		}
	}

	return folders
}

// MoveFolder moves a folder with 'name' and all its children
// to a different parent folder
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == dst {
		return []Folder{}, errors.New("error: cannot move a folder to itself")
	}

	srcFolder, srcID := FindFolder(name, f.orgs)
	dstFolder, dstID := FindFolder(dst, f.orgs)
	if srcFolder == nil {
		return []Folder{}, errors.New("error: source folder does not exist")
	}
	if dstFolder == nil {
		return []Folder{}, errors.New("error: destination folder does not exist")
	}
	if srcID != dstID {
		return []Folder{}, errors.New("error: cannot move a folder to a different organization")
	}
	if CheckIsChild(srcFolder, dstFolder) {
		return []Folder{}, errors.New("error: cannot move a folder to a child of itself")
	}

	// Change parent of source file to new parent, and remove source file
	// from the children of old parent node
	srcParent := srcFolder.parent
	if srcParent != nil {
		srcParent.children = RemoveChild(srcParent, srcFolder.file.Name)
	}

	// Set parent of source file to new parent, add source file to destination
	// node children.
	srcFolder.parent = dstFolder
	dstFolder.children = append(dstFolder.children, srcFolder)

	// Change file paths according to new parent file path for all children
	// in the subtree that has been moved
	srcFolder.file.Paths = srcFolder.parent.file.Paths + "." + srcFolder.file.Name
	ChangeChildPaths(srcFolder)

	folderList := CreateFolderSlice(f.orgs)
	return folderList, nil
}
