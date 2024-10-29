// Icons
import { GrBook } from "react-icons/gr";
import { FiPaperclip } from "react-icons/fi";
import { LuPartyPopper } from "react-icons/lu";
import { FaRegPenToSquare } from "react-icons/fa6";
import { IoNewspaperOutline } from "react-icons/io5";
import { BiMessageSquareDetail } from "react-icons/bi";
import { AiOutlineQuestionCircle } from "react-icons/ai";
import { IoMdHelpCircle, IoIosBarcode } from "react-icons/io";
import { FaGithub, FaShareAlt, FaChalkboardTeacher } from "react-icons/fa";

// Types
import type {
    ISideMenuLinkButton,
    ISideMenuItems
} from '../types/propsButtons';

// Features
import { styled } from '@mui/material'

// Components
import { Box } from '@mui/material'

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

export const sidemenu_data_quickaccess: { title: string, items: ISideMenuLinkButton[] } = {
    title: "Acesso Rápido",
    items: [
        {
            title: "Ir para o perfil do criador",
            description: "Aproveite para ver outros projetos!",

            icon: <BorderIcon><FaGithub /></BorderIcon>,
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 1",
            description: "Caso não tenha achado o que procurava.",
            icon: <BorderIcon><LuPartyPopper /></BorderIcon>,
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 2",
            description: "A sorte vem na terceira tentativa.",

            icon: <BorderIcon><FaShareAlt /></BorderIcon>,
            type: "link",
            href: "/",
        },
    ]
};

export const sidemenu_data_classroom: ISideMenuItems = {
    title: "Sala de aula",
    items: [
        {
            title: "Biblioteca de conteúdos",
            icon: <GrBook />,
            type: "link",
            href: "/home/bookshelf",
        },
        {
            title: "Atividades",
            icon: <FaRegPenToSquare />, // might need the other version for hover effect
            type: "link",
            href: "/",
        },
        {
            title: "Aulas digitais",
            icon: <FaChalkboardTeacher />,
            type: "link",
            href: "/",
        },
        {
            title: "Simulados e Provas",
            icon: <IoNewspaperOutline />,
            type: "group",
            items: [
                {
                    title: "Avaliações",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Resultados de Avaliações",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
            ],
        },
        {
            title: "Dúvidas e materiais",
            icon: <AiOutlineQuestionCircle />,
            type: "group",
            items: [
                {
                    title: "Ver materiais e tirar dúvidas",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
                {
                    title: "Minhas dúvidas",
                    icon: <div>I</div>,
                    type: "link",
                    href: "/",
                },
            ],
        },
    ],
};

export const sidemenu_data_utilities: { title: string, items: ISideMenuLinkButton[] } = {
    title: "Utilidades",
    items: [
        {
            title: "Mensagens",
            type: "link",
            href: "/admin/students",
            icon: <BiMessageSquareDetail/>,
        },
        {
            title: "Tutoriais",
            type: "link",
            href: "/admin/students",
            icon: <IoMdHelpCircle/>,
        },
        {
            title: "Dúvidas e materiais",
            type: "link",
            href: "/admin/students",
            icon: <FiPaperclip />,
        },
        {
            title: "Código de acesso",
            type: "link",
            href: "/admin/students",
            icon: <IoIosBarcode />,
        },
    ],
}