package main

import (
	"fmt"
)

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerinfo owner
}

type electricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerinfo owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e gasEngine) giveName() string {
	return e.ownerinfo.name
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e electricEngine) giveName() string {
	return e.ownerinfo.name
}

type engine interface {
	milesLeft() uint8
	giveName() string
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Printf("You can make it %s!\n", e.giveName())
	} else {
		fmt.Printf("You need to fuel %s!\n", e.giveName())
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{10, 10}
	var myEngine3 electricEngine = electricEngine{25, 5, owner{"Chavis"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerinfo.name)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Printf("Miles for %s left: %v\n", myEngine.ownerinfo.name, myEngine.milesLeft())
	fmt.Printf("Miles for %s left: %v\n", myEngine3.ownerinfo.name, myEngine3.milesLeft())
	var miles2go uint8 = 120
	canMakeIt(myEngine, miles2go)
	canMakeIt(myEngine3, miles2go)
}
