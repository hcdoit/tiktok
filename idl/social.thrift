namespace go social
include "user.thrift"

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
    3: list<user.User> user_list
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

struct Message{
    1: i64 id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string content
    5: string create_time
}

struct MessageChatRequest{
    1: string token
    2: i64 to_user_id
}

struct MessageChatResponse{
    1:  i32 status_code
    2:  string status_msg
    3:  list<Message> message_list
}

struct MessageActionRequest{
    1: string token
    2: i64 to_user_id
    3: i32 action_type
    4: string content
}

struct MessageActionResponse{
    1:  i32 status_code
    2:  string status_msg
}



service SocialService{
    RelationActionResponse RelationAction(1:RelationActionRequest req)
    RelationListResponse GetFollowList(1:RelationListRequest req)
    RelationListResponse GetFollowerList(1:RelationListRequest req)
    RelationListResponse GetFriendList(1:RelationListRequest req)
    RelationInfoResponse GetRelationInfo(1:RelationInfoRequest req)
    MessageChatResponse GetMessageChat(1:MessageChatRequest req)
    MessageActionResponse MessageAction(1:MessageActionRequest req)

}

