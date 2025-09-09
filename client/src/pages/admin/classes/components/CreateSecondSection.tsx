// Components
import Stack from '@mui/material/Stack';

import FormControl from '@mui/material/FormControl';

import StudentTable from '../components/StudentTable';

// Types
import type { TableData } from '../components/StudentTable'

export interface ISecondSectionData {
    studentSheet: TableData[]
}

/**
 * A input wrapped in a FormControl. Display years as selectable options
 * starting from 2024;
 * 
 * Meant to be used under a CreateFormContext. 
 */
const SecondFormSection = (): JSX.Element => {
    return (
        <FormControl sx={{ padding: '10px 0px 10px 10px' }}>
            <StudentTable />
            <Stack direction="row" gap={2} sx={{
                padding: 2
            }}>

            </Stack>
        </FormControl>
    )
}


export default SecondFormSection;