// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tflint

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	lsctx "github.com/hashicorp/terraform-ls/internal/context"
)

const defaultExecTimeout = 5 * time.Second

func ExecPath(ctx context.Context) (string, error) {
	linterOptions, err := lsctx.LinterOptions(ctx)
	if err != nil {
		return "", err
	}

	if linterOptions.TFLint.Path != "" {
		return linterOptions.TFLint.Path, nil
	}

	path, err := exec.LookPath(defaultExecutableName)
	if err != nil {
		return "", fmt.Errorf("unable to find tflint: %s", err)
	}
	return path, nil
}

func Timeout(ctx context.Context) (time.Duration, error) {
	linterOptions, err := lsctx.LinterOptions(ctx)
	if err != nil {
		return 0, err
	}

	if linterOptions.TFLint.Timeout != "" {
		return time.ParseDuration(linterOptions.TFLint.Timeout)
	}

	return defaultExecTimeout, nil
}

func ConfigPath(ctx context.Context) (string, error) {
	linterOptions, err := lsctx.LinterOptions(ctx)
	return linterOptions.TFLint.ConfigPath, err
}
