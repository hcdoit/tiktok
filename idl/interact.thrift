namespace go interact
include "video.thrift"
include "user.thrift"

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
    3:  list<video.Video> video_list
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
    2:  user.User user
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

struct VideoInteractRequest {
    1:  i64 user_id
    2:  i64 video_id
}

struct VideoInteractResponse{
    1:  i32 status_code
    2:  string status_msg
    3:  i64 favorite_count
    4:  i64 comment_count
    5:  bool is_favorite
    
}

service InteractService {
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    FavoriteListResponse GetFavoriteList(1:FavoriteListRequest req)
    CommentActionResponse CommentAction(1:CommentActionRequest req)
    CommentListResponse GetCommentList(1:CommentListRequest req)
    VideoInteractResponse GetVideoInteract(1:VideoInteractRequest req)
}