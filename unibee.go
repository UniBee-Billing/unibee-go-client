package unibee

const (
	WaitingAuthorized = 0
	Authorized        = 1
	CaptureRequest    = 2
)

const (
	PaymentCreated   = 10
	PaymentSuccess   = 20
	PaymentFailed    = 30
	PaymentCancelled = 40
)

const (
	RefundCreated   = 10
	RefundSuccess   = 20
	RefundFailed    = 30
	RefundCancelled = 40
	RefundReverse   = 50
)

const (
	PlanTypeMain          = 1
	PlanTypeAddon         = 2
	PlanIntervalUnitDay   = "day"
	PlanIntervalUnitMonth = "month"
	PlanIntervalUnitYear  = "year"
	PlanIntervalUnitWeek  = "week"
)

const (
	UNIBEE_WEBHOOK_EVENT_SUBSCRIPTION_CREATED   = "subscription.created"
	UNIBEE_WEBHOOK_EVENT_SUBSCRIPTION_UPDATED   = "subscription.updated"
	UNIBEE_WEBHOOK_EVENT_SUBSCRIPTION_CANCELLED = "subscription.cancelled"
	UNIBEE_WEBHOOK_EVENT_SUBSCRIPTION_EXPIRED   = "subscription.expired"
	UNIBEE_WEBHOOK_EVENT_USER_METRIC_UPDATED    = "user.metric.update"
	UNIBEE_WEBHOOK_EVENT_PAYMENT_CREATED        = "payment.created"
	UNIBEE_WEBHOOK_EVENT_PAYMENT_NEEDAUTHORISED = "payment.authorised.need"
	UNIBEE_WEBHOOK_EVENT_PAYMENT_SUCCESS        = "payment.success"
	UNIBEE_WEBHOOK_EVENT_PAYMENT_CANCEL         = "payment.cancelled"
	UNIBEE_WEBHOOK_EVENT_PAYMENT_FAILURE        = "payment.failure"
	UNIBEE_WEBHOOK_EVENT_REFUND_CREATED         = "refund.created"
	UNIBEE_WEBHOOK_EVENT_REFUND_SUCCESS         = "refund.success"
	UNIBEE_WEBHOOK_EVENT_REFUND_FAILURE         = "refund.failure"
	UNIBEE_WEBHOOK_EVENT_REFUND_CANCELLED       = "refund.cancelled"
	UNIBEE_WEBHOOK_EVENT_REFUND_REVERSE         = "refund.reverse"
)

// AppInfo contains information about the "app" which this integration belongs
// to. This should be reserved for plugins that wish to identify themselves
// with Stripe.
type AppInfo struct {
	Name      string `json:"name"`
	PartnerID string `json:"partner_id"`
	URL       string `json:"url"`
	Version   string `json:"version"`
}

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolSlice returns a slice of bool pointers given a slice of bools.
func BoolSlice(v []bool) []*bool {
	out := make([]*bool, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Slice returns a slice of float64 pointers given a slice of float64s.
func Float64Slice(v []float64) []*float64 {
	out := make([]*float64, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// Int32 returns a pointer to the int64 value passed in.
func Int32(v int32) *int32 {
	return &v
}

// Int32Value returns the value of the int32 pointer passed in or
// 0 if the pointer is nil.
func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Slice returns a slice of int64 pointers given a slice of int64s.
func Int64Slice(v []int64) []*int64 {
	out := make([]*int64, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlice returns a slice of string pointers given a slice of strings.
func StringSlice(v []string) []*string {
	out := make([]*string, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}
