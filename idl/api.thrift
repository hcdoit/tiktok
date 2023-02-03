namespace go api

struct UserRegisterRequest {
    1:  string username (api.query="username")
    2:  string password (api.query="password")
}

struct UserRegisterResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  i64 user_id
    4:  string token
}

struct UserLoginRequest {
    1:  string username (api.query="username")
    2:  string password (api.query="password")
}

struct UserLoginResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  i64 user_id
    4:  string token
}

struct User {
    1:  i64 id
    2:  string name
    3:  i64 follow_count
    4:  i64 follower_count
    5:  bool is_follow
}

struct GetUserRequest {
    1:  i64 user_id (api.query="user_id")
    2:  string token (api.query="token")
}

struct GetUserResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  User user
}
struct Video {
    1:  i64 id
    2:  User author
    3:  string play_url
    4:  string cover_url
    5:  i64 favorite_count
    6:  i64 comment_count
    7:  bool is_favorite
    8:  string title
}

struct FeedRequest {
    1:  i64 latest_time
    2:  string token
}

struct FeedResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  list<Video> video_list
    4:  i64 next_time
}

struct PublishActionRequest {
    1:  string token
    2:  binary data
    3:  string title
}

struct PublishActionResponse {
    1:  i32 status_code
    2:  string status_msg
}

struct PublishListRequest {
    1:  i64 user_id
    2:  string token
}

struct PublishListResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  list<Video> video_list
}
struct FavoriteActionRequest {
    1:  string token
    2:  i64 video_id
    3:  i32 action_type
}

struct FavoriteActionResponse {
    1:  i32 status_code
    2:  string status_msg
}

struct FavoriteListRequest {
    1:  i64 user_id
    2:  string token
}

struct FavoriteListResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  list<Video> video_list
}

struct CommentActionRequest {
    1:  string token
    2:  i64 video_id
    3:  i32 action_type
    4:  string comment_text
    5:  i64 comment_id
}

struct Comment {
    1:  i64 id
    2:  User user
    3:  string content
    4:  string create_date
}

struct CommentActionResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  Comment comment
}

struct CommentListRequest {
    1:  string token
    2:  i64 video_id
}

struct CommentListResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  list<Comment> comment_list
}

struct RelationActionRequest{
    1: string token
    2: i64 to_user_id
    3: i32 action_type
}

struct RelationActionResponse{
    1: i32 status_code
    2: string status_msg
}

struct RelationListRequest{
    1: string token
    2: i64 user_id
}

struct RelationListResponse{
    1: i32 status_code
    2: string status_msg
    3: list<User> user_list
}

struct RelationInfoRequest {
    1:  i64 my_id
    2:  i64 user_id
}

struct RelationInfoResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  i64 follow_count
    4:  i64 follower_count
    5:  bool is_follow
}

service ApiService{
     UserRegisterResponse Register(1: UserRegisterRequest req) (api.post="/douyin/user/register/")
     UserLoginResponse Login(1: UserLoginRequest req) (api.post="/douyin/user/login/")
     GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user/")
     FeedResponse GetFeed(1: FeedRequest req)(api.get="/douyin/feed/")
     PublishActionResponse PublishAction(1:PublishActionRequest req)(api.post="/douyin/publish/action/")
     PublishListResponse GetPublishList(1:PublishListRequest req)(api.get="/douyin/publish/list/")
     FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)(api.post="/douyin/favorite/action/")
     FavoriteListResponse GetFavoriteList(1:FavoriteListRequest req)(api.get="/douyin/favorite/list/")
     CommentActionResponse CommentAction(1:CommentActionRequest req)(api.post="/douyin/comment/action/")
     CommentListResponse GetCommentList(1:CommentListRequest req)(api.get="/douyin/comment/list/")
     RelationActionResponse RelationAction(1:RelationActionRequest req)(api.post="/douyin/relation/action/")
     RelationListResponse GetFollowList(1:RelationListRequest req)(api.get="/douyin/relation/follow/list/")
     RelationListResponse GetFollowerList(1:RelationListRequest req)(api.get="/douyin/relation/follower/list/")
     RelationListResponse GetFriendList(1:RelationListRequest req)(api.get="/douyin/relation/friend/list/")
}