import { DefaultFields } from ".";

export interface QuestionListDifficulty {
    difficulty_level: number
    difficulty_name: string
}

export interface Question extends DefaultFields {
    status: string // [Incomplete, Completed, Started]
    title: string
    description: string
    
    difficulty_level: number
}