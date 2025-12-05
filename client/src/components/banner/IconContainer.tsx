// Components
import Stack from "@mui/material/Stack"

// Utilities
import { styled } from '@mui/material';

// Models
import type { StackProps } from "@mui/material/Stack"

export default styled(({ children, ...props }: StackProps) => (
    <Stack direction="row" {...props}>{children}</Stack>
))(() => ({
    position: 'absolute',
    left: '-12%',

    justifyContent: "space-between",

    width: '124%',

    transition: 'all 180ms ease-in-out'
}))