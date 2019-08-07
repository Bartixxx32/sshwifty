package static_pages

// This file is part of Sshwifty Project
//
// Copyright (C) 2019 Rui NI (nirui@gmx.com)
//
// https://github.com/niruix/sshwifty
//
// This file is generated at Wed, 07 Aug 2019 15:54:10 CST
// by "go generate", DO NOT EDIT! Also, do not open this file, it maybe too large
// for your editor. You've been warned.
//
// This file may contain third-party binaries. See DEPENDENCES for detail.

import (
	"time"
	"encoding/hex"
)


// Static34 returns static file
func Static34() (
	int,        // FileStart
	int,        // FileEnd
	int,        // CompressedStart
	int,        // CompressedEnd
	string,     // ContentHash
	string,     // CompressedHash
	time.Time,  // Time of creation
	[]byte,     // Data
) {
	created, createErr := time.Parse(
		time.RFC1123, "Wed, 07 Aug 2019 15:54:10 CST")

	if createErr != nil {
		panic(createErr)
	}

	data, dataErr := hex.DecodeString(rawStatic34Data)
	
	rawStatic34Data = ""

	if dataErr != nil {
		panic(dataErr)
	}

	return 0, 601867, 
		601867, 1017175,
		"Tis6I5ztTXw=", "wYn+4mGWsdc=", created, data
}