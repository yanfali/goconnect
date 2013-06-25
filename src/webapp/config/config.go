package config

import (
	"path/filepath"
	"sync"
)

var (
	mutex  sync.Mutex
	Values = make(map[int]interface{})
)

const (
	BASEDIR = iota
	APPDIR
)

func SetBaseDir(path string) error {
	mutex.Lock()
	defer mutex.Unlock()
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	Values[BASEDIR] = filepath.Dir(absPath)
	Values[APPDIR] = filepath.Clean(Values[BASEDIR].(string) + "/../webapp")
	return nil
}

func GetBaseDir() string {
	return getValue(BASEDIR).(string)
}

func GetAppDir() string {
	return getValue(APPDIR).(string)
}

func getValue(key int) interface{} {
	mutex.Lock()
	defer mutex.Unlock()
	return Values[key]
}

//TODO Placeholder - load it from a file which never gets checked in
func GetHashKey() []byte {
	return []byte("0123456789ABCDEF0123456789ABCDEF")
}

//TODO Placeholder - load it from a file which never gets checked in
func GetBlockKey() []byte {
	return nil
}
