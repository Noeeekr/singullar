// Components
import SectionHeader from '@components/SectionHeader';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';

import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';

// Features
import { useState, useMemo, useEffect } from 'react';

// Types 
interface IFirstSectionData {
    currentYear: number,
    segment: number,
    series: number,
}

interface ISecondSectionData {

}

interface IThirdSectionData {
    
}
/**
 * A input wrapped in a FormControl. Display years as selectable options
 * starting from 2024;
 */
const FirstSectionInputs = ({ updateFormCb }: { updateFormCb: Function }): JSX.Element => {
    // Handle year value for first input
    const time = new Date().getFullYear();
    
    const years = [];
    for (let i = 2024; i <= time; i++) {
        years.push(i)
    }
    
    // First input handler
    const handleYearChange = (e: SelectChangeEvent<number>) => {
        setFormData(formData => ({ ...formData, currentYear: Number(e.target.value) }))
    }
    
    // Second input handler
    const handleSeriesChange = (e: SelectChangeEvent<number>) => {
        setFormData(formData => ({ ...formData, series: Number(e.target.value) }))
    }

    // Third input handler
    const handleSegmentChange = (e: SelectChangeEvent<number>) => {
        setFormData(formData => ({ ...formData, segment: Number(e.target.value)}))
    }
    
    const [formData, setFormData] = useState<IFirstSectionData>({
        currentYear: -1, // PT: Ano letivo
        segment: -1, // PT: segmento
        series: -1, // PT: série/ano
    })
    
    useEffect(() => {
        if (Object.values(formData).every(field => field != -1)) {
            updateFormCb(formData,2)
        }
    },[JSON.stringify(formData)])
    
    return (
        <Stack direction="row" gap={2} padding={4}>
            <FormControl>
                <div style={{ position: "relative", width: "100%"}}>
                    <InputLabel id="class-creation-year-input-label">
                        Ano letivo
                    </InputLabel>
                    <Select 
                        labelId="class-creation-year-input-label" 
                        label="Ano letivo" 
                        value={formData.currentYear}
                        aria-label="bal"
                        onChange={(e) => { handleYearChange(e) }}
                    >
                        <MenuItem value={-1}>Escolha uma opção</MenuItem>
                        { 
                            years.map((year) => {
                                return(
                                    <MenuItem
                                        value={year}
                                    >
                                        { year }
                                    </MenuItem>
                                )
                            }) 
                        }                      
                    </Select>
                </div>
            </FormControl>
            <FormControl>
                <div style={{ position: "relative", width: "100%"}}>
                    <InputLabel id="class-creation-segment-input-label">
                        Segmento
                    </InputLabel>
                    <Select 
                        labelId="class-creation-segment-input-label" 
                        label="Segmento" 
                        value={formData.segment}
                        disabled={formData.currentYear === -1}
                        aria-label="Segmento"
                        onChange={(e) => { handleSegmentChange(e) }}
                    >
                        <MenuItem value={-1}>Escolha uma opção</MenuItem>
                        <MenuItem value={1}>Ensino Fundamental 1</MenuItem>
                        <MenuItem value={2}>Ensino Fundamental 2</MenuItem>
                        <MenuItem value={3}>Ensino Médio</MenuItem>
                    </Select>
                </div>
            </FormControl>
            <FormControl>
                <div style={{ position: "relative", width: "100%"}}>
                    <InputLabel id="class-creation-series-input-label">
                        Série/Ano
                    </InputLabel>
                    {
                        formData.segment === -1
                        ? <Select 
                            labelId="class-creation-series-input-label" 
                            label="Série/Ano" 
                            value={formData.series}
                            disabled={formData.segment === -1}
                            aria-label="Série/Ano"
                            onChange={(e) => { handleSeriesChange(e) }}
                            MenuProps={{
                                PaperProps: {
                                    sx: {
                                        maxHeight: '200px',
                                        marginTop: '10px',
                                    }
                                }
                            }}
                        >
                            <MenuItem value={-1}>...</MenuItem>
                        </Select> 
                        : <></>
                    }
                    {
                        formData.segment === 1
                        ? <Select 
                            labelId="class-creation-series-input-label" 
                            label="Série/Ano" 
                            value={formData.series}
                            aria-label="Série/Ano"
                            onChange={(e) => { handleSeriesChange(e) }}
                            MenuProps={{
                                PaperProps: {
                                    sx: {
                                        maxHeight: '200px',
                                        marginTop: '10px',
                                    }
                                }
                            }}
                        >
                            <MenuItem value={0}>C.A (Classe de alfabetização)</MenuItem>
                            <MenuItem value={1}>1º Ano - Ensino Fundamental I</MenuItem>
                            <MenuItem value={2}>2º Ano - Ensino Fundamental I</MenuItem>
                            <MenuItem value={3}>3º Ano - Ensino Fundamental I</MenuItem>
                            <MenuItem value={4}>4º Ano - Ensino Fundamental I</MenuItem>
                            <MenuItem value={5}>5º Ano - Ensino Fundamental I</MenuItem>
                        </Select>
                        : <></>
                    }
                    {
                        formData.segment === 2 
                        ? <Select 
                            labelId="class-creation-series-input-label" 
                            label="Série/Ano" 
                            value={formData.series}
                            aria-label="Série/Ano"
                            onChange={(e) => { handleSeriesChange(e) }}
                            MenuProps={{
                                PaperProps: {
                                    sx: {
                                        maxHeight: '200px',
                                        marginTop: '10px',
                                    }
                                }
                            }}
                        >
                            <MenuItem value={6}>6º Ano - Ensino Fundamental II</MenuItem>
                            <MenuItem value={7}>7º Ano - Ensino Fundamental II</MenuItem>
                            <MenuItem value={8}>8º Ano - Ensino Fundamental II</MenuItem>
                            <MenuItem value={9}>9º Ano - Ensino Fundamental II</MenuItem>
                        </Select>
                        : <></>
                    }
                    {
                        formData.segment === 3
                        ? <Select 
                            labelId="class-creation-series-input-label" 
                            label="Série/Ano" 
                            value={formData.series}
                            aria-label="Série/Ano"
                            onChange={(e) => { handleSeriesChange(e) }}
                            MenuProps={{
                                PaperProps: {
                                    sx: {
                                        maxHeight: '200px',
                                        marginTop: '10px',
                                    }
                                }
                            }}
                        >   
                            <MenuItem value={10}>1º Ano - Ensino Médio</MenuItem>
                            <MenuItem value={11}>2º Ano - Ensino Médio</MenuItem>
                            <MenuItem value={12}>3º Ano - Ensino Médio</MenuItem>
                        </Select>
                        :<></>
                    }
                </div>
            </FormControl>
        </Stack>
    )
}

const SecondFormSection = ({ updateFormCb }: { updateFormCb: Function }): JSX.Element => {
    return (
        <FormControl>
            <Stack direction="row" gap={2} sx={{
                padding: 2
            }}>

            </Stack>
        </FormControl>
    )
}

const ThirdFormSection = ({ updateFormCb }: { updateFormCb: Function }): JSX.Element => {
    return (
        <FormControl>
            <Stack direction="row" gap={2} sx={{
                padding: 2
            }}>
                <Select>

                </Select>
            </Stack>
        </FormControl>
    )
}

const Create = (): JSX.Element => {
    // Form logic : Union of all inputs data
    const [formData, setFormData] = useState<IFirstSectionData & ISecondSectionData & IThirdSectionData>({
        currentYear: -1,
        segment: -1,
        series: -1,
    });

    // updates
    const updateFormData = (data: IFirstSectionData | ISecondSectionData | IThirdSectionData, status: 1 | 2 | 3) => {
        setFormData(formData => ({ ...formData, ...data}))
        setSectionStatus(status)
    }
    
    // define what stage of form is open
    let [sectionsStatus, setSectionStatus] = useState<1 | 2 | 3>(1)
    
    const form = useMemo(() => ([
        {
            title: "1. Dados gerais",
            subtitle: "Defina as Turmas que serão criadas",
            content: <FirstSectionInputs updateFormCb={updateFormData} />,
        },
        {
            title: "2. Adição de estudantes",
            subtitle: "Preencha a planilha de informações",
            content: <SecondFormSection updateFormCb={updateFormData} />,
        },
        {
            title: "3. Seleção de materiais",
            subtitle: "",
            content: <ThirdFormSection updateFormCb={updateFormData} />,
        }
    ]), []);

    return (
        <Stack gap={5}>
            <SectionHeader
                title="Cadastro de turmas"
            />
            <form>
                {
                    form.map((data, i) =>
                        <>
                            <Stack marginBottom={4}>
                                <Typography
                                    variant="h5"
                                    component="h4"
                                    color="primary.purpleDark"
                                    fontWeight={600}
                                >
                                    {data.title}
                                </Typography>
                                <Box sx={{
                                    height: sectionsStatus === i + 1 ? "100%" : 0,
                                    
                                    overflow: "hidden",

                                    marginTop: 0.5,
                                }}>
                                    <Typography
                                        variant="body1"
                                        component="p"
                                        color="rgb(145,145,145)"
                                        fontWeight={600}
                                    >
                                        {data.subtitle}
                                    </Typography>
                                    {data.content}
                                </Box>
                            </Stack>
                        </>
                    )
                }
            </form>
        </Stack>

    )
}

export default Create;          