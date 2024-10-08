namespace go launch_screen
include"model.thrift"

struct CreateImageRequest {
    1: required i64 pic_type,//1为空，2为页面跳转，3为app跳转
    2: optional i64 duration,
    3: optional string href,//连接
    4: required binary image,
    5: required i64 start_at,
    6: required i64 end_at,
    7: required i64 s_type,
    8: required i64 frequency,//单日最大展示次数
    9: required i64 start_time,//比如6表示6点
    10:required i64 end_time,
    11:required string text,//描述图片
    12:string regex,//正则匹配项

    13:string token,//get by token
}

struct CreateImageResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

struct GetImageRequest{
    1:required i64 picture_id,

    2:string token,
}

struct GetImageResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

struct GetImagesByUserIdRequest{
    1:string token,
}

struct GetImagesByUserIdResponse{
    1:model.BaseResp base,
    2:optional list<model.Picture> picture_list,
}

struct ChangeImagePropertyRequest {
    1: required i64 pic_type,//1为空，2为页面跳转，3为app跳转
    2: optional i64 duration,
    3: optional string href,//连接
    4: required i64 start_at,
    5: required i64 end_at,
    6: required i64 s_type,
    7: required i64 frequency,
    8: required i64 start_time,//比如6表示6点
    9:required i64 end_time,
    10:required string text,//描述图片
    11:string regex,//正则匹配项

    12:string token,//get by token
}

struct ChangeImagePropertyResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

struct ChangeImageRequest {
    1:required i64 picture_id,
    2:required binary image,

    3:string token,
}

struct ChangeImageResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

struct DeleteImageRequest{
    1:required i64 picture_id,

    2:string token,
}

struct DeleteImageResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

struct MobileGetImageRequest{
    1:required i64 type,
    2:required i64 student_id,
    3:required string college,
    4:required string device,
}

struct MobileGetImageResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

struct AddImagePointTimeRequest{
    1:required i64 picture_id,
}

struct AddImagePointTimeResponse{
    1:model.BaseResp base,
    2:optional model.Picture picture,
}

service LaunchScreenService{
    CreateImageResponse CreateImage(1:CreateImageRequest req)(api.post="/launch_screen/api/image"),
    GetImagesByUserIdResponse GetImage(1:GetImageRequest req)(api.get="/launch_screen/api/image"),
    GetImagesByUserIdResponse GetImagesByUserId(1:GetImagesByUserIdRequest req)(api.get="/launch_screen/api/images"),
    ChangeImagePropertyResponse ChangeImageProperty(1:ChangeImagePropertyRequest req)(api.put="/launch_screen/api/image"),
    ChangeImageResponse ChangeImage(1:ChangeImageRequest req)(api.put="/launch_screen/api/image/img"),
    DeleteImageResponse DeleteImage(1:DeleteImageRequest req)(api.delete="/launch_screen/api/image"),
    MobileGetImageResponse MobileGetImage(1:MobileGetImageRequest req)(api.get="/launch_screen/api/screen"),
    AddImagePointTimeResponse AddImagePointTime(1:AddImagePointTimeRequest req)(api.get="/launch_screen/api/image/point"),
}
