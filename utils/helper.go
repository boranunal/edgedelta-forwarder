package utils

import (
	"bytes"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	X86Architecture = "x86_64"
	ArmArchitecture = "arm64"
	AmdArchitecture = "amd64"
)

func GetRuntimeArchitecture() string {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return AmdArchitecture
	}

	switch string(uname.Machine[:bytes.IndexByte(uname.Machine[:], 0)]) {
	case "aarch64":
		return ArmArchitecture
	default:
		return X86Architecture
	}
}

func SetMapKeyAndAppendIfExists(m map[string]string, k, v string) {
	var sb strings.Builder
	if val, ok := m[k]; ok {
		sb.WriteString(v)
		sb.WriteString(",")
		sb.WriteString(val)
		m[k] = sb.String()
	}
	m[k] = v
}

func SetMapKeyAndDuplicateWithPrefixIfExists(m map[string]string, prefix, k, v string) {
	var sb strings.Builder
	if _, ok := m[k]; ok {
		sb.WriteString(prefix)
		sb.WriteString(k)
		m[k] = v
	}
	m[k] = v
}
