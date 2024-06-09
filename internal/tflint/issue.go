// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tflint

import "github.com/hashicorp/hcl/v2"

type Issue struct {
	Rule    Rule    `json:"rule"`
	Message string  `json:"message"`
	Range   Range   `json:"range"`
	Callers []Range `json:"callers"`
}

type Rule struct {
	Name     string `json:"name"`
	Severity string `json:"severity"`
	Link     string `json:"link"`
}

type Range struct {
	Filename string `json:"filename"`
	Start    Pos    `json:"start"`
	End      Pos    `json:"end"`
}

type Pos struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Error struct {
	Summary  string `json:"summary,omitempty"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Range    *Range `json:"range,omitempty"`
}

type JSON struct {
	Issues []Issue `json:"issues"`
	Errors []Error `json:"errors"`
}

func (j *JSON) AsHCLDiagsMap() map[string]hcl.Diagnostics {
	diags := map[string]hcl.Diagnostics{}
	for _, issue := range j.Issues {
		diags[issue.Range.Filename] = append(diags[issue.Range.Filename], &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  issue.Message,
			Subject: &hcl.Range{
				Start: hcl.Pos{Line: issue.Range.Start.Line, Column: issue.Range.Start.Column},
				End:   hcl.Pos{Line: issue.Range.End.Line, Column: issue.Range.End.Column},
			},
		})
	}
	return diags
}
