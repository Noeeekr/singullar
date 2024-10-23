// icons

import { GrBook } from "react-icons/gr";
import { VscAccount } from "react-icons/vsc";
import { BiDirections } from "react-icons/bi";
import { GoPerson, GoLock } from "react-icons/go";
import { FaRegPenToSquare } from "react-icons/fa6";
import { AiOutlineQuestionCircle } from "react-icons/ai";
import { LuPartyPopper, LuScanFace } from "react-icons/lu";
import { MdOutlineNotificationsNone } from "react-icons/md";
import { TbSmartHome,  TbSpeakerphone, TbGridDots } from "react-icons/tb";
import { FaChalkboardTeacher, FaGithub, FaShareAlt } from "react-icons/fa";
import { IoIosHelpCircleOutline, IoIosNotificationsOutline } from "react-icons/io";
import { IoNewspaperOutline, IoPeopleOutline, IoExitOutline, IoBarcodeOutline } from "react-icons/io5";

// components;

import NotificationPopupContent from '../../components/NotificationPopupContent'

import {
    Box,
    styled,
} from '@mui/material'

import { 
    ISideMenuItems, 
    ISideMenuLinkButton, 
    ISideMenuButton 
} from '../../types/propsButtons'

import MyAccountPopupContent from '../../components/MyAccountPopupContent'
import QuickAccessPopupContent from './components/QuickAccessPopupContent'

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

const students_sidemenu_data_myaccount: (ISideMenuLinkButton | ISideMenuButton)[] = [
    {
        title: "Dados pessoais e acesso",
        icon: <LuScanFace />,
        type: "link",
        href: "/",
    },
    {
        title: "Responsáveis vinculados",
        icon: <IoPeopleOutline />,
        type: "link",
        href: "/",
    },
    {
        title: "Código de acesso",
        icon: <IoBarcodeOutline />,
        type: "link",
        href: "/",
    },
    {
        title: "Comunicações",
        icon: <TbSpeakerphone />,
        type: "link",
        href: "/",
    },
    {
        title: "Privacidade",
        icon: <GoLock />,
        type: "link",
        href: "/",
    },
    {
        title: "Sair",
        icon: <IoExitOutline />,
        type: "button",
    },
];
const students_sidemenu_data_main: ISideMenuItems = {
    title: "Principal",
    items: [
        {
            title: "Início",
            icon: <TbSmartHome />,
            type: "link",
            href: "/home",
        },
        {
            title: "Notificações",
            icon: <MdOutlineNotificationsNone />, // might need the other version for hover effect
            type: "popup",
            content: <NotificationPopupContent />,
            notifications: true,
        },
        {
            title: "Ajuda",
            icon: <BiDirections />,
            type: "popup",
            content: <NotificationPopupContent />,
        },
        {
            title: "Minha conta",
            icon: <GoPerson />,
            type: "group",
            items: students_sidemenu_data_myaccount
        },
    ],
};
const students_sidemenu_data_classroom: ISideMenuItems = {
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
const students_sidemenu_data_quickaccess: { title: string, items: ISideMenuLinkButton[] } = {
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

export const students_sidemenu_data: ISideMenuItems[] = [
    students_sidemenu_data_main,
    students_sidemenu_data_classroom,
    students_sidemenu_data_quickaccess,
]

export type StudentNavbarPopupId = "help" | "notifications" | "myaccount" | "quickaccess" | "";

export interface INavbarPopupData {
    structure: "side" | "bubble",
    title: string,
    id: StudentNavbarPopupId,
    icon: JSX.Element,
    content: JSX.Element,
}

export const students_navbar_popup_data: INavbarPopupData[] = [
    {
        structure:"side",
        title:"Central de ajuda",
        id:"help",
        icon: <IoIosHelpCircleOutline color="white" fontSize={26} />,
        content: <div>Lorem Ipsum</div>,
    },
    {
        structure:"side",
        title:"Notificações",
        id:"notifications",
        icon: <IoIosNotificationsOutline color="white" fontSize={26} />,
        content: <NotificationPopupContent/>,
    },
    {
        structure: "bubble",
        title:"Minha conta",
        id:"myaccount",
        icon: <VscAccount color="white" fontSize={22} />,
        content: <MyAccountPopupContent items={students_sidemenu_data_myaccount} />,
    },
    {
        structure: "bubble",
        title:"Acesso Rápido",
        id:"quickaccess",
        icon: <TbGridDots color="white" fontSize={23} />,
        content: <QuickAccessPopupContent items={students_sidemenu_data_quickaccess.items} />,
    }
]
