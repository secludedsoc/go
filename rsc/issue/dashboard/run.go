// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"net/http"

	"trident.li/go/rsc/appfs/fs"
	"trident.li/go/rsc/issue/dashboard"
)

func main() {
	log.SetFlags(0)
	ctxt := fs.NewContext(new(http.Request))
	if err := dashboard.Update(ctxt, nil, "Go 1.2"); err != nil {
		log.Fatal(err)
	}
	log.Print("OK")
}
