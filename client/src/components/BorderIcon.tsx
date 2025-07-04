import { 
    Box,
    styled 
} from '@mui/material'

const BorderIcon = styled(({ children, style = {}, ...other }: { children: JSX.Element, style?: object }) => (
    <Box style={{ ...style, color: 'black' }} {...other}>
        {
            children
                ? children
                : <Box sx={{
                    borderRadius: 20,
                    backgroundColor: "rgb(110,110,110)",
                    width: 28,
                    height: 28,
                }} />
        }
    </Box>
))(() => ({
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    width: 44,
    height: 44,
    borderRadius: 8,
    border: 'solid 1px gray',
    overflow: 'hidden',
}))

export default BorderIcon;