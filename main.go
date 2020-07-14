package msidb

import (
	"fmt"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func CoInit() func() {
	ole.CoInitializeEx(0, 0)

	return func() { ole.CoUninitialize() }
}

func Query(msipath string) (map[string]string, error) {
	msifullpath, err := filepath.Abs(msipath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", msipath, err)
	}
	_installer, err := oleutil.CreateObject("WindowsInstaller.Installer")
	if err != nil {
		return nil, fmt.Errorf("CreateObject(WindowsInstaller.Installer): %w", err)
	}
	installer, err := _installer.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, fmt.Errorf("_installer.QueryInterface: %w", err)
	}
	defer installer.Release()

	_db, err := installer.CallMethod("OpenDatabase", msifullpath, 0)
	if err != nil {
		return nil, fmt.Errorf("installer.OpenDatabase: %w", err)
	}

	db := _db.ToIDispatch()
	defer db.Release()

	_view, err := db.CallMethod("OpenView", "SELECT `Property`,`Value` FROM `Property`")
	if err != nil {
		return nil, fmt.Errorf("OpenView: %w", err)
	}
	view := _view.ToIDispatch()
	defer view.Release()

	_, err = view.CallMethod("Execute")
	if err != nil {
		return nil, fmt.Errorf("view.Execute: %w", err)
	}

	result := map[string]string{}
	for {
		_record, err := view.CallMethod("Fetch")
		if err != nil || _record == nil {
			break
		}
		record := _record.ToIDispatch()
		if record == nil {
			break
		}
		key, err1 := record.GetProperty("StringData", 1)
		val, err2 := record.GetProperty("StringData", 2)
		if err1 == nil && err2 == nil {
			result[key.ToString()] = val.ToString()
		}
		record.Release()
	}
	return result, nil
}
