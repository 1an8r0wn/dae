/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2022-2025, daeuniverse Organization <dae@v2raya.org>
 */

package config

import (
	"testing"
)

func TestExportOutline(t *testing.T) {
	t.Log(ExportOutlineJson("test"))
}
