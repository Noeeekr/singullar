import OutlinedInput from "@mui/material/OutlinedInput";
import InputLabel from "@mui/material/InputLabel";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import SectionHeader from "@components/SectionHeader";
import SectionTitle from "@components/SectionTitle";
import Grid from "@mui/material/Grid2";
import Button from "@components/ButtonSolid";
import ButtonLink from "@components/ButtonLink";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography"

import { useEffect, useState } from "react";
import useFetchUsersByInstitutionID from "@hooks/useFetchUsersByInstitutionID";

import { useAppSelector } from "@slices/store";

const CreatePage = (): JSX.Element => {
    const [filterForm, setFilterForm] = useState<{
        name: string | -1;
        segment: number;
        id: number;
    }>({
        name: -1,
        segment: -1,
        id: -1,
    });

    const institution = useAppSelector(state => state.institution)
    let institutionID: number = 0;
    if (institution != null) {
        institutionID = institution.id
    }
    const [users, isLoading, error, fetchUsers] = useFetchUsersByInstitutionID(institutionID)

    useEffect(() => {
        console.log("Child called from func")
        fetchUsers()
    }, [fetchUsers])

    return (
        <div>
            <SectionHeader title="Selecione um estudante">
                <>
                    <Button sx={{ margin: "0 0 0 auto" }}>Filtrar</Button>
                    <ButtonLink
                        icon={<></>}
                        type="link"
                        title="doesnt-matter"
                        href="/admin/students/create"
                    >
                        Criar estudante
                    </ButtonLink>
                </>
            </SectionHeader>
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
                                <InputLabel id="student-search-filter-segment-label">
                                    Digite o segmento do estudante
                                </InputLabel>
                                <Select
                                    labelId="student-search-filter-segment-label"
                                    name="student-search-filter-segment-input"
                                    label="digite-o-segmento-do-estudante"
                                    value={filterForm.name}
                                    onChange={(e) => {
                                        setFilterForm((prevVal) => ({
                                            ...prevVal,
                                            name: e.target.value,
                                        }));
                                    }}
                                >
                                    <MenuItem value={-1}>Escolha uma opção</MenuItem>
                                    <MenuItem value={1}>Ensino Fundamental 1</MenuItem>
                                    <MenuItem value={2}>Ensino Fundamental 2</MenuItem>
                                    <MenuItem value={3}>Ensino Médio</MenuItem>
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
                                    label="digite-o-ID-do-estudante"
                                    id="student-search-filter-id-input"
                                />
                            </FormControl>
                        </Grid>
                    </Grid>
                </Stack>
                <Grid container paddingY={2} spacing={2}>
                    {
                        isLoading
                            ?
                            <Box>
                                <p>Um momento.. carregando usuários</p>
                            </Box>
                            : users.length == 0 && <p>Nenhum usuário encontrado</p>
                    }
                    {
                        error != null && <p>{error}</p>
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
