export interface IUser {
    "id": number
    "created_at": Date,
    "updated_at": Date,
    "deleted_at": Date | null,
    "email": string,
    "role": "institution" | "student" | "teacher",
    "name": string,
    "profile_img_url": string,
    "institution_id"?: number,
}

export interface IUserRequest {
    data: null | IUser,
    error: null | string,
}