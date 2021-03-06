// Code generated by go-swagger; DO NOT EDIT.

package sms_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewUpdateSMSCampaignStatusParams creates a new UpdateSMSCampaignStatusParams object
// with the default values initialized.
func NewUpdateSMSCampaignStatusParams() *UpdateSMSCampaignStatusParams {
	var ()
	return &UpdateSMSCampaignStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSMSCampaignStatusParamsWithTimeout creates a new UpdateSMSCampaignStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSMSCampaignStatusParamsWithTimeout(timeout time.Duration) *UpdateSMSCampaignStatusParams {
	var ()
	return &UpdateSMSCampaignStatusParams{

		timeout: timeout,
	}
}

// NewUpdateSMSCampaignStatusParamsWithContext creates a new UpdateSMSCampaignStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSMSCampaignStatusParamsWithContext(ctx context.Context) *UpdateSMSCampaignStatusParams {
	var ()
	return &UpdateSMSCampaignStatusParams{

		Context: ctx,
	}
}

// NewUpdateSMSCampaignStatusParamsWithHTTPClient creates a new UpdateSMSCampaignStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSMSCampaignStatusParamsWithHTTPClient(client *http.Client) *UpdateSMSCampaignStatusParams {
	var ()
	return &UpdateSMSCampaignStatusParams{
		HTTPClient: client,
	}
}

/*UpdateSMSCampaignStatusParams contains all the parameters to send to the API endpoint
for the update SMS campaign status operation typically these are written to a http.Request
*/
type UpdateSMSCampaignStatusParams struct {

	/*CampaignID
	  id of the campaign

	*/
	CampaignID int64
	/*Status
	  Status of the campaign.

	*/
	Status *models.UpdateCampaignStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) WithTimeout(timeout time.Duration) *UpdateSMSCampaignStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) WithContext(ctx context.Context) *UpdateSMSCampaignStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) WithHTTPClient(client *http.Client) *UpdateSMSCampaignStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) WithCampaignID(campaignID int64) *UpdateSMSCampaignStatusParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithStatus adds the status to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) WithStatus(status *models.UpdateCampaignStatus) *UpdateSMSCampaignStatusParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the update SMS campaign status params
func (o *UpdateSMSCampaignStatusParams) SetStatus(status *models.UpdateCampaignStatus) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSMSCampaignStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.Status != nil {
		if err := r.SetBodyParam(o.Status); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
