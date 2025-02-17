//package main

/*
#cgo LDFLAGS: -lstdc++ -lm
#cgo CFLAGS: -I.
#include "keycodes.h"
*/

//import "C"
/*
import (
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"




	"golang.org/x/term"
)

func FindKeyboardDevice() string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 255; i++ {
		buff, err := os.ReadFile(fmt.Sprintf(path, i))
		if err != nil {
			continue
		}

		deviceName := strings.ToLower(string(buff))

		fmt.Printf("%s, %s\n", deviceName, fmt.Sprintf(resolved, i))
	}

	return ""
}

func main() {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	file, err := os.Open("/dev/input/event3")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	b := make([]byte, 24)
	//FindKeyboardDevice()

	fmt.Print(C.KEY_NUMERIC_4)

	// Keymap derived from here: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
	keyMap := map[int]int{
		30: 0,  // KEY_A
		48: 1,  // KEY_B
		46: 2,  // KEY_C
		32: 3,  // KEY_D
		18: 4,  // KEY_E
		33: 5,  // KEY_F
		34: 6,  // KEY_G
		35: 7,  // KEY_H
		23: 8,  // KEY_I
		36: 9,  // KEY_J
		37: 10, // KEY_K
		38: 11, // KEY_L
		50: 12, // KEY_M
		49: 13, // KEY_N
		24: 14, // KEY_O
		25: 15, // KEY_P
		16: 16, // KEY_Q
		19: 17, // KEY_R
		31: 18, // KEY_S
		20: 19, // KEY_T
		22: 20, // KEY_U
		47: 21, // KEY_V
		17: 22, // KEY_W
		45: 23, // KEY_X
		21: 24, // KEY_Y
		44: 25, // KEY_Z
	}
	for {
		file.Read(b)
		//fmt.Printf("%b\n", b)

		//sec := binary.LittleEndian.Uint64(b[0:8])
		//usec := binary.LittleEndian.Uint64(b[8:16])
		//t := time.Unix(int64(sec), int64(usec)*1000)
		//fmt.Println(t)

		//var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		//fmt.Printf("CODE: %d TYPE: %d\n\r", code, typ)
		val := 'a' + keyMap[int(code)]
		if string(val) == "w" {
			fmt.Printf("String: %s TYPE: %d TIME: %s\r\n", string(val), typ, time.Now())
		}

		//binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		//fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, code, value)
	}
}*/

package main

import (
	"bytes"
	"fmt"
)

type Player struct {
	x, y   int
	sprite rune
}

func generateBG() [45 * 95 * 2]byte {
	const width = 95 * 2
	const height = 45

	var bg [width * height]byte
	for i := 0; i < len(bg); i += 2 {
		bg[i] = '#'
		bg[i+1] = ' '
	}
	fmt.Println(len(bg))
	emptySpace := bytes.Repeat([]byte{' ', ' '}, 40)
	offset := (width - len(emptySpace)) / 2
	for i := offset + (2 * width); i <= len(bg)-2*width; i += width {
		copy(bg[i:], emptySpace)
	}

	return bg
}

func arrayBytesToStr(arr [45 * 95 * 2]byte) string {
	var str string
	for i := 0; i < len(arr); i++ {
		if i%(95*2) == 0 {
			str += "\r\n"
		}
		str += string(arr[i])
	}
	return str
}

func print(buffer []byte, str []byte) []byte {
	copy(buffer, str)
	return buffer
}

func main() {
	var x [10]rune
	copy(x[:], y[:])
	fmt.Print("Hi\x7F")
}
