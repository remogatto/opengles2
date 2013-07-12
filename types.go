/*
Copyright © 2013 mortdeus <mortdeus@gocos2d.org>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
“Software”), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package gles2

/*
#include <stdlib.h>
#include <GLES2/gl2.h>
#include <GLES2/gl2ext.h>
#include <GLES2/gl2platform.h>
*/
import "C"
import "unsafe"

type (
	Void     unsafe.Pointer
	Enum     uint32
	Bitfield uint32
	Sizei    int32
	Clampf   float32
	Fixed    uintptr
	IntPtr   int
	SizeiPtr int
)

func glString(s string) *C.GLchar {
	return (*C.GLchar)(C.CString(s))
}
func goString(s *C.GLchar) *string {
	gs := C.GoString((*C.char)(s))
	return &gs
}
func goBoolean(n C.GLboolean) *bool {
	b := n == 1
	return &b
}
func glBoolean(n bool) C.GLboolean {
	var b int
	if n == true {
		b = 1
	}
	return C.GLboolean(b)
}
