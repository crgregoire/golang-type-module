package types

//
// GoogleHomeRequest is used when google home
// attempts to fulfill an action
//
type GoogleHomeRequest struct {
	// Session                     string                   `json:"session,omitempty"`
	Inputs    []GoogleHomeRequestInput `json:"inputs,omitempty"`
	RequestID string                   `json:"requestID,omitempty"`
	// QueryResult                 googleHomeQueryResult    `json:"queryResult,omitempty"`
	// OriginalDetectIntentRequest googleHomeIntentRequest  `json:"originalDetectIntentRequest,omitempty"`
}

//
// GoogleHomeRequestInput is a struct
// that supports google home integration
//
type GoogleHomeRequestInput struct {
	Commands []GoogleHomeInputCommand `json:"commands"`
	Intent   string                   `json:"intent"`
}

//
// GoogleHomeInputCommand is a struct
// that supports google home integration
//
type GoogleHomeInputCommand struct {
	Devices   []GoogleHomeDevice               `json:"devices"`
	Execution []GoogleHomeInputCommandExuction `json:"execution"`
}

//
// GoogleHomeInputCommandExuction is a struct
// that supports google home integration
//
type GoogleHomeInputCommandExuction struct {
	Command string                 `json:"command"`
	Params  map[string]interface{} `json:"params"`
}

type googleHomeQueryResult struct {
	QueryText                   string                            `json:"QueryText,omitempty"`
	LanguageCode                string                            `json:"LanguageCode,omitempty"`
	SpeechRecognitionConfidence float32                           `json:"SpeechRecognitionConfidence,omitempty"`
	Action                      string                            `json:"Action,omitempty"`
	Parameters                  map[string]interface{}            `json:"Parameters,omitempty"`
	AllRequiredParamsPresent    bool                              `json:"AllRequiredParamsPresent,omitempty"`
	FulfillmentText             string                            `json:"FulfillmentText,omitempty"`
	FulfillmentMessages         []map[string]interface{}          `json:"FulfillmentMessages,omitempty"`
	WebhookSource               string                            `json:"WebhookSource,omitempty"`
	WebhookPayload              map[string]interface{}            `json:"WebhookPayload,omitempty"`
	OutputContexts              []googleHomeOutputContext         `json:"OutputContexts,omitempty"`
	Intent                      googleHomeIntent                  `json:"Intent,omitempty"`
	IntentDetectionConfidence   float32                           `json:"IntentDetectionConfidence,omitempty"`
	DiagnosticInfo              map[string]interface{}            `json:"DiagnosticInfo,omitempty"`
	SentimentAnalysisResult     googleHomeSentimentAnalysisResult `json:"SentimentAnalysisResult,omitempty"`
}

type googleHomeIntentRequest struct {
	Source  string                 `json:"source,omitempty"`
	Version string                 `json:"version,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
}

type googleHomeOutputContext struct {
	Name          string                 `json:"name,omitempty"`
	LifespanCount int32                  `json:"lifespan_count,omitempty"`
	Parameters    map[string]interface{} `json:"parameters,omitempty"`
}

type googleHomeIntent struct {
	Name              string                 `json:"name,omitempty"`
	DisplayName       string                 `json:"display_name,omitempty"`
	WebhookState      string                 `json:"webhook_state,omitempty"`
	Priority          int32                  `json:"priority,omitempty"`
	IsFallback        bool                   `json:"is_fallback,omitempty"`
	MlDisabled        bool                   `json:"ml_disabled,omitempty"`
	InputContextNames string                 `json:"input_context_names,omitempty"`
	Events            string                 `json:"events,omitempty"`
	Action            string                 `json:"action,omitempty"`
	Parameters        map[string]interface{} `json:"parameters,omitempty"`
}

type googleHomeSentimentAnalysisResult struct {
	Sentiment sentiment `json:"sentiment,omitempty"`
}

type sentiment struct {
	Score     float32
	magnitude float32
}

//
// GoogleHomeResponse is used to reply to
// google home when fulfilling an intent
//
type GoogleHomeResponse struct {
	RequestID           string                    `json:"requestId,omitempty"`
	FulfillmentText     string                    `json:"fulfillmentText,omitempty"`
	FulfillmentMessages []map[string]interface{}  `json:"fulfillmentMessages,omitempty"`
	Source              string                    `json:"source,omitempty"`
	Payload             ResponsePayload           `json:"payload,omitempty"`
	OutputContexts      []googleHomeOutputContext `json:"outputContexts,omitempty"`
}

//
// ResponsePayload is a struct
// that supports google home integration
//
type ResponsePayload struct {
	AgentUserID string             `json:"agentUserId"`
	Devices     []GoogleHomeDevice `json:"devices"`
}

//
// GoogleHomeRequestUser is the user for
// which google home is fulfilling an intent
//
type GoogleHomeRequestUser struct {
	AccessToken            string `json:"accessToken,omitempty"`
	Locale                 string `json:"locale,omitempty"`
	LastSeen               string `json:"lastSeen,omitempty"`
	UserVerificationStatus string `json:"userVerificationStatus,omitempty"`
}

//
// GoogleHomeDevice is a struct
// that supports google home integration
//
type GoogleHomeDevice struct {
	ID              string                     `json:"id,omitempty"`
	Type            string                     `json:"type,omitempty"`
	Traits          []string                   `json:"traits,omitempty"`
	Name            GoogleHomeDeviceName       `json:"name,omitempty"`
	DeviceInfo      string                     `json:"deviceinfo,omitempty"`
	WillReportState bool                       `json:"willReportState"`
	Attributes      GoogleHomeDeviceAttributes `json:"attributes"`
}

//
// GoogleHomeDeviceAttributes is a struct
// that supports google home integration
//
type GoogleHomeDeviceAttributes struct {
	Pausable       bool     `json:"pausable"`
	AvailableZones []string `json:"availableZones"`
}

//
// GoogleHomeDeviceName is a struct
// that supports google home integration
//
type GoogleHomeDeviceName struct {
	DefaultNames []string `json:"defaultNames,omitempty"`
	Name         string   `json:"name,omitempty"`
	Nicknames    []string `json:"nicknames,omitempty"`
}

//
// GoogleHomeDeviceInfo is a struct
// that supports google home integration
//
type GoogleHomeDeviceInfo struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
}
