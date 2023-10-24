// ipfs_integration_test.go

package ipfs

import (
	"testing"
)

func TestRetrieveFileFromIPFS(t *testing.T) {
	// Replace with the actual IPFS hash you want to test
	ipfsHash := "QmSomeHash"

	data, err := RetrieveFileFromIPFS(ipfsHash)
	if err != nil {
		t.Errorf("Failed to retrieve file from IPFS: %v", err)
	}

	// Write your assertions here to validate the retrieved data.
	// For example, you can check if the data length is as expected.
	if len(data) == 0 {
		t.Errorf("Retrieved data is empty")
	}
}

func TestListIPFSDirectory(t *testing.T) {
	// Replace with the actual IPFS directory hash you want to test
	ipfsHash := "QmSomeDirectoryHash"

	files, err := ListIPFSDirectory(ipfsHash)
	if err != nil {
		t.Errorf("Failed to list IPFS directory: %v", err)
	}

	// Write your assertions here to validate the list of files.
	// For example, you can check if the number of files is as expected.
	if len(files) == 0 {
		t.Errorf("No files found in the IPFS directory")
	}
}
