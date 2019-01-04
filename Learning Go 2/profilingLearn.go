package main

import "github.com/pkg/profile"

func main() {

	defer profile.Start().Stop()

}
