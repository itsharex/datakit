// Code generated by 'ccgo time/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o time/time_linux_amd64.go -pkgname time', DO NOT EDIT.

package time

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	CLOCK_BOOTTIME            = 7      // time.h:60:1:
	CLOCK_BOOTTIME_ALARM      = 9      // time.h:64:1:
	CLOCK_MONOTONIC           = 1      // time.h:48:1:
	CLOCK_MONOTONIC_COARSE    = 6      // time.h:58:1:
	CLOCK_MONOTONIC_RAW       = 4      // time.h:54:1:
	CLOCK_PROCESS_CPUTIME_ID  = 2      // time.h:50:1:
	CLOCK_REALTIME            = 0      // time.h:46:1:
	CLOCK_REALTIME_ALARM      = 8      // time.h:62:1:
	CLOCK_REALTIME_COARSE     = 5      // time.h:56:1:
	CLOCK_TAI                 = 11     // time.h:66:1:
	CLOCK_THREAD_CPUTIME_ID   = 3      // time.h:52:1:
	TIMER_ABSTIME             = 1      // time.h:69:1:
	TIME_UTC                  = 1      // time.h:65:1:
	X_ATFILE_SOURCE           = 1      // features.h:342:1:
	X_BITS_ENDIANNESS_H       = 1      // endianness.h:2:1:
	X_BITS_ENDIAN_H           = 1      // endian.h:20:1:
	X_BITS_TIME64_H           = 1      // time64.h:24:1:
	X_BITS_TIME_H             = 1      // time.h:24:1:
	X_BITS_TYPESIZES_H        = 1      // typesizes.h:24:1:
	X_BITS_TYPES_H            = 1      // types.h:24:1:
	X_BITS_TYPES_LOCALE_T_H   = 1      // locale_t.h:20:1:
	X_BITS_TYPES___LOCALE_T_H = 1      // __locale_t.h:21:1:
	X_BSD_SIZE_T_             = 0      // stddef.h:189:1:
	X_BSD_SIZE_T_DEFINED_     = 0      // stddef.h:192:1:
	X_DEFAULT_SOURCE          = 1      // features.h:227:1:
	X_FEATURES_H              = 1      // features.h:19:1:
	X_FILE_OFFSET_BITS        = 64     // <builtin>:25:1:
	X_GCC_SIZE_T              = 0      // stddef.h:195:1:
	X_LP64                    = 1      // <predefined>:284:1:
	X_POSIX_C_SOURCE          = 200809 // features.h:281:1:
	X_POSIX_SOURCE            = 1      // features.h:279:1:
	X_SIZET_                  = 0      // stddef.h:196:1:
	X_SIZE_T                  = 0      // stddef.h:183:1:
	X_SIZE_T_                 = 0      // stddef.h:188:1:
	X_SIZE_T_DECLARED         = 0      // stddef.h:193:1:
	X_SIZE_T_DEFINED          = 0      // stddef.h:191:1:
	X_SIZE_T_DEFINED_         = 0      // stddef.h:190:1:
	X_STDC_PREDEF_H           = 1      // <predefined>:162:1:
	X_STRUCT_TIMESPEC         = 1      // struct_timespec.h:3:1:
	X_SYS_CDEFS_H             = 1      // cdefs.h:19:1:
	X_SYS_SIZE_T_H            = 0      // stddef.h:184:1:
	X_TIME_H                  = 1      // time.h:23:1:
	X_T_SIZE                  = 0      // stddef.h:186:1:
	X_T_SIZE_                 = 0      // stddef.h:185:1:
	Linux                     = 1      // <predefined>:231:1:
	Unix                      = 1      // <predefined>:177:1:
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// Wide character type.
//    Locale-writers should change this as necessary to
//    be big enough to hold unique values not between 0 and 127,
//    and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

// A null pointer constant.

// This defines CLOCKS_PER_SEC, which is the number of processor clock
//    ticks per second, and possibly a number of other constants.
// System-dependent timing definitions.  Linux version.
//    Copyright (C) 1996-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <time.h> instead.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Copyright (C) 1991-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Determine the wordsize from the preprocessor defines.

// Both x86-64 and x32 use the 64-bit system call interface.
// Bit size of the time_t type at glibc build time, x86-64 and x32 case.
//    Copyright (C) 2018-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// For others, time size is word size.

// Convenience types.
type X__u_char = uint8   /* types.h:31:23 */
type X__u_short = uint16 /* types.h:32:28 */
type X__u_int = uint32   /* types.h:33:22 */
type X__u_long = uint64  /* types.h:34:27 */

// Fixed-size types, underlying types depend on word size and compiler.
type X__int8_t = int8     /* types.h:37:21 */
type X__uint8_t = uint8   /* types.h:38:23 */
type X__int16_t = int16   /* types.h:39:26 */
type X__uint16_t = uint16 /* types.h:40:28 */
type X__int32_t = int32   /* types.h:41:20 */
type X__uint32_t = uint32 /* types.h:42:22 */
type X__int64_t = int64   /* types.h:44:25 */
type X__uint64_t = uint64 /* types.h:45:27 */

// Smallest types with at least a given width.
type X__int_least8_t = X__int8_t     /* types.h:52:18 */
type X__uint_least8_t = X__uint8_t   /* types.h:53:19 */
type X__int_least16_t = X__int16_t   /* types.h:54:19 */
type X__uint_least16_t = X__uint16_t /* types.h:55:20 */
type X__int_least32_t = X__int32_t   /* types.h:56:19 */
type X__uint_least32_t = X__uint32_t /* types.h:57:20 */
type X__int_least64_t = X__int64_t   /* types.h:58:19 */
type X__uint_least64_t = X__uint64_t /* types.h:59:20 */

// quad_t is also 64 bits.
type X__quad_t = int64    /* types.h:63:18 */
type X__u_quad_t = uint64 /* types.h:64:27 */

// Largest integral types.
type X__intmax_t = int64   /* types.h:72:18 */
type X__uintmax_t = uint64 /* types.h:73:27 */

// The machine-dependent file <bits/typesizes.h> defines __*_T_TYPE
//    macros for each of the OS types we define below.  The definitions
//    of those macros must use the following macros for underlying types.
//    We define __S<SIZE>_TYPE and __U<SIZE>_TYPE for the signed and unsigned
//    variants of each of the following integer types on this machine.
//
// 	16		-- "natural" 16-bit type (always short)
// 	32		-- "natural" 32-bit type (always int)
// 	64		-- "natural" 64-bit type (long or long long)
// 	LONG32		-- 32-bit type, traditionally long
// 	QUAD		-- 64-bit type, traditionally long long
// 	WORD		-- natural type of __WORDSIZE bits (int or long)
// 	LONGWORD	-- type of __WORDSIZE bits, traditionally long
//
//    We distinguish WORD/LONGWORD, 32/LONG32, and 64/QUAD so that the
//    conventional uses of `long' or `long long' type modifiers match the
//    types we define, even when a less-adorned type would be the same size.
//    This matters for (somewhat) portably writing printf/scanf formats for
//    these types, where using the appropriate l or ll format modifiers can
//    make the typedefs and the formats match up across all GNU platforms.  If
//    we used `long' when it's 64 bits where `long long' is expected, then the
//    compiler would warn about the formats not matching the argument types,
//    and the programmer changing them to shut up the compiler would break the
//    program's portability.
//
//    Here we assume what is presently the case in all the GCC configurations
//    we support: long long is always 64 bits, long is always word/address size,
//    and int is always 32 bits.

// No need to mark the typedef with __extension__.
// bits/typesizes.h -- underlying types for *_t.  Linux/x86-64 version.
//    Copyright (C) 2012-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// See <bits/types.h> for the meaning of these macros.  This file exists so
//    that <bits/types.h> need not vary across different GNU platforms.

// X32 kernel interface is 64-bit.

// Tell the libc code that off_t and off64_t are actually the same type
//    for all ABI purposes, even if possibly expressed as different base types
//    for C type-checking purposes.

// Same for ino_t and ino64_t.

// And for __rlim_t and __rlim64_t.

// And for fsblkcnt_t, fsblkcnt64_t, fsfilcnt_t and fsfilcnt64_t.

// Number of descriptors that can fit in an `fd_set'.

// bits/time64.h -- underlying types for __time64_t.  Generic version.
//    Copyright (C) 2018-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Define __TIME64_T_TYPE so that it is always a 64-bit type.

// If we already have 64-bit time type then use it.

type X__dev_t = uint64                     /* types.h:145:25 */ // Type of device numbers.
type X__uid_t = uint32                     /* types.h:146:25 */ // Type of user identifications.
type X__gid_t = uint32                     /* types.h:147:25 */ // Type of group identifications.
type X__ino_t = uint64                     /* types.h:148:25 */ // Type of file serial numbers.
type X__ino64_t = uint64                   /* types.h:149:27 */ // Type of file serial numbers (LFS).
type X__mode_t = uint32                    /* types.h:150:26 */ // Type of file attribute bitmasks.
type X__nlink_t = uint64                   /* types.h:151:27 */ // Type of file link counts.
type X__off_t = int64                      /* types.h:152:25 */ // Type of file sizes and offsets.
type X__off64_t = int64                    /* types.h:153:27 */ // Type of file sizes and offsets (LFS).
type X__pid_t = int32                      /* types.h:154:25 */ // Type of process identifications.
type X__fsid_t = struct{ F__val [2]int32 } /* types.h:155:26 */ // Type of file system IDs.
type X__clock_t = int64                    /* types.h:156:27 */ // Type of CPU usage counts.
type X__rlim_t = uint64                    /* types.h:157:26 */ // Type for resource measurement.
type X__rlim64_t = uint64                  /* types.h:158:28 */ // Type for resource measurement (LFS).
type X__id_t = uint32                      /* types.h:159:24 */ // General type for IDs.
type X__time_t = int64                     /* types.h:160:26 */ // Seconds since the Epoch.
type X__useconds_t = uint32                /* types.h:161:30 */ // Count of microseconds.
type X__suseconds_t = int64                /* types.h:162:31 */ // Signed count of microseconds.

type X__daddr_t = int32 /* types.h:164:27 */ // The type of a disk address.
type X__key_t = int32   /* types.h:165:25 */ // Type of an IPC key.

// Clock ID used in clock and timer functions.
type X__clockid_t = int32 /* types.h:168:29 */

// Timer ID returned by `timer_create'.
type X__timer_t = uintptr /* types.h:171:12 */

// Type to represent block size.
type X__blksize_t = int64 /* types.h:174:29 */

// Types from the Large File Support interface.

// Type to count number of disk blocks.
type X__blkcnt_t = int64   /* types.h:179:28 */
type X__blkcnt64_t = int64 /* types.h:180:30 */

// Type to count file system blocks.
type X__fsblkcnt_t = uint64   /* types.h:183:30 */
type X__fsblkcnt64_t = uint64 /* types.h:184:32 */

// Type to count file system nodes.
type X__fsfilcnt_t = uint64   /* types.h:187:30 */
type X__fsfilcnt64_t = uint64 /* types.h:188:32 */

// Type of miscellaneous file system fields.
type X__fsword_t = int64 /* types.h:191:28 */

type X__ssize_t = int64 /* types.h:193:27 */ // Type of a byte count, or error.

// Signed long type used in system calls.
type X__syscall_slong_t = int64 /* types.h:196:33 */
// Unsigned long type used in system calls.
type X__syscall_ulong_t = uint64 /* types.h:198:33 */

// These few don't really vary by system, they always correspond
//
//	to one of the other defined types.
type X__loff_t = X__off64_t /* types.h:202:19 */ // Type of file sizes and offsets (LFS).
type X__caddr_t = uintptr   /* types.h:203:14 */

// Duplicates info from stdint.h but this is used in unistd.h.
type X__intptr_t = int64 /* types.h:206:25 */

// Duplicate info from sys/socket.h.
type X__socklen_t = uint32 /* types.h:209:23 */

// C99: An integer type that can be accessed as an atomic entity,
//
//	even in the presence of asynchronous interrupts.
//	It is not currently necessary for this to be machine-specific.
type X__sig_atomic_t = int32 /* types.h:214:13 */

// Seconds since the Epoch, visible to user code when time_t is too
//    narrow only for consistency with the old way of widening too-narrow
//    types.  User code should never use __time64_t.

// ISO/IEC 9899:1999 7.23.1: Components of time
//    The macro `CLOCKS_PER_SEC' is an expression with type `clock_t' that is
//    the number per second of the value returned by the `clock' function.
// CAE XSH, Issue 4, Version 2: <time.h>
//    The value of CLOCKS_PER_SEC is required to be 1 million on all
//    XSI-conformant systems.

// Identifier for system-wide realtime clock.
// Monotonic system-wide clock.
// High-resolution timer from the CPU.
// Thread-specific CPU-time clock.
// Monotonic system-wide clock, not adjusted for frequency scaling.
// Identifier for system-wide realtime clock, updated only on ticks.
// Monotonic system-wide clock, updated only on ticks.
// Monotonic system-wide clock that includes time spent in suspension.
// Like CLOCK_REALTIME but also wakes suspended system.
// Like CLOCK_BOOTTIME but also wakes suspended system.
// Like CLOCK_REALTIME but in International Atomic Time.

// Flag to indicate time is absolute.

// Many of the typedefs and structs whose official home is this header
//    may also need to be defined by other headers.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Returned by `clock'.
type Clock_t = X__clock_t /* clock_t.h:7:19 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Returned by `time'.
type Time_t = X__time_t /* time_t.h:7:18 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// ISO C `broken-down time' structure.
type Tm = struct {
	Ftm_sec      int32
	Ftm_min      int32
	Ftm_hour     int32
	Ftm_mday     int32
	Ftm_mon      int32
	Ftm_year     int32
	Ftm_wday     int32
	Ftm_yday     int32
	Ftm_isdst    int32
	F__ccgo_pad1 [4]byte
	Ftm_gmtoff   int64
	Ftm_zone     uintptr
} /* struct_tm.h:7:1 */

// NB: Include guard matches what <linux/time.h> uses.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Endian macros for string.h functions
//    Copyright (C) 1992-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Definitions for byte order, according to significance of bytes,
//    from low addresses to high addresses.  The value is what you get by
//    putting '4' in the most significant byte, '3' in the second most
//    significant byte, '2' in the second least significant byte, and '1'
//    in the least significant byte, and then writing down one digit for
//    each byte, starting with the byte at the lowest address at the left,
//    and proceeding to the byte with the highest address at the right.

// This file defines `__BYTE_ORDER' for the particular machine.

// i386/x86_64 are little-endian.

// Some machines may need to use a different endianness for floating point
//    values.

// POSIX.1b structure for a time value.  This is like a `struct timeval' but
//
//	has nanoseconds instead of microseconds.
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec X__syscall_slong_t
} /* struct_timespec.h:10:1 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Clock ID used in clock and timer functions.
type Clockid_t = X__clockid_t /* clockid_t.h:7:21 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Timer ID returned by `timer_create'.
type Timer_t = X__timer_t /* timer_t.h:7:19 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// NB: Include guard matches what <linux/time.h> uses.

// POSIX.1b structure for timer start values and intervals.
type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
	Fit_value struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
} /* struct_itimerspec.h:8:1 */

type Pid_t = X__pid_t /* time.h:54:17 */

// Definition of locale_t.
//    Copyright (C) 2017-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Definition of struct __locale_struct and __locale_t.
//    Copyright (C) 1997-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//    Contributed by Ulrich Drepper <drepper@cygnus.com>, 1997.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// POSIX.1-2008: the locale_t type, representing a locale context
//    (implementation-namespace version).  This type should be treated
//    as opaque by applications; some details are exposed for the sake of
//    efficiency in e.g. ctype functions.

type X__locale_struct = struct {
	F__locales       [13]uintptr
	F__ctype_b       uintptr
	F__ctype_tolower uintptr
	F__ctype_toupper uintptr
	F__names         [13]uintptr
} /* __locale_t.h:28:1 */

type X__locale_t = uintptr /* __locale_t.h:42:32 */

type Locale_t = X__locale_t /* locale_t.h:24:20 */

var _ int8 /* gen.c:2:13: */
