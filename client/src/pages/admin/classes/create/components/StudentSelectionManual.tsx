import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";
import SectionTitle from "@components/headers/sectionHeader/SectionTitle";
import ButtonSolid from "@components/buttons/Button/Solid";

import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import OutlinedInput from "@mui/material/OutlinedInput";
import Select from "@mui/material/Select";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import { useContext, useEffect, useState, useCallback } from "react";
import { Controller, useForm } from "react-hook-form";
import useContextAwareFetch from "@hooks/useContextAwareFetch";

import { EF1, EF2, EM } from "../../../../../models/server/server";
import { SERVER_ADDR } from "../../../../../configs";
import { FormContext } from "./Form";

import type { User } from "@models/server/server";
import type { StudentFilters } from "@hooks/useFetchStudents";

const StudentSelectionManual = ({ returnButtonCallback }: { returnButtonCallback: () => void }): JSX.Element => {
    const { setFormData, formSections } = useContext(FormContext);
    const { register, getValues, control } = useForm<StudentFilters>({
        defaultValues: {
            name: "",
            email: "",
            segment: formSections.firstSection.segment,
            id: undefined,
            class_id: undefined,
        }
    });

    const { response, isLoading, error, send } = useContextAwareFetch<User[], StudentFilters[]>(
        `${SERVER_ADDR}/api/students`,
        {
            method: 'POST',
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            cache: 'no-cache',
        },
    )

    useEffect(() => {
        console.log(getValues())
        send([getValues()]);
    }, [])

    const [id, setId] = useState<number | null>(null)

    const toggleStudent = useCallback((id: number) => {
        let requestStudents = formSections.secondSection.students.filter((idA) => idA != id)
        if (formSections.secondSection.students.length == requestStudents.length) {
            requestStudents.push(id)
            setFormData("secondSection.students", requestStudents)
        } else {
            setFormData("secondSection.students", requestStudents)
        }
    }, [formSections, setFormData])
    return <Box>
        <Stack flexDirection="row" flexWrap="wrap" alignItems="center" gap={2}>
            <Stack gap={1}>
                <ButtonSolid
                    sx={{ backgroundColor: "white", boxShadow: "0px 0px 2px 3px rgb(0,0,0,0.01)" }}
                    color="primary.purpleDark"
                    onClick={returnButtonCallback}
                >
                    Voltar
                </ButtonSolid>
                <ButtonSolid
                    onClick={() => send([getValues()])}
                >
                    Procurar
                </ButtonSolid>
            </Stack>
            <Box flexGrow={1}>
                <SectionTitle
                    sx={{
                        marginBottom: "10px",
                    }}
                >
                    Procurar usando segmento
                </SectionTitle>
                <FormControl>
                    <InputLabel id="segment">
                        Insira o segmento do estudante
                    </InputLabel>
                    <Controller
                        name="segment"
                        control={control}
                        render={({ field }) =>
                            <Select
                                {...field}
                                {...register("segment")}
                                labelId="student-search-filter-segment-label"
                                label="digite-o-segmento-do-estudante"
                                value={formSections.firstSection.segment}
                                disabled={true}
                            >
                                <MenuItem value={""}>Nenhum</MenuItem>
                                <MenuItem value={EF1}>Ensino Fundamental 1</MenuItem>
                                <MenuItem value={EF2}>Ensino Fundamental 2</MenuItem>
                                <MenuItem value={EM}>Ensino Médio</MenuItem>
                            </Select>
                        }
                    />
                </FormControl>
            </Box>
            <Box flexGrow={1}>
                <SectionTitle
                    sx={{
                        marginBottom: "10px",
                    }}
                >
                    Procurar com Plataforma ID (opcional)
                </SectionTitle>
                <FormControl>
                    <InputLabel htmlFor="outlined-input-form-sheet-id">
                        Insira o id
                    </InputLabel>
                    <OutlinedInput
                        {...register("id", {
                            setValueAs(value) {
                                if (value == "") return null;
                                return Number(value)
                            },
                        })}
                        label="Insira o id (opcional)"
                        type="number"
                        id="outlined-input-form-sheet-id"
                        value={id}
                        onChange={(e) => { setId(Number(e.target.value) > -1 ? null : 0) }}
                    />
                </FormControl>
            </Box>
            <Box flexGrow={1}>
                <SectionTitle
                    sx={{
                        marginBottom: "10px",
                    }}
                >
                    Procurar com nome (opcional)
                </SectionTitle>
                <FormControl>
                    <InputLabel htmlFor="outlined-input-form-sheet-name">
                        Insira o nome
                    </InputLabel>
                    <OutlinedInput
                        {...register("name")}
                        label="Insira o nome (opcional)"
                        id="outlined-input-form-sheet-name"
                        value={""}
                    />
                </FormControl>
            </Box>
        </Stack>
        <Stack flexDirection="column" alignItems="center" marginTop={2}>
            {
                isLoading
                    ? <Typography variant="subtitle1">Procurando estudantes..</Typography>
                    : <></>
            }
            {
                error == ""
                    ? <></>
                    : <ErrorBubble err={error} />
            }
            <Typography component="p" marginY={2}>
                {
                    `${formSections.secondSection.students.length} estudantes selecionados. `
                }
                <Typography
                    onClick={() => { setFormData("secondSection.students", []) }}
                    component="span"
                    sx={{
                        cursor: "pointer",
                        color: "primary.purpleLight",
                        textDecoration: "underline"
                    }}
                >
                    Limpar escolha
                </Typography>
            </Typography>
            {
                response == null
                    ? <></>
                    : <Stack gap={1} width="100%">
                        {
                            response.map((student) => {
                                const selected = formSections?.secondSection.students?.find((sId) => sId == student.id)
                                return (
                                    <Box
                                        onClick={() => toggleStudent(student.id)}
                                        key={student.id}
                                        sx={{
                                            cursor: "pointer",
                                            backgroundColor: selected ? "primary.purpleLight" : "white",
                                            transition: "background-color 150ms ease-in-out",
                                            width: "100%",
                                            padding: "0.5rem 1rem",
                                            borderRadius: "0.5rem",
                                            boxShadow: "0px 0px 1px 4px rgb(150,150,150,0.05)",

                                            flex: 1,
                                        }}
                                    >
                                        <Stack direction="row" justifyContent="space-between" alignItems="center">
                                            <Typography sx={{ color: selected ? "white" : "black" }} variant="body2" component="p">{student.name}</Typography>
                                            <Typography sx={{ color: selected ? "white" : "black" }} variant="body1" component="p">PLATAFORMA ID {student.id}</Typography>
                                        </Stack>
                                        <Stack direction="row" justifyContent="space-between" alignItems="center">
                                            <Typography sx={{ color: selected ? "white" : "black" }} variant="subtitle1" component="p">EMAIL: {student.email}</Typography>
                                            <Typography sx={{ color: selected ? "white" : "black" }} variant="subtitle1" component="p">{student.segment}</Typography>
                                        </Stack>
                                    </Box>
                                )
                            })
                        }
                    </Stack>
            }
        </Stack>
    </Box>
}
export default StudentSelectionManual;