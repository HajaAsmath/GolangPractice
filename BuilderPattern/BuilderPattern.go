package main

import (
	"fmt"
	"strings"
)

type shield struct {
	left  bool
	right bool
	front bool
	back  bool
}

type shieldBuilder struct {
	code string
}

func createShieldBuilder() *shieldBuilder {
	return new(shieldBuilder)
}

func (builder *shieldBuilder) raiseLeftShield() *shieldBuilder {
	builder.code += "L"
	return builder
}

func (builder *shieldBuilder) raiseRightShield() *shieldBuilder {
	builder.code += "R"
	return builder
}

func (builder *shieldBuilder) raiseFrontShield() *shieldBuilder {
	builder.code += "F"
	return builder
}

func (builder *shieldBuilder) raiseBackShield() *shieldBuilder {
	builder.code += "B"
	return builder
}

func (builder *shieldBuilder) build() *shield {
	code := builder.code

	return &shield{
		left:  strings.Contains(code, "L"),
		right: strings.Contains(code, "R"),
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
	}
}

func main() {
	shieldBuilder := createShieldBuilder()
	builder := shieldBuilder.raiseBackShield().raiseFrontShield().build()
	fmt.Printf("%+v", *builder)
}
