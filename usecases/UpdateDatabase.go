package usecases

import "github.com/jclebreton/opensirene/domain"

type UpdateDatabaseRequest struct {
	DPath string
}

// UpdateDatabase will grab and parse update files from open data to the database
func (i *Interactor) UpdateDatabase(r UpdateDatabaseRequest) (bool, error) {

	// Retrieve the list of update files stored in database
	// Search updates
	// Files to grab

	_, err := r.getRemoteFiles(i)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UpdateDatabaseRequest) getRemoteFiles(i *Interactor) ([]domain.RemoteFile, error) {
	rfs, err := i.SireneR.GetRemoteFiles(r.DPath)
	if err != nil {
		return nil, err
	}

	return rfs, nil
}

//
//func (r *UpdateDatabaseRequest) getDatabaseStatus(i *Interactor) ([]string, error) {
//	hh, err := i.DBStatusRW.GetSuccessList()
//	if err != nil {
//		return nil, err
//	}
//
//	if hh == nil || len(hh) == 0 {
//		return nil, errors.New("nothing found")
//	}
//
//	return hh, nil
//}
//
//func (r *UpdateDatabaseRequest) getFilesToImport(localFiles []string, remoteFiles sirene.RemoteFiles) sirene.RemoteFiles {
//	result := sirene.RemoteFiles{}
//	for _, remoteFile := range remoteFiles {
//		if !r.isAlreadyImported(localFiles, remoteFile.FileName) {
//			result = append(result, remoteFile)
//		}
//	}
//	return result
//}
//
//func (r *UpdateDatabaseRequest) isAlreadyImported(localFiles []string, remoteFile string) bool {
//	for _, localFile := range localFiles {
//		if remoteFile == localFile {
//			return true
//		}
//	}
//	return false
//}
//
//func (r *UpdateDatabaseRequest) startUpdate() {}
