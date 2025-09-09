// COmponents
import FormControl from '@mui/material/FormControl'
import Select from '@mui/material/Select'
import MenuItem from '@mui/material/MenuItem'
import InputLabel from '@mui/material/InputLabel'
import Grid from '@mui/material/Grid2'
import LinkButton from '@components/ButtonLink'
import SectionHeader from '@components/SectionHeader';

// Features
import { useAppSelector } from '@slices/store';

const Search = (): JSX.Element => {
    const institution = useAppSelector((store) => store.institution)

    return (
        <div>
            <SectionHeader
                title="Turmas"
                subtitle={institution?.name || "Nome desconhecido"}
            >
                <div style={{ margin: '0px 0px 0px auto' }}>
                    <LinkButton
                        title=""
                        icon={<></>}
                        type="link"
                        href="/admin/classes/create"
                    >
                        Cadastrar turmas
                    </LinkButton>
                </div>
            </SectionHeader>
            <form style={{ marginTop: '60px', display: "flex", gap: '10px', flexDirection: "row", minWidth: "100%" }}>
                <Grid container spacing={2} sx={{ width: "100%"}}>
                    <Grid size={4}>
                        <FormControl
                            aria-labelledby="school-select-label-id-ano-letivo"
                        >
                            <InputLabel id="school-select-label-id-ano-letivo">
                                Ano letivo
                            </InputLabel>
                            <Select
                                label="Ano letivo" 
                                labelId="school-select-label-id-ano-letivo"
                            >
                                <MenuItem>

                                </MenuItem>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid size={4}>
                        <FormControl
                            aria-labelledby="school-select-label-id-segmento"
                        >
                            <InputLabel id="school-select-label-id-segmento">
                                Segmento
                            </InputLabel>
                            <Select
                                label="Segmento"
                                labelId="school-select-label-id-segmento"
                            >
                                <MenuItem>

                                </MenuItem>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid size={4}>
                        <FormControl
                            aria-labelledby="school-select-label-id-serieano"
                        >
                            <InputLabel
                                id="school-select-label-id-serieano"
                            >
                                Série/Ano
                            </InputLabel>
                            <Select
                                label="Série/Ano"
                                labelId="school-select-label-id-serieano"                            
                            >
                                <MenuItem>

                                </MenuItem>
                            </Select>
                        </FormControl>
                    </Grid>
                </Grid>
            </form>
        </div>
    )
}

export default Search;