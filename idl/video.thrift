namespace go video
include "user.thrift"
struct Video {
    1:  i64 id
    2:  user.User author
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

struct VideoRequest{
    1:  i64 user_id
    2:  i64 video_id
}

struct VideoResponse{
    1:  i32 status_code
    2:  string status_msg
    3:  Video video
}

service VideoService{
    FeedResponse GetFeed(1: FeedRequest req)
    PublishActionResponse PublishAction(1:PublishActionRequest req)
    PublishListResponse GetPublishList(1:PublishListRequest req)
    VideoResponse GetVideo(1:VideoRequest req)
}