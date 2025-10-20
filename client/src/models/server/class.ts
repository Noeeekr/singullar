import type { DefaultFields } from "."
import type { UserSegments } from "./user"

export interface Class extends DefaultFields {
    "name": string,
    "segment": UserSegments,
    "series": string,
    "institution_id": number
    "teacher_id": number
}
