// Code generated by "stringer -type=OpType -output=op_type_string.go"; DO NOT EDIT.

package operation

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpTypeUnknown-0]
	_ = x[OpTypeGetTerraformVersion-1]
	_ = x[OpTypeObtainSchema-2]
	_ = x[OpTypeParseModuleConfiguration-3]
	_ = x[OpTypeParseVariables-4]
	_ = x[OpTypeParseModuleManifest-5]
	_ = x[OpTypeLoadModuleMetadata-6]
	_ = x[OpTypeDecodeReferences-7]
}

const _OpType_name = "OpTypeUnknownOpTypeGetTerraformVersionOpTypeObtainSchemaOpTypeParseModuleConfigurationOpTypeParseVariablesOpTypeParseModuleManifestOpTypeLoadModuleMetadataOpTypeDecodeReferences"

var _OpType_index = [...]uint8{0, 13, 38, 56, 86, 106, 131, 155, 177}

func (i OpType) String() string {
	if i >= OpType(len(_OpType_index)-1) {
		return "OpType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpType_name[_OpType_index[i]:_OpType_index[i+1]]
}
