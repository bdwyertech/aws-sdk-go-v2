// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Contains information about the contexts that a user is using in a session. You
// can configure Amazon Lex V2 to set a context when an intent is fulfilled, or you
// can set a context using the , , or operations. Use a context to indicate to
// Amazon Lex V2 intents that should be used as follow-up intents. For example, if
// the active context is order-fulfilled, only intents that have order-fulfilled
// configured as a trigger are considered for follow up.
type ActiveContext struct {

	// A lis tof contexts active for the request. A context can be activated when a
	// previous intent is fulfilled, or by including the context in the request. If you
	// don't specify a list of contexts, Amazon Lex will use the current list of
	// contexts for the session. If you specify an empty list, all contexts for the
	// session are cleared.
	//
	// This member is required.
	ContextAttributes map[string]string

	// The name of the context.
	//
	// This member is required.
	Name *string

	// Indicates the number of turns or seconds that the context is active. Once the
	// time to live expires, the context is no longer returned in a response.
	//
	// This member is required.
	TimeToLive *ActiveContextTimeToLive
}

// The time that a context is active. You can specify the time to live in seconds
// or in conversation turns.
type ActiveContextTimeToLive struct {

	// The number of seconds that the context is active. You can specify between 5 and
	// 86400 seconds (24 hours).
	//
	// This member is required.
	TimeToLiveInSeconds *int32

	// The number of turns that the context is active. You can specify up to 20 turns.
	// Each request and response from the bot is a turn.
	//
	// This member is required.
	TurnsToLive *int32
}

// A button that appears on a response card show to the user.
type Button struct {

	// The text that is displayed on the button.
	//
	// This member is required.
	Text *string

	// The value returned to Amazon Lex V2 when a user chooses the button.
	//
	// This member is required.
	Value *string
}

// Provides a score that indicates the confidence that Amazon Lex V2 has that an
// intent is the one that satisfies the user's intent.
type ConfidenceScore struct {

	// A score that indicates how confident Amazon Lex V2 is that an intent satisfies
	// the user's intent. Ranges between 0.00 and 1.00. Higher scores indicate higher
	// confidence.
	Score float64
}

// The next action that Amazon Lex V2 should take.
type DialogAction struct {

	// The next action that the bot should take in its interaction with the user. The
	// possible values are:
	//
	// * Close - Indicates that there will not be a response from
	// the user. For example, the statement "Your order has been placed" does not
	// require a response.
	//
	// * ConfirmIntent - The next action is asking the user if the
	// intent is complete and ready to be fulfilled. This is a yes/no question such as
	// "Place the order?"
	//
	// * Delegate - The next action is determined by Amazon Lex
	// V2.
	//
	// * ElicitSlot - The next action is to elicit a slot value from the user.
	//
	// This member is required.
	Type DialogActionType

	// The name of the slot that should be elicited from the user.
	SlotToElicit *string
}

// A card that is shown to the user by a messaging platform. You define the
// contents of the card, the card is displayed by the platform. When you use a
// response card, the response from the user is constrained to the text associated
// with a button on the card.
type ImageResponseCard struct {

	// The title to display on the response card. The format of the title is determined
	// by the platform displaying the response card.
	//
	// This member is required.
	Title *string

	// A list of buttons that should be displayed on the response card. The arrangement
	// of the buttons is determined by the platform that displays the button.
	Buttons []Button

	// The URL of an image to display on the response card. The image URL must be
	// publicly available so that the platform displaying the response card has access
	// to the image.
	ImageUrl *string

	// The subtitle to display on the response card. The format of the subtitle is
	// determined by the platform displaying the response card.
	Subtitle *string
}

// The current intent that Amazon Lex V2 is attempting to fulfill.
type Intent struct {

	// The name of the intent.
	//
	// This member is required.
	Name *string

	// Contains information about whether fulfillment of the intent has been confirmed.
	ConfirmationState ConfirmationState

	// A map of all of the slots for the intent. The name of the slot maps to the value
	// of the slot. If a slot has not been filled, the value is null.
	Slots map[string]Slot

	// Contains fulfillment information for the intent.
	State IntentState
}

// An intent that Amazon Lex V2 determined might satisfy the user's utterance. The
// intents are ordered by the confidence score.
type Interpretation struct {

	// A list of intents that might satisfy the user's utterance. The intents are
	// ordered by the confidence score.
	Intent *Intent

	// Determines the threshold where Amazon Lex V2 will insert the
	// AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning
	// alternative intents in a response. AMAZON.FallbackIntent and
	// AMAZON.KendraSearchIntent are only inserted if they are configured for the bot.
	NluConfidence *ConfidenceScore

	// The sentiment expressed in an utterance. When the bot is configured to send
	// utterances to Amazon Comprehend for sentiment analysis, this field contains the
	// result of the analysis.
	SentimentResponse *SentimentResponse
}

// Container for text that is returned to the customer..
type Message struct {

	// Indicates the type of response.
	//
	// This member is required.
	ContentType MessageContentType

	// The text of the message.
	Content *string

	// A card that is shown to the user by a messaging platform. You define the
	// contents of the card, the card is displayed by the platform. When you use a
	// response card, the response from the user is constrained to the text associated
	// with a button on the card.
	ImageResponseCard *ImageResponseCard
}

// Provides information about the sentiment expressed in a user's response in a
// conversation. Sentiments are determined using Amazon Comprehend. Sentiments are
// only returned if they are enabled for the bot. For more information, see
// Determine Sentiment
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-sentiment.html) in the
// Amazon Comprehend developer guide.
type SentimentResponse struct {

	// The overall sentiment expressed in the user's response. This is the sentiment
	// most likely expressed by the user based on the analysis by Amazon Comprehend.
	Sentiment SentimentType

	// The individual sentiment responses for the utterance.
	SentimentScore *SentimentScore
}

// The individual sentiment responses for the utterance.
type SentimentScore struct {

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the MIXED sentiment.
	Mixed float64

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the NEGATIVE sentiment.
	Negative float64

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the NEUTRAL sentiment.
	Neutral float64

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the POSITIVE sentiment.
	Positive float64
}

// The state of the user's session with Amazon Lex V2.
type SessionState struct {

	// One or more contexts that indicate to Amazon Lex V2 the context of a request.
	// When a context is active, Amazon Lex V2 considers intents with the matching
	// context as a trigger as the next intent in a session.
	ActiveContexts []ActiveContext

	// The next step that Amazon Lex V2 should take in the conversation with a user.
	DialogAction *DialogAction

	// The active intent that Amazon Lex V2 is processing.
	Intent *Intent

	//
	OriginatingRequestId *string

	// Map of key/value pairs representing session-specific context information. It
	// contains application information passed between Amazon Lex V2 and a client
	// application.
	SessionAttributes map[string]string
}

// A value that Amazon Lex V2 uses to fulfill an intent.
type Slot struct {

	// When the shape value is List, it indicates that the values field contains a list
	// of slot values. When the value is Scalar, it indicates that the value field
	// contains a single value.
	Shape Shape

	// The current value of the slot.
	Value *Value

	// A list of one or more values that the user provided for the slot. For example,
	// if a for a slot that elicits pizza toppings, the values might be "pepperoni" and
	// "pineapple."
	Values []Slot
}

// The value of a slot.
type Value struct {

	// The value that Amazon Lex V2 determines for the slot. The actual value depends
	// on the setting of the value selection strategy for the bot. You can choose to
	// use the value entered by the user, or you can have Amazon Lex V2 choose the
	// first value in the resolvedValues list.
	//
	// This member is required.
	InterpretedValue *string

	// The text of the utterance from the user that was entered for the slot.
	OriginalValue *string

	// A list of additional values that have been recognized for the slot.
	ResolvedValues []string
}
