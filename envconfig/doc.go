// Copyright (c) 2013 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package envconfig implements decoding of environment variables based on a user
// defined specification. A typical use is using environment variables for
// configuration settings.
//
// In addition, this modified version of envconfig supports a few "extra" types,
// such as comma separated []strings, plus some app specific ones.
package envconfig
