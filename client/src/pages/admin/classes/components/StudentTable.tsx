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

// Features
import { styled } from '@mui/material';
import { useAppSelector } from '@slices/store';

import { useEffect, useState } from 'react';

// Types
import type { ChangeEvent, SyntheticEvent, Dispatch, SetStateAction } from 'react';

interface TableRow {
    name: string | number,
    id: string | number,
    email: string | number,
}

export type TableData = TableRow[];

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
const parseFileData = (rawData: string) => {
    const rowDividerRegExp = /[\u00C0-\u00FFa-zA-Z \d]+,\d{7,},[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)/g;

    const rows = [...rawData.matchAll(rowDividerRegExp)];

    const data = []
    for (const row of rows) {
        const [
            name,
            id,
            email,
        ]: (number | string)[] = String(row).split(",");

        data.push({
            name,
            id: Number(id),
            email,
        })
    }

    return data     
}

const FileGuideDialog = ({ isOpen, setIsOpen }: { isOpen: boolean, setIsOpen: Dispatch<SetStateAction<boolean>> }): JSX.Element => {
    const [accordionOpen, setAccordionOpen] = useState<'panel1' | 'panel2' | 'panel3' | false>('panel1')
    
    const handleAccordionChange =
    (panel: 'panel1' | 'panel2' | 'panel3') => (_event: SyntheticEvent, isExpanded: boolean) => {
        setAccordionOpen(isExpanded ? panel : false);
    };

    return(
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
                    <Typography component="h5" variant="h6" sx={{ marginRight: '30px'}}>
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
                    expandIcon={<div>d</div>}
                >
                    <Typography component="h5" variant="h6" sx={{ marginRight: '30px'}}>
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
                    <ul style={{ marginLeft: '20px'}}>
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
                    expandIcon={<div>d</div>}
                >
                    <Typography component="h5" variant="h6" sx={{ marginRight: '30px'}}>
                        Exemplos de panilhas em formato CSV
                    </Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <p style={{ marginLeft: '5px', fontFamily: 'Verdana', fontSize: '12px'}}>
                        Minha_planilha.txt
                    </p>
                        <div style={{ fontFamily: 'Verdana', fontSize: '12px',boxSizing: 'border-box', color: 'rgb(220,220,220)', backgroundColor: 'rgb(30,30,40)', borderRadius: '4px', border: 'solid 2px rgb(140,140,165)', padding: '5px'}}>
                            <p style={{ margin: '1px'}}>                                
                                John Doe,1000003,johndoe@gmail.com
                            </p>
                            <p style={{ margin: '1px'}}>
                                Kathrine Moe,1000006,kathmoe@hotmail.com
                            </p>
                            <p style={{ margin: '1px'}}>
                                Philiphs Vicent,1000005,philippps@outlook.com 
                            </p>
                        </div>
                </AccordionDetails>
            </Accordion>
        </Dialog>
    )
}
/**
    * Turns a number into its text equivalent of school year;
    * 
    * @param Index The column where the text should be relative to.
    * @param Num The number to parse.
    */
export const parseNumIntoText = (index: number, num: number | string): string => {
    if (index == 1) {
        switch(num) { 
            case 1:
                return "Ensino Fundamental I";
            case 2:
                return "Ensino Fundamental II";
            case 3:
                return "Ensino Médio";
            default:
                return "Erro 69"
        }
    } else {
        return num + "º Ano"
    }
}

const StudentTable = (): JSX.Element => {
    const [data, setData] = useState<null | TableData>(null);
    const [isFetching, setIsFetching] = useState(false);
    const [error, setError] = useState("");
    const [isOpen, setIsOpen] = useState(false);

    const institution = useAppSelector(store => store.institution)

    // Send CSV file to server
    useEffect(() => {
        if (data === null) {
            return;
        }

        setIsFetching(true);

        fetch('http://localhost:8000/api/user/students/sheet',{
            method: 'POST',
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            cache: 'no-cache',
            body: JSON.stringify({
                students: data,
                institutionId: institution?.id,
            })

        })
        .then(res => {
            if (res.status > 299) {
                setError("um erro aconteceu e não cataloga no set")
                setData(null)

                return
            }

            return res.json()
        })
        .then(res => {
            console.log(res)
            if (res.error) {
                setError(res.error)
                setData(null)  
            }
        })
        .catch(() => {
            setError("Falha ao enviar os dados ao servidor.")
            setData(null)
        })
        .finally(() => {
            setIsFetching(false)
        })
    }, [data, institution?.id]);

    // Process the file
    const handleFileInput = async (e: ChangeEvent<HTMLInputElement>) => {
        const files = e.target.files;
        
        if (!files) return;

        const text = await files[0].text();

        const data = parseFileData(text);

        setData(data);
        // TODO: Turn file data into sheet
    }

    if (isFetching) {
        return(
            <p>
                Processando a planilha CSV. Isso pode demorar alguns instantes.
            </p>
        )
    }

    if (!data) {
        return (
            <Stack direction="column" gap={1}>
                <Typography>
                Carregue uma planilha no formato CSV. <span style={{ color: 'rgb(140,110,210)', cursor: 'pointer', textDecoration: 'underline' }} onClick={() => (setIsOpen(bool => !bool))}>Aprenda sobre o formato CSV.</span>
            </Typography>
            <FileGuideDialog isOpen={isOpen} setIsOpen={setIsOpen}/>
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
                                onChange={handleFileInput}
                            />
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
            {
                error && <ErrorBubble err={error} />
            }
            </Stack>
        )
    }

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
                                <TableCol sx={{ flex: '0 0 50px', justifyContent: 'center' }}/>
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

                data.map((_d, i) => (
                    <>
                        {[i, ...Object.values(data[i])].map((val, i) => {
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
                                                : parseNumIntoText(i,val)
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

export default StudentTable