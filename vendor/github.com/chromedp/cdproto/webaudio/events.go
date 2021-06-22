package webaudio

// Code generated by cdproto-gen. DO NOT EDIT.

// EventContextCreated notifies that a new BaseAudioContext has been created.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-contextCreated
type EventContextCreated struct {
	Context *BaseAudioContext `json:"context"`
}

// EventContextWillBeDestroyed notifies that an existing BaseAudioContext
// will be destroyed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-contextWillBeDestroyed
type EventContextWillBeDestroyed struct {
	ContextID GraphObjectID `json:"contextId"`
}

// EventContextChanged notifies that existing BaseAudioContext has changed
// some properties (id stays the same)..
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-contextChanged
type EventContextChanged struct {
	Context *BaseAudioContext `json:"context"`
}

// EventAudioListenerCreated notifies that the construction of an
// AudioListener has finished.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-audioListenerCreated
type EventAudioListenerCreated struct {
	Listener *AudioListener `json:"listener"`
}

// EventAudioListenerWillBeDestroyed notifies that a new AudioListener has
// been created.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-audioListenerWillBeDestroyed
type EventAudioListenerWillBeDestroyed struct {
	ContextID  GraphObjectID `json:"contextId"`
	ListenerID GraphObjectID `json:"listenerId"`
}

// EventAudioNodeCreated notifies that a new AudioNode has been created.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-audioNodeCreated
type EventAudioNodeCreated struct {
	Node *AudioNode `json:"node"`
}

// EventAudioNodeWillBeDestroyed notifies that an existing AudioNode has been
// destroyed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-audioNodeWillBeDestroyed
type EventAudioNodeWillBeDestroyed struct {
	ContextID GraphObjectID `json:"contextId"`
	NodeID    GraphObjectID `json:"nodeId"`
}

// EventAudioParamCreated notifies that a new AudioParam has been created.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-audioParamCreated
type EventAudioParamCreated struct {
	Param *AudioParam `json:"param"`
}

// EventAudioParamWillBeDestroyed notifies that an existing AudioParam has
// been destroyed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-audioParamWillBeDestroyed
type EventAudioParamWillBeDestroyed struct {
	ContextID GraphObjectID `json:"contextId"`
	NodeID    GraphObjectID `json:"nodeId"`
	ParamID   GraphObjectID `json:"paramId"`
}

// EventNodesConnected notifies that two AudioNodes are connected.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-nodesConnected
type EventNodesConnected struct {
	ContextID             GraphObjectID `json:"contextId"`
	SourceID              GraphObjectID `json:"sourceId"`
	DestinationID         GraphObjectID `json:"destinationId"`
	SourceOutputIndex     float64       `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64       `json:"destinationInputIndex,omitempty"`
}

// EventNodesDisconnected notifies that AudioNodes are disconnected. The
// destination can be null, and it means all the outgoing connections from the
// source are disconnected.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-nodesDisconnected
type EventNodesDisconnected struct {
	ContextID             GraphObjectID `json:"contextId"`
	SourceID              GraphObjectID `json:"sourceId"`
	DestinationID         GraphObjectID `json:"destinationId"`
	SourceOutputIndex     float64       `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64       `json:"destinationInputIndex,omitempty"`
}

// EventNodeParamConnected notifies that an AudioNode is connected to an
// AudioParam.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-nodeParamConnected
type EventNodeParamConnected struct {
	ContextID         GraphObjectID `json:"contextId"`
	SourceID          GraphObjectID `json:"sourceId"`
	DestinationID     GraphObjectID `json:"destinationId"`
	SourceOutputIndex float64       `json:"sourceOutputIndex,omitempty"`
}

// EventNodeParamDisconnected notifies that an AudioNode is disconnected to
// an AudioParam.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-nodeParamDisconnected
type EventNodeParamDisconnected struct {
	ContextID         GraphObjectID `json:"contextId"`
	SourceID          GraphObjectID `json:"sourceId"`
	DestinationID     GraphObjectID `json:"destinationId"`
	SourceOutputIndex float64       `json:"sourceOutputIndex,omitempty"`
}
