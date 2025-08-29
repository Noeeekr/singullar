// Server Response Objects
export interface DefaultResponse<ResponseData> {
    data: ResponseData,
    error: string,
}

// Default models 
type ID = number
export interface IDField {
    "id": ID
}
export interface DefaultFields extends IDField {
    created_at: Date,
    deleted_at: Date | null,
    updated_at: Date,
}

export interface Institution extends DefaultFields {
    name: string,
}
export interface User extends DefaultFields {
    "email": string,
    "role": IUserRoles,
    "name": string,
    "profile_picture": string,
    "institution_id"?: number,
}

export type IUserRoles = "admin" | "student" | "teacher" | "supervisor";

