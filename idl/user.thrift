namespace go user

struct UserRegisterRequest {
    1:  string username
    2:  string password
}

struct UserRegisterResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  i64 user_id
    4:  string token
}

struct UserLoginRequest {
    1:  string username
    2:  string password
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
    1:  i64 user_id
    2:  string token
}

struct GetUserResponse {
    1:  i32 status_code
    2:  string status_msg
    3:  User user
}

service UserService {
    UserRegisterResponse Register(1: UserRegisterRequest req)
    UserLoginResponse Login(1: UserLoginRequest req)
    GetUserResponse GetUser(1: GetUserRequest req)
}