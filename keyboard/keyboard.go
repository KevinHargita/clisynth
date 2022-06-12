package keyboard

import (
	"bytes"
	"fmt"
	"math"
	"syscall"
	"time"
)

type Keyboard struct {
	CenterFreq       float64
	halfstepRatioMap map[rune]float64
	PressedKeys      chan [4]float64
	ExitSig          bool
}

const halfstep float64 = 1.0595

var keys []rune = []rune("awsedftgyhujkolp;'")

func New(centerFreq float64) *Keyboard {
	halfstepRatioMap := make(map[rune]float64)
	for i, key := range keys {
		halfstepRatioMap[key] = math.Pow(halfstep, float64(i-9))
	}

	newKeyboard := &Keyboard{
		centerFreq,
		halfstepRatioMap,
		make(chan [4]float64),
		false,
	}

	go newKeyboard.listen()

	PrintKeyMap()

	return newKeyboard
}

func (k *Keyboard) listen() {
	fmt.Println("Keyboard Started...")
	user32DLL := syscall.NewLazyDLL("User32.dll")
	getKeyboardState := user32DLL.NewProc("GetAsyncKeyState")
	var pressedKeys string
	for {
		pressedKeysByteBuffer := bytes.NewBufferString("")

		//Check for escape press
		r1, _, err := getKeyboardState.Call(
			uintptr(0x1B))
		if err != syscall.Errno(0) {
			fmt.Println(err.Error())
			return
		}
		if r1 != 0 {
			fmt.Println("Keyboard Stopped...")
			k.ExitSig = true
			break
		}

		//Check for press of all character a-z
		for i := 0; i < 26; i++ {
			r2, _, err := getKeyboardState.Call(
				uintptr(i + 65))
			if err != syscall.Errno(0) {
				fmt.Println(err.Error())
			}

			if r2 != 0 {
				pressedKeysByteBuffer.WriteByte(byte(i + 97))
			}
		}

		currentPressedKeys := pressedKeysByteBuffer.String()
		if pressedKeys != currentPressedKeys {
			pressedKeys = currentPressedKeys
			var pressedKeysFreq [4]float64
			for i := 0; i < 4 && i < len(pressedKeys); i++ {
				pressedKeysFreq[i] = k.halfstepRatioMap[rune(pressedKeys[i])]
			}
			//fmt.Println(pressedKeysFreq)
			k.PressedKeys <- pressedKeysFreq
		}

		time.Sleep(time.Duration(time.Second / 50))
	}
}

func PrintKeyMap() {
	fmt.Println(
		`
		 _____________________________________________________________________________________
		 |     |     |  |     |     |     |     |  |     |  |     |     |     |     |  |     |
		 |     |     |  |     |     |     |     |  |     |  |     |     |     |     |  |     |
		 |     |  W  |  |  E  |     |     |  T  |  |  Y  |  |  U  |     |     |  O  |  |  P  |
		 |     |_____|  |_____|     |     |_____|  |_____|  |_____|     |     |_____|  |_____|
		 |        |        |        |        |        |        |        |        |        |
		 |   A    |   S    |   D    |   F    |   G    |   H    |   J    |   K    |   L    |
		 |________|________|________|________|________|________|________|________|________|
		
		 `)
}
