import { DefaultFields } from ".";

export interface QuestionListDifficulty {
    question_difficulty_name: string
    question_difficulty_level: number
}

export interface Question extends DefaultFields {
    // Necessary to identify it's status as [right, wrong, not_done]
    question_id: string
    question_title: string
    question_description: string
    question_institution_id: number
    question_difficulty_level: number
    question_short_description: string
}

export interface QuestionList extends DefaultFields {
    id: number,
    title: string,
    subject: string,
    institution_id: number,
    difficulty_level: number,
}
export interface ExpandedQuestionList extends DefaultFields {
    id: number,
    title: string,
    subject: string,
    institution_id: number,
    difficulty_level: number,
    question_quantity: number,
}