package posts

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

const (
	OpWeightMsgCreatePost           = simulation.OpWeightMsgCreatePost
	OpWeightMsgEditPost             = simulation.OpWeightMsgEditPost
	OpWeightMsgAddReaction          = simulation.OpWeightMsgAddReaction
	OpWeightMsgRemoveReaction       = simulation.OpWeightMsgRemoveReaction
	OpWeightMsgAnswerPoll           = simulation.OpWeightMsgAnswerPoll
	OpWeightMsgRegisterReaction     = simulation.OpWeightMsgRegisterReaction
	DefaultGasValue                 = simulation.DefaultGasValue
	EventTypePostCreated            = types.EventTypePostCreated
	EventTypePostEdited             = types.EventTypePostEdited
	EventTypePostReactionAdded      = types.EventTypePostReactionAdded
	EventTypePostReactionRemoved    = types.EventTypePostReactionRemoved
	EventTypeAnsweredPoll           = types.EventTypeAnsweredPoll
	EventTypeClosePoll              = types.EventTypeClosePoll
	EventTypeRegisterReaction       = types.EventTypeRegisterReaction
	AttributeKeyPostID              = types.AttributeKeyPostID
	AttributeKeyPostParentID        = types.AttributeKeyPostParentID
	AttributeKeyPostOwner           = types.AttributeKeyPostOwner
	AttributeKeyPostEditTime        = types.AttributeKeyPostEditTime
	AttributeKeyPollAnswerer        = types.AttributeKeyPollAnswerer
	AttributeKeyPostReactionOwner   = types.AttributeKeyPostReactionOwner
	AttributeKeyPostReactionValue   = types.AttributeKeyPostReactionValue
	AttributeKeyReactionCreator     = types.AttributeKeyReactionCreator
	AttributeKeyReactionShortCode   = types.AttributeKeyReactionShortCode
	AttributeKeyReactionSubSpace    = types.AttributeKeyReactionSubSpace
	AttributeKeyCreationTime        = types.AttributeKeyCreationTime
	PostSortByID                    = types.PostSortByID
	PostSortByCreationDate          = types.PostSortByCreationDate
	PostSortOrderAscending          = types.PostSortOrderAscending
	PostSortOrderDescending         = types.PostSortOrderDescending
	ModuleName                      = types.ModuleName
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	MaxPostMessageLength            = types.MaxPostMessageLength
	MaxOptionalDataFieldsNumber     = types.MaxOptionalDataFieldsNumber
	MaxOptionalDataFieldValueLength = types.MaxOptionalDataFieldValueLength
	ActionCreatePost                = types.ActionCreatePost
	ActionEditPost                  = types.ActionEditPost
	ActionAnswerPoll                = types.ActionAnswerPoll
	ActionAddPostReaction           = types.ActionAddPostReaction
	ActionRemovePostReaction        = types.ActionRemovePostReaction
	ActionRegisterReaction          = types.ActionRegisterReaction
	QuerierRoute                    = types.QuerierRoute
	QueryPost                       = types.QueryPost
	QueryPosts                      = types.QueryPosts
	QueryPollAnswers                = types.QueryPollAnswers
	QueryRegisteredReactions        = types.QueryRegisteredReactions
)

var (
	// functions aliases
	NewHandler                    = keeper.NewHandler
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	DecodeStore                   = simulation.DecodeStore
	RandomizedGenState            = simulation.RandomizedGenState
	WeightedOperations            = simulation.WeightedOperations
	SimulateMsgAnswerToPoll       = simulation.SimulateMsgAnswerToPoll
	SimulateMsgCreatePost         = simulation.SimulateMsgCreatePost
	SimulateMsgEditPost           = simulation.SimulateMsgEditPost
	SimulateMsgAddPostReaction    = simulation.SimulateMsgAddPostReaction
	SimulateMsgRemovePostReaction = simulation.SimulateMsgRemovePostReaction
	SimulateMsgRegisterReaction   = simulation.SimulateMsgRegisterReaction
	RandomPost                    = simulation.RandomPost
	RandomPostData                = simulation.RandomPostData
	RandomPostReactionData        = simulation.RandomPostReactionData
	RandomPostReactionValue       = simulation.RandomPostReactionValue
	RandomPostID                  = simulation.RandomPostID
	RandomMessage                 = simulation.RandomMessage
	RandomSubspace                = simulation.RandomSubspace
	RandomHashtag                 = simulation.RandomHashtag
	RandomMedias                  = simulation.RandomMedias
	RandomPollData                = simulation.RandomPollData
	GetAccount                    = simulation.GetAccount
	RandomReactionValue           = simulation.RandomReactionValue
	RandomReactionShortCode       = simulation.RandomReactionShortCode
	RandomReactionData            = simulation.RandomReactionData
	RegisteredReactionsData       = simulation.RegisteredReactionsData
	NewReaction                   = types.NewReaction
	IsEmoji                       = types.IsEmoji
	RegisterCodec                 = types.RegisterCodec
	NewPostCreationPocketData     = types.NewPostCreationPocketData
	NewPollData                   = types.NewPollData
	ArePollDataEquals             = types.ArePollDataEquals
	NewUserAnswer                 = types.NewUserAnswer
	NewPostReaction               = types.NewPostReaction
	NewMsgRegisterReaction        = types.NewMsgRegisterReaction
	PostStoreKey                  = types.PostStoreKey
	PostCommentsStoreKey          = types.PostCommentsStoreKey
	PostReactionsStoreKey         = types.PostReactionsStoreKey
	ReactionsStoreKey             = types.ReactionsStoreKey
	PollAnswersStoreKey           = types.PollAnswersStoreKey
	NewMsgCreatePost              = types.NewMsgCreatePost
	NewMsgEditPost                = types.NewMsgEditPost
	NewMsgAnswerPoll              = types.NewMsgAnswerPoll
	NewPostMedia                  = types.NewPostMedia
	ValidateURI                   = types.ValidateURI
	NewPostMedias                 = types.NewPostMedias
	NewPostResponse               = types.NewPostResponse
	DefaultQueryPostsParams       = types.DefaultQueryPostsParams
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	ValidateGenesis               = types.ValidateGenesis
	ParsePostID                   = types.ParsePostID
	NewPost                       = types.NewPost
	NewMsgAddPostReaction         = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction      = types.NewMsgRemovePostReaction
	ParseAnswerID                 = types.ParseAnswerID
	NewPollAnswer                 = types.NewPollAnswer
	NewPollAnswers                = types.NewPollAnswers

	// variable aliases
	RandomMimeTypes          = simulation.RandomMimeTypes
	RandomHosts              = simulation.RandomHosts
	ModuleCdc                = types.ModuleCdc
	SubspaceRegEx            = types.SubspaceRegEx
	HashtagRegEx             = types.HashtagRegEx
	ShortCodeRegEx           = types.ShortCodeRegEx
	URIRegEx                 = types.URIRegEx
	LastPostIDStoreKey       = types.LastPostIDStoreKey
	PostStorePrefix          = types.PostStorePrefix
	PostCommentsStorePrefix  = types.PostCommentsStorePrefix
	PostReactionsStorePrefix = types.PostReactionsStorePrefix
	ReactionsStorePrefix     = types.ReactionsStorePrefix
	PollAnswersStorePrefix   = types.PollAnswersStorePrefix
)

type (
	Keeper                   = keeper.Keeper
	PostData                 = simulation.PostData
	PostReactionData         = simulation.PostReactionData
	ReactionData             = simulation.ReactionData
	PollAnswersQueryResponse = types.PollAnswersQueryResponse
	Reaction                 = types.Reaction
	Reactions                = types.Reactions
	PostCreationPacketData   = types.PostCreationPacketData
	PollData                 = types.PollData
	UserAnswer               = types.UserAnswer
	UserAnswers              = types.UserAnswers
	PostReaction             = types.PostReaction
	PostReactions            = types.PostReactions
	CreatePostPacketData     = types.CreatePostPacketData
	AckDataCreation          = types.AckDataCreation
	MsgRegisterReaction      = types.MsgRegisterReaction
	MsgCreatePost            = types.MsgCreatePost
	MsgEditPost              = types.MsgEditPost
	MsgAnswerPoll            = types.MsgAnswerPoll
	PostMedia                = types.PostMedia
	PostMedias               = types.PostMedias
	PostQueryResponse        = types.PostQueryResponse
	QueryPostsParams         = types.QueryPostsParams
	GenesisState             = types.GenesisState
	PostID                   = types.PostID
	PostIDs                  = types.PostIDs
	Post                     = types.Post
	Posts                    = types.Posts
	OptionalData             = types.OptionalData
	KeyValue                 = types.KeyValue
	MsgAddPostReaction       = types.MsgAddPostReaction
	MsgRemovePostReaction    = types.MsgRemovePostReaction
	AnswerID                 = types.AnswerID
	PollAnswer               = types.PollAnswer
	PollAnswers              = types.PollAnswers
)
