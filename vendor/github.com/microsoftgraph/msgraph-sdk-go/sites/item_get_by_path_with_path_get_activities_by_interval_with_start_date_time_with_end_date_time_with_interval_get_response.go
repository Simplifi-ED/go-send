package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse 
type ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}
// NewItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instantiates a new ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse and sets the default values.
func NewItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse()(*ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) {
    m := &ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemActivityStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable 
type ItemGetByPathWithPathGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable)
    SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable)()
}