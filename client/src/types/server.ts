export interface IUser {
    "id": number
    "created_at": Date,
    "updated_at": Date,
    "deleted_at": Date | null,
    "email": string,
    "role": "admin" | "student" | "teacher" | "supervisor",
    "name": string,
    "profile_img_url": string,
    "institution_id"?: number,
}

export interface IDefaultRequest {
    data: null | IUser,
    error: null | string,
}

export interface IInstitutions {
    id: number,
    
    name: string,
    segment: string,
    series: string,
    
    created_at: Date,
    deleted_at: Date | null,
    updated_at: Date,
}