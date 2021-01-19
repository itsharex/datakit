/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type MemoryInformationForVm struct {
	VirtualNodeCount uint32 `json:"VirtualNodeCount,omitempty"`

	VirtualMachineMemory *VmMemory `json:"VirtualMachineMemory,omitempty"`

	VirtualNodes []VirtualNodeInfo `json:"VirtualNodes,omitempty"`
}
