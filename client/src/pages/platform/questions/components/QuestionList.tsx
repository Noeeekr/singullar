import Grid from "@mui/material/Grid2";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography"
import QuestionNumberDisplay from "./NumberDisplay";

import { styled } from "@mui/material";

import type { Grid2Props } from "@mui/material"
import type { StackProps } from "@mui/material"

export interface FilterOptions {
    name?: string,
    subject?: string,
    difficultyName?: string,
    difficultyLevel?: number,
}
export interface QuestionList {
    title: string,
    subject: string,
    questionQuantity: number
    difficultyLevel: number
    difficultyName: string
}
export interface QuestionListProps extends StackProps {
    filters?: FilterOptions
} 


const QuestionContainer = styled(({ children, ...props }: Grid2Props) => (
    <Grid size={4} {...props}>{ children }</Grid>
))(({ theme }) => ({
    backgroundColor: "rgba(255,255,255,0.5)",
    height: "100%",
        minHeight: "200px",
    width: "100%",
        minWidth: "350px",
    padding: "20px",
    border: `solid 1px ${theme.palette.grey[400]}`,
    borderRadius: "1rem",

    cursor: "pointer",

    transition: "150ms ease-in-out border, 150ms ease-in-out background-color",
    
    "&:hover": {
        backgroundColor: `rgba(138, 59, 217, 0.19)`,
        border: `solid 1px ${theme.palette.primary.purpleLight}`
    }
}))

export default function ({ filters }: QuestionListProps): JSX.Element {
    const questionLists: QuestionList[] = [
        {
            title: "Fundamentos da matematica",
            subject: "matematica",
            difficultyLevel: 0,
            difficultyName: "fundamental",
            questionQuantity: 3,
        },
        {
            title: "Matematica básica",
            subject: "matematica",
            difficultyLevel: 1,
            difficultyName: "iniciante",
            questionQuantity: 6,
        },
        {
            title: "Matematica básica (Extras)",
            subject: "matematica",
            difficultyLevel: 2,
            difficultyName: "intermediário",
            questionQuantity: 3,
        },
        {
            title: "Matematica avançada.",
            subject: "matematica",
            difficultyLevel: 3,
            difficultyName: "avançado",
            questionQuantity: 3,
        }, 
    ]
    
    return(
        <Grid container spacing={2}>
            { 
            questionLists.map((questionList, index) => {
                const questionDisplay: JSX.Element[] = new Array(questionList.questionQuantity)
                for (let i = 0; i < questionList.questionQuantity; i++) {
                    questionDisplay[i] = <QuestionNumberDisplay level={questionList.difficultyLevel}>{ i + 1 }</QuestionNumberDisplay>
                }
                return(
                    <QuestionContainer key={index}>
                        <Stack direction="row" justifyContent="space-between">
                            <Typography variant="h6" fontWeight="bold" component="p" textTransform="capitalize">
                                { questionList.title }
                            </Typography>
                            <Typography textAlign="end" textTransform="capitalize">
                                <span style={{ fontWeight: "bold" }}>Matéria:</span> { questionList.subject }
                            </Typography>
                        </Stack>
                        <Typography textTransform="capitalize">
                            <span style={{ fontWeight: "bold" }}>Dificuldade:</span> { questionList.difficultyName }
                        </Typography>
                        <Stack direction="row" gap={1} marginY={2}>
                            { ...questionDisplay }
                        </Stack>
                    </QuestionContainer>
                )
            })
            }
        </Grid>
    )
}