package msidb

import (
	"golang.org/x/sys/windows/registry"
	"path/filepath"
)

var localMachineRegistryList = []string{
	`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
	`SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
}

func IsInstalled(productCode string) bool {
	for _, dir := range localMachineRegistryList {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE,
			filepath.Join(dir, productCode), registry.READ)
		if err == nil {
			k.Close()
			return true
		}
	}
	return false
}
