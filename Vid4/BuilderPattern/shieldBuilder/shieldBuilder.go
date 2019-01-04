package shieldBuilder

import (
	"strings"
	"fmt"
)

type Shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type ShieldBuilder struct {
	code string
}

//Constructor for Shield builder
func NewShieldBuilder() *ShieldBuilder {
	return new(ShieldBuilder)
}

func (sh *ShieldBuilder) RaiseFront() *ShieldBuilder {
	sh.code += "F"
	return sh
}

func (sh *ShieldBuilder) RaiseBack() *ShieldBuilder {
	sh.code += "B"
	return sh
}

func (sh *ShieldBuilder) RaiseLeft() *ShieldBuilder {
	sh.code += "L"
	return sh
}

func (sh *ShieldBuilder) RaiseRight() *ShieldBuilder {
	sh.code += "R"
	return sh
}

func (sh *ShieldBuilder) Build() *Shield {
	code := sh.code
	return &Shield{
		front: strings.Contains(code, "F"),
		back: strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left: strings.Contains(code, "L"),
	}
}


//djfjsjdsksdnsdnfj
func (sh *Shield) UseShield() {

	if sh.left || sh.right || sh.front || sh.back {
		fmt.Println("Using Shield now !")
	}else {
		fmt.Println("No Shield available !")
	}
}