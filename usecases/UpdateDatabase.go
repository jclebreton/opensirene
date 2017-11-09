package usecases

import (
	"errors"

	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

type UpdateDatabaseRequest struct{}

// UpdateDatabase will grab and parse update files from open data to the database
func (i *Interactor) UpdateDatabase(r GetDBStatusRequest) (bool, error) {

	// Retrieve the list of update files stored in database
	// Search updates
	// Files to grab

	return true, nil
}

func (r *UpdateDatabaseRequest) getDatabaseStatus(i *Interactor) ([]string, error) {
	hh, err := i.DBStatusRW.GetSuccessList()
	if err != nil {
		return nil, err
	}

	if hh == nil || len(hh) == 0 {
		return nil, errors.New("nothing found")
	}

	return hh, nil
}

func (r *UpdateDatabaseRequest) getFilesToImport(localFiles []string, remoteFiles sirene.RemoteFiles) sirene.RemoteFiles {
	result := sirene.RemoteFiles{}
	for _, remoteFile := range remoteFiles {
		if !r.isAlreadyImported(localFiles, remoteFile.FileName) {
			result = append(result, remoteFile)
		}
	}
	return result
}

func (r *UpdateDatabaseRequest) isAlreadyImported(localFiles []string, remoteFile string) bool {
	for _, localFile := range localFiles {
		if remoteFile == localFile {
			return true
		}
	}
	return false
}

func (r *UpdateDatabaseRequest) startUpdate() {}
