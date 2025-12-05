// Components
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";
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
import type { Institution } from "@models/server";

export default () => {
    const { response, isLoading, error, send } = useContextAwareFetch<Institution, {}>(
        `${SERVER_ADDR}/api/dashboard`,
        defaultRequestInit,
    )

    useEffect(() => {
        send()
    }, [])

    return (
        <Paper elevation={2} component="section" variant="outlined" sx={{ padding: 2 }}>
            <BoldTitle> Instituição</BoldTitle>
            <Stack gap={0.5} marginTop={2}>
                {
                    !response || error
                        ? <ErrorBubble message={error} />
                        : <>
                            <Display title="Nome da instituição">
                                {response.name}
                            </Display>
                            <Display title="Plataforma ID">
                                {response.id * 1_000_000}
                            </Display>
                            <Display title="Data de criação">
                                {new Date(response.created_at).toLocaleDateString()}
                            </Display>
                        </>
                }
                {
                    isLoading
                        ?
                        <Typography variant="subtitle1" fontWeight="bold">
                            Carregando informações
                        </Typography>
                        : <></>
                }
            </Stack>
        </Paper>
    )
}