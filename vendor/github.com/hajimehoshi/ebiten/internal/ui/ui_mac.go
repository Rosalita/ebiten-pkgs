// Copyright 2016 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build darwin
// +build !js
// +build !ios

package ui

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
//
// #import <AppKit/AppKit.h>
//
// static void currentMonitorPos(void* windowPtr, int* x, int* y) {
//   NSScreen* screen = [NSScreen mainScreen];
//   if (windowPtr) {
//     NSWindow* window = (NSWindow*)windowPtr;
//     if ([window isVisible]) {
//       // When the window is visible, the window is already initialized.
//       // [NSScreen mainScreen] sometimes tells a lie when the window is put across monitors (#703).
//       screen = [window screen];
//     }
//   }
//   NSDictionary* screenDictionary = [screen deviceDescription];
//   NSNumber* screenID = [screenDictionary objectForKey:@"NSScreenNumber"];
//   CGDirectDisplayID aID = [screenID unsignedIntValue];
//   const CGRect bounds = CGDisplayBounds(aID);
//   *x = bounds.origin.x;
//   *y = bounds.origin.y;
// }
import "C"

import (
	"unsafe"

	"github.com/hajimehoshi/ebiten/internal/glfw"
)

func glfwScale() float64 {
	return 1
}

func adjustWindowPosition(x, y int) (int, int) {
	return x, y
}

func (u *userInterface) currentMonitorFromPosition() *glfw.Monitor {
	x := C.int(0)
	y := C.int(0)
	// Note: [NSApp mainWindow] is nil when it doesn't have its border. Use u.window here.
	win := unsafe.Pointer(u.window.GetCocoaWindow())
	C.currentMonitorPos(win, &x, &y)
	for _, m := range glfw.GetMonitors() {
		mx, my := m.GetPos()
		if int(x) == mx && int(y) == my {
			return m
		}
	}
	return glfw.GetPrimaryMonitor()
}

func (u *userInterface) nativeWindow() uintptr {
	return u.window.GetCocoaWindow()
}
