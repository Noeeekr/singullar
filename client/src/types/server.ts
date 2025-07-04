export interface IUser {
    "id": number
    "created_at": Date,
    "updated_at": Date,
    "deleted_at": Date | null,
    "email": string,
    "role": IUserRoles,
    "name": string,
    "profile_picture": string,
    "institution_id"?: number,
}

export type IUserRoles = "admin" | "student" | "teacher" | "supervisor";

export interface IDefaultRequest {
    data: null | IUser,
    error: null | string,
}

export interface IInstitutions {
    id: number,
    
    name: string,
    
    created_at: Date,
    deleted_at: Date | null,
    updated_at: Date,
}