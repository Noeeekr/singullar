// Components
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

// Features
import { styled } from '@mui/material';

// Types

interface TableRow {
    segment: string | number,
    series: string | number,
    className: string | number,
    studentName: string | number,
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
    const rowDividerRegExp = /\d,\d,[\u00C0-\u00FFa-zA-Z \d]+,[\u00C0-\u00FFa-zA-Z \d]+,\d{7,},[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)/g;

    const rows = [...rawData.matchAll(rowDividerRegExp)];

    const data = []
    for (let row of rows) {
        let [
            segment,
            series,
            className,
            studentName,
            id,
            email,
        ]: (number | string)[] = String(row).split(",");

        // "If it passed regex it is confirmed to be a number, right?"
        segment = Number(segment)
        series = Number(series)
        id = Number(id)

        data.push({
            segment,
            series,
            className,
            studentName,
            id,
            email,
        })
    }

    return data     
}

const TestComponentForSheet = ({ data }: { data: null | TableData }): JSX.Element => {
    if (!data) {
        return <></>
    }

    const ColsDescription = [
        "Segmento",
        "Série/Ano",
        "Nome da Turma",
        "Matrícula",
        "Plataforma ID",
        "Email",
    ];

    /**
    * Turns a number into its text equivalent of school year;
    * 
    * @param Index The column where the text should be relative to.
    * @param Num The number to parse.
    */
    const parseNumIntoText = (index: number, num: number | string): string => {
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

export default TestComponentForSheet