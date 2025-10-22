import { DefaultFields } from ".";

export interface Question extends DefaultFields {
    status: string // [Incomplete, Completed, Started]
    title: string
    description: string
    
    difficultyLevel: number
    difficultyName: string
}