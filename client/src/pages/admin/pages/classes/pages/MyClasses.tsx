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
import { useParams } from 'react-router-dom';

const PageClasses = (): JSX.Element => {
    const insts = useAppSelector((store) => store.institutions)

    const { id } = useParams()
    
    return (
        <div>
            <SectionHeader
                title="Turmas"
                subtitle={insts.find(i => i.id.toString() == id)?.name || "Nome desconhecido"}
            >
                <div style={{ margin: '0px 0px 0px auto' }}>
                    <LinkButton
                        type="link"
                        icon={<div>D</div>}
                        title="Cadastrar turmas"
                        variant="solid"
                        href="/admin/classes/create"
                    />
                </div>
            </SectionHeader>
            <form style={{ marginTop: '60px' }}>
                <Grid container spacing={2}>
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

export default PageClasses;