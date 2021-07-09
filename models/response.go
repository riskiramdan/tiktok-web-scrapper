package models

type Data struct {
	Props Props `json:"props"`
}

type InitialProps struct {
	StatusCode         int    `json:"statusCode"`
	Wid                string `json:"$wid"`
	Host               string `json:"$host"`
	PageURL            string `json:"$pageUrl"`
	FullURL            string `json:"$fullUrl"`
	IsMobile           bool   `json:"$isMobile"`
	IsAndroid          bool   `json:"$isAndroid"`
	IsIOS              bool   `json:"$isIOS"`
	NeedFix            bool   `json:"$needFix"`
	RequestID          string `json:"$requestId"`
	CsrfToken          string `json:"$csrfToken"`
	LogID              string `json:"$logId"`
	Region             string `json:"$region"`
	AppType            string `json:"$appType"`
	ReflowType         string `json:"$reflowType"`
	SgOpen             bool   `json:"$sgOpen"`
	BaseURL            string `json:"$baseURL"`
	ShowLiveSuggest    bool   `json:"$showLiveSuggest"`
	LiveDiscoverSwitch bool   `json:"$liveDiscoverSwitch"`
	Nonce              string `json:"nonce"`
}

type Video struct {
	ID            string   `json:"id"`
	Height        int      `json:"height"`
	Width         int      `json:"width"`
	Duration      int      `json:"duration"`
	Ratio         string   `json:"ratio"`
	Cover         string   `json:"cover"`
	OriginCover   string   `json:"originCover"`
	DynamicCover  string   `json:"dynamicCover"`
	PlayAddr      string   `json:"playAddr"`
	DownloadAddr  string   `json:"downloadAddr"`
	ShareCover    []string `json:"shareCover"`
	ReflowCover   string   `json:"reflowCover"`
	Bitrate       int      `json:"bitrate"`
	EncodedType   string   `json:"encodedType"`
	Format        string   `json:"format"`
	VideoQuality  string   `json:"videoQuality"`
	EncodeUserTag string   `json:"encodeUserTag"`
}

type Author struct {
	ID             string `json:"id"`
	ShortID        string `json:"shortId"`
	UniqueID       string `json:"uniqueId"`
	Nickname       string `json:"nickname"`
	AvatarLarger   string `json:"avatarLarger"`
	AvatarMedium   string `json:"avatarMedium"`
	AvatarThumb    string `json:"avatarThumb"`
	Signature      string `json:"signature"`
	CreateTime     int    `json:"createTime"`
	Verified       bool   `json:"verified"`
	SecUID         string `json:"secUid"`
	Ftc            bool   `json:"ftc"`
	Relation       int    `json:"relation"`
	OpenFavorite   bool   `json:"openFavorite"`
	CommentSetting int    `json:"commentSetting"`
	DuetSetting    int    `json:"duetSetting"`
	StitchSetting  int    `json:"stitchSetting"`
	PrivateAccount bool   `json:"privateAccount"`
	Secret         bool   `json:"secret"`
	IsADVirtual    bool   `json:"isADVirtual"`
	RoomID         string `json:"roomId"`
}

type Music struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	PlayURL            string `json:"playUrl"`
	CoverLarge         string `json:"coverLarge"`
	CoverMedium        string `json:"coverMedium"`
	CoverThumb         string `json:"coverThumb"`
	AuthorName         string `json:"authorName"`
	Original           bool   `json:"original"`
	Duration           int    `json:"duration"`
	Album              string `json:"album"`
	ScheduleSearchTime int    `json:"scheduleSearchTime"`
}

type VideoStats struct {
	DiggCount    int `json:"diggCount"`
	ShareCount   int `json:"shareCount"`
	CommentCount int `json:"commentCount"`
	PlayCount    int `json:"playCount"`
}

type DuetInfo struct {
	DuetFromID string `json:"duetFromId"`
}

type AuthorStats struct {
	FollowerCount  int `json:"followerCount"`
	FollowingCount int `json:"followingCount"`
	Heart          int `json:"heart"`
	HeartCount     int `json:"heartCount"`
	VideoCount     int `json:"videoCount"`
	DiggCount      int `json:"diggCount"`
}

type ItemStruct struct {
	ID             string       `json:"id"`
	Desc           string       `json:"desc"`
	CreateTime     int          `json:"createTime"`
	ScheduleTime   int          `json:"scheduleTime"`
	Video          *Video       `json:"video"`
	Author         *Author      `json:"author"`
	Music          *Music       `json:"music"`
	Stats          *VideoStats  `json:"stats"`
	IsActivityItem bool         `json:"isActivityItem"`
	DuetInfo       *DuetInfo    `json:"duetInfo"`
	AuthorStats    *AuthorStats `json:"authorStats"`
}

type ItemInfo struct {
	ItemStruct *ItemStruct `json:"itemStruct"`
}

type PageProps struct {
	Key        string    `json:"key"`
	ServerCode int       `json:"serverCode"`
	StatusCode int       `json:"statusCode"`
	StatusMsg  string    `json:"statusMsg"`
	ItemInfo   *ItemInfo `json:"itemInfo"`
}

type Props struct {
	InitialProps *InitialProps `json:"initialProps"`
	PageProps    *PageProps    `json:"pageProps"`
}

type TiktokResponseData struct {
	Author      *Author      `json:"author"`
	AuthorStats *AuthorStats `json:"authorStats"`
	Stats       *VideoStats  `json:"stats"`
	Video       *Video       `json:"video"`
}
