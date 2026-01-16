package keyboard

import (	
	"unicode"
)

type Keyboard struct {
	Keys map[int]int
}

func NewKeyboard(dure int) *Keyboard {
	dict := make(map[int]int)
	for i := 0; i <= 256; i++{
		dict[i] = dure
	}
	return &Keyboard{Keys:dict}
}

func (keyboard *Keyboard) Enter(inp string) string {
	var output string
	for _, r := range inp{
		t:= false  
		if r>= 'A' && r <= 'Z' && keyboard.Keys[256] > 0 {
			keyboard.Keys[256]-=1
			t = true
		}

		r=unicode.ToLower(r)
		assci := int(r)
		
		if keyboard.Keys[assci] > 0 {
			keyboard.Keys[assci] = keyboard.Keys[assci]-1
			if t == true{
				r-=32
			}
			output+=string(r)			
		} 
	}
	return output
}