package tsubato

// CatalogPost https://github.com/4chan/4chan-API/blob/master/pages/Catalog.md
type CatalogPost struct {
	Post
	SemanticURL   string `json:"semantic_url"`
	Replies       int    `json:"replies"`
	Images        int    `json:"images"`
	OmittedPosts  int    `json:"omitted_posts"`
	OmittedImages int    `json:"omitted_images"`
	LastReplies   []Post `json:"last_replies"`
	LastModified  int    `json:"last_modified"`
	Bumplimit     int    `json:"bumplimit,omitempty"`
	Imagelimit    int    `json:"imagelimit,omitempty"`
}

// Post is the common fields between the post representations
// of all API responses.
type Post struct {
	ID              int    `json:"no"`
	Now             string `json:"now"`
	Name            string `json:"name"`
	Comment         string `json:"com"`
	Filename        string `json:"filename"`
	Ext             string `json:"ext"`
	Width           int    `json:"w"`
	Height          int    `json:"h"`
	ThumbnailWidth  int    `json:"tn_w"`
	ThumbnailHeight int    `json:"tn_h"`
	ImageUploadTime int64  `json:"tim"`
	UnixTime        int    `json:"time"`
	Md5             string `json:"md5"`
	Filesize        int    `json:"fsize"`
	ReplyingTo      int    `json:"resto"`
	Capcode         string `json:"capcode"`
}

// ThreadPost https://github.com/4chan/4chan-API/blob/master/pages/Threads.md
type ThreadPost struct {
	Post
	SemanticURL string `json:"semantic_url,omitempty"`
	LastReplies []Post
	Replies     int `json:"replies,omitempty"`
	Images      int `json:"images,omitempty"`
	UniqueIps   int `json:"unique_ips,omitempty"`
}

// Board https://github.com/4chan/4chan-API/blob/master/pages/Boards.md
type Board struct {
	Board           string `json:"board"`
	Title           string `json:"title"`
	Worksafe        int    `json:"ws_board"`
	PerPage         int    `json:"per_page"`
	Pages           int    `json:"pages"`
	MaxFilesize     int    `json:"max_filesize"`
	MaxWebmFilesize int    `json:"max_webm_filesize"`
	MaxCommentChars int    `json:"max_comment_chars"`
	MaxWebmDuration int    `json:"max_webm_duration"`
	BumpLimit       int    `json:"bump_limit"`
	ImageLimit      int    `json:"image_limit"`
	MetaDescription string `json:"meta_description"`
	Spoilers        int    `json:"spoilers,omitempty"`
	CustomSpoilers  int    `json:"custom_spoilers,omitempty"`
	IsArchived      int    `json:"is_archived,omitempty"`
	ForcedAnon      int    `json:"forced_anon,omitempty"`
}
