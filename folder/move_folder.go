package folder

import (
	"errors"

	"github.com/gofrs/uuid"
)

func findFolder(name string, orgs map[uuid.UUID]Organization) (*FileNode, string) {
	for orgID, org := range orgs {
		for _, fileNode := range org.folders {
			if fileNode.file.Name == name {
				return fileNode, orgID.String()
			}
		}
	}

	return nil, ""
}

func checkIsChild(src *FileNode, dst *FileNode) bool {
	for _, childNode := range src.children {
		if childNode.file.Name == dst.file.Name {
			return true
		} else if checkIsChild(childNode, dst) {
			return true
		}
	}

	return false
}

func removeChild(parentNode *FileNode, fileToRemove string) []*FileNode {
	for i, fileNode := range parentNode.children {
		if fileNode.file.Name == fileToRemove {
			return append(parentNode.children[:i], parentNode.children[:i+1]...)
		}
	}

	return parentNode.children
}

func changeChildPaths(parentNode *FileNode) {
	for _, childNode := range parentNode.children {
		childNode.file.Paths = parentNode.file.Paths + "." + childNode.file.Name
		changeChildPaths(childNode)
	}
}

func createFolderSlice(orgs map[uuid.UUID]Organization) []Folder {
	folders := []Folder{}
	for _, org := range orgs {
		for _, fileNode := range org.folders {
			folders = append(folders, fileNode.file)
		}
	}

	return folders
}

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == dst {
		return []Folder{}, errors.New("error: cannot move a folder to itself")
	}

	srcFolder, srcID := findFolder(name, f.orgs)
	dstFolder, dstID := findFolder(dst, f.orgs)
	if srcFolder == nil {
		return []Folder{}, errors.New("error: source folder does not exist")
	}
	if dstFolder == nil {
		return []Folder{}, errors.New("error: destination folder does not exist")
	}
	if srcID != dstID {
		return []Folder{}, errors.New("error: cannot move a folder to a different organization")
	}
	if checkIsChild(srcFolder, dstFolder) {
		return []Folder{}, errors.New("error: cannot move a folder to a child of itself")
	}

	srcParent := srcFolder.parent
	if srcParent != nil {
		srcParent.children = removeChild(srcParent, srcFolder.file.Name)
	}

	srcFolder.parent = dstFolder
	dstFolder.children = append(dstFolder.children, srcFolder)
	srcFolder.file.Paths = srcFolder.parent.file.Paths + "." + srcFolder.file.Name

	changeChildPaths(srcFolder)

	folderList := createFolderSlice(f.orgs)
	return folderList, nil
}
