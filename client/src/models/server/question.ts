import { DefaultFields } from ".";

export interface QuestionListDifficulty {
    question_difficulty_level: number
    question_difficulty_name: string
}

export interface Question extends DefaultFields {
    status: string // [Incomplete, Completed, Started]
    title: string
    description: string
    
    difficulty_level: number
}