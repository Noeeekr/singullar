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
    "role": UserRoles,
    "name": string,
    "profile_picture": string,
    "institution_id"?: number,
}

export const ROLE_ADMIN = "admin"
export const ROLE_STUDENT = "student"
export const ROLE_TEACHER = "teacher"
export const ROLE_SUPERVISOR = "supervisor"
export type UserRoles = typeof ROLE_ADMIN | typeof ROLE_STUDENT | typeof ROLE_TEACHER | typeof ROLE_SUPERVISOR;

