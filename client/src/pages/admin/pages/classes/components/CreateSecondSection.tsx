// Components
import Dialog from '@mui/material/Dialog';

import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Box from '@mui/material/Box';

import FormControl from '@mui/material/FormControl';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid2';

import StudentTable from '../components/StudentTable';

// Features
import { useState } from 'react';
import { parseFileData } from '../components/StudentTable';

// Types
import type { ChangeEvent } from 'react';
import type { TableData } from '../components/StudentTable'

export interface ISecondSectionData {
    studentSheet: object | -1
}

/**
 * A input wrapped in a FormControl. Display years as selectable options
 * starting from 2024;
 * 
 * Meant to be used under a CreateFormContext. 
 */
const SecondFormSection = (): JSX.Element => {
    const [isOpen, setIsOpen] = useState(false);
    const [tableData, setTableData] = useState<null | TableData>(null);

    const handleFileInput = async (e: ChangeEvent<HTMLInputElement>) => {
        let files = e.target.files;
        
        if (!files) return;

        let text = await files[0].text();

        let data = parseFileData(text);

        setTableData(data);
        // TODO: Turn file data into sheet
    }

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
                        Vanessa, 1000001, 1 (Para ensino fundamental I)
                    </Typography>
                    <Typography variant="body2" component="p">
                        Lucas, 1000002, 2 (Para ensino fundamental II)
                    </Typography>
                    <Typography variant="body2" component="p">
                        João, 1000003, 3 (Para ensino médio)
                    </Typography>
                </Box>
            </Dialog>
            {
                tableData === null &&
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
            }
            <StudentTable data={tableData}/>
            <Stack direction="row" gap={2} sx={{
                padding: 2
            }}>

            </Stack>
        </FormControl>
    )
}


export default SecondFormSection;