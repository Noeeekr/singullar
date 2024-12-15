// Components
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

// Features
import { styled } from '@mui/material';

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

export default TestComponentForSheet