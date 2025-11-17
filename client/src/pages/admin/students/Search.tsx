import OutlinedInput from "@mui/material/OutlinedInput";
import InputLabel from "@mui/material/InputLabel";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import SectionHeader from "@components/headers/sectionHeader/SectionHeader";
import SectionTitle from "@components/headers/sectionHeader/SectionTitle";
import Grid from "@mui/material/Grid2";
import Button from "@components/buttons/Default/Solid";
import LinkButton from "@components/buttons/Link";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography"

import { useEffect } from "react";
import useFetchUsers from "@hooks/useFetchUsers";

import { EF1, EF2, EM, ROLE_STUDENT, User, UserRoles } from "../../../models/server"
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";
import { useForm } from "react-hook-form";

const roles: UserRoles[] = [ROLE_STUDENT]

type Inputs = Pick<User, "segment" | "name" | "id">

const CreatePage = (): JSX.Element => {
    const { register, watch } = useForm<Inputs>({
        defaultValues: {
            name: "",
            segment: null,
            id: 0,
        }
    });

    const name = watch("name")
    const segment = watch("segment")
    const id = watch("id")

    const { response, isLoading, error, send } = useFetchUsers()

    useEffect(() => {
        send({"accepted_roles": roles })
    }, [])

    return (
        <div>
            { /* Action Buttons */}

            <SectionHeader title="Selecione um estudante">
                <Box sx={{ margin: "0 0 0 auto" }}>
                    <LinkButton
                        icon={{ component: <></> }}
                        type="link"
                        title="Criar Estudante"
                        variant="solid"
                        href="/admin/students/create"
                    >
                        Criar estudante
                    </LinkButton>
                </Box>
            </SectionHeader>

            { /* Filter Student Form */}

            <Box display="flex" flexDirection="column" gap={3}>
                <Stack direction="row" gap={2}>
                    <Grid container spacing={{ mobile: 0, xss: 4 }} sx={{ flex: 1 }}>
                        <Grid size={{ mobile: 12, xss: 4 }}>
                            <SectionTitle
                                sx={{
                                    margin: "30px 0px 10px 0px",
                                }}
                            >
                                Filtrar por nome
                            </SectionTitle>
                            <FormControl>
                                <InputLabel htmlFor="student-search-filter-name-input">
                                    Digite o nome do estudante
                                </InputLabel>
                                <OutlinedInput
                                    label="digite-o-nome-do-estudante"
                                    id="student-search-filter-name-input"
                                    {...register("name")}
                                />
                            </FormControl>
                        </Grid>
                        <Grid size={{ mobile: 12, xss: 4 }}>
                            <SectionTitle
                                sx={{
                                    margin: "30px 0px 10px 0px",
                                }}
                            >
                                Filtrar por segmento
                            </SectionTitle>
                            <FormControl>
                                <InputLabel id="segment">
                                    Digite o segmento do estudante
                                </InputLabel>
                                <Select
                                    labelId="student-search-filter-segment-label"
                                    label="digite-o-segmento-do-estudante"
                                    {...register("segment")}
                                >
                                    <MenuItem value={""}>Escolha uma opção</MenuItem>
                                    <MenuItem value={EF1}>Ensino Fundamental 1</MenuItem>
                                    <MenuItem value={EF2}>Ensino Fundamental 2</MenuItem>
                                    <MenuItem value={EM}>Ensino Médio</MenuItem>
                                </Select>
                            </FormControl>
                        </Grid>
                        <Grid size={{ mobile: 12, xss: 4 }}>
                            <SectionTitle
                                sx={{
                                    margin: "30px 0px 10px 0px",
                                }}
                            >
                                Filtrar por matricula/ID
                            </SectionTitle>
                            <FormControl>
                                <InputLabel htmlFor="student-search-filter-id-input">
                                    Digite o ID do estudante
                                </InputLabel>
                                <OutlinedInput
                                    type="number"
                                    label="digite-o-ID-do-estudante"
                                    id="student-search-filter-id-input"
                                    {...register("id", {
                                        valueAsNumber: true,
                                    })}
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
                                        ? <ErrorBubble message="Ocorreu um erro ao procurar os usuários. Por favor tente novamente mais tarde." />
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
                                    response != null && response.length == 0 && !isLoading ?
                                        <>
                                            <Typography variant="subtitle1">
                                                Nenhum usuário encontrado
                                            </Typography>
                                            <Button title="Recarregar" onClick={() => { send({"accepted_roles": roles})}}>Recarregar</Button>
                                        </>
                                        : <></>
                                }
                            </Box>
                            : <></>
                    }
                    {
                        response?.filter((user) => {
                            if (name && user.name.toLowerCase().search(name.toLowerCase())) {
                                return false
                            }
                            if (id && user.id != id) {
                                return false
                            }
                            if (segment && user.segment != segment) {
                                return false
                            }
                            return true
                        }).sort((usr1, usr2) => {
                            const name1 = usr1.name.toLowerCase()
                            const name2 = usr2.name.toLowerCase()

                            if (name1 == name2) {
                                return 0
                            } else if (name1 > name2) {
                                return 1
                            } else {
                                return -1
                            }
                        }).map((user, i) => {
                            return <Grid size={6} key={i}>
                                <Box sx={{
                                    flex: 1,
                                    bgcolor: "white",
                                    padding: 2,
                                    borderRadius: 2,
                                    boxShadow: "1px 1px 4px 1px rgb(190,190,190,0.3)",
                                }}>
                                    <Typography variant="body1" fontWeight="bold" key={user.name}>{user.name}</Typography>
                                    <Typography variant="body2" color="gray" key={user.email}>{user.email}</Typography>
                                </Box>
                            </Grid>
                        })
                    }
                </Grid>
            </Box>
        </div>
    );
};

export default CreatePage;