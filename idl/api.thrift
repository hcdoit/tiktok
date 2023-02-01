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

service ApiService{
     UserRegisterResponse Register(1: UserRegisterRequest req) (api.post="/douyin/user/register/")
     UserLoginResponse Login(1: UserLoginRequest req) (api.post="/douyin/user/login/")
     GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user/")
     FeedResponse GetFeed(1: FeedRequest req)(api.get="/douyin/feed/")
     PublishActionResponse PublishAction(1:PublishActionRequest req)(api.post="/douyin/publish/action/")
     PublishListResponse GetPublishList(1:PublishListRequest req)(api.get="/douyin/publish/list/")
}