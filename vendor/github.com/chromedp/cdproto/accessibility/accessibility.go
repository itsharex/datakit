// Package accessibility provides the Chrome DevTools Protocol
// commands, types, and events for the Accessibility domain.
//
// Generated by the cdproto-gen command.
package accessibility

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
)

// DisableParams disables the accessibility domain.
type DisableParams struct{}

// Disable disables the accessibility domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Accessibility.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enables the accessibility domain which causes AXNodeIds to
// remain consistent between method calls. This turns on accessibility for the
// page, which can impact performance until accessibility is disabled.
type EnableParams struct{}

// Enable enables the accessibility domain which causes AXNodeIds to remain
// consistent between method calls. This turns on accessibility for the page,
// which can impact performance until accessibility is disabled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Accessibility.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// GetPartialAXTreeParams fetches the accessibility node and partial
// accessibility tree for this DOM node, if it exists.
type GetPartialAXTreeParams struct {
	NodeID         cdp.NodeID             `json:"nodeId,omitempty"`         // Identifier of the node to get the partial accessibility tree for.
	BackendNodeID  cdp.BackendNodeID      `json:"backendNodeId,omitempty"`  // Identifier of the backend node to get the partial accessibility tree for.
	ObjectID       runtime.RemoteObjectID `json:"objectId,omitempty"`       // JavaScript object id of the node wrapper to get the partial accessibility tree for.
	FetchRelatives bool                   `json:"fetchRelatives,omitempty"` // Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
}

// GetPartialAXTree fetches the accessibility node and partial accessibility
// tree for this DOM node, if it exists.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-getPartialAXTree
//
// parameters:
func GetPartialAXTree() *GetPartialAXTreeParams {
	return &GetPartialAXTreeParams{}
}

// WithNodeID identifier of the node to get the partial accessibility tree
// for.
func (p GetPartialAXTreeParams) WithNodeID(nodeID cdp.NodeID) *GetPartialAXTreeParams {
	p.NodeID = nodeID
	return &p
}

// WithBackendNodeID identifier of the backend node to get the partial
// accessibility tree for.
func (p GetPartialAXTreeParams) WithBackendNodeID(backendNodeID cdp.BackendNodeID) *GetPartialAXTreeParams {
	p.BackendNodeID = backendNodeID
	return &p
}

// WithObjectID JavaScript object id of the node wrapper to get the partial
// accessibility tree for.
func (p GetPartialAXTreeParams) WithObjectID(objectID runtime.RemoteObjectID) *GetPartialAXTreeParams {
	p.ObjectID = objectID
	return &p
}

// WithFetchRelatives whether to fetch this nodes ancestors, siblings and
// children. Defaults to true.
func (p GetPartialAXTreeParams) WithFetchRelatives(fetchRelatives bool) *GetPartialAXTreeParams {
	p.FetchRelatives = fetchRelatives
	return &p
}

// GetPartialAXTreeReturns return values.
type GetPartialAXTreeReturns struct {
	Nodes []*Node `json:"nodes,omitempty"` // The Accessibility.AXNode for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
}

// Do executes Accessibility.getPartialAXTree against the provided context.
//
// returns:
//
//	nodes - The Accessibility.AXNode for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
func (p *GetPartialAXTreeParams) Do(ctx context.Context) (nodes []*Node, err error) {
	// execute
	var res GetPartialAXTreeReturns
	err = cdp.Execute(ctx, CommandGetPartialAXTree, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// GetFullAXTreeParams fetches the entire accessibility tree for the root
// Document.
type GetFullAXTreeParams struct {
	Depth   int64       `json:"depth,omitempty"`   // The maximum depth at which descendants of the root node should be retrieved. If omitted, the full tree is returned.
	FrameID cdp.FrameID `json:"frameId,omitempty"` // The frame for whose document the AX tree should be retrieved. If omitted, the root frame is used.
}

// GetFullAXTree fetches the entire accessibility tree for the root Document.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-getFullAXTree
//
// parameters:
func GetFullAXTree() *GetFullAXTreeParams {
	return &GetFullAXTreeParams{}
}

// WithDepth the maximum depth at which descendants of the root node should
// be retrieved. If omitted, the full tree is returned.
func (p GetFullAXTreeParams) WithDepth(depth int64) *GetFullAXTreeParams {
	p.Depth = depth
	return &p
}

// WithFrameID the frame for whose document the AX tree should be retrieved.
// If omitted, the root frame is used.
func (p GetFullAXTreeParams) WithFrameID(frameID cdp.FrameID) *GetFullAXTreeParams {
	p.FrameID = frameID
	return &p
}

// GetFullAXTreeReturns return values.
type GetFullAXTreeReturns struct {
	Nodes []*Node `json:"nodes,omitempty"`
}

// Do executes Accessibility.getFullAXTree against the provided context.
//
// returns:
//
//	nodes
func (p *GetFullAXTreeParams) Do(ctx context.Context) (nodes []*Node, err error) {
	// execute
	var res GetFullAXTreeReturns
	err = cdp.Execute(ctx, CommandGetFullAXTree, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// GetRootAXNodeParams fetches the root node. Requires enable() to have been
// called previously.
type GetRootAXNodeParams struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // The frame in whose document the node resides. If omitted, the root frame is used.
}

// GetRootAXNode fetches the root node. Requires enable() to have been called
// previously.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-getRootAXNode
//
// parameters:
func GetRootAXNode() *GetRootAXNodeParams {
	return &GetRootAXNodeParams{}
}

// WithFrameID the frame in whose document the node resides. If omitted, the
// root frame is used.
func (p GetRootAXNodeParams) WithFrameID(frameID cdp.FrameID) *GetRootAXNodeParams {
	p.FrameID = frameID
	return &p
}

// GetRootAXNodeReturns return values.
type GetRootAXNodeReturns struct {
	Node *Node `json:"node,omitempty"`
}

// Do executes Accessibility.getRootAXNode against the provided context.
//
// returns:
//
//	node
func (p *GetRootAXNodeParams) Do(ctx context.Context) (node *Node, err error) {
	// execute
	var res GetRootAXNodeReturns
	err = cdp.Execute(ctx, CommandGetRootAXNode, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Node, nil
}

// GetAXNodeAndAncestorsParams fetches a node and all ancestors up to and
// including the root. Requires enable() to have been called previously.
type GetAXNodeAndAncestorsParams struct {
	NodeID        cdp.NodeID             `json:"nodeId,omitempty"`        // Identifier of the node to get.
	BackendNodeID cdp.BackendNodeID      `json:"backendNodeId,omitempty"` // Identifier of the backend node to get.
	ObjectID      runtime.RemoteObjectID `json:"objectId,omitempty"`      // JavaScript object id of the node wrapper to get.
}

// GetAXNodeAndAncestors fetches a node and all ancestors up to and including
// the root. Requires enable() to have been called previously.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-getAXNodeAndAncestors
//
// parameters:
func GetAXNodeAndAncestors() *GetAXNodeAndAncestorsParams {
	return &GetAXNodeAndAncestorsParams{}
}

// WithNodeID identifier of the node to get.
func (p GetAXNodeAndAncestorsParams) WithNodeID(nodeID cdp.NodeID) *GetAXNodeAndAncestorsParams {
	p.NodeID = nodeID
	return &p
}

// WithBackendNodeID identifier of the backend node to get.
func (p GetAXNodeAndAncestorsParams) WithBackendNodeID(backendNodeID cdp.BackendNodeID) *GetAXNodeAndAncestorsParams {
	p.BackendNodeID = backendNodeID
	return &p
}

// WithObjectID JavaScript object id of the node wrapper to get.
func (p GetAXNodeAndAncestorsParams) WithObjectID(objectID runtime.RemoteObjectID) *GetAXNodeAndAncestorsParams {
	p.ObjectID = objectID
	return &p
}

// GetAXNodeAndAncestorsReturns return values.
type GetAXNodeAndAncestorsReturns struct {
	Nodes []*Node `json:"nodes,omitempty"`
}

// Do executes Accessibility.getAXNodeAndAncestors against the provided context.
//
// returns:
//
//	nodes
func (p *GetAXNodeAndAncestorsParams) Do(ctx context.Context) (nodes []*Node, err error) {
	// execute
	var res GetAXNodeAndAncestorsReturns
	err = cdp.Execute(ctx, CommandGetAXNodeAndAncestors, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// GetChildAXNodesParams fetches a particular accessibility node by AXNodeId.
// Requires enable() to have been called previously.
type GetChildAXNodesParams struct {
	ID      NodeID      `json:"id"`
	FrameID cdp.FrameID `json:"frameId,omitempty"` // The frame in whose document the node resides. If omitted, the root frame is used.
}

// GetChildAXNodes fetches a particular accessibility node by AXNodeId.
// Requires enable() to have been called previously.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-getChildAXNodes
//
// parameters:
//
//	id
func GetChildAXNodes(id NodeID) *GetChildAXNodesParams {
	return &GetChildAXNodesParams{
		ID: id,
	}
}

// WithFrameID the frame in whose document the node resides. If omitted, the
// root frame is used.
func (p GetChildAXNodesParams) WithFrameID(frameID cdp.FrameID) *GetChildAXNodesParams {
	p.FrameID = frameID
	return &p
}

// GetChildAXNodesReturns return values.
type GetChildAXNodesReturns struct {
	Nodes []*Node `json:"nodes,omitempty"`
}

// Do executes Accessibility.getChildAXNodes against the provided context.
//
// returns:
//
//	nodes
func (p *GetChildAXNodesParams) Do(ctx context.Context) (nodes []*Node, err error) {
	// execute
	var res GetChildAXNodesReturns
	err = cdp.Execute(ctx, CommandGetChildAXNodes, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// QueryAXTreeParams query a DOM node's accessibility subtree for accessible
// name and role. This command computes the name and role for all nodes in the
// subtree, including those that are ignored for accessibility, and returns
// those that mactch the specified name and role. If no DOM node is specified,
// or the DOM node does not exist, the command returns an error. If neither
// accessibleName or role is specified, it returns all the accessibility nodes
// in the subtree.
type QueryAXTreeParams struct {
	NodeID         cdp.NodeID             `json:"nodeId,omitempty"`         // Identifier of the node for the root to query.
	BackendNodeID  cdp.BackendNodeID      `json:"backendNodeId,omitempty"`  // Identifier of the backend node for the root to query.
	ObjectID       runtime.RemoteObjectID `json:"objectId,omitempty"`       // JavaScript object id of the node wrapper for the root to query.
	AccessibleName string                 `json:"accessibleName,omitempty"` // Find nodes with this computed name.
	Role           string                 `json:"role,omitempty"`           // Find nodes with this computed role.
}

// QueryAXTree query a DOM node's accessibility subtree for accessible name
// and role. This command computes the name and role for all nodes in the
// subtree, including those that are ignored for accessibility, and returns
// those that mactch the specified name and role. If no DOM node is specified,
// or the DOM node does not exist, the command returns an error. If neither
// accessibleName or role is specified, it returns all the accessibility nodes
// in the subtree.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#method-queryAXTree
//
// parameters:
func QueryAXTree() *QueryAXTreeParams {
	return &QueryAXTreeParams{}
}

// WithNodeID identifier of the node for the root to query.
func (p QueryAXTreeParams) WithNodeID(nodeID cdp.NodeID) *QueryAXTreeParams {
	p.NodeID = nodeID
	return &p
}

// WithBackendNodeID identifier of the backend node for the root to query.
func (p QueryAXTreeParams) WithBackendNodeID(backendNodeID cdp.BackendNodeID) *QueryAXTreeParams {
	p.BackendNodeID = backendNodeID
	return &p
}

// WithObjectID JavaScript object id of the node wrapper for the root to
// query.
func (p QueryAXTreeParams) WithObjectID(objectID runtime.RemoteObjectID) *QueryAXTreeParams {
	p.ObjectID = objectID
	return &p
}

// WithAccessibleName find nodes with this computed name.
func (p QueryAXTreeParams) WithAccessibleName(accessibleName string) *QueryAXTreeParams {
	p.AccessibleName = accessibleName
	return &p
}

// WithRole find nodes with this computed role.
func (p QueryAXTreeParams) WithRole(role string) *QueryAXTreeParams {
	p.Role = role
	return &p
}

// QueryAXTreeReturns return values.
type QueryAXTreeReturns struct {
	Nodes []*Node `json:"nodes,omitempty"` // A list of Accessibility.AXNode matching the specified attributes, including nodes that are ignored for accessibility.
}

// Do executes Accessibility.queryAXTree against the provided context.
//
// returns:
//
//	nodes - A list of Accessibility.AXNode matching the specified attributes, including nodes that are ignored for accessibility.
func (p *QueryAXTreeParams) Do(ctx context.Context) (nodes []*Node, err error) {
	// execute
	var res QueryAXTreeReturns
	err = cdp.Execute(ctx, CommandQueryAXTree, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// Command names.
const (
	CommandDisable               = "Accessibility.disable"
	CommandEnable                = "Accessibility.enable"
	CommandGetPartialAXTree      = "Accessibility.getPartialAXTree"
	CommandGetFullAXTree         = "Accessibility.getFullAXTree"
	CommandGetRootAXNode         = "Accessibility.getRootAXNode"
	CommandGetAXNodeAndAncestors = "Accessibility.getAXNodeAndAncestors"
	CommandGetChildAXNodes       = "Accessibility.getChildAXNodes"
	CommandQueryAXTree           = "Accessibility.queryAXTree"
)
