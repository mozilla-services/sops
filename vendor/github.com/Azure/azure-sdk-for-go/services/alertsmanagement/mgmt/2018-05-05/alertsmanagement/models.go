package alertsmanagement

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
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/alertsmanagement/mgmt/2018-05-05/alertsmanagement"

// AlertModificationEvent enumerates the values for alert modification event.
type AlertModificationEvent string

const (
	// AlertCreated ...
	AlertCreated AlertModificationEvent = "AlertCreated"
	// MonitorConditionChange ...
	MonitorConditionChange AlertModificationEvent = "MonitorConditionChange"
	// StateChange ...
	StateChange AlertModificationEvent = "StateChange"
)

// PossibleAlertModificationEventValues returns an array of possible values for the AlertModificationEvent const type.
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return []AlertModificationEvent{AlertCreated, MonitorConditionChange, StateChange}
}

// AlertsSortByFields enumerates the values for alerts sort by fields.
type AlertsSortByFields string

const (
	// AlertsSortByFieldsAlertState ...
	AlertsSortByFieldsAlertState AlertsSortByFields = "alertState"
	// AlertsSortByFieldsLastModifiedDateTime ...
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = "lastModifiedDateTime"
	// AlertsSortByFieldsMonitorCondition ...
	AlertsSortByFieldsMonitorCondition AlertsSortByFields = "monitorCondition"
	// AlertsSortByFieldsName ...
	AlertsSortByFieldsName AlertsSortByFields = "name"
	// AlertsSortByFieldsSeverity ...
	AlertsSortByFieldsSeverity AlertsSortByFields = "severity"
	// AlertsSortByFieldsStartDateTime ...
	AlertsSortByFieldsStartDateTime AlertsSortByFields = "startDateTime"
	// AlertsSortByFieldsTargetResource ...
	AlertsSortByFieldsTargetResource AlertsSortByFields = "targetResource"
	// AlertsSortByFieldsTargetResourceGroup ...
	AlertsSortByFieldsTargetResourceGroup AlertsSortByFields = "targetResourceGroup"
	// AlertsSortByFieldsTargetResourceName ...
	AlertsSortByFieldsTargetResourceName AlertsSortByFields = "targetResourceName"
	// AlertsSortByFieldsTargetResourceType ...
	AlertsSortByFieldsTargetResourceType AlertsSortByFields = "targetResourceType"
)

// PossibleAlertsSortByFieldsValues returns an array of possible values for the AlertsSortByFields const type.
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return []AlertsSortByFields{AlertsSortByFieldsAlertState, AlertsSortByFieldsLastModifiedDateTime, AlertsSortByFieldsMonitorCondition, AlertsSortByFieldsName, AlertsSortByFieldsSeverity, AlertsSortByFieldsStartDateTime, AlertsSortByFieldsTargetResource, AlertsSortByFieldsTargetResourceGroup, AlertsSortByFieldsTargetResourceName, AlertsSortByFieldsTargetResourceType}
}

// AlertsSummaryGroupByFields enumerates the values for alerts summary group by fields.
type AlertsSummaryGroupByFields string

const (
	// AlertsSummaryGroupByFieldsAlertRule ...
	AlertsSummaryGroupByFieldsAlertRule AlertsSummaryGroupByFields = "alertRule"
	// AlertsSummaryGroupByFieldsAlertState ...
	AlertsSummaryGroupByFieldsAlertState AlertsSummaryGroupByFields = "alertState"
	// AlertsSummaryGroupByFieldsMonitorCondition ...
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = "monitorCondition"
	// AlertsSummaryGroupByFieldsMonitorService ...
	AlertsSummaryGroupByFieldsMonitorService AlertsSummaryGroupByFields = "monitorService"
	// AlertsSummaryGroupByFieldsSeverity ...
	AlertsSummaryGroupByFieldsSeverity AlertsSummaryGroupByFields = "severity"
	// AlertsSummaryGroupByFieldsSignalType ...
	AlertsSummaryGroupByFieldsSignalType AlertsSummaryGroupByFields = "signalType"
)

// PossibleAlertsSummaryGroupByFieldsValues returns an array of possible values for the AlertsSummaryGroupByFields const type.
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return []AlertsSummaryGroupByFields{AlertsSummaryGroupByFieldsAlertRule, AlertsSummaryGroupByFieldsAlertState, AlertsSummaryGroupByFieldsMonitorCondition, AlertsSummaryGroupByFieldsMonitorService, AlertsSummaryGroupByFieldsSeverity, AlertsSummaryGroupByFieldsSignalType}
}

// AlertState enumerates the values for alert state.
type AlertState string

const (
	// AlertStateAcknowledged ...
	AlertStateAcknowledged AlertState = "Acknowledged"
	// AlertStateClosed ...
	AlertStateClosed AlertState = "Closed"
	// AlertStateNew ...
	AlertStateNew AlertState = "New"
)

// PossibleAlertStateValues returns an array of possible values for the AlertState const type.
func PossibleAlertStateValues() []AlertState {
	return []AlertState{AlertStateAcknowledged, AlertStateClosed, AlertStateNew}
}

// MonitorCondition enumerates the values for monitor condition.
type MonitorCondition string

const (
	// Fired ...
	Fired MonitorCondition = "Fired"
	// Resolved ...
	Resolved MonitorCondition = "Resolved"
)

// PossibleMonitorConditionValues returns an array of possible values for the MonitorCondition const type.
func PossibleMonitorConditionValues() []MonitorCondition {
	return []MonitorCondition{Fired, Resolved}
}

// MonitorService enumerates the values for monitor service.
type MonitorService string

const (
	// ActivityLogAdministrative ...
	ActivityLogAdministrative MonitorService = "ActivityLog Administrative"
	// ActivityLogAutoscale ...
	ActivityLogAutoscale MonitorService = "ActivityLog Autoscale"
	// ActivityLogPolicy ...
	ActivityLogPolicy MonitorService = "ActivityLog Policy"
	// ActivityLogRecommendation ...
	ActivityLogRecommendation MonitorService = "ActivityLog Recommendation"
	// ActivityLogSecurity ...
	ActivityLogSecurity MonitorService = "ActivityLog Security"
	// ApplicationInsights ...
	ApplicationInsights MonitorService = "Application Insights"
	// LogAnalytics ...
	LogAnalytics MonitorService = "Log Analytics"
	// Nagios ...
	Nagios MonitorService = "Nagios"
	// Platform ...
	Platform MonitorService = "Platform"
	// SCOM ...
	SCOM MonitorService = "SCOM"
	// ServiceHealth ...
	ServiceHealth MonitorService = "ServiceHealth"
	// SmartDetector ...
	SmartDetector MonitorService = "SmartDetector"
	// VMInsights ...
	VMInsights MonitorService = "VM Insights"
	// Zabbix ...
	Zabbix MonitorService = "Zabbix"
)

// PossibleMonitorServiceValues returns an array of possible values for the MonitorService const type.
func PossibleMonitorServiceValues() []MonitorService {
	return []MonitorService{ActivityLogAdministrative, ActivityLogAutoscale, ActivityLogPolicy, ActivityLogRecommendation, ActivityLogSecurity, ApplicationInsights, LogAnalytics, Nagios, Platform, SCOM, ServiceHealth, SmartDetector, VMInsights, Zabbix}
}

// Severity enumerates the values for severity.
type Severity string

const (
	// Sev0 ...
	Sev0 Severity = "Sev0"
	// Sev1 ...
	Sev1 Severity = "Sev1"
	// Sev2 ...
	Sev2 Severity = "Sev2"
	// Sev3 ...
	Sev3 Severity = "Sev3"
	// Sev4 ...
	Sev4 Severity = "Sev4"
)

// PossibleSeverityValues returns an array of possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{Sev0, Sev1, Sev2, Sev3, Sev4}
}

// SignalType enumerates the values for signal type.
type SignalType string

const (
	// Log ...
	Log SignalType = "Log"
	// Metric ...
	Metric SignalType = "Metric"
	// Unknown ...
	Unknown SignalType = "Unknown"
)

// PossibleSignalTypeValues returns an array of possible values for the SignalType const type.
func PossibleSignalTypeValues() []SignalType {
	return []SignalType{Log, Metric, Unknown}
}

// SmartGroupModificationEvent enumerates the values for smart group modification event.
type SmartGroupModificationEvent string

const (
	// SmartGroupModificationEventAlertAdded ...
	SmartGroupModificationEventAlertAdded SmartGroupModificationEvent = "AlertAdded"
	// SmartGroupModificationEventAlertRemoved ...
	SmartGroupModificationEventAlertRemoved SmartGroupModificationEvent = "AlertRemoved"
	// SmartGroupModificationEventSmartGroupCreated ...
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = "SmartGroupCreated"
	// SmartGroupModificationEventStateChange ...
	SmartGroupModificationEventStateChange SmartGroupModificationEvent = "StateChange"
)

// PossibleSmartGroupModificationEventValues returns an array of possible values for the SmartGroupModificationEvent const type.
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return []SmartGroupModificationEvent{SmartGroupModificationEventAlertAdded, SmartGroupModificationEventAlertRemoved, SmartGroupModificationEventSmartGroupCreated, SmartGroupModificationEventStateChange}
}

// SmartGroupsSortByFields enumerates the values for smart groups sort by fields.
type SmartGroupsSortByFields string

const (
	// SmartGroupsSortByFieldsAlertsCount ...
	SmartGroupsSortByFieldsAlertsCount SmartGroupsSortByFields = "alertsCount"
	// SmartGroupsSortByFieldsLastModifiedDateTime ...
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = "lastModifiedDateTime"
	// SmartGroupsSortByFieldsSeverity ...
	SmartGroupsSortByFieldsSeverity SmartGroupsSortByFields = "severity"
	// SmartGroupsSortByFieldsStartDateTime ...
	SmartGroupsSortByFieldsStartDateTime SmartGroupsSortByFields = "startDateTime"
	// SmartGroupsSortByFieldsState ...
	SmartGroupsSortByFieldsState SmartGroupsSortByFields = "state"
)

// PossibleSmartGroupsSortByFieldsValues returns an array of possible values for the SmartGroupsSortByFields const type.
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return []SmartGroupsSortByFields{SmartGroupsSortByFieldsAlertsCount, SmartGroupsSortByFieldsLastModifiedDateTime, SmartGroupsSortByFieldsSeverity, SmartGroupsSortByFieldsStartDateTime, SmartGroupsSortByFieldsState}
}

// State enumerates the values for state.
type State string

const (
	// StateAcknowledged ...
	StateAcknowledged State = "Acknowledged"
	// StateClosed ...
	StateClosed State = "Closed"
	// StateNew ...
	StateNew State = "New"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateAcknowledged, StateClosed, StateNew}
}

// TimeRange enumerates the values for time range.
type TimeRange string

const (
	// Oned ...
	Oned TimeRange = "1d"
	// Oneh ...
	Oneh TimeRange = "1h"
	// Sevend ...
	Sevend TimeRange = "7d"
	// ThreeZerod ...
	ThreeZerod TimeRange = "30d"
)

// PossibleTimeRangeValues returns an array of possible values for the TimeRange const type.
func PossibleTimeRangeValues() []TimeRange {
	return []TimeRange{Oned, Oneh, Sevend, ThreeZerod}
}

// Alert an alert created in alert management service.
type Alert struct {
	autorest.Response `json:"-"`
	Properties        *AlertProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
}

// AlertModification alert Modification details
type AlertModification struct {
	autorest.Response `json:"-"`
	Properties        *AlertModificationProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
}

// AlertModificationItem alert modification item.
type AlertModificationItem struct {
	// ModificationEvent - Reason for the modification. Possible values include: 'AlertCreated', 'StateChange', 'MonitorConditionChange'
	ModificationEvent AlertModificationEvent `json:"modificationEvent,omitempty"`
	// OldValue - Old value
	OldValue *string `json:"oldValue,omitempty"`
	// NewValue - New value
	NewValue *string `json:"newValue,omitempty"`
	// ModifiedAt - Modified date and time
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// ModifiedBy - Modified user details (Principal client name)
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// Comments - Modification comments
	Comments *string `json:"comments,omitempty"`
	// Description - Description of the modification
	Description *string `json:"description,omitempty"`
}

// AlertModificationProperties properties of the alert modification item.
type AlertModificationProperties struct {
	// AlertID - READ-ONLY; Unique Id of the alert for which the history is being retrieved
	AlertID *string `json:"alertId,omitempty"`
	// Modifications - Modification details
	Modifications *[]AlertModificationItem `json:"modifications,omitempty"`
}

// AlertProperties alert property bag
type AlertProperties struct {
	Essentials   *Essentials `json:"essentials,omitempty"`
	Context      interface{} `json:"context,omitempty"`
	EgressConfig interface{} `json:"egressConfig,omitempty"`
}

// AlertsList list the alerts.
type AlertsList struct {
	autorest.Response `json:"-"`
	// NextLink - URL to fetch the next set of alerts.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - List of alerts
	Value *[]Alert `json:"value,omitempty"`
}

// AlertsListIterator provides access to a complete listing of Alert values.
type AlertsListIterator struct {
	i    int
	page AlertsListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AlertsListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AlertsListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AlertsListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AlertsListIterator) Response() AlertsList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AlertsListIterator) Value() Alert {
	if !iter.page.NotDone() {
		return Alert{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AlertsListIterator type.
func NewAlertsListIterator(page AlertsListPage) AlertsListIterator {
	return AlertsListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (al AlertsList) IsEmpty() bool {
	return al.Value == nil || len(*al.Value) == 0
}

// alertsListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (al AlertsList) alertsListPreparer(ctx context.Context) (*http.Request, error) {
	if al.NextLink == nil || len(to.String(al.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(al.NextLink)))
}

// AlertsListPage contains a page of Alert values.
type AlertsListPage struct {
	fn func(context.Context, AlertsList) (AlertsList, error)
	al AlertsList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AlertsListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertsListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.al)
	if err != nil {
		return err
	}
	page.al = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AlertsListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AlertsListPage) NotDone() bool {
	return !page.al.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AlertsListPage) Response() AlertsList {
	return page.al
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AlertsListPage) Values() []Alert {
	if page.al.IsEmpty() {
		return nil
	}
	return *page.al.Value
}

// Creates a new instance of the AlertsListPage type.
func NewAlertsListPage(getNextPage func(context.Context, AlertsList) (AlertsList, error)) AlertsListPage {
	return AlertsListPage{fn: getNextPage}
}

// AlertsSummary summary of alerts based on the input filters and 'groupby' parameters.
type AlertsSummary struct {
	autorest.Response `json:"-"`
	Properties        *AlertsSummaryGroup `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
}

// AlertsSummaryGroup group the result set.
type AlertsSummaryGroup struct {
	// Total - Total count of the result set.
	Total *int32 `json:"total,omitempty"`
	// SmartGroupsCount - Total count of the smart groups.
	SmartGroupsCount *int32 `json:"smartGroupsCount,omitempty"`
	// Groupedby - Name of the field aggregated
	Groupedby *string `json:"groupedby,omitempty"`
	// Values - List of the items
	Values *[]AlertsSummaryGroupItem `json:"values,omitempty"`
}

// AlertsSummaryGroupItem alerts summary group item
type AlertsSummaryGroupItem struct {
	// Name - Value of the aggregated field
	Name *string `json:"name,omitempty"`
	// Count - Count of the aggregated field
	Count *int32 `json:"count,omitempty"`
	// Groupedby - Name of the field aggregated
	Groupedby *string `json:"groupedby,omitempty"`
	// Values - List of the items
	Values *[]AlertsSummaryGroupItem `json:"values,omitempty"`
}

// ErrorResponse an error response from the service.
type ErrorResponse struct {
	Error *ErrorResponseBody `json:"error,omitempty"`
}

// ErrorResponseBody details of error response.
type ErrorResponseBody struct {
	// Code - Error code, intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - Description of the error, intended for display in user interface.
	Message *string `json:"message,omitempty"`
	// Target - Target of the particular error, for example name of the property.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]ErrorResponseBody `json:"details,omitempty"`
}

// Essentials this object contains consistent fields across different monitor services.
type Essentials struct {
	// Severity - READ-ONLY; Severity of alert Sev0 being highest and Sev4 being lowest. Possible values include: 'Sev0', 'Sev1', 'Sev2', 'Sev3', 'Sev4'
	Severity Severity `json:"severity,omitempty"`
	// SignalType - READ-ONLY; The type of signal the alert is based on, which could be metrics, logs or activity logs. Possible values include: 'Metric', 'Log', 'Unknown'
	SignalType SignalType `json:"signalType,omitempty"`
	// AlertState - READ-ONLY; Alert object state, which can be modified by the user. Possible values include: 'AlertStateNew', 'AlertStateAcknowledged', 'AlertStateClosed'
	AlertState AlertState `json:"alertState,omitempty"`
	// MonitorCondition - READ-ONLY; Can be 'Fired' or 'Resolved', which represents whether the underlying conditions have crossed the defined alert rule thresholds. Possible values include: 'Fired', 'Resolved'
	MonitorCondition MonitorCondition `json:"monitorCondition,omitempty"`
	// TargetResource - Target ARM resource, on which alert got created.
	TargetResource *string `json:"targetResource,omitempty"`
	// TargetResourceName - Name of the target ARM resource name, on which alert got created.
	TargetResourceName *string `json:"targetResourceName,omitempty"`
	// TargetResourceGroup - Resource group of target ARM resource, on which alert got created.
	TargetResourceGroup *string `json:"targetResourceGroup,omitempty"`
	// TargetResourceType - Resource type of target ARM resource, on which alert got created.
	TargetResourceType *string `json:"targetResourceType,omitempty"`
	// MonitorService - READ-ONLY; Monitor service on which the rule(monitor) is set. Possible values include: 'ApplicationInsights', 'ActivityLogAdministrative', 'ActivityLogSecurity', 'ActivityLogRecommendation', 'ActivityLogPolicy', 'ActivityLogAutoscale', 'LogAnalytics', 'Nagios', 'Platform', 'SCOM', 'ServiceHealth', 'SmartDetector', 'VMInsights', 'Zabbix'
	MonitorService MonitorService `json:"monitorService,omitempty"`
	// AlertRule - READ-ONLY; Rule(monitor) which fired alert instance. Depending on the monitor service,  this would be ARM id or name of the rule.
	AlertRule *string `json:"alertRule,omitempty"`
	// SourceCreatedID - READ-ONLY; Unique Id created by monitor service for each alert instance. This could be used to track the issue at the monitor service, in case of Nagios, Zabbix, SCOM etc.
	SourceCreatedID *string `json:"sourceCreatedId,omitempty"`
	// SmartGroupID - READ-ONLY; Unique Id of the smart group
	SmartGroupID *string `json:"smartGroupId,omitempty"`
	// SmartGroupingReason - READ-ONLY; Verbose reason describing the reason why this alert instance is added to a smart group
	SmartGroupingReason *string `json:"smartGroupingReason,omitempty"`
	// StartDateTime - READ-ONLY; Creation time(ISO-8601 format) of alert instance.
	StartDateTime *date.Time `json:"startDateTime,omitempty"`
	// LastModifiedDateTime - READ-ONLY; Last modification time(ISO-8601 format) of alert instance.
	LastModifiedDateTime *date.Time `json:"lastModifiedDateTime,omitempty"`
	// MonitorConditionResolvedDateTime - READ-ONLY; Resolved time(ISO-8601 format) of alert instance. This will be updated when monitor service resolves the alert instance because the rule condition is no longer met.
	MonitorConditionResolvedDateTime *date.Time `json:"monitorConditionResolvedDateTime,omitempty"`
	// LastModifiedUserName - READ-ONLY; User who last modified the alert, in case of monitor service updates user would be 'system', otherwise name of the user.
	LastModifiedUserName *string `json:"lastModifiedUserName,omitempty"`
}

// Operation operation provided by provider
type Operation struct {
	// Name - Name of the operation
	Name *string `json:"name,omitempty"`
	// Display - Properties of the operation
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay properties of the operation
type OperationDisplay struct {
	// Provider - Provider name
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation name
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation
	Description *string `json:"description,omitempty"`
}

// OperationsList lists the operations available in the AlertsManagement RP.
type OperationsList struct {
	autorest.Response `json:"-"`
	// NextLink - URL to fetch the next set of alerts.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - Array of operations
	Value *[]Operation `json:"value,omitempty"`
}

// OperationsListIterator provides access to a complete listing of Operation values.
type OperationsListIterator struct {
	i    int
	page OperationsListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationsListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationsListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationsListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationsListIterator) Response() OperationsList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationsListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationsListIterator type.
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return OperationsListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationsList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationsListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationsList) operationsListPreparer(ctx context.Context) (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationsListPage contains a page of Operation values.
type OperationsListPage struct {
	fn func(context.Context, OperationsList) (OperationsList, error)
	ol OperationsList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationsListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationsListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationsListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationsListPage) Response() OperationsList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationsListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationsListPage type.
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return OperationsListPage{fn: getNextPage}
}

// Resource an azure resource object
type Resource struct {
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
}

// SmartGroup set of related alerts grouped together smartly by AMS.
type SmartGroup struct {
	autorest.Response     `json:"-"`
	*SmartGroupProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for SmartGroup.
func (sg SmartGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sg.SmartGroupProperties != nil {
		objectMap["properties"] = sg.SmartGroupProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SmartGroup struct.
func (sg *SmartGroup) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var smartGroupProperties SmartGroupProperties
				err = json.Unmarshal(*v, &smartGroupProperties)
				if err != nil {
					return err
				}
				sg.SmartGroupProperties = &smartGroupProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				sg.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				sg.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				sg.Name = &name
			}
		}
	}

	return nil
}

// SmartGroupAggregatedProperty aggregated property of each type
type SmartGroupAggregatedProperty struct {
	// Name - Name of the type.
	Name *string `json:"name,omitempty"`
	// Count - Total number of items of type.
	Count *int32 `json:"count,omitempty"`
}

// SmartGroupModification alert Modification details
type SmartGroupModification struct {
	autorest.Response `json:"-"`
	Properties        *SmartGroupModificationProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
}

// SmartGroupModificationItem smartGroup modification item.
type SmartGroupModificationItem struct {
	// ModificationEvent - Reason for the modification. Possible values include: 'SmartGroupModificationEventSmartGroupCreated', 'SmartGroupModificationEventStateChange', 'SmartGroupModificationEventAlertAdded', 'SmartGroupModificationEventAlertRemoved'
	ModificationEvent SmartGroupModificationEvent `json:"modificationEvent,omitempty"`
	// OldValue - Old value
	OldValue *string `json:"oldValue,omitempty"`
	// NewValue - New value
	NewValue *string `json:"newValue,omitempty"`
	// ModifiedAt - Modified date and time
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// ModifiedBy - Modified user details (Principal client name)
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// Comments - Modification comments
	Comments *string `json:"comments,omitempty"`
	// Description - Description of the modification
	Description *string `json:"description,omitempty"`
}

// SmartGroupModificationProperties properties of the smartGroup modification item.
type SmartGroupModificationProperties struct {
	// SmartGroupID - READ-ONLY; Unique Id of the smartGroup for which the history is being retrieved
	SmartGroupID *string `json:"smartGroupId,omitempty"`
	// Modifications - Modification details
	Modifications *[]SmartGroupModificationItem `json:"modifications,omitempty"`
	// NextLink - URL to fetch the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// SmartGroupProperties properties of smart group.
type SmartGroupProperties struct {
	// AlertsCount - Total number of alerts in smart group
	AlertsCount *int32 `json:"alertsCount,omitempty"`
	// SmartGroupState - READ-ONLY; Smart group state. Possible values include: 'StateNew', 'StateAcknowledged', 'StateClosed'
	SmartGroupState State `json:"smartGroupState,omitempty"`
	// Severity - READ-ONLY; Severity of smart group is the highest(Sev0 >... > Sev4) severity of all the alerts in the group. Possible values include: 'Sev0', 'Sev1', 'Sev2', 'Sev3', 'Sev4'
	Severity Severity `json:"severity,omitempty"`
	// StartDateTime - READ-ONLY; Creation time of smart group. Date-Time in ISO-8601 format.
	StartDateTime *date.Time `json:"startDateTime,omitempty"`
	// LastModifiedDateTime - READ-ONLY; Last updated time of smart group. Date-Time in ISO-8601 format.
	LastModifiedDateTime *date.Time `json:"lastModifiedDateTime,omitempty"`
	// LastModifiedUserName - READ-ONLY; Last modified by user name.
	LastModifiedUserName *string `json:"lastModifiedUserName,omitempty"`
	// Resources - Summary of target resources in the smart group
	Resources *[]SmartGroupAggregatedProperty `json:"resources,omitempty"`
	// ResourceTypes - Summary of target resource types in the smart group
	ResourceTypes *[]SmartGroupAggregatedProperty `json:"resourceTypes,omitempty"`
	// ResourceGroups - Summary of target resource groups in the smart group
	ResourceGroups *[]SmartGroupAggregatedProperty `json:"resourceGroups,omitempty"`
	// MonitorServices - Summary of monitorServices in the smart group
	MonitorServices *[]SmartGroupAggregatedProperty `json:"monitorServices,omitempty"`
	// MonitorConditions - Summary of monitorConditions in the smart group
	MonitorConditions *[]SmartGroupAggregatedProperty `json:"monitorConditions,omitempty"`
	// AlertStates - Summary of alertStates in the smart group
	AlertStates *[]SmartGroupAggregatedProperty `json:"alertStates,omitempty"`
	// AlertSeverities - Summary of alertSeverities in the smart group
	AlertSeverities *[]SmartGroupAggregatedProperty `json:"alertSeverities,omitempty"`
	// NextLink - The URI to fetch the next page of alerts. Call ListNext() with this URI to fetch the next page alerts.
	NextLink *string `json:"nextLink,omitempty"`
}

// SmartGroupsList list the alerts.
type SmartGroupsList struct {
	autorest.Response `json:"-"`
	// NextLink - URL to fetch the next set of alerts.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - List of alerts
	Value *[]SmartGroup `json:"value,omitempty"`
}
