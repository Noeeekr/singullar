// COmponents
import Stack from '@mui/material/Stack'
import Typography from '@mui/material/Typography'
import CircularButton from '../../../components/ButtonCircular';
import FormControl from '@mui/material/FormControl'
import Select from '@mui/material/Select'
import MenuItem from '@mui/material/MenuItem'
import InputLabel from '@mui/material/InputLabel'
import Grid from '@mui/material/Grid2'
import LinkButton from '../../../components/ButtonLink'

// Features
import { useAppSelector } from '../../../slices/store';
import { useParams, useNavigate } from 'react-router-dom';

const PageClasses = (): JSX.Element => {
    const insts = useAppSelector((store) => store.institutions)

    const { id } = useParams()

    const navigate = useNavigate()
    return (
        <div>
            <Stack direction="row" alignItems="start" width="100%" gap={1}>
                <CircularButton
                    onClickCb={() => { navigate(-1) }}
                >
                    &lt;
                </CircularButton>
                <Stack gap={1.5}>
                    <Typography component="h4" variant="h3" fontWeight="600">
                        Turmas
                    </Typography>
                    <Typography component="p" variant="body1" fontWeight="600" color="primary.whiteLow">
                        {insts.find(i => i.id.toString() == id)?.name || "Nome desconhecido"}
                    </Typography>
                </Stack>
                <div style={{ margin: '0px 0px 0px auto' }}>
                    <LinkButton 
                        type="link" 
                        icon={<div>D</div>} 
                        title="Cadastrar turmas" 
                        variant="solid"
                        href="/admin/classes/create"
                    />
                </div>
            </Stack>
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