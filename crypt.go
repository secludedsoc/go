// Copyright 2013, Jonas mg
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// Package crypt provides interface for password crypt functions and collects
// common constants.
package crypt

import (
	"github.com/tridentli/osutil-crypt/apr1_crypt"
	"github.com/tridentli/osutil-crypt/common"
	"github.com/tridentli/osutil-crypt/md5_crypt"
	"github.com/tridentli/osutil-crypt/sha256_crypt"
	"github.com/tridentli/osutil-crypt/sha512_crypt"
)

func init() {
	crypt.RegisterCrypt(crypt.APR1, apr1_crypt.New, apr1_crypt.MagicPrefix)
	crypt.RegisterCrypt(crypt.MD5, md5_crypt.New, md5_crypt.MagicPrefix)
	crypt.RegisterCrypt(crypt.SHA256, sha256_crypt.New, sha256_crypt.MagicPrefix)
	crypt.RegisterCrypt(crypt.SHA512, sha512_crypt.New, sha512_crypt.MagicPrefix)
}

func NewFromHash(hashedKey string) (crypt.Crypter, error) {
	return crypt.NewFromHash(hashedKey)
}
