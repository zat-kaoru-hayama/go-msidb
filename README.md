go-msidb
=========

The golang library to get the version of the Windows Installer.

msiver.exe &amp; gmsiver.exe
========================

The tool to get the version of the Windows Installer.

CUI Version:
```
msiver [FILEPATH]
```

GUI Version:

```
gmsiver [FILEPATH]
```

FILEPATH is one of these.

- .MSI file's path
- The Directory path containing .MSI Files
- .ZIP file's path containing .MSI Files

Install
-------

Download the binary package from [Releases](https://github.com/zat-kaoru-hayama/go-msidb/releases) and extract the executable.

### for scoop-installer

```
scoop install https://raw.githubusercontent.com/zat-kaoru-hayama/go-msidb/master/msiver.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install msiver
```
