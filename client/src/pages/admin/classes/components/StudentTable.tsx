// Components
import Dialog from '@mui/material/Dialog';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';

import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Grid from '@mui/material/Grid2';
import Box from '@mui/material/Box';
import ErrorBubble from '@components/ErrorBubble';
import ButtonSolid from '@components/ButtonSolid';

// Features
import { InputLabel, MenuItem, OutlinedInput, Select, styled } from '@mui/material';

import { useEffect, useState } from 'react';

// Types
import type { ChangeEvent, SyntheticEvent, Dispatch, SetStateAction } from 'react';
import useContextAwareFetch from '@hooks/useContextAwareFetch';
import SectionTitle from '@components/SectionTitle';
import { EF1, EF2, EM, User, UserSegments } from '../../../../types/server';
import { SERVER_ADDR } from '../../../../configs';
import { useForm } from 'react-hook-form';

export interface TableRow {
    name: string,
    id: string,
    email: string,
    segment: UserSegments,
}

const TableCol = styled(Box)(() => ({
    display: 'flex',
    alignItems: 'center',

    backgroundColor: 'rgb(245,245,245)',
    boxSizing: 'border-box',
    border: 'solid 0.5px rgb(230,230,230)',
    padding: 10,
}))

/**
 * Parses the text of a .txt file in CSV format to an object 
 */
const parseFileData = (
    formData: string
): {
    tableData: TableRow[],
    error: string,
} => {
    const rowDividerRegExp = /[\u00C0-\u00FFa-zA-Z \d]+,\d{7,},[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)/g;

    const rows = [...formData.matchAll(rowDividerRegExp)];

    const tableData: TableRow[] = []
    for (const row of rows) {
        const [
            name,
            id,
            email,
        ]: (number | string)[] = String(row).trim().split(",");

        tableData.push({
            name,
            id,
            email,
            segment: null,
        })
    }

    let error = ""
    if (tableData.length == 0) {
        error = "Nenhuma linha válida encontrada"
    }
    return { tableData: tableData, error: error }
}

/**
    * Turns a number into its text equivalent of school year;
    * 
    * @param Index The column where the text should be relative to.
    * @param Num The number to parse.
    */
export const parseNumIntoText = (index: number, num: number | string): string => {
    if (index == 1) {
        switch (num) {
            case 1:
                return "Ensino Fundamental I";
            case 2:
                return "Ensino Fundamental II";
            case 3:
                return "Ensino Médio";
            default:
                return "Segmento não encontrado"
        }
    } else {
        return num + "º Ano"
    }
}

const FileGuideDialog = ({ isOpen, setIsOpen }: { isOpen: boolean, setIsOpen: Dispatch<SetStateAction<boolean>> }): JSX.Element => {
    const [accordionOpen, setAccordionOpen] = useState<'panel1' | 'panel2' | 'panel3' | false>('panel1')

    const handleAccordionChange =
        (panel: 'panel1' | 'panel2' | 'panel3') => (_event: SyntheticEvent, isExpanded: boolean) => {
            setAccordionOpen(isExpanded ? panel : false);
        };

    return (
        <Dialog open={isOpen} onClose={() => (setIsOpen(bool => !bool))} sx={{
            '& .MuiDialog-paper': {
                maxWidth: '550px',
                minWidth: '200px',
            }
        }}>
            <Accordion
                disableGutters
                elevation={0}
                square
                expanded={accordionOpen == "panel1"}
                onChange={handleAccordionChange("panel1")}
            >
                <AccordionSummary
                    aria-controls="studentTable-panel1-content"
                    id="studentTable-panel1-header"
                    expandIcon={<div>v</div>}
                >
                    <Typography component="h5" variant="h6" sx={{ marginRight: '30px' }}>
                        Descubra o formato Comma Separated Values
                    </Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <Typography>
                        O formato Comma-Separated Values (Valores separados por virgula em português) é um tipo de arquivo de texto que armazena dados em uma estrutura de tabela, separando os valores por vírgulas.
                    </Typography>
                    <Typography>
                        Os arquivos CSV são simples e funcionam na maioria das aplicações que lidam com dados estruturados.
                    </Typography>
                </AccordionDetails>
            </Accordion>
            <Accordion disableGutters elevation={0} square expanded={accordionOpen == "panel2"} onChange={handleAccordionChange("panel2")}>
                <AccordionSummary
                    aria-controls="studentTable-panel2-content"
                    id="studentTable-panel2-header"
                    expandIcon={<div>v</div>}
                >
                    <Typography component="h5" variant="h6" sx={{ marginRight: '30px' }}>
                        Como criar uma panilha de estudantes
                    </Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <Typography>
                        Crie um arquivo de texto (com o final .txt). Dentro dele insira os valores para o nome do estudante, sua matrícula e seu email. Todas essas informações estão disponíveis na area de visão do aluno.
                    </Typography>
                    <Typography>
                        As informações devem estar na seguinte ordem, separados por virgula:
                    </Typography>
                    <ul style={{ marginLeft: '20px' }}>
                        <li>
                            <Typography>
                                Nome do estudante
                            </Typography>

                        </li>
                        <li>
                            <Typography>
                                Plataforma ID (Caso não seja possível identificar por nome, será utilizado para identificar o estudante)
                            </Typography>

                        </li>
                        <li>
                            <Typography>
                                Email (Caso não seja possível identificar por ID, será utilizado para identificar o estudante)
                            </Typography>

                        </li>
                    </ul>
                </AccordionDetails>
            </Accordion>
            <Accordion disableGutters elevation={0} square expanded={accordionOpen == "panel3"} onChange={handleAccordionChange("panel3")}>
                <AccordionSummary
                    aria-controls="studentTable-panel2-content"
                    id="studentTable-panel2-header"
                    expandIcon={<div>v</div>}
                >
                    <Typography component="h5" variant="h6" sx={{ marginRight: '30px' }}>
                        Exemplos de panilhas em formato CSV
                    </Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <p style={{ marginLeft: '5px', fontFamily: 'Verdana', fontSize: '12px' }}>
                        Minha_planilha.txt
                    </p>
                    <div style={{ fontFamily: 'Verdana', fontSize: '12px', boxSizing: 'border-box', color: 'rgb(220,220,220)', backgroundColor: 'rgb(30,30,40)', borderRadius: '4px', border: 'solid 2px rgb(140,140,165)', padding: '5px' }}>
                        <p style={{ margin: '1px' }}>
                            John Doe,1000003,johndoe@gmail.com
                        </p>
                        <p style={{ margin: '1px' }}>
                            Kathrine Moe,1000006,kathmoe@hotmail.com
                        </p>
                        <p style={{ margin: '1px' }}>
                            Philiphs Vicent,1000005,philippps@outlook.com
                        </p>
                    </div>
                </AccordionDetails>
            </Accordion>
        </Dialog>
    )
}

const TableOptions = (): JSX.Element => {
    const [tableRequest, setTableRequest] = useState<TableRow[]>([])
    const [isOpen, setIsOpen] = useState(false);
    const [error, setError] = useState<string>("");
    const [isManualSetup, setIsManualSetup] = useState(false)

    const [response, isProcessingFile, fetchError, send] = useContextAwareFetch<User[], TableRow[]>(
        `${SERVER_ADDR}/api/institution/students`,
        {
            method: 'POST',
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            cache: 'no-cache',
        },
    )

    // Parses the file, sets the table request to trigger the fetch
    const handleFileInput = async (e: ChangeEvent<HTMLInputElement>) => {
        const files = e.target.files;
        if (!files) {
            setError("Falha ao processar o arquivo.");
            return
        };

        const text = await files[0].text();
        const { tableData, error } = parseFileData(text);
        if (error) {
            setError(error)
            return
        }

        setTableRequest(tableData);
    }

    // When triggered, checks if there is data to send and then sends to server
    useEffect(() => {
        if (tableRequest.length == 0) {
            return
        }

        send(tableRequest)
    }, [tableRequest, send])

    if (isManualSetup) {
        return <TableForm
            returnButtonCallback={() => { setIsManualSetup(false) }}
        />
    }
    if (isProcessingFile) {
        return <Typography variant="subtitle2" component="p" sx={{ marginTop: 2, marginX: "auto" }}>
            Processando a planilha CSV. Isso pode demorar alguns instantes.
        </Typography>
    }
    if (response != null) {
        return <div>Resposta recebida e diferente de [null]</div>
    }

    return <Stack direction="column" gap={1}>
        <Typography>
            Carregue uma planilha no formato CSV. <span style={{ color: 'rgb(140,110,210)', cursor: 'pointer', textDecoration: 'underline' }} onClick={() => (setIsOpen(bool => !bool))}>Aprenda sobre o formato CSV.</span>
        </Typography>
        <FileGuideDialog isOpen={isOpen} setIsOpen={setIsOpen} />
        <Grid container spacing={1} sx={{ marginTop: 2 }}>
            <Grid size={{ mobile: 12, xs: 6 }}>
                <FormControl>
                    <Button type="button" sx={{
                        position: 'relative',
                        background: 'rgb(200,200,200,0.1)',
                        minHeight: '35px',
                        padding: 0,
                        color: 'primary.purpleDark',
                        overflow: 'none',
                    }}>
                        <input
                            type="file"
                            accept=".txt"
                            style={{
                                cursor: 'pointer',
                                width: '100%',
                                height: '100%',
                                position: 'absolute',
                                opacity: 0,
                            }}
                            onInput={handleFileInput}
                        />
                        <Typography color="primary.purpleDark" variant="body2" component="p">
                            Enviar uma planilha
                        </Typography>
                    </Button>
                </FormControl>
            </Grid>
            <Grid size={{ mobile: 12, xs: 6 }}>
                <Button
                    onClick={() => { setIsManualSetup(true) }}
                    type="button"
                    sx={{
                        width: '100%',
                        padding: 1,

                        minHeight: '35px',
                        borderRadius: 2,

                        color: 'primary.purpleDark',

                        overflow: 'none',
                    }}
                >
                    <Typography color="primary.purpleDark" variant="body2" component="p">
                        Adicionar estudantes manualmente
                    </Typography>
                </Button>
            </Grid>
        </Grid>
        {
            (error || fetchError) && <ErrorBubble err={fetchError || error || "Erro desconhecido"} />
        }
    </Stack>
}
const TableForm = ({ returnButtonCallback }: { returnButtonCallback: () => void }): JSX.Element => {
    const { register, getValues } = useForm<TableRow>({
        defaultValues: {
            name: "",
            email: "",
            segment: null,
            id: undefined,
        }
    });

    const [students, isLoading, error, send] = useContextAwareFetch<User[], TableRow[]>(
        `${SERVER_ADDR}/api/institution/students`,
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
        send([getValues()]);
    }, [])

    const [id, setId] = useState<number | null>(null)

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
                    onClick={() => { console.log("Delete later", [getValues()]); send([getValues()]) }}
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
                    <Select
                        {...register("segment")}
                        labelId="student-search-filter-segment-label"
                        label="digite-o-segmento-do-estudante"
                    >
                        <MenuItem value={""}>Nenhum</MenuItem>
                        <MenuItem value={EF1}>Ensino Fundamental 1</MenuItem>
                        <MenuItem value={EF2}>Ensino Fundamental 2</MenuItem>
                        <MenuItem value={EM}>Ensino Médio</MenuItem>
                    </Select>
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
                        {...register("id")}
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
                        id="outlined-input-form-sheet-name" />
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
            {
                students == null
                    ? <></>
                    : <Stack gap={1} width="100%">
                        {
                            students.map((student) => {
                                return (
                                    <Box
                                        key={student.id}
                                        sx={{
                                            backgroundColor: "white",
                                            width: "100%",
                                            padding: "0.5rem 1rem",
                                            borderRadius: "1rem",
                                            boxShadow: "0px 0px 1px 4px rgb(150,150,150,0.05)",

                                            flex: 1,
                                        }}
                                    >
                                        <Stack direction="row" justifyContent="space-between" alignItems="center">
                                            <Typography variant="body2" component="p">{student.name}</Typography>
                                            <Typography variant="body1" component="p">PLATAFORMA ID {student.id}</Typography>
                                        </Stack>
                                        <Stack direction="row" justifyContent="space-between" alignItems="center">
                                            <Typography variant="subtitle1" component="p">EMAIL: {student.email}</Typography>
                                            <Typography variant="subtitle1" component="p">{student.segment}</Typography>
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
export const StudentTable = ({ tableData }: { tableData: TableRow[] }): JSX.Element => {
    const ColsDescription = [
        "Segmento",
        "Série/Ano",
        "Nome da Turma",
        "Matrícula",
        "Plataforma ID",
        "Email",
    ]; // some of those will come from db

    return (
        <div style={{
            boxSizing: 'border-box',
            border: 'solid 0.5px rgb(190,190,190)',
            margin: '10px 0px',

            overflow: 'scroll',

            maxWidth: '100%',
            maxHeight: '600px',
        }}>

            <div style={{
                display: 'grid',
                gridTemplateColumns: 'repeat(7, auto)',
                gridTemplateRows: 'auto',
            }}>
                {["", ...ColsDescription].map((descriptionText, i) => {
                    if (!i) {
                        return (
                            <TableCol sx={{ flex: '0 0 50px', justifyContent: 'center' }} />
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
                                {
                                    descriptionText
                                }
                            </Typography>
                        </div>
                    )
                })}
                {
                    // Data columns

                    tableData.map((_d, i) => (
                        <>
                            {[i, ...Object.values(tableData[i])].map((val, i) => {
                                if (!i) {
                                    return (
                                        <TableCol sx={{ flex: '0 0 50px', justifyContent: 'center' }}>
                                            <Typography variant="body2" component="p">
                                                {val}
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
                                            {
                                                i !== 1 && i !== 2
                                                    ? val
                                                    : parseNumIntoText(i, val)
                                            }
                                        </Typography>
                                    </TableCol>
                                )
                            })}
                        </>
                    ))

                }
            </div>
        </div>
    )
}

export { parseFileData };

export default TableOptions