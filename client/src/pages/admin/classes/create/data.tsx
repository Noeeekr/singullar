import FirstSectionInputs, { ButtonAddClass } from './components/FormFirstSection';
import SecondSectionInputs from './components/FormSecondSection';
import ThirdSectionInputs from './components/FormThirdSection';
import FourthSectionInputs from './components/FormFourthSection';

export type FormSection = {
    title: string,
    subtitle: string,
    content: JSX.Element,
    button?: JSX.Element,
}

export const formSections: FormSection[] = [
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
    },
    {
        title: "3. Escolha do professor",
        subtitle: "Escolha o professor para essa turma",
        content: <ThirdSectionInputs />,
    },
    {
        title: "4. Seleção de materiais",
        subtitle: "Escolha os livros a serem disponibilizados para os alunos",
        content: <FourthSectionInputs />,
    },
]