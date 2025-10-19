import FormControl from "@mui/material/FormControl"
import InputLabel from "@mui/material/InputLabel"
import MenuItem from "@mui/material/MenuItem"
import Select from "@mui/material/Select"
import Grid from "@mui/material/Grid2"
import ButtonSolid from "@components/buttons/Button/Solid"
import OutlinedInput from "@mui/material/OutlinedInput"
import Stack from "@mui/material/Stack"
import Typography from "@mui/material/Typography"

// Features
import { useForm } from "react-hook-form"
import { useAppSelector } from "@slices/store"

// Models
import type { UserSegments } from "@models/server"
import { EF1, EF2, EM } from "../../../../models/server"
import { useState } from "react"
import { SearchClassFilters } from "@hooks/useFetchClasses"

export default function SearchFilters({ send }: { send: (body: SearchClassFilters[]) => void }): JSX.Element {
    const institution = useAppSelector((store) => store.institution)
    const [showMore, setShowMore] = useState(false);
    const { setValue, getValues, watch, register } = useForm<SearchClassFilters>()

    const segment = watch("segment")

    return (
        <form style={{ marginTop: '60px', display: "flex", gap: '10px', flexDirection: "row", minWidth: "100%" }}>
            <Grid container spacing={2} sx={{ width: "100%" }}>
                <Grid size={12}>
                    <FormControl>
                        <InputLabel htmlFor="class-name-outlined-input">
                            Digite o nome da turma
                        </InputLabel>
                        <OutlinedInput {...register("class_name", { setValueAs: (value) => value ? value : undefined } )} id="class-name-outlined-input" label="Digite o nome do estudante" />
                    </FormControl>
                </Grid>
                {
                    !showMore
                        ? <></>
                        : <>
                            <Grid size={4}>
                                <FormControl aria-labelledby="school-select-label-id-ano-letivo">
                                    <InputLabel id="school-select-label-id-ano-letivo">
                                        Ano letivo
                                    </InputLabel>
                                    <Select
                                        label="Ano letivo"
                                        labelId="school-select-label-id-ano-letivo"
                                        onChange={(e) => {
                                            const currentYear = new Date()
                                            currentYear.setFullYear(e.target.value as number)
                                            setValue("creation_year", currentYear)
                                        }}
                                    >
                                        {
                                            (() => {
                                                if (institution == null) {
                                                    return <MenuItem value={0}>...</MenuItem>
                                                }

                                                let years: number[] = []
                                                let currentYear = new Date(Date.now()).getFullYear()
                                                let createdYear = new Date(institution.created_at).getFullYear()
                                                while (createdYear < currentYear) {
                                                    years.push(createdYear)
                                                    createdYear++
                                                }
                                                return years.map((year) => <MenuItem value={year}>{year}</MenuItem>)
                                            })()
                                        }
                                    </Select>
                                </FormControl>
                            </Grid>
                            <Grid size={4}>
                                <FormControl aria-labelledby="school-select-label-id-segmento">
                                    <InputLabel id="school-select-label-id-segmento">
                                        Segmento
                                    </InputLabel>
                                    <Select
                                        label="Segmento"
                                        labelId="school-select-label-id-segmento"
                                        onChange={(e) => setValue("segment", e.target.value as UserSegments)}
                                    >
                                        <MenuItem value={undefined}>Nenhum</MenuItem>
                                        <MenuItem value={EF1}>Ensino Fundamental 1</MenuItem>
                                        <MenuItem value={EF2}>Ensino Fundamental 2</MenuItem>
                                        <MenuItem value={EM}>Ensino Médio</MenuItem>
                                    </Select>
                                </FormControl>
                            </Grid>
                            <Grid size={4}>
                                <FormControl aria-labelledby="school-select-label-id-serieano">
                                    <InputLabel id="school-select-label-id-serieano">
                                        Série/Ano
                                    </InputLabel>
                                    {
                                        (() => {
                                            switch (segment) {
                                                case "":
                                                    return (
                                                        <Select
                                                            label="Série/Ano"
                                                            labelId="school-select-label-id-serieano"
                                                            disabled={!segment}
                                                            onChange={(e) => setValue("series", e.target.value as string)}
                                                        >
                                                            <MenuItem value={""}>...</MenuItem>
                                                        </Select>
                                                    )
                                                case EF1:
                                                    return (
                                                        <Select
                                                            label="Série/Ano"
                                                            labelId="school-select-label-id-serieano"
                                                            disabled={!segment}
                                                            onChange={(e) => setValue("series", e.target.value as string)}
                                                        >
                                                            <MenuItem value={"alfabetização"}>C.A (Classe de alfabetização)</MenuItem>
                                                            <MenuItem value={"1º Ano"}>1º Ano - Ensino Fundamental I</MenuItem>
                                                            <MenuItem value={"2º Ano"}>2º Ano - Ensino Fundamental I</MenuItem>
                                                            <MenuItem value={"3º Ano"}>3º Ano - Ensino Fundamental I</MenuItem>
                                                            <MenuItem value={"4º Ano"}>4º Ano - Ensino Fundamental I</MenuItem>
                                                            <MenuItem value={"5º Ano"}>5º Ano - Ensino Fundamental I</MenuItem>
                                                        </Select>
                                                    )
                                                case EF2:
                                                    return (
                                                        <Select
                                                            label="Série/Ano"
                                                            labelId="school-select-label-id-serieano"
                                                            disabled={!segment}
                                                            onChange={(e) => setValue("series", e.target.value as string)}
                                                        >
                                                            <MenuItem value={"6º Ano"}>6º Ano - Ensino Fundamental II</MenuItem>
                                                            <MenuItem value={"7º Ano"}>7º Ano - Ensino Fundamental II</MenuItem>
                                                            <MenuItem value={"8º Ano"}>8º Ano - Ensino Fundamental II</MenuItem>
                                                            <MenuItem value={"9º Ano"}>9º Ano - Ensino Fundamental II</MenuItem>
                                                        </Select>
                                                    )
                                                case EM:
                                                    return (
                                                        <Select
                                                            label="Série/Ano"
                                                            labelId="school-select-label-id-serieano"
                                                            disabled={!segment}
                                                            onChange={(e) => setValue("series", e.target.value as string)}
                                                        >
                                                            <MenuItem value={"1º Ano"}>1º Ano - Ensino Médio</MenuItem>
                                                            <MenuItem value={"2º Ano"}>2º Ano - Ensino Médio</MenuItem>
                                                            <MenuItem value={"3º Ano"}>3º Ano - Ensino Médio</MenuItem>
                                                        </Select>
                                                    )
                                            }
                                        })()
                                    }
                                </FormControl>
                            </Grid>
                            <Grid size={6}>
                                <FormControl>
                                    <InputLabel htmlFor="student-name-outlined-input">
                                        Digite o nome do estudante
                                    </InputLabel>
                                    <OutlinedInput {...register("student_name",  { setValueAs: (value) => value ? value : undefined } )} id="student-name-outlined-input" label="Digite o nome do professor" />
                                </FormControl>
                            </Grid>
                            <Grid size={6}>
                                <FormControl>
                                    <InputLabel htmlFor="teacher-name-outlined-input">
                                        Digite o nome do professor
                                    </InputLabel>
                                    <OutlinedInput {...register("teacher_name",  { setValueAs: (value) => value ? value : undefined } )} id="teacher-name-outlined-input" label="Digite o nome do estudante" />
                                </FormControl>
                            </Grid>
                        </>
                }
            </Grid>
            <Stack flexDirection="column" gap={2}>
                <ButtonSolid title="Pesquisar" onClick={() => send([getValues()])}/>
                <Typography
                    color="primary.purpleLight"
                    sx={{
                        textDecoration: "underline",
                        cursor: "pointer",
                    }}
                    onClick={() => setShowMore((prev) => !prev)}
                    fontWeight="bold"
                >
                    Ver {showMore ? "menos" : "mais"} filtros
                </Typography>
            </Stack>
        </form>
    )
}