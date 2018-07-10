package client

import (
	"path"
	"strings"
)

func isRoot(key string) bool {
	return key == "/"
}

// ensure key, return (realKey, parentKey)
func (clt *EtcdHRCHYClient) ensureKey(key string) (string, string, error) {
	if !strings.HasPrefix(key, "/") {
		return "", "", ErrorInvalidKey
	}

	if isRoot(key) {
		return clt.rootKey, clt.rootKey, nil
	}

	return key, path.Clean(key + "/../"), nil

}
