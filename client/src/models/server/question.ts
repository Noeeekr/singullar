import { DefaultFields } from ".";

export interface QuestionListDifficulty {
    question_difficulty_level: number
    question_difficulty_name: string
}

export interface Question extends DefaultFields {
    status: string // [Incomplete, Completed, Started]
    question_title: string
    question_short_description: string
    question_description: string
    question_difficulty_level: number
    question_institution_id: number
    difficulty_level: number
}