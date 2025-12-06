import {
    styled,
    Typography,
} from '@mui/material'

const SectionTitle = styled(({ children, ...other }: { children: string }) => (
    <Typography variant={"subtitle1"} {...other}>{children}</Typography>
))(({ theme }) => ({
    margin: '0px 10px',
    textTransform: 'uppercase',
    textWrap: 'nowrap',
    color: theme.palette.primary.whiteSemiLow,
}));

export default SectionTitle