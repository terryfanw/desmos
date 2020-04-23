package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types/models"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/polls"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/reactions"
	"github.com/desmos-labs/desmos/x/posts/internal/types/msgs"
)

const (
	ModuleName                      = common.ModuleName
	RouterKey                       = common.RouterKey
	StoreKey                        = common.StoreKey
	MaxPostMessageLength            = common.MaxPostMessageLength
	MaxOptionalDataFieldsNumber     = common.MaxOptionalDataFieldsNumber
	MaxOptionalDataFieldValueLength = common.MaxOptionalDataFieldValueLength
	ActionCreatePost                = common.ActionCreatePost
	ActionEditPost                  = common.ActionEditPost
	ActionAnswerPoll                = common.ActionAnswerPoll
	ActionAddPostReaction           = common.ActionAddPostReaction
	ActionRemovePostReaction        = common.ActionRemovePostReaction
	ActionRegisterReaction          = common.ActionRegisterReaction
	QuerierRoute                    = common.QuerierRoute
	QueryPost                       = common.QueryPost
	QueryPosts                      = common.QueryPosts
	QueryPollAnswers                = common.QueryPollAnswers
	QueryRegisteredReactions        = common.QueryRegisteredReactions
	PostSortByID                    = common.PostSortByID
	PostSortByCreationDate          = common.PostSortByCreationDate
	PostSortOrderAscending          = common.PostSortOrderAscending
	PostSortOrderDescending         = common.PostSortOrderDescending
)

var (
	// functions aliases
	NewPostResponse          = models.NewPostResponse
	RegisterModelsCodec      = models.RegisterModelsCodec
	NewPostCreationData      = models.NewPostCreationData
	PostStoreKey             = models.PostStoreKey
	PostCommentsStoreKey     = models.PostCommentsStoreKey
	PostReactionsStoreKey    = models.PostReactionsStoreKey
	ReactionsStoreKey        = models.ReactionsStoreKey
	PollAnswersStoreKey      = models.PollAnswersStoreKey
	NewPost                  = models.NewPost
	ParsePostID              = models.ParsePostID
	NewPostMedia             = common.NewPostMedia
	ValidateURI              = common.ValidateURI
	NewPostMedias            = common.NewPostMedias
	ParseAnswerID            = polls.ParseAnswerID
	NewPollAnswer            = polls.NewPollAnswer
	NewPollAnswers           = polls.NewPollAnswers
	NewPollData              = polls.NewPollData
	ArePollDataEquals        = polls.ArePollDataEquals
	NewUserAnswer            = polls.NewUserAnswer
	NewPostReaction          = reactions.NewPostReaction
	NewReaction              = reactions.NewReaction
	IsEmoji                  = reactions.IsEmoji
	RegisterMessagesCodec    = msgs.RegisterMessagesCodec
	NewMsgCreatePost         = msgs.NewMsgCreatePost
	NewMsgEditPost           = msgs.NewMsgEditPost
	NewMsgAnswerPoll         = msgs.NewMsgAnswerPoll
	NewMsgAddPostReaction    = msgs.NewMsgAddPostReaction
	NewMsgRemovePostReaction = msgs.NewMsgRemovePostReaction
	NewMsgRegisterReaction   = msgs.NewMsgRegisterReaction

	// variable aliases
	ModelsCdc                = models.ModelsCdc
	SubspaceRegEx            = common.SubspaceRegEx
	HashtagRegEx             = common.HashtagRegEx
	ShortCodeRegEx           = common.ShortCodeRegEx
	URIRegEx                 = common.URIRegEx
	LastPostIDStoreKey       = common.LastPostIDStoreKey
	PostStorePrefix          = common.PostStorePrefix
	PostCommentsStorePrefix  = common.PostCommentsStorePrefix
	PostReactionsStorePrefix = common.PostReactionsStorePrefix
	ReactionsStorePrefix     = common.ReactionsStorePrefix
	PollAnswersStorePrefix   = common.PollAnswersStorePrefix
	MsgsCodec                = msgs.MsgsCodec
)

type (
	PostReaction             = reactions.PostReaction
	PostReactions            = reactions.PostReactions
	Reaction                 = reactions.Reaction
	Reactions                = reactions.Reactions
	MsgCreatePost            = msgs.MsgCreatePost
	MsgEditPost              = msgs.MsgEditPost
	MsgAnswerPoll            = msgs.MsgAnswerPoll
	MsgAddPostReaction       = msgs.MsgAddPostReaction
	MsgRemovePostReaction    = msgs.MsgRemovePostReaction
	MsgRegisterReaction      = msgs.MsgRegisterReaction
	PostQueryResponse        = models.PostQueryResponse
	PostCreationData         = models.PostCreationData
	PollAnswersQueryResponse = models.PollAnswersQueryResponse
	Post                     = models.Post
	Posts                    = models.Posts
	PostID                   = models.PostID
	PostIDs                  = models.PostIDs
	PostMedia                = common.PostMedia
	PostMedias               = common.PostMedias
	OptionalData             = common.OptionalData
	KeyValue                 = common.KeyValue
	AnswerID                 = polls.AnswerID
	PollAnswer               = polls.PollAnswer
	PollAnswers              = polls.PollAnswers
	PollData                 = polls.PollData
	UserAnswer               = polls.UserAnswer
	UserAnswers              = polls.UserAnswers
)
