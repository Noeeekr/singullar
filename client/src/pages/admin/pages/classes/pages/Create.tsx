// Components
import Dialog from '@mui/material/Dialog';

import SectionHeader from '@components/SectionHeader';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';

import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid2';
import ButtonSolid from '@components/ButtonSolid';

// Features
import { useState, useMemo, useEffect, useContext, createContext } from 'react';

import { styled } from '@mui/material';

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

const FormDataContext = createContext<{
    formData: (IFirstSectionData & ISecondSectionData & IThirdSectionData),
    setFormData: Function,
    setSectionStatus: Function,
}>({
    formData: {
        currentYear: -1,
        segment: -1,
        series: -1,
    },
    setFormData: () => { console.log("If this is showing, it means a context function that's being used is not implemented. This might end in code being executed incorrectly. Please spend some time debugging this page component") },
    setSectionStatus: () => { console.log("If this is showing, it means a context function that's being used is not implemented in some piece of code. This might end in code being executed incorrectly. Please spend some time debugging this page component") }
});

const TableCol = styled(Box)(() => ({
    display: 'flex',
    alignItems: 'center',

    flex: 1,

    backgroundColor: 'rgb(245,245,245)',
    width: '50px',
    height: '50px',
    boxSizing: 'border-box',
    border: 'solid 0.5px rgb(230,230,230)',
    padding: 10,
}))
/**
 * A input wrapped in a FormControl. Display years as selectable options
 * starting from 2024;
 * 
 * Meant to be used under a FormDataContext
 */
const FirstSectionInputs = (): JSX.Element => {
    // Handle year value for first input
    const time = new Date().getFullYear();
    const { setFormData } = useContext(FormDataContext);

    const years = [];
    for (let i = 2024; i <= time; i++) {
        years.push(i)
    }

    // First input handler
    const handleYearChange = (e: SelectChangeEvent<number>) => {
        setFormPiece(formData => ({ ...formData, currentYear: Number(e.target.value) }))
    }

    // Second input handler
    const handleSeriesChange = (e: SelectChangeEvent<number>) => {
        setFormPiece(formData => ({ ...formData, series: Number(e.target.value) }))
    }

    // Third input handler
    const handleSegmentChange = (e: SelectChangeEvent<number>) => {
        setFormPiece(formData => ({ ...formData, segment: Number(e.target.value) }))
    }

    const [formPiece, setFormPiece] = useState<IFirstSectionData>({
        currentYear: -1, // PT: Ano letivo
        segment: -1, // PT: segmento
        series: -1, // PT: série/ano
    })

    useEffect(() => {
        if (Object.values(formPiece).every(field => field != -1)) {
            setFormData(formPiece, 2)
        }
    }, [JSON.stringify(formPiece)])

    return (
        <Stack direction="row" gap={2} padding={4}>
            <FormControl>
                <div style={{ position: "relative", width: "100%" }}>
                    <InputLabel id="class-creation-year-input-label">
                        Ano letivo
                    </InputLabel>
                    <Select
                        labelId="class-creation-year-input-label"
                        label="Ano letivo"
                        value={formPiece.currentYear}
                        aria-label="bal"
                        onChange={(e) => { handleYearChange(e) }}
                    >
                        <MenuItem value={-1}>Escolha uma opção</MenuItem>
                        {
                            years.map((year) => {
                                return (
                                    <MenuItem
                                        value={year}
                                    >
                                        {year}
                                    </MenuItem>
                                )
                            })
                        }
                    </Select>
                </div>
            </FormControl>
            <FormControl>
                <div style={{ position: "relative", width: "100%" }}>
                    <InputLabel id="class-creation-segment-input-label">
                        Segmento
                    </InputLabel>
                    <Select
                        labelId="class-creation-segment-input-label"
                        label="Segmento"
                        value={formPiece.segment}
                        disabled={formPiece.currentYear === -1}
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
                <div style={{ position: "relative", width: "100%" }}>
                    <InputLabel id="class-creation-series-input-label">
                        Série/Ano
                    </InputLabel>
                    {
                        formPiece.segment === -1
                            ? <Select
                                labelId="class-creation-series-input-label"
                                label="Série/Ano"
                                value={formPiece.series}
                                disabled={formPiece.segment === -1}
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
                        formPiece.segment === 1
                            ? <Select
                                labelId="class-creation-series-input-label"
                                label="Série/Ano"
                                value={formPiece.series}
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
                        formPiece.segment === 2
                            ? <Select
                                labelId="class-creation-series-input-label"
                                label="Série/Ano"
                                value={formPiece.series}
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
                        formPiece.segment === 3
                            ? <Select
                                labelId="class-creation-series-input-label"
                                label="Série/Ano"
                                value={formPiece.series}
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
                            : <></>
                    }
                </div>
            </FormControl>
        </Stack>
    )
}

const SecondFormSection = (): JSX.Element => {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <FormControl sx={{ padding: '10px 0px 10px 10px' }}>
            <Typography>
                Carregue uma planilha no formato CSV. <span style={{ color: 'rgb(140,110,210)', cursor: 'pointer', textDecoration: 'underline' }} onClick={() => (setIsOpen(bool => !bool))}>Aprenda sobre o formato CSV.</span>
            </Typography>
            <Dialog open={isOpen} onClose={() => (setIsOpen(bool => !bool))} sx={{
                '& .MuiDialog-paper': {
                    maxWidth: '550px',
                }
            }}>
                <Box sx={{
                    padding: 2,
                }}>
                    <Typography variant="h4" component="h4">
                        Entenda o formato CSV.
                    </Typography>
                    <Typography variant="subtitle2" component="h6" sx={{ marginBottom: 2 }}>
                        Uma demonstração rápida de como usar as planilhas CSV para adicionar alunos.
                    </Typography>
                    <Typography variant="body2" component="p" sx={{ marginBottom: 2 }}>
                        O formato de arquivo CSV, abreviação de Comma Separated Values (em português brasileiro, valores separados por vírgula), é comumente usado para incluir dados de forma padronizada. O padrão CSV é composto por valores em formato de texto separados por 1 (uma) vírgula.
                    </Typography>
                    <Typography variant="body2" component="p" sx={{ marginBottom: 2 }}>
                        Para enviar uma planilha CSV crie um arquivo de texto (com final .txt) e envie no local apropriado.
                    </Typography>
                    <Typography variant="body2" component="p" sx={{ marginBottom: 2 }}>
                        O formato da planilha deve atender ao padrão abaixo, separado por virgula, na ordem :
                    </Typography>
                    <Typography variant="body2" component="p">
                        - Nome
                    </Typography>
                    <Typography variant="body2" component="p">
                        - Platarforma ID
                    </Typography>
                    <Typography variant="body2" component="p">
                        - Segmento
                    </Typography>
                    <Typography variant="subtitle2" component="h6" sx={{ margin: '20px 0px 10px 0px' }}>
                        Exemplo de uma planilha estruturada corretamente:
                    </Typography>
                    <Typography variant="body2" component="p">
                        Vanessa, 1000001, E.F.II
                    </Typography>
                    <Typography variant="body2" component="p">
                        Lucas, 1000002, E.F.I
                    </Typography>
                    <Typography variant="body2" component="p">
                        João, 1000003, E.M
                    </Typography>
                </Box>
            </Dialog>
            <Grid container spacing={1} sx={{ marginTop: 2 }}>
                <Grid size={{ mobile: 12, xs: 6 }}>
                    <FormControl>
                        <Button type="button" sx={{
                            position: 'relative',

                            background: 'rgb(50,50,100,0.05)',
                            minHeight: '35px',
                            borderRadius: 2,
                            padding: 0,

                            color: 'primary.purpleDark',

                            overflow: 'none',
                        }}>
                            <input type="file" accept=".txt" style={{
                                cursor: 'pointer',
                                width: '100%',
                                height: '100%',
                                position: 'absolute',
                                opacity: 0,
                            }} />
                            <Typography color="primary.purpleDark" variant="body2" component="p">
                                Enviar uma planilha
                            </Typography>
                        </Button>
                    </FormControl>
                </Grid>
                <Grid size={{ mobile: 12, xs: 6 }}>
                    <Button type="button" sx={{
                        width: '100%',
                        padding: 1,

                        minHeight: '35px',
                        borderRadius: 2,

                        color: 'primary.purpleDark',

                        overflow: 'none',
                    }}>
                        <Typography color="primary.purpleDark" variant="body2" component="p">
                            Adicionar estudantes manualmente
                        </Typography>
                    </Button>
                </Grid>
            </Grid>
            <Stack direction="row" gap={2} sx={{
                padding: 2
            }}>

            </Stack>
        </FormControl>
    )
}

const ThirdFormSection = (): JSX.Element => {
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

const TestComponentForSheet = (): JSX.Element => {
    const randColNumber = [0, 1, 2, 3, 4, 5]

    const ColsDescription = ["Nome do aluno", "Matricula do aluno", "Segmento", "Turma"]
    const ColsData = ["Pedro", "12303201312", "E.F.II", "6º Ano"]

    return (
        <div style={{
            boxSizing: 'border-box',
            border: 'solid 0.5px rgb(190,190,190)',
        }}>
            {
                // Description Cols
                <div style={{ display: 'flex' }}>
                    {["", ...ColsDescription].map((descriptionText, i) => {
                        if (!i) {
                            return (
                                <div style={{
                                    flexGrow: 0,
                                    flexShrink: 0,
                                    boxSizing: 'border-box',
                                    border: 'solid 0.5px rgb(230,230,230)',
                                    backgroundColor: 'rgb(245,245,245)',
                                    width: '50px',
                                    height: '50px',
                                }}>
                                    <Typography variant="body2" component="p">
                                        {descriptionText}
                                    </Typography>
                                </div>
                            )
                        }

                        return (
                            <div
                                style={{
                                    flex: 1,
                                    boxSizing: 'border-box',
                                    border: 'solid 0.5px rgb(230,230,230)',
                                    backgroundColor: 'rgb(245,245,245)',
                                    padding: 10,
                                }}
                            >
                                <Typography variant="body2" component="p">
                                    {descriptionText}
                                </Typography>
                            </div>
                        )
                    })}
                </div>
            }
            {
                // Data columns

                randColNumber.map((_d, i) => (
                    <div style={{ display: 'flex' }}>
                        {[i, ...ColsData].map((textData, i) => {
                            if (!i) {
                                return (
                                    <TableCol sx={{ flex: '0 0 50px', justifyContent: 'center' }}>
                                        <Typography variant="body2" component="p">
                                            {textData}
                                        </Typography>
                                    </TableCol>
                                )
                            }

                            return (
                                <TableCol sx={{
                                    '&:hover': {
                                        border: 'solid 0.5px rgb(80,80,80)',
                                        backgroundColor: 'rgb(255,215,90)',
                                    }
                                }}>
                                    <Typography variant="body2" component="p" sx={{
                                        '&:hover': {
                                            textDecoration: 'underline',
                                            cursor: 'pointer',
                                        }
                                    }}>
                                        {textData}
                                    </Typography>
                                </TableCol>
                            )
                        })}
                    </div>
                ))

            }
        </div>
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
        setFormData(formData => ({ ...formData, ...data }))
        setSectionStatus(status)
    }

    // define what stage of form is open
    let [sectionsStatus, setSectionStatus] = useState<1 | 2 | 3>(1)

    const form = useMemo(() => ([
        {
            title: "1. Dados gerais",
            subtitle: "Defina as Turmas que serão criadas",
            content: <FirstSectionInputs />,
        },
        {
            title: "2. Adição de estudantes",
            subtitle: "Preencha a planilha de informações",
            content: <SecondFormSection />,
        },
        {
            title: "3. Seleção de materiais",
            subtitle: "",
            content: <ThirdFormSection />,
        }
    ]), []);

    return (
        <Stack gap={5}>
            <SectionHeader
                title="Cadastro de turmas"
            />
            <form>
                <FormDataContext.Provider value={{
                    formData: formData,
                    setFormData: updateFormData,
                    setSectionStatus: setSectionStatus,
                }}>
                    {
                        form.map((data, i) =>
                            <>
                                <Stack marginBottom={4}>
                                    <Stack direction="row">
                                        <Stack direction="column">
                                        <Typography
                                            variant="h5"
                                            component="h4"
                                            color="primary.purpleDark"
                                            fontWeight={600}
                                        >
                                            {data.title}
                                        </Typography>
                                        {
                                            sectionsStatus === i + 1 && <Typography
                                                variant="body1"
                                                component="p"
                                                color="rgb(145,145,145)"
                                                fontWeight={600}
                                            >
                                                {data.subtitle}
                                            </Typography>
                                        }
                                        </Stack>
                                        {
                                            i + 1 > 1 && sectionsStatus === i + 1 && (
                                            <ButtonSolid 
                                                onClick={() => (setSectionStatus(prev => (prev - 1) as 1 | 2 ))}
                                                sx={{ 
                                                    margin: '0px 0px 0px auto', 
                                                }}
                                            >
                                                Refazer última etápa
                                            </ButtonSolid>
                                            )
                                        }
                                    </Stack>
                                    <Box sx={{
                                        height: sectionsStatus === i + 1 ? "100%" : 0,

                                        overflow: "hidden",

                                        marginTop: 0.5,
                                    }}>
                                        {data.content}
                                    </Box>
                                </Stack>
                            </>
                        )
                    }
                </FormDataContext.Provider>
            </form>
        </Stack>

    )
}

export default Create;          