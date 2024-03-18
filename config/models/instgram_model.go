// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    instagramPayload, err := UnmarshalInstagramPayload(bytes)
//    bytes, err = instagramPayload.Marshal()

package main

import "encoding/json"

func UnmarshalInstagramPayload(data []byte) (InstagramPayload, error) {
	var r InstagramPayload
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramPayload) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InstagramPayload struct {
	Graphql     Graphql `json:"graphql"`
	ShowQRModal bool    `json:"showQRModal"`
}

type Graphql struct {
	ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
}

type ShortcodeMedia struct {
	Typename                    string                        `json:"__typename"`
	ID                          string                        `json:"id"`
	Shortcode                   string                        `json:"shortcode"`
	Dimensions                  Dimensions                    `json:"dimensions"`
	GatingInfo                  interface{}                   `json:"gating_info"`
	FactCheckOverallRating      interface{}                   `json:"fact_check_overall_rating"`
	FactCheckInformation        interface{}                   `json:"fact_check_information"`
	SensitivityFrictionInfo     interface{}                   `json:"sensitivity_friction_info"`
	SharingFrictionInfo         SharingFrictionInfo           `json:"sharing_friction_info"`
	MediaOverlayInfo            interface{}                   `json:"media_overlay_info"`
	MediaPreview                string                        `json:"media_preview"`
	DisplayURL                  string                        `json:"display_url"`
	DisplayResources            []DisplayResource             `json:"display_resources"`
	AccessibilityCaption        interface{}                   `json:"accessibility_caption"`
	DashInfo                    DashInfo                      `json:"dash_info"`
	HasAudio                    bool                          `json:"has_audio"`
	VideoURL                    string                        `json:"video_url"`
	VideoViewCount              int64                         `json:"video_view_count"`
	VideoPlayCount              interface{}                   `json:"video_play_count"`
	IsVideo                     bool                          `json:"is_video"`
	TrackingToken               string                        `json:"tracking_token"`
	UpcomingEvent               interface{}                   `json:"upcoming_event"`
	EdgeMediaToTaggedUser       EdgeMediaToCaptionClass       `json:"edge_media_to_tagged_user"`
	EdgeMediaToCaption          EdgeMediaToCaptionClass       `json:"edge_media_to_caption"`
	CanSeeInsightsAsBrand       bool                          `json:"can_see_insights_as_brand"`
	CaptionIsEdited             bool                          `json:"caption_is_edited"`
	HasRankedComments           bool                          `json:"has_ranked_comments"`
	LikeAndViewCountsDisabled   bool                          `json:"like_and_view_counts_disabled"`
	EdgeMediaToParentComment    EdgeMediaToParentCommentClass `json:"edge_media_to_parent_comment"`
	EdgeMediaToHoistedComment   EdgeMediaToCaptionClass       `json:"edge_media_to_hoisted_comment"`
	EdgeMediaPreviewComment     EdgeMediaPreview              `json:"edge_media_preview_comment"`
	CommentsDisabled            bool                          `json:"comments_disabled"`
	CommentingDisabledForViewer bool                          `json:"commenting_disabled_for_viewer"`
	TakenAtTimestamp            int64                         `json:"taken_at_timestamp"`
	EdgeMediaPreviewLike        EdgeMediaPreview              `json:"edge_media_preview_like"`
	EdgeMediaToSponsorUser      EdgeMediaToCaptionClass       `json:"edge_media_to_sponsor_user"`
	IsAffiliate                 bool                          `json:"is_affiliate"`
	IsPaidPartnership           bool                          `json:"is_paid_partnership"`
	Location                    interface{}                   `json:"location"`
	NftAssetInfo                interface{}                   `json:"nft_asset_info"`
	ViewerHasLiked              bool                          `json:"viewer_has_liked"`
	ViewerHasSaved              bool                          `json:"viewer_has_saved"`
	ViewerHasSavedToCollection  bool                          `json:"viewer_has_saved_to_collection"`
	ViewerInPhotoOfYou          bool                          `json:"viewer_in_photo_of_you"`
	ViewerCanReshare            bool                          `json:"viewer_can_reshare"`
	Owner                       ShortcodeMediaOwner           `json:"owner"`
	IsAd                        bool                          `json:"is_ad"`
	EdgeWebMediaToRelatedMedia  EdgeMediaToCaptionClass       `json:"edge_web_media_to_related_media"`
	CoauthorProducers           []interface{}                 `json:"coauthor_producers"`
	PinnedForUsers              []interface{}                 `json:"pinned_for_users"`
	EncodingStatus              interface{}                   `json:"encoding_status"`
	IsPublished                 bool                          `json:"is_published"`
	ProductType                 string                        `json:"product_type"`
	Title                       string                        `json:"title"`
	VideoDuration               float64                       `json:"video_duration"`
	ThumbnailSrc                string                        `json:"thumbnail_src"`
	ClipsMusicAttributionInfo   interface{}                   `json:"clips_music_attribution_info"`
	EdgeRelatedProfiles         EdgeMediaToCaptionClass       `json:"edge_related_profiles"`
}

type DashInfo struct {
	IsDashEligible    bool        `json:"is_dash_eligible"`
	VideoDashManifest interface{} `json:"video_dash_manifest"`
	NumberOfQualities int64       `json:"number_of_qualities"`
}

type Dimensions struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
}

type DisplayResource struct {
	Src          string `json:"src"`
	ConfigWidth  int64  `json:"config_width"`
	ConfigHeight int64  `json:"config_height"`
}

type EdgeMediaPreview struct {
	Count int64                         `json:"count"`
	Edges []EdgeMediaPreviewCommentEdge `json:"edges"`
}

type EdgeMediaToParentCommentClass struct {
	Count    int64                         `json:"count"`
	PageInfo PageInfo                      `json:"page_info"`
	Edges    []EdgeMediaPreviewCommentEdge `json:"edges"`
}

type PurpleNode struct {
	ID                   string                         `json:"id"`
	Text                 string                         `json:"text"`
	CreatedAt            int64                          `json:"created_at"`
	DidReportAsSpam      bool                           `json:"did_report_as_spam"`
	Owner                NodeOwner                      `json:"owner"`
	ViewerHasLiked       bool                           `json:"viewer_has_liked"`
	EdgeLikedBy          EdgeFollowedByClass            `json:"edge_liked_by"`
	IsRestrictedPending  bool                           `json:"is_restricted_pending"`
	EdgeThreadedComments *EdgeMediaToParentCommentClass `json:"edge_threaded_comments,omitempty"`
}

type EdgeMediaPreviewCommentEdge struct {
	Node PurpleNode `json:"node"`
}

type PageInfo struct {
	HasNextPage bool        `json:"has_next_page"`
	EndCursor   interface{} `json:"end_cursor"`
}

type EdgeFollowedByClass struct {
	Count int64 `json:"count"`
}

type NodeOwner struct {
	ID            string `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	ProfilePicURL string `json:"profile_pic_url"`
	Username      string `json:"username"`
}

type EdgeMediaToCaptionClass struct {
	Edges []EdgeMediaToCaptionEdge `json:"edges"`
}

type EdgeMediaToCaptionEdge struct {
	Node FluffyNode `json:"node"`
}

type FluffyNode struct {
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
}

type ShortcodeMediaOwner struct {
	ID                        string              `json:"id"`
	IsVerified                bool                `json:"is_verified"`
	ProfilePicURL             string              `json:"profile_pic_url"`
	Username                  string              `json:"username"`
	BlockedByViewer           bool                `json:"blocked_by_viewer"`
	RestrictedByViewer        interface{}         `json:"restricted_by_viewer"`
	FollowedByViewer          bool                `json:"followed_by_viewer"`
	FullName                  string              `json:"full_name"`
	HasBlockedViewer          bool                `json:"has_blocked_viewer"`
	IsEmbedsDisabled          bool                `json:"is_embeds_disabled"`
	IsPrivate                 bool                `json:"is_private"`
	IsUnpublished             bool                `json:"is_unpublished"`
	RequestedByViewer         bool                `json:"requested_by_viewer"`
	PassTieringRecommendation bool                `json:"pass_tiering_recommendation"`
	EdgeOwnerToTimelineMedia  EdgeFollowedByClass `json:"edge_owner_to_timeline_media"`
	EdgeFollowedBy            EdgeFollowedByClass `json:"edge_followed_by"`
}

type SharingFrictionInfo struct {
	ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
	BloksAppURL               interface{} `json:"bloks_app_url"`
}
