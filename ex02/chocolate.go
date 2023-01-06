package main

/*
	#cgo CFLAGS: -x objective-c -DGL_SILENCE_DEPRECATION
	#cgo LDFLAGS: -framework Cocoa -framework OpenGL
	#include <OpenGL/gl3.h>
	#import <Carbon/Carbon.h> // for HIToolbox/Events.h
	#import <Cocoa/Cocoa.h>
	#include <pthread.h>
	#include <stdint.h>
	#include <stdlib.h>
	void startDriver();
	void stopDriver();
	void makeCurrentContext(uintptr_t ctx);
	void flushContext(uintptr_t ctx);
	uintptr_t doNewWindow(int width, int height, char* title);
	void doShowWindow(uintptr_t id);
	void doCloseWindow(uintptr_t id);
	uint64_t threadID();
	#include "application.h"
	#include "window.h"
*/
import "C"

func main() {
	C.InitApplication()
	w := C.Window_Create(150, 150, 300, 200, C.CString("School 21"))
	C.Window_MakeKeyAndOrderFront(w)
	C.RunApplication()
}
