// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tflint

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl/v2"
)

type Executor struct {
	execPath   string
	configPath string
	timeout    time.Duration
	modPath    string
}

func TFLintExecutorForModule(ctx context.Context, modPath string) (*Executor, error) {
	execPath, err := ExecPath(ctx)
	if err != nil {
		return nil, err
	}

	timeout, err := Timeout(ctx)
	if err != nil {
		return nil, err
	}

	configPath, err := ConfigPath(ctx)
	if err != nil {
		return nil, err
	}

	return &Executor{
		execPath:   execPath,
		configPath: configPath,
		timeout:    timeout,
		modPath:    modPath,
	}, nil
}

func (e *Executor) Inspect(ctx context.Context) (map[string]hcl.Diagnostics, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	ctx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	args := []string{"--format", "json", "--force", "--no-color"}
	if e.configPath != "" {
		args = append(args, "--config", e.configPath)
	}
	cmd := exec.CommandContext(ctx, e.execPath, args...)
	cmd.Dir = e.modPath
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}
	cmd.WaitDelay = 3 * time.Second

	if err := cmd.Run(); err != nil {
		if stderr.String() != "" {
			return nil, fmt.Errorf("%w\n%s", err, stderr.String())
		}
	}

	var output JSON
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		return nil, err
	}
	var errs *multierror.Error
	for _, err := range output.Errors {
		errs = multierror.Append(errs, fmt.Errorf("%s; %s", err.Summary, err.Message))
	}

	return output.AsHCLDiagsMap(), errs.ErrorOrNil()
}
