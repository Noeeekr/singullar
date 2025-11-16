// Components
import SolidButton from "@components/buttons/Default/Solid"
import SectionHeader from "@components/headers/sectionHeader"
import BoldTitle from "@components/titles/BoldTitle"
import Box from "@mui/material/Box"
import FormControl from "@mui/material/FormControl"
import FormHelperText from "@mui/material/FormHelperText"
import InputLabel from "@mui/material/InputLabel"
import OutlinedInput from "@mui/material/OutlinedInput"
import Stack from "@mui/material/Stack"


// Utilities
import { SubmitHandler, useForm } from "react-hook-form"
import { useState } from "react"
import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch"
import { SERVER_ADDR } from "../../../../configs"
import { Subject } from "@models/server"

import ErrorBubble from "@components/bubbles/ErrorBubble"
import SuccessBubble from "@components/bubbles/SuccessBubble"

// Models
export interface CreateSubjectRequest {
    subject_name?: string
}

export default (): JSX.Element => {
    const { register, handleSubmit, reset, getValues } = useForm<CreateSubjectRequest>()
    const [isValidName, setIsValidName] = useState(false)

    const { response, error, send } = useContextAwareFetch<Subject, CreateSubjectRequest>(
        `${SERVER_ADDR}/api/subjects/create`,
        {
            ...defaultRequestInit,
            method: "POST",
        },
    )

    const handleFormSubmit: SubmitHandler<CreateSubjectRequest> = (data) => {
        send(data)
        setIsValidName(false)
        reset({ subject_name: "" })
    }

    return (
        <Box component="section">
            <SectionHeader
                title="Adicionar matéria"
                subtitle="Preencha as informações da matéria"
            >
                <SolidButton
                    button={{ type: "submit" }}
                    onClick={() => handleFormSubmit(getValues())}
                    sx={{ textTransform: "capitalize" }}
                    title="Criar matéria"
                    disabled={!isValidName} />
            </SectionHeader>
            <Stack gap={1} marginY={2} component="form" onSubmit={handleSubmit(handleFormSubmit)}>
                <BoldTitle>
                    Nome da matéria
                </BoldTitle>
                <FormControl>
                    <InputLabel htmlFor="create-subject-outlined-input-subject-name">
                        Insira o nome da matéria
                    </InputLabel>
                    <OutlinedInput
                        {...register("subject_name", { required: true, minLength: 4 })}
                        id="create-subject-outlined-input-subject-name"
                        label="Insira o nome da matéria"
                        onChange={(e) => {
                            setIsValidName(e.target.value.length >= 4)
                        }}
                    />
                    {
                        isValidName
                            ? <></>
                            : <FormHelperText error={true}>*Por favor, insira no mínimo 4 caractéres.</FormHelperText>
                    }
                </FormControl>
            </Stack>
            <ErrorBubble message={error} />
            <SuccessBubble message={response != null ? "Matéria criada com sucesso!" : ""} />
        </Box>
    )
}