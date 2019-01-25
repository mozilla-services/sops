package textanalytics

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/textanalytics"

// BatchInput ...
type BatchInput struct {
	Documents *[]Input `json:"documents,omitempty"`
}

// DetectedLanguage ...
type DetectedLanguage struct {
	// Name - Long name of a detected language (e.g. English, French).
	Name *string `json:"name,omitempty"`
	// Iso6391Name - A two letter representation of the detected language according to the ISO 639-1 standard (e.g. en, fr).
	Iso6391Name *string `json:"iso6391Name,omitempty"`
	// Score - A confidence score between 0 and 1. Scores close to 1 indicate 100% certainty that the identified language is true.
	Score *float64 `json:"score,omitempty"`
}

// EntitiesBatchResultItemV2dot1 ...
type EntitiesBatchResultItemV2dot1 struct {
	// ID - Unique document identifier.
	ID *string `json:"id,omitempty"`
	// Entities - Recognized entities in the document.
	Entities *[]EntityRecordV2dot1 `json:"entities,omitempty"`
}

// EntitiesBatchResultV2dot1 ...
type EntitiesBatchResultV2dot1 struct {
	autorest.Response `json:"-"`
	Documents         *[]EntitiesBatchResultItemV2dot1 `json:"documents,omitempty"`
	Errors            *[]ErrorRecord                   `json:"errors,omitempty"`
}

// EntityRecordV2dot1 ...
type EntityRecordV2dot1 struct {
	// Name - Entity formal name.
	Name *string `json:"name,omitempty"`
	// Matches - List of instances this entity appears in the text.
	Matches *[]MatchRecordV2dot1 `json:"matches,omitempty"`
	// WikipediaLanguage - Wikipedia language for which the WikipediaId and WikipediaUrl refers to.
	WikipediaLanguage *string `json:"wikipediaLanguage,omitempty"`
	// WikipediaID - Wikipedia unique identifier of the recognized entity.
	WikipediaID *string `json:"wikipediaId,omitempty"`
	// WikipediaURL - URL for the entity's English Wikipedia page.
	WikipediaURL *string `json:"wikipediaUrl,omitempty"`
	// BingID - Bing unique identifier of the recognized entity. Use in conjunction with the Bing Entity Search API to fetch additional relevant information.
	BingID *string `json:"bingId,omitempty"`
	// Type - Entity type from Named Entity Recognition model
	Type *string `json:"type,omitempty"`
	// SubType - Entity sub type from Named Entity Recognition model
	SubType *string `json:"subType,omitempty"`
}

// ErrorRecord ...
type ErrorRecord struct {
	// ID - Input document unique identifier the error refers to.
	ID *string `json:"id,omitempty"`
	// Message - Error message.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Code       *string        `json:"code,omitempty"`
	Message    *string        `json:"message,omitempty"`
	Target     *string        `json:"target,omitempty"`
	InnerError *InternalError `json:"innerError,omitempty"`
}

// Input ...
type Input struct {
	// ID - Unique, non-empty document identifier.
	ID   *string `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
}

// InternalError ...
type InternalError struct {
	Code       *string        `json:"code,omitempty"`
	Message    *string        `json:"message,omitempty"`
	InnerError *InternalError `json:"innerError,omitempty"`
}

// KeyPhraseBatchResult ...
type KeyPhraseBatchResult struct {
	autorest.Response `json:"-"`
	Documents         *[]KeyPhraseBatchResultItem `json:"documents,omitempty"`
	Errors            *[]ErrorRecord              `json:"errors,omitempty"`
}

// KeyPhraseBatchResultItem ...
type KeyPhraseBatchResultItem struct {
	// KeyPhrases - A list of representative words or phrases. The number of key phrases returned is proportional to the number of words in the input document.
	KeyPhrases *[]string `json:"keyPhrases,omitempty"`
	// ID - Unique document identifier.
	ID *string `json:"id,omitempty"`
}

// LanguageBatchResult ...
type LanguageBatchResult struct {
	autorest.Response `json:"-"`
	Documents         *[]LanguageBatchResultItem `json:"documents,omitempty"`
	Errors            *[]ErrorRecord             `json:"errors,omitempty"`
}

// LanguageBatchResultItem ...
type LanguageBatchResultItem struct {
	// ID - Unique document identifier.
	ID *string `json:"id,omitempty"`
	// DetectedLanguages - A list of extracted languages.
	DetectedLanguages *[]DetectedLanguage `json:"detectedLanguages,omitempty"`
}

// MatchRecordV2dot1 ...
type MatchRecordV2dot1 struct {
	// Text - Entity text as appears in the request.
	Text *string `json:"text,omitempty"`
	// Offset - Start position (in Unicode characters) for the entity match text.
	Offset *int32 `json:"offset,omitempty"`
	// Length - Length (in Unicode characters) for the entity match text.
	Length *int32 `json:"length,omitempty"`
}

// MultiLanguageBatchInput ...
type MultiLanguageBatchInput struct {
	Documents *[]MultiLanguageInput `json:"documents,omitempty"`
}

// MultiLanguageInput ...
type MultiLanguageInput struct {
	// Language - This is the 2 letter ISO 639-1 representation of a language. For example, use "en" for English; "es" for Spanish etc.,
	Language *string `json:"language,omitempty"`
	// ID - Unique, non-empty document identifier.
	ID   *string `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
}

// SentimentBatchResult ...
type SentimentBatchResult struct {
	autorest.Response `json:"-"`
	Documents         *[]SentimentBatchResultItem `json:"documents,omitempty"`
	Errors            *[]ErrorRecord              `json:"errors,omitempty"`
}

// SentimentBatchResultItem ...
type SentimentBatchResultItem struct {
	// Score - A decimal number between 0 and 1 denoting the sentiment of the document. A score above 0.7 usually refers to a positive document while a score below 0.3 normally has a negative connotation. Mid values refer to neutral text.
	Score *float64 `json:"score,omitempty"`
	// ID - Unique document identifier.
	ID *string `json:"id,omitempty"`
}
