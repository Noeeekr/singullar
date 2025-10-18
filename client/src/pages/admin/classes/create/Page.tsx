// Components
import SectionHeader from '@components/headers/sectionHeader/SectionHeader';
import Stack from '@mui/material/Stack';

import CreateForm from './components/Form';

/**
 * 
 * Class Creation Page for admin role
 * 
 * Has a CreateFormContext to comunicate between form sections.. 
 * 
 */
const Create = (): JSX.Element => {
    return (
        <Stack gap={5}>
            <SectionHeader
                title="Cadastro de turmas"
                subtitle="Defina as informações necessárias para criar novas turmas"
            />
            <CreateForm />
        </Stack>

    )
}

export default Create;          