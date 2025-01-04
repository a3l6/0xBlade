package main

/*
#cgo LDFLAGS: -lstdc++ -lm
#cgo CFLAGS: -I.
#include "keycodes.h"
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"
	/*"os/signal"
	"strings"
	"syscall"

	"golang.org/x/term"*/)

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

	/*oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
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
	*/
	file, err := os.Open("/dev/input/event3")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	b := make([]byte, 24)
	//FindKeyboardDevice()

	fmt.Print(C.KEY_NUMERIC_4)

	for {
		file.Read(b)
		fmt.Printf("%b\n", b)

		sec := binary.LittleEndian.Uint64(b[0:8])
		usec := binary.LittleEndian.Uint64(b[8:16])
		t := time.Unix(int64(sec), int64(usec)*1000)
		fmt.Println(t)

		var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])

		if code == 16 {
			fmt.Print("DELETE")
			os.Exit(0)
		}
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, code, value)
	}

}
