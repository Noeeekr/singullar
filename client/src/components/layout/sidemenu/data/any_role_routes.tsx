import BorderIcon from '@components/BorderIcon.tsx'

import { BiMessageSquareDetail } from 'react-icons/bi';
import { LuPartyPopper } from 'react-icons/lu';
import { FaChalkboardTeacher, FaGithub, FaShareAlt } from 'react-icons/fa';
import { IoNewspaperOutline } from 'react-icons/io5';
import { GrBook } from 'react-icons/gr';
import { FiPaperclip } from 'react-icons/fi';
import { FaRegPenToSquare } from 'react-icons/fa6';
import { AiOutlineQuestionCircle } from 'react-icons/ai';
import { IoMdHelpCircle, IoIosBarcode } from 'react-icons/io';
import { LinkButtonProps } from '@components/buttons/Link';
import { SideMenuButtons } from '../SideMenuButtons';

export const sidemenu_data_quickaccess: { title: string, menus: LinkButtonProps[], type: "group" } = {
    type: "group",
    title: "Acesso Rápido",
    menus: [
        {
            title: "Ir para o perfil do criador",
            description: "Aproveite para ver outros projetos!",

            icon: {
                component: <BorderIcon><FaGithub /></BorderIcon>,
            },
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 1",
            description: "Caso não tenha achado o que procurava.",
            icon: {
                component: <BorderIcon><LuPartyPopper /></BorderIcon>,
            },
            type: "link",
            href: "/",
        },
        {
            title: "Ir para o perfil do parceiro 2",
            description: "A sorte vem na terceira tentativa.",

            icon: {
                component: <BorderIcon><FaShareAlt /></BorderIcon>,
            },
            type: "link",
            href: "/",
        },
    ]
};

export const sidemenu_data_classroom: SideMenuButtons = {
    title: "Sala de aula",
    menus: [
        {
            title: "Biblioteca de conteúdos",
            icon: {
                component: <GrBook />,
            },
            type: "link",
            href: "/classroom/bookshelf",
        },
        {
            title: "Atividades",
            icon: {
                component: <FaRegPenToSquare />, // might need the other version for hover effect
            },
            type: "link",
            href: "/classroom/exercises",
        },
        {
            title: "Aulas digitais",
            icon: {
                component: <FaChalkboardTeacher />,
            },
            type: "link",
            href: "/classroom/meetings",
        },
        {
            title: "Simulados e Provas",
            icon: {
                component: <IoNewspaperOutline />,
            },
            type: "group",
            menus: [
                {
                    title: "Avaliações",
                    icon: {
                        component: <></>,
                    },
                    type: "link",
                    href: "/classroom/tests",
                },
                {
                    title: "Resultados de Avaliações",
                    icon: {
                        component: <></>,
                    },
                    type: "link",
                    href: "/classroom/tests/results",
                },
            ],
        },
        {
            title: "Dúvidas e materiais",
            icon: {
                component: <AiOutlineQuestionCircle />,
            },
            type: "group",
            menus: [
                {
                    title: "Ver materiais e tirar dúvidas",
                    icon: {
                        component: <></>,
                    },
                    type: "link",
                    href: "/help",
                },
                {
                    title: "Minhas dúvidas",
                    icon: {
                        component: <></>,
                    },
                    type: "link",
                    href: "/help/questions",
                },
            ],
        },
    ],
    type: "group",
};

export const sidemenu_data_utilities: SideMenuButtons = {
    title: "Utilidades",
    menus: [
        {
            title: "Mensagens",
            type: "link",
            href: "/admin/students",
            icon: {
                component: <BiMessageSquareDetail />,
            },
        },
        {
            title: "Tutoriais",
            type: "link",
            href: "/admin/students",
            icon: {
                component: <IoMdHelpCircle />,
            },
        },
        {
            title: "Dúvidas e materiais",
            type: "link",
            href: "/admin/students",
            icon: {
                component: <FiPaperclip />,
            },
        },
        {
            title: "Código de acesso",
            type: "link",
            href: "/admin/students",
            icon: {
                component: <IoIosBarcode />,
            },
        },
    ],
    type: "group",
}