namespace go video
include "user.thrift"
struct Video {
    1: required i64 id
    2: required user.User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}

struct FeedRequest {
    1: optional i64 latest_time
    2: optional string token
}

struct FeedResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Video> video_list
}

struct PublishActionRequest {
    1: required string token
    2: required string data
    3: required string title
}

struct PublishActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct PublishListRequest {
    1: required i64 user_id
    2: required string token
}

struct PublishListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Video> video_list
}

service VideoService{
    FeedResponse GetFeed(1: FeedRequest req)
    PublishActionResponse PublishAction(1:PublishActionRequest req)
    PublishListResponse GetPublishList(1:PublishListRequest req)
    
}