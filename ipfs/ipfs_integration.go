package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

// initIPFS initializes an IPFS shell and returns it
func initIPFS() *shell.Shell {
	// Create an IPFS shell instance connected to a local node
	sh := shell.NewShell("localhost:5001") // You should provide the correct IPFS node address

	return sh
}

// RetrieveFileFromIPFS retrieves a file from IPFS by its hash
func RetrieveFileFromIPFS(hash string) ([]byte, error) {
	sh := initIPFS()

	// Get the file from IPFS by its hash
	reader, err := sh.Cat(hash)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Read the file data
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ListIPFSDirectory lists the contents of an IPFS directory by its hash
func ListIPFSDirectory(hash string) ([]string, error) {
	sh := initIPFS()

	// List the contents of the directory
	entries, err := sh.List(hash)
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, entry := range entries {
		fileList = append(fileList, entry.Name)
	}

	return fileList, nil
}
