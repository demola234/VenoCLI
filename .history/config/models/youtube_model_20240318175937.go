package main

import "encoding/json"

func UnmarshalYoutubePayload(data []byte) (YoutubePayload, error) {
	var r YoutubePayload
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *YoutubePayload) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type YoutubePayload struct {
	ResponseContext        ResponseContext   `json:"responseContext"`
	PlayabilityStatus      PlayabilityStatus `json:"playabilityStatus"`
	StreamingData          StreamingData     `json:"streamingData"`
	PlaybackTracking       PlaybackTracking  `json:"playbackTracking"`
	Captions               Captions          `json:"captions"`
	VideoDetails           VideoDetails      `json:"videoDetails"`
	PlayerConfig           PlayerConfig      `json:"playerConfig"`
	Storyboards            Storyboards       `json:"storyboards"`
	Microformat            Microformat       `json:"microformat"`
	TrackingParams         string            `json:"trackingParams"`
	Attestation            Attestation       `json:"attestation"`
	Messages               []Message         `json:"messages"`
	AdBreakHeartbeatParams string            `json:"adBreakHeartbeatParams"`
	FrameworkUpdates       FrameworkUpdates  `json:"frameworkUpdates"`
}

type Attestation struct {
	PlayerAttestationRenderer PlayerAttestationRenderer `json:"playerAttestationRenderer"`
}

type PlayerAttestationRenderer struct {
	Challenge    string       `json:"challenge"`
	BotguardData BotguardData `json:"botguardData"`
}

type BotguardData struct {
	Program            string             `json:"program"`
	InterpreterSafeURL InterpreterSafeURL `json:"interpreterSafeUrl"`
	ServerEnvironment  int64              `json:"serverEnvironment"`
}

type InterpreterSafeURL struct {
	PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
}

type Captions struct {
	PlayerCaptionsTracklistRenderer PlayerCaptionsTracklistRenderer `json:"playerCaptionsTracklistRenderer"`
}

type PlayerCaptionsTracklistRenderer struct {
	CaptionTracks          []CaptionTrack        `json:"captionTracks"`
	AudioTracks            []AudioTrack          `json:"audioTracks"`
	TranslationLanguages   []TranslationLanguage `json:"translationLanguages"`
	DefaultAudioTrackIndex int64                 `json:"defaultAudioTrackIndex"`
}

type AudioTrack struct {
	CaptionTrackIndices      []int64 `json:"captionTrackIndices"`
	DefaultCaptionTrackIndex int64   `json:"defaultCaptionTrackIndex"`
	Visibility               string  `json:"visibility"`
	HasDefaultTrack          bool    `json:"hasDefaultTrack"`
	CaptionsInitialState     string  `json:"captionsInitialState"`
}

type CaptionTrack struct {
	BaseURL        string      `json:"baseUrl"`
	Name           Description `json:"name"`
	VssID          string      `json:"vssId"`
	LanguageCode   string      `json:"languageCode"`
	RTL            *bool       `json:"rtl,omitempty"`
	IsTranslatable bool        `json:"isTranslatable"`
	TrackName      string      `json:"trackName"`
	Kind           *string     `json:"kind,omitempty"`
}

type Description struct {
	SimpleText string `json:"simpleText"`
}

type TranslationLanguage struct {
	LanguageCode string      `json:"languageCode"`
	LanguageName Description `json:"languageName"`
}

type FrameworkUpdates struct {
	EntityBatchUpdate EntityBatchUpdate `json:"entityBatchUpdate"`
}

type EntityBatchUpdate struct {
	Mutations []Mutation `json:"mutations"`
	Timestamp Timestamp  `json:"timestamp"`
}

type Mutation struct {
	EntityKey string  `json:"entityKey"`
	Type      string  `json:"type"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	OfflineabilityEntity OfflineabilityEntity `json:"offlineabilityEntity"`
}

type OfflineabilityEntity struct {
	Key                     string `json:"key"`
	AddToOfflineButtonState string `json:"addToOfflineButtonState"`
}

type Timestamp struct {
	Seconds string `json:"seconds"`
	Nanos   int64  `json:"nanos"`
}

type Message struct {
	MealbarPromoRenderer MealbarPromoRenderer `json:"mealbarPromoRenderer"`
}

type MealbarPromoRenderer struct {
	Icon                IconClass            `json:"icon"`
	MessageTexts        []MessageTitle       `json:"messageTexts"`
	ActionButton        ActionButton         `json:"actionButton"`
	DismissButton       DismissButton        `json:"dismissButton"`
	TriggerCondition    string               `json:"triggerCondition"`
	Style               string               `json:"style"`
	TrackingParams      string               `json:"trackingParams"`
	ImpressionEndpoints []ImpressionEndpoint `json:"impressionEndpoints"`
	IsVisible           bool                 `json:"isVisible"`
	MessageTitle        MessageTitle         `json:"messageTitle"`
}

type ActionButton struct {
	ButtonRenderer ActionButtonButtonRenderer `json:"buttonRenderer"`
}

type ActionButtonButtonRenderer struct {
	Style          string        `json:"style"`
	Size           string        `json:"size"`
	Text           MessageTitle  `json:"text"`
	TrackingParams string        `json:"trackingParams"`
	Command        PurpleCommand `json:"command"`
}

type PurpleCommand struct {
	ClickTrackingParams    string                       `json:"clickTrackingParams"`
	CommandExecutorCommand PurpleCommandExecutorCommand `json:"commandExecutorCommand"`
}

type PurpleCommandExecutorCommand struct {
	Commands []CommandElement `json:"commands"`
}

type CommandElement struct {
	ClickTrackingParams *string               `json:"clickTrackingParams,omitempty"`
	CommandMetadata     PurpleCommandMetadata `json:"commandMetadata"`
	BrowseEndpoint      *BrowseEndpoint       `json:"browseEndpoint,omitempty"`
	FeedbackEndpoint    *FeedbackEndpoint     `json:"feedbackEndpoint,omitempty"`
}

type BrowseEndpoint struct {
	BrowseID string `json:"browseId"`
	Params   string `json:"params"`
}

type PurpleCommandMetadata struct {
	WebCommandMetadata PurpleWebCommandMetadata `json:"webCommandMetadata"`
}

type PurpleWebCommandMetadata struct {
	URL         *string `json:"url,omitempty"`
	WebPageType *string `json:"webPageType,omitempty"`
	RootVe      *int64  `json:"rootVe,omitempty"`
	APIURL      string  `json:"apiUrl"`
	SendPost    *bool   `json:"sendPost,omitempty"`
}

type FeedbackEndpoint struct {
	FeedbackToken string    `json:"feedbackToken"`
	UIActions     UIActions `json:"uiActions"`
}

type UIActions struct {
	HideEnclosingContainer bool `json:"hideEnclosingContainer"`
}

type MessageTitle struct {
	Runs []Run `json:"runs"`
}

type Run struct {
	Text string `json:"text"`
}

type DismissButton struct {
	ButtonRenderer DismissButtonButtonRenderer `json:"buttonRenderer"`
}

type DismissButtonButtonRenderer struct {
	Style          string        `json:"style"`
	Size           string        `json:"size"`
	Text           MessageTitle  `json:"text"`
	TrackingParams string        `json:"trackingParams"`
	Command        FluffyCommand `json:"command"`
}

type FluffyCommand struct {
	ClickTrackingParams    string                       `json:"clickTrackingParams"`
	CommandExecutorCommand FluffyCommandExecutorCommand `json:"commandExecutorCommand"`
}

type FluffyCommandExecutorCommand struct {
	Commands []ImpressionEndpoint `json:"commands"`
}

type ImpressionEndpoint struct {
	ClickTrackingParams string                            `json:"clickTrackingParams"`
	CommandMetadata     ImpressionEndpointCommandMetadata `json:"commandMetadata"`
	FeedbackEndpoint    FeedbackEndpoint                  `json:"feedbackEndpoint"`
}

type ImpressionEndpointCommandMetadata struct {
	WebCommandMetadata FluffyWebCommandMetadata `json:"webCommandMetadata"`
}

type FluffyWebCommandMetadata struct {
	SendPost bool   `json:"sendPost"`
	APIURL   string `json:"apiUrl"`
}

type IconClass struct {
	Thumbnails []ThumbnailElement `json:"thumbnails"`
}

type ThumbnailElement struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Microformat struct {
	PlayerMicroformatRenderer PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
	Thumbnail          IconClass   `json:"thumbnail"`
	Embed              Embed       `json:"embed"`
	Title              Description `json:"title"`
	Description        Description `json:"description"`
	LengthSeconds      string      `json:"lengthSeconds"`
	OwnerProfileURL    string      `json:"ownerProfileUrl"`
	ExternalChannelID  string      `json:"externalChannelId"`
	IsFamilySafe       bool        `json:"isFamilySafe"`
	AvailableCountries []string    `json:"availableCountries"`
	IsUnlisted         bool        `json:"isUnlisted"`
	HasYpcMetadata     bool        `json:"hasYpcMetadata"`
	ViewCount          string      `json:"viewCount"`
	Category           string      `json:"category"`
	PublishDate        string      `json:"publishDate"`
	OwnerChannelName   string      `json:"ownerChannelName"`
	UploadDate         string      `json:"uploadDate"`
}

type Embed struct {
	IframeURL string `json:"iframeUrl"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
}

type PlayabilityStatus struct {
	Status          string     `json:"status"`
	PlayableInEmbed bool       `json:"playableInEmbed"`
	Miniplayer      Miniplayer `json:"miniplayer"`
	ContextParams   string     `json:"contextParams"`
}

type Miniplayer struct {
	MiniplayerRenderer MiniplayerRenderer `json:"miniplayerRenderer"`
}

type MiniplayerRenderer struct {
	PlaybackMode string `json:"playbackMode"`
}

type PlaybackTracking struct {
	VideostatsPlaybackURL                   PtrackingURLClass `json:"videostatsPlaybackUrl"`
	VideostatsDelayplayURL                  PtrackingURLClass `json:"videostatsDelayplayUrl"`
	VideostatsWatchtimeURL                  PtrackingURLClass `json:"videostatsWatchtimeUrl"`
	PtrackingURL                            PtrackingURLClass `json:"ptrackingUrl"`
	QoeURL                                  PtrackingURLClass `json:"qoeUrl"`
	AtrURL                                  AtrURLClass       `json:"atrUrl"`
	VideostatsScheduledFlushWalltimeSeconds []int64           `json:"videostatsScheduledFlushWalltimeSeconds"`
	VideostatsDefaultFlushIntervalSeconds   int64             `json:"videostatsDefaultFlushIntervalSeconds"`
	YoutubeRemarketingURL                   AtrURLClass       `json:"youtubeRemarketingUrl"`
}

type AtrURLClass struct {
	BaseURL                 string `json:"baseUrl"`
	ElapsedMediaTimeSeconds int64  `json:"elapsedMediaTimeSeconds"`
}

type PtrackingURLClass struct {
	BaseURL string `json:"baseUrl"`
}

type PlayerConfig struct {
	AudioConfig           AudioConfig           `json:"audioConfig"`
	StreamSelectionConfig StreamSelectionConfig `json:"streamSelectionConfig"`
	MediaCommonConfig     MediaCommonConfig     `json:"mediaCommonConfig"`
	WebPlayerConfig       WebPlayerConfig       `json:"webPlayerConfig"`
}

type AudioConfig struct {
	LoudnessDB              float64 `json:"loudnessDb"`
	PerceptualLoudnessDB    float64 `json:"perceptualLoudnessDb"`
	EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
}

type MediaCommonConfig struct {
	DynamicReadaheadConfig DynamicReadaheadConfig `json:"dynamicReadaheadConfig"`
}

type DynamicReadaheadConfig struct {
	MaxReadAheadMediaTimeMS int64 `json:"maxReadAheadMediaTimeMs"`
	MinReadAheadMediaTimeMS int64 `json:"minReadAheadMediaTimeMs"`
	ReadAheadGrowthRateMS   int64 `json:"readAheadGrowthRateMs"`
}

type StreamSelectionConfig struct {
	MaxBitrate string `json:"maxBitrate"`
}

type WebPlayerConfig struct {
	UseCobaltTvosDash       bool                    `json:"useCobaltTvosDash"`
	WebPlayerActionsPorting WebPlayerActionsPorting `json:"webPlayerActionsPorting"`
}

type WebPlayerActionsPorting struct {
	GetSharePanelCommand        GetSharePanelCommand        `json:"getSharePanelCommand"`
	SubscribeCommand            SubscribeCommand            `json:"subscribeCommand"`
	UnsubscribeCommand          UnsubscribeCommand          `json:"unsubscribeCommand"`
	AddToWatchLaterCommand      AddToWatchLaterCommand      `json:"addToWatchLaterCommand"`
	RemoveFromWatchLaterCommand RemoveFromWatchLaterCommand `json:"removeFromWatchLaterCommand"`
}

type AddToWatchLaterCommand struct {
	ClickTrackingParams  string                                     `json:"clickTrackingParams"`
	CommandMetadata      ImpressionEndpointCommandMetadata          `json:"commandMetadata"`
	PlaylistEditEndpoint AddToWatchLaterCommandPlaylistEditEndpoint `json:"playlistEditEndpoint"`
}

type AddToWatchLaterCommandPlaylistEditEndpoint struct {
	PlaylistID string         `json:"playlistId"`
	Actions    []PurpleAction `json:"actions"`
}

type PurpleAction struct {
	AddedVideoID string `json:"addedVideoId"`
	Action       string `json:"action"`
}

type GetSharePanelCommand struct {
	ClickTrackingParams                 string                              `json:"clickTrackingParams"`
	CommandMetadata                     ImpressionEndpointCommandMetadata   `json:"commandMetadata"`
	WebPlayerShareEntityServiceEndpoint WebPlayerShareEntityServiceEndpoint `json:"webPlayerShareEntityServiceEndpoint"`
}

type WebPlayerShareEntityServiceEndpoint struct {
	SerializedShareEntity string `json:"serializedShareEntity"`
}

type RemoveFromWatchLaterCommand struct {
	ClickTrackingParams  string                                          `json:"clickTrackingParams"`
	CommandMetadata      ImpressionEndpointCommandMetadata               `json:"commandMetadata"`
	PlaylistEditEndpoint RemoveFromWatchLaterCommandPlaylistEditEndpoint `json:"playlistEditEndpoint"`
}

type RemoveFromWatchLaterCommandPlaylistEditEndpoint struct {
	PlaylistID string         `json:"playlistId"`
	Actions    []FluffyAction `json:"actions"`
}

type FluffyAction struct {
	Action         string `json:"action"`
	RemovedVideoID string `json:"removedVideoId"`
}

type SubscribeCommand struct {
	ClickTrackingParams string                            `json:"clickTrackingParams"`
	CommandMetadata     ImpressionEndpointCommandMetadata `json:"commandMetadata"`
	SubscribeEndpoint   SubscribeEndpoint                 `json:"subscribeEndpoint"`
}

type SubscribeEndpoint struct {
	ChannelIDS []string `json:"channelIds"`
	Params     string   `json:"params"`
}

type UnsubscribeCommand struct {
	ClickTrackingParams string                            `json:"clickTrackingParams"`
	CommandMetadata     ImpressionEndpointCommandMetadata `json:"commandMetadata"`
	UnsubscribeEndpoint SubscribeEndpoint                 `json:"unsubscribeEndpoint"`
}

type ResponseContext struct {
	VisitorData                     string                          `json:"visitorData"`
	ServiceTrackingParams           []ServiceTrackingParam          `json:"serviceTrackingParams"`
	MaxAgeSeconds                   int64                           `json:"maxAgeSeconds"`
	MainAppWebResponseContext       MainAppWebResponseContext       `json:"mainAppWebResponseContext"`
	WebResponseContextExtensionData WebResponseContextExtensionData `json:"webResponseContextExtensionData"`
}

type MainAppWebResponseContext struct {
	LoggedOut     bool   `json:"loggedOut"`
	TrackingParam string `json:"trackingParam"`
}

type ServiceTrackingParam struct {
	Service string  `json:"service"`
	Params  []Param `json:"params"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type WebResponseContextExtensionData struct {
	HasDecorated bool `json:"hasDecorated"`
}

type Storyboards struct {
	PlayerStoryboardSpecRenderer PlayerStoryboardSpecRenderer `json:"playerStoryboardSpecRenderer"`
}

type PlayerStoryboardSpecRenderer struct {
	Spec                           string `json:"spec"`
	RecommendedLevel               int64  `json:"recommendedLevel"`
	HighResolutionRecommendedLevel int64  `json:"highResolutionRecommendedLevel"`
}

type StreamingData struct {
	ExpiresInSeconds string   `json:"expiresInSeconds"`
	Formats          []Format `json:"formats"`
	AdaptiveFormats  []Format `json:"adaptiveFormats"`
}

type Format struct {
	Itag             int64          `json:"itag"`
	URL              string         `json:"url"`
	MIMEType         string         `json:"mimeType"`
	Bitrate          int64          `json:"bitrate"`
	Width            *int64         `json:"width,omitempty"`
	Height           *int64         `json:"height,omitempty"`
	InitRange        *Range         `json:"initRange,omitempty"`
	IndexRange       *Range         `json:"indexRange,omitempty"`
	LastModified     string         `json:"lastModified"`
	ContentLength    string         `json:"contentLength"`
	Quality          Quality        `json:"quality"`
	FPS              *int64         `json:"fps,omitempty"`
	QualityLabel     *string        `json:"qualityLabel,omitempty"`
	ProjectionType   ProjectionType `json:"projectionType"`
	AverageBitrate   int64          `json:"averageBitrate"`
	ApproxDurationMS string         `json:"approxDurationMs"`
	ColorInfo        *ColorInfo     `json:"colorInfo,omitempty"`
	HighReplication  *bool          `json:"highReplication,omitempty"`
	AudioQuality     *string        `json:"audioQuality,omitempty"`
	AudioSampleRate  *string        `json:"audioSampleRate,omitempty"`
	AudioChannels    *int64         `json:"audioChannels,omitempty"`
	LoudnessDB       *float64       `json:"loudnessDb,omitempty"`
}

type ColorInfo struct {
	TransferCharacteristics string `json:"transferCharacteristics"`
}

type Range struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type VideoDetails struct {
	VideoID           string    `json:"videoId"`
	Title             string    `json:"title"`
	LengthSeconds     string    `json:"lengthSeconds"`
	Keywords          []string  `json:"keywords"`
	ChannelID         string    `json:"channelId"`
	IsOwnerViewing    bool      `json:"isOwnerViewing"`
	ShortDescription  string    `json:"shortDescription"`
	IsCrawlable       bool      `json:"isCrawlable"`
	Thumbnail         IconClass `json:"thumbnail"`
	AllowRatings      bool      `json:"allowRatings"`
	ViewCount         string    `json:"viewCount"`
	Author            string    `json:"author"`
	IsPrivate         bool      `json:"isPrivate"`
	IsUnpluggedCorpus bool      `json:"isUnpluggedCorpus"`
	IsLiveContent     bool      `json:"isLiveContent"`
}

type ProjectionType string

const (
	Rectangular ProjectionType = "RECTANGULAR"
)

type Quality string

const (
	Small Quality = "small"
	Tiny  Quality = "tiny"
)
