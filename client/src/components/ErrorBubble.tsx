// Components 
import Box from '@mui/material/Box';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import { 
    HighlightOff
} from '@mui/icons-material'

const ErrorBubble = (
    { err }: { err: string }
): JSX.Element => {
    return (
        <Box
            sx={{
                marginY: 2,
                padding: 2,
                borderRadius: 3.5,
                backgroundColor: (theme) => `${theme.palette.error.light}`,
            }}>
            <Stack
                direction="row"
                spacing={1}
                sx={{
                    alignItems: 'center',
                }}
            >
                <Box
                    sx={{
                        position: 'relative',
                        display: 'flex',
                        '&:before': {
                            position: 'absolute',
                            top: -16,
                            content: '""',
                            backgroundColor: (theme) => theme.palette.error.main,
                            width: '80%',
                            height: 5,
                            marginLeft: '2px',
                            borderRadius: 2,
                        }
                    }}
                >

                    <HighlightOff
                        color="error"
                    />
                </Box>
                <Typography
                    color="error.dark"
                    variant="body1"
                    sx={{
                        fontWeight: 'bold',
                    }}
                >
                    {err}
                </Typography>
            </Stack>
        </Box>
    )
}

export default ErrorBubble