import {
    styled,
    Divider,
} from '@mui/material'

const CustomDivider = styled(({ ...props }) => (
    <Divider {...props} />
))(() => ({
    marginX: 1,
    borderBottomWidth: 2,
    borderColor: 'rgb(235,235,235)',
}))

export default CustomDivider