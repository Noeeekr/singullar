// Components
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";
import ButtonLink from "@components/buttons/Link";
import Typography from "@mui/material/Typography"
import BoldTitle from "@components/titles/BoldTitle";
import Display from "./Display";
import Stack from "@mui/material/Stack"
import Paper from "@mui/material/Paper";

// Utilities
import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch";
import { SERVER_ADDR } from "../../../../configs";
import { useEffect } from "react";

// Models
import type { Subject } from "@models/server"

export default () => {
    const { response, isLoading, error, send } = useContextAwareFetch<Subject[], {}>(
        `${SERVER_ADDR}/api/subjects`,
        defaultRequestInit,
    )

    useEffect(() => {
        send()
    }, [])

    return (
        <Paper elevation={2} component="section" variant="outlined" sx={{ padding: 2 }}>
            <BoldTitle>Matérias</BoldTitle>
            <Stack gap={0.5} marginTop={2}>
                {
                    error
                        ? <ErrorBubble message={error} />
                        : response?.length
                            ? response.map((subject, i) => <Display key={i} title="Matéria">{subject.subject_name}</Display>)
                            : <ErrorBubble message="Nenhuma matéria encontrada" />
                }
                <Stack alignItems="center" justifyContent="center">
                    <Typography variant="body2" fontWeight="300" fontSize={13}>
                        Faltando alguma matéria?
                    </Typography>
                    <ButtonLink variant="solid" href="/platform/subjects/create" title="Adicionar matéria" />
                </Stack>
            </Stack>
            {
                isLoading
                    ?
                    <Typography variant="subtitle1" fontWeight="bold">
                        Carregando informações
                    </Typography>
                    : <></>
            }
        </Paper>
    )
}