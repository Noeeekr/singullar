import { DefaultFields } from ".";

export interface QuestionListDifficulty {
    question_difficulty_name: string
    question_difficulty_level: number
}

export interface Question extends DefaultFields {
    // Necessary to identify it's status as [right, wrong, not_done]
    question_id: string
    question_title: string
    question_status: string
    question_description: string
    question_institution_id: number
    question_difficulty_level: number
    question_short_description: string
}

export interface QuestionList extends DefaultFields {
    id: number,
    subject_name: string,
    institution_id: number,
    question_amount: number,
    question_list_title: string,
    question_list_subject_id: string,
    question_list_difficulty_level: number,
}
export interface ExpandedQuestionList extends DefaultFields {
    id: number,
    title: string,
    subject: string,
    institution_id: number,
    difficulty_level: number,
    question_quantity: number,
}

const difficulties: {[index: number]: string} = {
    1: "Fundamental",
    2: "Iniciante",
    3: "Intermediario",
    4: "Intermediario+",
    5: "Intermediario++",
    6: "Avançado",
    7: "Avançado+",
    8: "Avançado++",
    9: "Profissional",
    10: "Expert",
}

// Parses the difficulty level (int) to its name (string) 
// **Not always realiable, if the level is not set will return "Dificuldade desconhecida" 
export function parseQuestionDifficulty(difficulty_level: number): string {
    const difficulty_level_name: string | undefined = difficulties[difficulty_level]
    return difficulty_level_name == undefined ? "Dificuldade desconhecida" : difficulty_level_name
}