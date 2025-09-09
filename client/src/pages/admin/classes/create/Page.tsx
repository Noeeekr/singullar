// Components
import SectionHeader from '@components/SectionHeader';
import Stack from '@mui/material/Stack';

import FirstSectionInputs, { ButtonAddClass } from '../components/CreateFirstSection';
import SecondSectionInputs from '../components/CreateSecondSection';
import ThirdSectionInputs from '../components/CreateThirdSection';
import CreateForm, { ButtonRedoLastSection } from '../components/CreateForm';

const form = [
    {
        title: "1. Dados gerais",
        subtitle: "Defina as informações básicas da turma",
        content: <FirstSectionInputs />,
        button: <ButtonAddClass />
    },
    {
        title: "2. Adição de estudantes",
        subtitle: "Preencha as informações dos alunos",
        content: <SecondSectionInputs />,
        button: <ButtonRedoLastSection />,
    },
    {
        title: "3. Seleção de materiais",
        subtitle: "",
        content: <ThirdSectionInputs />,
        button: <ButtonRedoLastSection />,
    }
]

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
            <CreateForm form={form}/>
        </Stack>

    )
}

export default Create;          