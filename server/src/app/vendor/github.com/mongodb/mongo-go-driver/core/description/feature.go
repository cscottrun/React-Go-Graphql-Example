// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package description

import (
	"fmt"
)

// MaxStalenessSupported returns an error if the given server version
// does not support max staleness.
func MaxStalenessSupported(serverVersion Version, wireVersion *VersionRange) error {
	if !serverVersion.AtLeast(3, 4, 0) || (wireVersion != nil && wireVersion.Max < 5) {
		return fmt.Errorf("max staleness is only supported for servers 3.4 or newer")
	}

	return nil
}

// ScramSHA1Supported returns an error if the given server version
// does not support scram-sha-1.
func ScramSHA1Supported(version Version) error {
	if !version.AtLeast(3, 0, 0) {
		return fmt.Errorf("SCRAM-SHA-1 is only supported for servers 3.0 or newer")
	}

	return nil
}
