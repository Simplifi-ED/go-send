package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instead.
type ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse struct {
    ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse
}
// NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse instantiates a new ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse and sets the default values.
func NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse()(*ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) {
    m := &ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse{
        ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse: *NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(),
    }
    return m
}
// CreateItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse(), nil
}
// ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instead.
type ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable interface {
    ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}