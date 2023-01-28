namespace go user

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct UserRegisterRequest {
    1: required string username
    2: required string password
}

struct UserRegisterResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct UserLoginRequest {
    1: required string username
    2: required string password
}

struct UserLoginResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct GetUserRequest {
    1: required i64 user_id
    2: required string token
}

struct GetUserResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required User user
}

service UserService {
    UserRegisterResponse Register(1: UserRegisterRequest req)
    UserLoginResponse Login(1: UserLoginRequest req)
    GetUserResponse GetUser(1: UserLoginRequest req)
}