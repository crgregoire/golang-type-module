package types

import "time"

//
// AlexaRequest is a struct
// that will support alexa integration
//
type AlexaRequest struct {
	Directive AlexaDirective `json:"directive"`
}

//
// AlexaDirective is a struct
// that will support alexa integration
//
type AlexaDirective struct {
	Header   AlexaHeader                               `json:"header"`
	Payload  AlexaPayload                              `json:"payload"`
	Endpoint AlexaDiscoverResponseEventPayloadEndpoint `json:"endpoint"`
}

//
// AlexaHeader is a struct
// that will support alexa integration
//
type AlexaHeader struct {
	Namespace        string    `json:"namespace"`
	Name             string    `json:"name"`
	MessageID        string    `json:"messageId"`
	PayloadVersion   string    `json:"payloadVersion"`
	Value            string    `json:"value,omitempty"`
	TimeOfSample     time.Time `json:"timeOfSample,omitempty"`
	CorrelationToken string    `json:"correlationToken,omitempty"`
}

//
// AlexaPayload is a struct
// that will support alexa integration
//
type AlexaPayload struct {
	Scope AlexaScope `json:"scope"`
}

//
// AlexaScope is a struct
// that will support alexa integration
//
type AlexaScope struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

//
// AlexaResponse is a struct
// that will support alexa integration
//
type AlexaResponse struct {
	Event   AlexaDiscoverResponseEvent `json:"event"`
	Context AlexaContext               `json:"context,omitempty"`
}

//
// AlexaContext is a struct
// that will support alexa integration
//
type AlexaContext struct {
	Properties []AlexaHeader `json:"properties"`
}

//
// AlexaDiscoverResponseEvent is a struct
// that will support alexa integration
//
type AlexaDiscoverResponseEvent struct {
	Payload AlexaDiscoverResponseEventPayload `json:"payload"`
	Header  AlexaHeader                       `json:"header"`
}

//
// AlexaDiscoverResponseEventPayload is a struct
// that will support alexa integration
//
type AlexaDiscoverResponseEventPayload struct {
	Endpoints []AlexaDiscoverResponseEventPayloadEndpoint `json:"endpoints"`
}

//
// AlexaDiscoverResponseEventPayloadEndpoint is a struct
// that will support alexa integration
//
type AlexaDiscoverResponseEventPayloadEndpoint struct {
	EndpointID           string                      `json:"endpointId"`
	FriendlyName         string                      `json:"friendlyName"`
	Description          string                      `json:"description"`
	ManufacturerName     string                      `json:"manufacturerName"`
	DisplayCategories    []string                    `json:"displayCategories"`
	AdditionalAttributes AlexaAdditionalAttributes   `json:"additionalAttributes"`
	Cookie               map[string]interface{}      `json:"cookie"`
	Capabilities         []AlexaEndpointCapabilities `json:"capabilities"`
	Scope                AlexaScope                  `json:"scope,omitempty"`
}

//
// AlexaEndpointCapabilities is a struct
// that will support alexa integration
//
type AlexaEndpointCapabilities struct {
	Type       string                    `json:"type"`
	Interface  string                    `json:"interface"`
	Version    string                    `json:"version"`
	Instance   string                    `json:"instance,omitempty"`
	Properties AlexaCapabilityProperites `json:"properties"`
}

//
// AlexaCapabilityResources is a struct
// that will support alexa integration
//
type AlexaCapabilityResources struct {
	FriendlyNames []AlexaCapabilityResourcesFriendlyNames `json:"friendlyNames"`
}

//
// AlexaCapabilityResourcesFriendlyNames is a struct
// that will support alexa integration
//
type AlexaCapabilityResourcesFriendlyNames struct {
	Type  string            `json:"@type"`
	Value map[string]string `json:"value"`
}

//
// AlexaCapabilityProperites is a struct
// that will support alexa integration
//
type AlexaCapabilityProperites struct {
	ProactivelyReported bool                `json:"proactivelyReported"`
	Retrievable         bool                `json:"retrievable"`
	Supported           []map[string]string `json:"supported"`
}

//
// AlexaAdditionalAttributes is a struct
// that will support alexa integration
//
type AlexaAdditionalAttributes struct {
	Manufacturer     string `json:"manufacturer"`
	Model            string `json:"model"`
	SerialNumber     string `json:"serialNumber"`
	FirmwareVersion  string `json:"firmwareVersion"`
	SoftwareVersion  string `json:"softwareVersion"`
	CustomIdentifier string `json:"customIdentifier"`
}
