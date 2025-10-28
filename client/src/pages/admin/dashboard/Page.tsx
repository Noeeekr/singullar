import Stack from "@mui/material/Stack"
import Divider from "@mui/material/Divider"
import Typography from "@mui/material/Typography"

import { useEffect } from "react";
import useContextAwareFetch from "@hooks/useContextAwareFetch";

import { Institution } from "@models/server";
import { SERVER_ADDR } from "../../../configs";
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";
import SectionHeader from "@components/headers/sectionHeader/SectionHeader";

import PopupButton from "@components/buttons/Popup";

const Dashboard = (): JSX.Element => {
    const { response, isLoading, error, send } = useContextAwareFetch<Institution, {}>(
        `${SERVER_ADDR}/api/dashboard`,
        {
            "method": "GET",
            "credentials": "include",
            "headers": {
                "Content-Type": "application/json",
            },
        }
    )

    useEffect(() => {
        send({})
    }, [])


    return (
        <Stack gap={2}>
            <SectionHeader title="Informações Escolares" subtitle="Ultima atualização: Agora" />
            <Divider />
            <div>
                {
                    !response || error
                        ? <ErrorBubble err={error} />
                        : <>
                            <Typography variant="body2">
                                <span style={{ fontWeight: "bold" }}>Nome da instituição:</span> {response.name}
                            </Typography>
                            <Typography variant="body2">
                                <span style={{ fontWeight: "bold" }}>Plataforma ID:</span> {response.id * 1_000_000}
                            </Typography>
                            <Typography variant="body2">
                                <span style={{ fontWeight: "bold" }}>Data de criação:</span> {new Date(response.created_at).toLocaleDateString()}
                            </Typography>
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
            </div>
            <Divider />
            <Typography variant="h5" component="h4" fontWeight="bold">
                Curriculo Escolar
            </Typography>
            <div>
                PopupButton
                <div style={{ background: "rgb(230,0,0, 0.05)", padding: 20, border: "solid 1px rgb(0,0,0,0.05)" }}>
                    <PopupButton
                        variant="side"
                        element={<div>ELeme</div>}
                        id="schooldata"
                        type="popup"
                        title="title"
                    >
                        Contentzao
                    </PopupButton>
                </div>
            </div>
        </Stack>
    )
}

export default Dashboard;
