import Box from "@mui/material/Box"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import Typography from "@mui/material/Typography"
import Button from "@mui/material/Button"
import InputLabel from "@mui/material/InputLabel"
import SectionTitle from "@components/headers/sectionHeader/SectionTitle"
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble"

import { useForm } from "react-hook-form"
import { useCallback, useEffect, useState } from "react";
import useFetchUsers from "@hooks/useFetchUsers";

import { ROLE_TEACHER } from "../../../../models/server" 
import type { UserRoles } from "@models/server"

const roles: UserRoles[] = [ROLE_TEACHER]

type Inputs = {
    name: string | null;
    id: number | null;
}

/**
 * 
 * @param selectable Specifies if you can select any teachers. The number providen is the amoutn of teachers that can be selected
 * @param onSelect calls the functions with all selected teachers ids as an array on first argument
 */
export default function Search(
    { selectable = -1, onSelect = () => { } }:
        {
            selectable?: number,
            onSelect?: (ids: number[]) => void
        }
): JSX.Element {
    const { register, watch } = useForm<Inputs>({
        defaultValues: {
            name: null,
            id: null,
        }
    });

    const [selectedUsersIds, setSelectedUsersIds] = useState<number[]>([]);
    const name = watch("name")
    const id = watch("id")

    const { response, isLoading, error, send } = useFetchUsers()

    const handleSelect = useCallback((id: number) => {
        if (!selectable) {
            return
        }
        let ids = selectedUsersIds.filter((sId) => sId != id)
        if (ids.length == selectedUsersIds.length) {
            ids.push(id)
            if (selectedUsersIds.length >= selectable) {
                return
            }
        }
        setSelectedUsersIds(ids)
    }, [selectedUsersIds, selectable])

    useEffect(() => {
        onSelect(selectedUsersIds)
    }, [selectedUsersIds, onSelect])

    useEffect(() => {
        send({ "accepted_roles": roles })
    }, [])

    return (
        <Box display="flex" flexDirection="column" gap={3}>
            <Stack direction="row" gap={2}>
                <Grid container spacing={{ mobile: 0, xss: 4 }} sx={{ flex: 1 }}>
                    <Grid size={{ mobile: 12, xss: 5 }}>
                        <SectionTitle
                            sx={{
                                margin: "30px 0px 10px 0px",
                            }}
                        >
                            Filtrar por nome
                        </SectionTitle>
                        <FormControl>
                            <InputLabel htmlFor="teacher-search-filter-name-input">
                                Digite o nome do professor
                            </InputLabel>
                            <OutlinedInput
                                {...register("name")}
                                label="digite-o-nome-do-professor"
                                id="teacher-search-filter-name-input"
                            />
                        </FormControl>
                    </Grid>
                    <Grid size={{ mobile: 12, xss: 5 }}>
                        <SectionTitle
                            sx={{
                                margin: "30px 0px 10px 0px",
                            }}
                        >
                            Filtrar por matricula/ID
                        </SectionTitle>
                        <FormControl>
                            <InputLabel htmlFor="teacher-search-filter-id-input">
                                Digite o ID do professor
                            </InputLabel>
                            <OutlinedInput
                                type="number"
                                {...register("id", {
                                    valueAsNumber: true,
                                })}
                                label="digite-o-ID-do-professor"
                                id="teacher-search-filter-id-input"
                            />
                        </FormControl>
                    </Grid>
                </Grid>
            </Stack>

            { /* Fetched Student List */}

            <Grid container paddingY={2} spacing={2}>
                {
                    response != null && response.length == 0
                        ?
                        <Box sx={{
                            borderRadius: 2,
                            padding: 2,
                            margin: "auto",
                            display: "flex",
                            flexDirection: "column",
                            gap: 2,
                            textAlign: "center",
                        }}>
                            {
                                error != ""
                                    ? <ErrorBubble err={`Ocorreu um erro ao procurar os usuários. Por favor tente novamente mais tarde.`} />
                                    : <></>
                            }
                            {
                                isLoading ?
                                    <Typography variant="subtitle1">
                                        Um momento.. carregando usuários
                                    </Typography>
                                    : <></>
                            }
                            {
                                response.length == 0 && !isLoading ?
                                    <>
                                        <Typography variant="subtitle1">
                                            Nenhum usuário encontrado
                                        </Typography>
                                        <Button onClick={() => { send({"accepted_roles": roles })}}>Recarregar</Button>
                                    </>
                                    : <></>
                            }
                        </Box>
                        : <></>
                }
                {
                    response != null && response.sort((usr1, usr2) => {
                        const name1 = usr1.name.toLowerCase()
                        const name2 = usr2.name.toLowerCase()

                        if (name1 == name2) {
                            return 0
                        } else if (name1 > name2) {
                            return 1
                        } else {
                            return -1
                        }
                    }).filter((user) => {
                        if (name && user.name.toLowerCase().search(name.toLowerCase())) {
                            return false
                        }
                        if (id && user.id != id) {
                            return false
                        }
                        return true
                    }).map((user, i) => {
                        const isSelected = selectedUsersIds.includes(user.id);
                        return <Grid size={6} key={i}>
                            <Box
                                sx={{
                                    cursor: selectable ? "pointer" : "initial",
                                    flex: 1,
                                    bgcolor: isSelected ? "primary.purpleLight" : selectable == selectedUsersIds.length ? "rgb(220,220,220)" : "white",
                                    border: selectable <= selectedUsersIds.length ? "solid 1px rgb(230,230,230)": "solid 2px rgb(200,200,200)",
                                    transition: "all 150ms ease-in-out",
                                    padding: 2,
                                    borderRadius: 2,
                                    boxShadow: "1px 1px 4px 1px rgb(190,190,190,0.3)",
                                }}
                                onClick={() => { handleSelect(user.id) }}
                            >
                                <Typography variant="body1" fontWeight="bold" color={isSelected ? "white" : "black"} key={user.name}>{user.name}</Typography>
                                <Typography variant="body2" color={isSelected ? "rgb(210,210,210)" : "gray"} key={user.email}>{user.email}</Typography>
                            </Box>
                        </Grid>
                    })
                }
            </Grid>
        </Box>
    )
}