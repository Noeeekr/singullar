import Box from "@mui/material/Box"
import Stack, { StackProps } from "@mui/material/Stack"
import Typography from "@mui/material/Typography"

import type { Question } from "@models/server"
import { styled } from "@mui/material/styles"

export interface QuestionBubbleProps extends StackProps {
    question: Question
}

const BubbleContainer = styled(Stack)(({ theme }) => ({
    flex: 1,

    background: "rgb(255,255,255, 0.9)",
    minHeight: 180,
    minWidth: 400,
    maxWidth: 700,
    padding: 16,
    border: "solid 1px rgba(190,190,190, 0.8)",
    borderRadius: 14,
    
    transition: "border 140ms ease-in-out, background-color 140ms ease-in-out, transform 140ms ease-in-out",
    cursor: true ? "pointer" : "initial",
    
    "&:hover": {
        transform: "scale(1.02)",
        backgroundColor: "rgb(160, 100, 240, 0.3)",
        borderColor: theme.palette.primary.purpleLight
    }
}))

export default function ({ question, ...props }: QuestionBubbleProps): JSX.Element {
    return (
        <BubbleContainer
            {...props}
            gap={2}
        >
            <Stack direction="row" gap={1} alignItems="center" justifyContent="space-between">
                <Typography component="p" fontWeight="bold" textTransform="capitalize">{ question.title }</Typography>
                <Stack direction="row" gap={1}>
                    <Box sx={{
                        backgroundColor: "rgb(130,230,130)",
                        paddingX: 1,
                        paddingY: 0.5,
                        borderRadius: 2,
                    }}>
                        <Typography fontWeight="bold" color="rgb(50, 150, 50)">{ question.difficultyName }</Typography>
                    </Box>
                    <Box sx={{
                        backgroundColor: "rgb(255,150,150)",
                        paddingX: 1,
                        paddingY: 0.5,
                        borderRadius: 2,
                    }}>
                        <Typography fontWeight="bold" color="rgb(190,50,50)">{ question.status }</Typography>
                    </Box>
                </Stack>
            </Stack>
            <Stack>
                <Typography component="p" fontWeight="bold" textTransform="capitalize">Descrição</Typography>
                <Typography textOverflow="ellipsis" sx={{ userSelect: "none", WebkitLineClamp: 3, WebkitBoxOrient: "vertical", display: "-webkit-box", overflow: "hidden" }}>
                    { question.description }
                </Typography>
            </Stack>
        </BubbleContainer>
    )
}