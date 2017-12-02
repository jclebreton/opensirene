package update

import (
	"github.com/jclebreton/opensirene/domain"
	"github.com/jclebreton/opensirene/interfaces/files"
	"github.com/pkg/errors"
)

func checksumMatch(rf *domain.RemoteFile) (bool, error) {
	sum, err := files.CalculateChecksum(rf.Path, rf.Checksum.Type)
	if err != nil {
		return false, errors.Wrap(err, "unable to calculate checksum")
	}

	return sum == rf.Checksum.Value, nil
}
