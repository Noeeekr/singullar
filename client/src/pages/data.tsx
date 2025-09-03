import BorderIcon from '@components/BorderIcon.tsx'

import { 
    FaChalkboardTeacher,
    FaGithub,   
    LuPartyPopper,
    FaShareAlt,
    GrBook,
    FaRegPenToSquare,
    IoNewspaperOutline,
    BiMessageSquareDetail,
    AiOutlineQuestionCircle,
    IoMdHelpCircle,
    FiPaperclip,
    IoIosBarcode,
} from "./icons"

// Types
import type {
    ISideMenuLinkButton,
    ISideMenuItems
} from '../types/buttonProps';

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
            icon: <BiMessageSquareDetail />,
        },
        {
            title: "Tutoriais",
            type: "link",
            href: "/admin/students",
            icon: <IoMdHelpCircle />,
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