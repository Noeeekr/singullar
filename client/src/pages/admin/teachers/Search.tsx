import OutlinedInput from "@mui/material/OutlinedInput";
import InputLabel from "@mui/material/InputLabel";
import FormControl from "@mui/material/FormControl";
import SectionHeader from "@components/SectionHeader";
import SectionTitle from "@components/SectionTitle";
import Grid from "@mui/material/Grid2";
import Button from "@components/ButtonSolid";
import ButtonLink from "@components/ButtonLink";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography"

import { useForm } from "react-hook-form"
import { useEffect } from "react";
import useFetchUsers from "@hooks/useFetchUsers";

import { ROLE_TEACHER, UserRoles } from "../../../types/server"
import ErrorBubble from "@components/ErrorBubble";

const roles: UserRoles[] = [ROLE_TEACHER]

type Inputs = {
    name: string | null;
    id: number | null;
}

const CreatePage = (): JSX.Element => {
    const { register, watch } = useForm<Inputs>({
        defaultValues: {
            name: null,
            id: null,
        }
    });

    const name = watch("name")
    const id = watch("id")

    const [users, isLoading, error, fetchUsers] = useFetchUsers(roles)

    useEffect(() => {
        fetchUsers()
    }, [fetchUsers])

    return (
        <div>
            { /* Action Buttons */}

            <SectionHeader title="Selecione um estudante">
                <>
                    <Button sx={{ margin: "0 0 0 auto" }}>Filtrar</Button>
                    <ButtonLink
                        icon={<></>}
                        type="link"
                        title="doesnt-matter"
                        href="/admin/students/create"
                    >
                        Criar professor
                    </ButtonLink>
                </>
            </SectionHeader>

            { /* Filter Student Form */}

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
                        users.length == 0
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
                                    users.length == 0 && !isLoading ?
                                        <>
                                            <Typography variant="subtitle1">
                                                Nenhum usuário encontrado
                                            </Typography>
                                            <Button onClick={() => { fetchUsers() }}>Recarregar</Button>
                                        </>
                                        : <></>
                                }
                            </Box>
                            : <></>
                    }
                    {
                        users.sort((usr1, usr2) => {
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
                        } ).map((user, i) => {
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