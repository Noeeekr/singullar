import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import SectionHeader from '@components/SectionHeader';
import SectionTitle from '@components/SectionTitle';
import Grid from '@mui/material/Grid2';
import Button from '@components/ButtonSolid';
import Stack from '@mui/material/Stack';

import { useState } from 'react';

const CreatePage = (): JSX.Element => {
    const [filterForm, setFilterForm] = useState<{
        name: string | -1,
        segment: number,
        id: number,
    }>({
        name: -1,
        segment: -1,
        id: -1,
    })

    return (
        <div>
            <SectionHeader 
                title="Selecione um estudante" 
            >
                <Button
                        sx={{
                            width: '160px',
                            margin: '10px 0px 0px auto',
                        }}
                    >
                    Filtrar
                </Button>
            </SectionHeader>
            <Stack direction="row" gap={2}>
                <Grid container spacing={{ mobile: 0, xs: 4 }} sx={{ flex: 1}}>
                    <Grid size={{ mobile: 12, xs: 4 }}>
                        <SectionTitle sx={{
                            margin: '30px 0px 10px 0px'
                        }}>
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
                    <Grid size={{ mobile: 12, xs: 4 }}>
                        <SectionTitle sx={{
                            margin: '30px 0px 10px 0px'
                        }}>
                            Filtrar por segmento
                        </SectionTitle>
                        <FormControl>
                            <InputLabel htmlFor="student-search-filter-name-input">
                                Digite o segmento do estudante
                            </InputLabel>
                            <Select
                                label="digite-o-segmento-do-estudante"
                                id="student-search-filter-name-input"
                                value={filterForm.name}
                                onChange={(e) => { setFilterForm(prevVal => ({ ...prevVal, name: e.target.value })) }}
                            >
                                <MenuItem value={-1}>Escolha uma opção</MenuItem>
                                <MenuItem value={1}>Ensino Fundamental 1</MenuItem>
                                <MenuItem value={2}>Ensino Fundamental 2</MenuItem>
                                <MenuItem value={3}>Ensino Médio</MenuItem>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid size={{ mobile: 12, xs: 4 }}>
                        <SectionTitle sx={{
                            margin: '30px 0px 10px 0px'
                        }}>
                            Filtrar por matricula/ID
                        </SectionTitle>
                        <FormControl>
                            <InputLabel htmlFor="student-search-filter-name-input">
                                Digite o ID do estudante
                            </InputLabel>
                            <OutlinedInput
                                label="digite-o-ID-do-estudante"
                                id="student-search-filter-name-input"
                            />
                        </FormControl>
                    </Grid>
                </Grid>
            </Stack>
        </div>
    )
};

export default CreatePage;