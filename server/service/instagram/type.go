package instagram

type Response struct {
	EntryData    struct {
		ProfilePage []struct {
			SeoCategoryInfos      [][]string `json:"seo_category_infos"`
			LoggingPageId         string     `json:"logging_page_id"`
			ShowSuggestedProfiles bool       `json:"show_suggested_profiles"`
			ShowFollowDialog      bool       `json:"show_follow_dialog"`
			Graphql               struct {
				User struct {
					EdgeFollowedBy         struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					FollowsViewer         bool        `json:"follows_viewer"`
					FullName              string      `json:"full_name"`
					HasArEffects          bool        `json:"has_ar_effects"`
					HasClips              bool        `json:"has_clips"`
					HasGuides             bool        `json:"has_guides"`
					HasChannel            bool        `json:"has_channel"`
					HasBlockedViewer      bool        `json:"has_blocked_viewer"`
					HighlightReelCount    int         `json:"highlight_reel_count"`
					HasRequestedViewer    bool        `json:"has_requested_viewer"`
					HideLikeAndViewCounts bool        `json:"hide_like_and_view_counts"`
					Id                    string      `json:"id"`
					IsBusinessAccount     bool        `json:"is_business_account"`
					IsProfessionalAccount bool        `json:"is_professional_account"`
					IsJoinedRecently      bool        `json:"is_joined_recently"`
					BusinessAddressJson   interface{} `json:"business_address_json"`
					BusinessContactMethod interface{} `json:"business_contact_method"`
					BusinessEmail         interface{} `json:"business_email"`
					BusinessPhoneNumber   interface{} `json:"business_phone_number"`
					BusinessCategoryName  string      `json:"business_category_name"`
					OverallCategoryName   interface{} `json:"overall_category_name"`
					CategoryEnum          string      `json:"category_enum"`
					CategoryName          string      `json:"category_name"`
					IsPrivate             bool        `json:"is_private"`
					IsVerified            bool        `json:"is_verified"`
					EdgeMutualFollowedBy  struct {
						Count int           `json:"count"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_mutual_followed_by"`
					ProfilePicUrl                string        `json:"profile_pic_url"`
					ProfilePicUrlHd              string        `json:"profile_pic_url_hd"`
					RequestedByViewer            bool          `json:"requested_by_viewer"`
					ShouldShowCategory           bool          `json:"should_show_category"`
					ShouldShowPublicContacts     bool          `json:"should_show_public_contacts"`
					Username                     string        `json:"username"`
					ConnectedFbPage              interface{}   `json:"connected_fb_page"`
					Pronouns                     []interface{} `json:"pronouns"`
					EdgeFelixCombinedPostUploads struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_felix_combined_post_uploads"`
					EdgeOwnerToTimelineMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename   string `json:"__typename"`
								Id         string `json:"id"`
								Shortcode  string `json:"shortcode"`
								Dimensions struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayUrl            string `json:"display_url"`
								EdgeMediaToTaggedUser struct {
									Edges []interface{} `json:"edges"`
								} `json:"edge_media_to_tagged_user"`
								FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
								FactCheckInformation   interface{} `json:"fact_check_information"`
								GatingInfo             interface{} `json:"gating_info"`
								SharingFrictionInfo    struct {
									ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
									BloksAppUrl               interface{} `json:"bloks_app_url"`
								} `json:"sharing_friction_info"`
								MediaOverlayInfo interface{} `json:"media_overlay_info"`
								MediaPreview     *string     `json:"media_preview"`
								Owner            struct {
									Id       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								IsVideo              bool   `json:"is_video"`
								HasUpcomingEvent     bool   `json:"has_upcoming_event"`
								AccessibilityCaption string `json:"accessibility_caption"`
								EdgeMediaToCaption   struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								EdgeLikedBy      struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								Location *struct {
									Id            string `json:"id"`
									HasPublicPage bool   `json:"has_public_page"`
									Name          string `json:"name"`
									Slug          string `json:"slug"`
								} `json:"location"`
								ThumbnailSrc       string `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								CoauthorProducers     []interface{} `json:"coauthor_producers"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}