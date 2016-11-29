# Snom300FirmwareChecker

[![CircleCI](https://circleci.com/gh/Chabare/Snom300FirmwareChecker.svg?style=svg)](https://circleci.com/gh/Chabare/Snom300FirmwareChecker)

Gets the newest firmware + rollup from: 'http://wiki.snom.com/Category:Firmware:snom300'

Writes a file with the version numbers to ~/.snom.

The versions in the file are checked against the versions on  the website, if they aren't the same, the new file will be downloaded. The file will be downloaded into the execution directory. The link will be shown in the console.

## Install
Run `go get -u github.com/chabare/Snom300FirmwareChecker` -> Placed into $GOPATH/bin

## Example output

New:
```
Firmware number: V8_7_5_35 (new)
Rollup number: V8_7_5_44 (new)
Rollup link: http://downloads.snom.com/fw/mru-preview/snom300-8.7.5.44-SIP-f.bin
Firmware link: http://downloads.snom.com/fw/snom300-8.7.5.35-SIP-f.bin
```
Old:
```
Firmware number: V8_7_5_35 (old)
Rollup number: V8_7_5_44 (old)
```
One new:
```
Firmware number: V8_7_5_35 (new)
Rollup number: V8_7_5_44 (old)
Firmware link: http://downloads.snom.com/fw/snom300-8.7.5.35-SIP-f.bin
```
