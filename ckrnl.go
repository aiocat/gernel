// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"errors"
	"syscall"
	"unsafe"
)

type Krnl struct {
	LibraryLoaded      bool
	Library            syscall.Handle
	InjectFunction     uintptr
	PipeActiveFunction uintptr
	SendToPipeFunction uintptr
}

func NewKrnl() *Krnl {
	return &Krnl{
		LibraryLoaded: false,
		Library:       syscall.InvalidHandle,
	}
}

func (krnl *Krnl) Load() error {
	lib, err := syscall.LoadLibrary(".\\ckrnl\\libckrnl.dll")
	if err != nil {
		krnl.LibraryLoaded = false
		return err
	}

	krnl.Library = lib
	krnl.LibraryLoaded = true

	krnl.InjectFunction, _ = syscall.GetProcAddress(krnl.Library, "Inject")
	krnl.PipeActiveFunction, _ = syscall.GetProcAddress(krnl.Library, "PipeActive")
	krnl.SendToPipeFunction, _ = syscall.GetProcAddress(krnl.Library, "SendToPipe")

	return nil
}

func (krnl *Krnl) Inject() error {
	if !krnl.LibraryLoaded {
		return errors.New("load the library first")
	}

	syscall.SyscallN(uintptr(krnl.InjectFunction), 0)
	return nil
}

func (krnl *Krnl) IsInjected() bool {
	if !krnl.LibraryLoaded {
		return false
	}

	value, _, _ := syscall.SyscallN(uintptr(krnl.PipeActiveFunction), 0)
	return value == 1
}

func (krnl *Krnl) Execute(lua string) bool {
	if !krnl.LibraryLoaded {
		return false
	}

	bytePtr, _ := syscall.BytePtrFromString(lua)
	unsafePtr := unsafe.Pointer(bytePtr)
	value, _, _ := syscall.SyscallN(uintptr(krnl.SendToPipeFunction), uintptr(unsafePtr))
	return value == 1
}

func (krnl *Krnl) Free() {
	syscall.FreeLibrary(krnl.Library)
}
