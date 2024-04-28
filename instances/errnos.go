// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import "syscall"

// ErrnoNames converts Linux errno values to their string representation.
func ErrnoNames() map[syscall.Errno]string {
	return map[syscall.Errno]string{
		syscall.Errno(0x7):  "E2BIG",
		syscall.Errno(0xd):  "EACCES",
		syscall.Errno(0x62): "EADDRINUSE",
		syscall.Errno(0x63): "EADDRNOTAVAIL",
		syscall.Errno(0x44): "EADV",
		syscall.Errno(0x61): "EAFNOSUPPORT",
		syscall.Errno(0xb):  "EAGAIN",
		syscall.Errno(0x72): "EALREADY",
		syscall.Errno(0x34): "EBADE",
		syscall.Errno(0x9):  "EBADF",
		syscall.Errno(0x4d): "EBADFD",
		syscall.Errno(0x4a): "EBADMSG",
		syscall.Errno(0x35): "EBADR",
		syscall.Errno(0x38): "EBADRQC",
		syscall.Errno(0x39): "EBADSLT",
		syscall.Errno(0x3b): "EBFONT",
		syscall.Errno(0x10): "EBUSY",
		syscall.Errno(0x7d): "ECANCELED",
		syscall.Errno(0xa):  "ECHILD",
		syscall.Errno(0x2c): "ECHRNG",
		syscall.Errno(0x46): "ECOMM",
		syscall.Errno(0x67): "ECONNABORTED",
		syscall.Errno(0x6f): "ECONNREFUSED",
		syscall.Errno(0x68): "ECONNRESET",
		syscall.Errno(0x23): "EDEADLOCK",
		syscall.Errno(0x59): "EDESTADDRREQ",
		syscall.Errno(0x21): "EDOM",
		syscall.Errno(0x49): "EDOTDOT",
		syscall.Errno(0x7a): "EDQUOT",
		syscall.Errno(0x11): "EEXIST",
		syscall.Errno(0xe):  "EFAULT",
		syscall.Errno(0x1b): "EFBIG",
		syscall.Errno(0x70): "EHOSTDOWN",
		syscall.Errno(0x71): "EHOSTUNREACH",
		syscall.Errno(0x85): "EHWPOISON",
		syscall.Errno(0x2b): "EIDRM",
		syscall.Errno(0x54): "EILSEQ",
		syscall.Errno(0x73): "EINPROGRESS",
		syscall.Errno(0x4):  "EINTR",
		syscall.Errno(0x16): "EINVAL",
		syscall.Errno(0x5):  "EIO",
		syscall.Errno(0x6a): "EISCONN",
		syscall.Errno(0x15): "EISDIR",
		syscall.Errno(0x78): "EISNAM",
		syscall.Errno(0x7f): "EKEYEXPIRED",
		syscall.Errno(0x81): "EKEYREJECTED",
		syscall.Errno(0x80): "EKEYREVOKED",
		syscall.Errno(0x33): "EL2HLT",
		syscall.Errno(0x2d): "EL2NSYNC",
		syscall.Errno(0x2e): "EL3HLT",
		syscall.Errno(0x2f): "EL3RST",
		syscall.Errno(0x4f): "ELIBACC",
		syscall.Errno(0x50): "ELIBBAD",
		syscall.Errno(0x53): "ELIBEXEC",
		syscall.Errno(0x52): "ELIBMAX",
		syscall.Errno(0x51): "ELIBSCN",
		syscall.Errno(0x30): "ELNRNG",
		syscall.Errno(0x28): "ELOOP",
		syscall.Errno(0x7c): "EMEDIUMTYPE",
		syscall.Errno(0x18): "EMFILE",
		syscall.Errno(0x1f): "EMLINK",
		syscall.Errno(0x5a): "EMSGSIZE",
		syscall.Errno(0x48): "EMULTIHOP",
		syscall.Errno(0x24): "ENAMETOOLONG",
		syscall.Errno(0x77): "ENAVAIL",
		syscall.Errno(0x64): "ENETDOWN",
		syscall.Errno(0x66): "ENETRESET",
		syscall.Errno(0x65): "ENETUNREACH",
		syscall.Errno(0x17): "ENFILE",
		syscall.Errno(0x37): "ENOANO",
		syscall.Errno(0x69): "ENOBUFS",
		syscall.Errno(0x32): "ENOCSI",
		syscall.Errno(0x3d): "ENODATA",
		syscall.Errno(0x13): "ENODEV",
		syscall.Errno(0x2):  "ENOENT",
		syscall.Errno(0x8):  "ENOEXEC",
		syscall.Errno(0x7e): "ENOKEY",
		syscall.Errno(0x25): "ENOLCK",
		syscall.Errno(0x43): "ENOLINK",
		syscall.Errno(0x7b): "ENOMEDIUM",
		syscall.Errno(0xc):  "ENOMEM",
		syscall.Errno(0x2a): "ENOMSG",
		syscall.Errno(0x40): "ENONET",
		syscall.Errno(0x41): "ENOPKG",
		syscall.Errno(0x5c): "ENOPROTOOPT",
		syscall.Errno(0x1c): "ENOSPC",
		syscall.Errno(0x3f): "ENOSR",
		syscall.Errno(0x3c): "ENOSTR",
		syscall.Errno(0x26): "ENOSYS",
		syscall.Errno(0xf):  "ENOTBLK",
		syscall.Errno(0x6b): "ENOTCONN",
		syscall.Errno(0x14): "ENOTDIR",
		syscall.Errno(0x27): "ENOTEMPTY",
		syscall.Errno(0x76): "ENOTNAM",
		syscall.Errno(0x83): "ENOTRECOVERABLE",
		syscall.Errno(0x58): "ENOTSOCK",
		syscall.Errno(0x5f): "ENOTSUP",
		syscall.Errno(0x19): "ENOTTY",
		syscall.Errno(0x4c): "ENOTUNIQ",
		syscall.Errno(0x6):  "ENXIO",
		syscall.Errno(0x4b): "EOVERFLOW",
		syscall.Errno(0x82): "EOWNERDEAD",
		syscall.Errno(0x1):  "EPERM",
		syscall.Errno(0x60): "EPFNOSUPPORT",
		syscall.Errno(0x20): "EPIPE",
		syscall.Errno(0x47): "EPROTO",
		syscall.Errno(0x5d): "EPROTONOSUPPORT",
		syscall.Errno(0x5b): "EPROTOTYPE",
		syscall.Errno(0x22): "ERANGE",
		syscall.Errno(0x4e): "EREMCHG",
		syscall.Errno(0x42): "EREMOTE",
		syscall.Errno(0x79): "EREMOTEIO",
		syscall.Errno(0x55): "ERESTART",
		syscall.Errno(0x84): "ERFKILL",
		syscall.Errno(0x1e): "EROFS",
		syscall.Errno(0x6c): "ESHUTDOWN",
		syscall.Errno(0x5e): "ESOCKTNOSUPPORT",
		syscall.Errno(0x1d): "ESPIPE",
		syscall.Errno(0x3):  "ESRCH",
		syscall.Errno(0x45): "ESRMNT",
		syscall.Errno(0x74): "ESTALE",
		syscall.Errno(0x56): "ESTRPIPE",
		syscall.Errno(0x3e): "ETIME",
		syscall.Errno(0x6e): "ETIMEDOUT",
		syscall.Errno(0x6d): "ETOOMANYREFS",
		syscall.Errno(0x1a): "ETXTBSY",
		syscall.Errno(0x75): "EUCLEAN",
		syscall.Errno(0x31): "EUNATCH",
		syscall.Errno(0x57): "EUSERS",
		syscall.Errno(0x12): "EXDEV",
		syscall.Errno(0x36): "EXFULL",
	}
}