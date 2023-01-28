namespace go interact

struct FavoriteActionRequest {
    1: required string token,
    2: required i64 video_id,
    3: required i32 action_type,
}

struct FavoriteActionResponse {
    1: required i32 status_code,
    2: optional string status_msg,
}

struct FavoriteListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct FavoriteListResponse {
    1: required i32 status_code,
    2: required string status_msg,
    3: required list<video.Video> video_list,
}

struct CommentActionRequest {
    1: required string token,
    2: required i64 video_id,
    3: required i32 action_type,
    4: optional string comment_text,
    5: optional i64 comment_id,
}

struct Comment {
    required i64 id,
    required user.User user,
    required string content,
    required string create_date,
}

struct CommentActionResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: optional Comment comment,
}

struct CommentListRequest {
    1: required string token,
    2: required i64 video_id,
}

struct CommentListResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<Comment> comment_list,
}

service InteractService {
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest)
    FavoriteListResponse GetFavoriteList(1:FavoriteListRequest)
    CommentActionResponse CommentAction(1:CommentActionRequest)
    CommentListResponse GetCommentList(1:CommentListRequest)
}